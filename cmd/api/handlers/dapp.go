package handlers

import (
	"net/http"

	"github.com/baking-bad/bcdhub/internal/helpers"
	"github.com/baking-bad/bcdhub/internal/models/dapp"
	"github.com/baking-bad/bcdhub/internal/models/tokenmetadata"
	"github.com/baking-bad/bcdhub/internal/models/types"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// GetDAppList -
func (ctx *Context) GetDAppList(c *gin.Context) {
	dapps, err := ctx.DApps.All()
	if err != nil {
		if ctx.Storage.IsRecordNotFound(err) {
			c.SecureJSON(http.StatusOK, []interface{}{})
			return
		}
		ctx.handleError(c, err, 0)
		return
	}

	results := make([]DApp, len(dapps))
	for i := range dapps {
		result, err := ctx.appendDAppInfo(dapps[i], false)
		if ctx.handleError(c, err, 0) {
			return
		}
		results[i] = result
	}

	c.SecureJSON(http.StatusOK, results)
}

// GetDApp -
func (ctx *Context) GetDApp(c *gin.Context) {
	var req getDappRequest
	if err := c.BindUri(&req); ctx.handleError(c, err, http.StatusBadRequest) {
		return
	}

	dapp, err := ctx.DApps.Get(req.Slug)
	if err != nil {
		if ctx.Storage.IsRecordNotFound(err) {
			c.SecureJSON(http.StatusNoContent, gin.H{})
			return
		}
		ctx.handleError(c, err, 0)
		return
	}

	response, err := ctx.appendDAppInfo(dapp, true)
	if ctx.handleError(c, err, 0) {
		return
	}

	c.SecureJSON(http.StatusOK, response)
}

// GetDexTokens -
func (ctx *Context) GetDexTokens(c *gin.Context) {
	var req getDappRequest
	if err := c.BindUri(&req); ctx.handleError(c, err, http.StatusBadRequest) {
		return
	}

	dapp, err := ctx.DApps.Get(req.Slug)
	if err != nil {
		if ctx.Storage.IsRecordNotFound(err) {
			c.SecureJSON(http.StatusNoContent, gin.H{})
			return
		}
		ctx.handleError(c, err, 0)
		return
	}
	if !helpers.StringInArray("DEX", dapp.Categories) {
		ctx.handleError(c, errors.New("dapp is not DEX"), http.StatusBadRequest)
		return
	}

	if len(dapp.DexTokens) == 0 {
		c.SecureJSON(http.StatusOK, dapp.DexTokens)
		return
	}

	dexTokens := make([]TokenMetadata, 0)

	for _, token := range dapp.DexTokens {
		tokenMetadata, err := ctx.TokenMetadata.GetAll(tokenmetadata.GetContext{
			Contract: token.Contract,
			Network:  types.Mainnet,
			TokenID:  &token.TokenID,
		})
		if err != nil {
			if ctx.Storage.IsRecordNotFound(err) {
				continue
			}
			ctx.handleError(c, err, 0)
			return
		}

		initiators := make(map[string]struct{})
		entrypoints := make(map[string]struct{})
		for _, c := range dapp.Contracts {
			initiators[c.Address] = struct{}{}
			for i := range c.Entrypoint {
				entrypoints[c.Entrypoint[i]] = struct{}{}
			}
		}

		initiatorsArr := make([]string, 0)
		for address := range initiators {
			initiatorsArr = append(initiatorsArr, address)
		}

		entrypointsArr := make([]string, 0)
		for entrypoint := range entrypoints {
			entrypointsArr = append(entrypointsArr, entrypoint)
		}

		vol, err := ctx.Transfers.GetToken24HoursVolume(types.Mainnet, token.Contract, initiatorsArr, entrypointsArr, token.TokenID)
		if err != nil {
			if ctx.Storage.IsRecordNotFound(err) {
				continue
			}
			ctx.handleError(c, err, 0)
			return
		}

		for i := range tokenMetadata {
			tm := TokenMetadataFromElasticModel(tokenMetadata[i], true)
			tm.Volume24Hours = &vol
			dexTokens = append(dexTokens, tm)
		}
	}
	c.SecureJSON(http.StatusOK, dexTokens)
}

// GetDexTezosVolume -
func (ctx *Context) GetDexTezosVolume(c *gin.Context) {
	var req getDappRequest
	if err := c.BindUri(&req); ctx.handleError(c, err, http.StatusBadRequest) {
		return
	}

	dapp, err := ctx.DApps.Get(req.Slug)
	if err != nil {
		if ctx.Storage.IsRecordNotFound(err) {
			c.SecureJSON(http.StatusNoContent, gin.H{})
			return
		}
		ctx.handleError(c, err, 0)
		return
	}

	if !helpers.StringInArray("DEX", dapp.Categories) {
		ctx.handleError(c, errors.New("dapp is not DEX"), http.StatusBadRequest)
		return
	}

	if len(dapp.Contracts) == 0 {
		c.SecureJSON(http.StatusOK, 0)
	}

	var volume float64
	for _, address := range dapp.Contracts {
		vol, err := ctx.Operations.GetContract24HoursVolume(types.Mainnet, address.Address, address.Entrypoint)
		if ctx.handleError(c, err, 0) {
			return
		}
		volume += vol
	}
	c.SecureJSON(http.StatusOK, volume)
}

func (ctx *Context) appendDAppInfo(dapp dapp.DApp, withDetails bool) (DApp, error) {
	result := DApp{
		Name:             dapp.Name,
		ShortDescription: dapp.ShortDescription,
		FullDescription:  dapp.FullDescription,
		WebSite:          dapp.WebSite,
		Slug:             dapp.Slug,
		Authors:          dapp.Authors,
		SocialLinks:      dapp.SocialLinks,
		Interfaces:       dapp.Interfaces,
		Categories:       dapp.Categories,
		Soon:             dapp.Soon,
	}

	if len(dapp.Pictures) > 0 {
		screenshots := make([]Screenshot, 0)
		for _, pic := range dapp.Pictures {
			switch pic.Type {
			case "logo":
				result.Logo = pic.Link
			case "cover":
				result.Cover = pic.Link
			default:
				screenshots = append(screenshots, Screenshot{
					Type: pic.Type,
					Link: pic.Link,
				})
			}
		}

		result.Screenshots = screenshots
	}

	if withDetails {
		if len(dapp.Contracts) > 0 {
			result.Contracts = make([]DAppContract, 0)

			for _, address := range dapp.Contracts {
				contract, err := ctx.Contracts.Get(types.Mainnet, address.Address)
				if err != nil {
					if ctx.Storage.IsRecordNotFound(err) {
						continue
					}
					return result, err
				}
				result.Contracts = append(result.Contracts, DAppContract{
					Network:     contract.Network.String(),
					Address:     contract.Account.Address,
					Alias:       contract.Account.Alias,
					ReleaseDate: contract.Timestamp.UTC(),
				})

				if address.WithTokens {
					metadata, err := ctx.TokenMetadata.GetAll(tokenmetadata.GetContext{
						Contract: address.Address,
						Network:  types.Mainnet,
						TokenID:  nil,
					})
					if err != nil {
						return result, err
					}
					tokens, err := ctx.addSupply(metadata)
					if err != nil {
						return result, err
					}
					result.Tokens = append(result.Tokens, tokens...)
				}
			}
		}
	}

	return result, nil
}
