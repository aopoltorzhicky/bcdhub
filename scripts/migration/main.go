package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/baking-bad/bcdhub/internal/config"
	"github.com/baking-bad/bcdhub/internal/contractparser/consts"
	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/baking-bad/bcdhub/scripts/migration/migrations"
)

func main() {
	migrationMap := map[string]migrations.Migration{
		"timestamp":                 &migrations.SetTimestamp{},
		"language":                  &migrations.SetLanguage{},
		"contract_alias":            &migrations.SetContractAlias{Network: consts.Mainnet},
		"operation_alias":           &migrations.SetOperationAlias{Network: consts.Mainnet},
		"bmd_strings":               &migrations.SetBMDStrings{},
		"bmd_timestamp":             &migrations.SetBMDTimestamp{},
		"fa1_tag":                   &migrations.SetFA1{},
		"operation_strings":         &migrations.SetOperationStrings{},
		"operation_burned":          &migrations.SetOperationBurned{},
		"total_withdrawn":           &migrations.SetTotalWithdrawn{},
		"set_migration_kind":        &migrations.SetMigrationKind{},
		"set_bmd_protocol":          &migrations.SetBMDProtocol{},
		"lost":                      &migrations.FindLostOperations{},
		"contract_migrations_count": &migrations.SetContractMigrationsCount{},
		"state_chain_id":            &migrations.SetStateChainID{},
		"set_alias_slug":            &migrations.SetAliasSlug{},
		"set_contract_fingerprint":  &migrations.SetContractFingerprint{},
		"set_operation_errors":      &migrations.SetOperationErrors{},
	}

	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		logger.Fatal(err)
	}

	env := os.Getenv("MIGRATION")

	if env == "" {
		fmt.Println("Set MIGRATION env variable. Available migrations:")
		for name, m := range migrationMap {
			fmt.Printf("- %s%s| %s\n", name, strings.Repeat(" ", 25-len(name)), m.Description())
		}
		return
	}

	if _, ok := migrationMap[env]; !ok {
		log.Fatal("Unknown migration key: ", env)
	}

	start := time.Now()

	ctx, err := migrations.NewContext(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer ctx.Close()

	logger.Info("Starting %v migration...", env)

	if err := migrationMap[env].Do(ctx); err != nil {
		log.Fatal(err)
	}

	logger.Success("%s migration done. Spent: %v", env, time.Since(start))
}
