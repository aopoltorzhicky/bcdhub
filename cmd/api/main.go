package main

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/tidwall/gjson"
	"gopkg.in/go-playground/validator.v9"

	"github.com/baking-bad/bcdhub/cmd/api/handlers"
	"github.com/baking-bad/bcdhub/cmd/api/oauth"
	"github.com/baking-bad/bcdhub/internal/contractparser/cerrors"
	"github.com/baking-bad/bcdhub/internal/database"
	"github.com/baking-bad/bcdhub/internal/elastic"
	"github.com/baking-bad/bcdhub/internal/helpers"
	"github.com/baking-bad/bcdhub/internal/jsonload"
	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/baking-bad/bcdhub/internal/noderpc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	gjson.AddModifier("upper", func(json, arg string) string {
		return strings.ToUpper(json)
	})
	gjson.AddModifier("lower", func(json, arg string) string {
		return strings.ToLower(json)
	})

	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = "config.json"
	}

	var cfg config
	if err := jsonload.StructFromFile(configFile, &cfg); err != nil {
		logger.Fatal(err)
	}

	helpers.InitSentry(cfg.Sentry.Debug)
	helpers.SetTagSentry("project", cfg.Sentry.Project)
	defer helpers.CatchPanicSentry()

	es := elastic.WaitNew([]string{cfg.Search.URI})
	if err := cerrors.LoadErrorDescriptions("data/errors.json"); err != nil {
		logger.Fatal(err)
	}

	db, err := database.New(cfg.DB.URI)
	if err != nil {
		logger.Error(err)
		helpers.CatchErrorSentry(err)
		return
	}
	defer db.Close()

	rpc := createRPC(cfg)

	oauth, err := oauth.New()
	if err != nil {
		logger.Error(err)
		helpers.CatchErrorSentry(err)
	}

	ctx, err := handlers.NewContext(es, rpc, cfg.Dir, db, oauth)
	if err != nil {
		logger.Error(err)
		helpers.CatchErrorSentry(err)
		return
	}

	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("address", handlers.AddressValidator); err != nil {
			logger.Fatal(err)
		}
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("opg", handlers.OPGValidator); err != nil {
			logger.Fatal(err)
		}
	}

	var networks []string
	for network := range cfg.NodeRPC {
		networks = append(networks, network)
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("network", handlers.MakeNetworkValidator(networks)); err != nil {
			logger.Fatal(err)
		}
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("status", handlers.StatusValidator); err != nil {
			logger.Fatal(err)
		}
	}

	r.Use(corsSettings())
	r.Use(helpers.SentryMiddleware())
	v1 := r.Group("v1")
	{
		v1.GET("search", ctx.Search)

		stats := v1.Group("stats")
		{
			stats.GET("", ctx.GetStats)
			networkStats := stats.Group(":network")
			{
				networkStats.GET("", ctx.GetNetworkStats)
				networkStats.GET("series", ctx.GetSeries)
			}
		}

		contract := v1.Group("contract")
		contract.Use(ctx.IsAuthenticated())
		{
			network := contract.Group(":network")
			{
				address := network.Group(":address")
				{
					address.GET("", ctx.GetContract)
					address.GET("code", ctx.GetContractCode)
					address.GET("operations", ctx.GetContractOperations)
					address.GET("migrations", ctx.GetContractMigrations)
					address.GET("storage", ctx.GetContractStorage)
					address.GET("raw_storage", ctx.GetContractStorageRaw)
					address.GET("rich_storage", ctx.GetContractStorageRich)
					address.GET("rating", ctx.GetContractRating)
					address.GET("mempool", ctx.GetMempool)
					address.GET("same", ctx.GetSameContracts)
					address.GET("similar", ctx.GetSimilarContracts)
					bigmap := address.Group("bigmap")
					{
						bigmap.GET(":ptr", ctx.GetBigMap)
						bigmap.GET(":ptr/:key_hash", ctx.GetBigMapByKeyHash)
					}
					entrypoints := address.Group("entrypoints")
					{
						entrypoints.GET("", ctx.GetEntrypoints)
						entrypoints.GET("schema", ctx.GetEntrypointSchema)
						entrypoints.POST("data", ctx.GetEntrypointData)
					}
				}
			}
		}

		fa12 := v1.Group("tokens")
		{
			network := fa12.Group(":network")
			{
				network.GET("", ctx.GetFA12)
				address := network.Group(":address")
				{
					address.GET("transfers", ctx.GetFA12OperationsForAddress)
				}
			}
		}

		v1.GET("pick_random", ctx.GetRandomContract)
		v1.POST("diff", ctx.GetDiff)
		v1.GET("opg/:hash", ctx.GetOperation)
		v1.GET("projects", ctx.GetProjects)
		v1.GET("formatter", ctx.GetFormatter)

		// PRIVATE
		// TODO - remove in prod
		// v1.POST("vote", ctx.Vote)
		// v1.POST("set_alias", ctx.SetAlias)

		oauth := v1.Group("oauth")
		{
			provider := oauth.Group(":provider")
			{
				provider.GET("login", ctx.OauthLogin)
				provider.GET("callback", ctx.OauthCallback)
			}
		}

		authorized := v1.Group("/")
		authorized.Use(ctx.AuthJWTRequired())
		{
			profile := authorized.Group("profile")
			{
				profile.GET("", ctx.GetUserProfile)
				subscriptions := profile.Group("subscriptions")
				{
					subscriptions.GET("", ctx.ListSubscriptions)
					subscriptions.POST("", ctx.CreateSubscription)
					subscriptions.DELETE("", ctx.DeleteSubscription)
					subscriptions.GET("recommended", ctx.Recommendations)
					subscriptions.GET("timeline", ctx.GetTimeline)
				}
			}
		}
	}
	if err := r.Run(cfg.Address); err != nil {
		logger.Error(err)
		helpers.CatchErrorSentry(err)
		return
	}
}

func corsSettings() gin.HandlerFunc {
	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = []string{"*"}
	cfg.AllowCredentials = true
	cfg.AddAllowHeaders("X-Requested-With", "Authorization")
	return cors.New(cfg)
}

func createRPC(cfg config) map[string]noderpc.Pool {
	rpc := make(map[string]noderpc.Pool)
	for network, hosts := range cfg.NodeRPC {
		rpc[network] = noderpc.NewPool(hosts, time.Second*30)
	}
	return rpc
}
