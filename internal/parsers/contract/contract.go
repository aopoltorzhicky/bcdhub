package contract

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/baking-bad/bcdhub/internal/bcd"
	astContract "github.com/baking-bad/bcdhub/internal/bcd/contract"
	"github.com/baking-bad/bcdhub/internal/config"
	"github.com/baking-bad/bcdhub/internal/models/contract"
	"github.com/baking-bad/bcdhub/internal/models/operation"
	"github.com/baking-bad/bcdhub/internal/models/types"
	"github.com/baking-bad/bcdhub/internal/parsers"
	"github.com/pkg/errors"
)

// Parser -
type Parser struct {
	ctx *config.Context
}

// NewParser -
func NewParser(ctx *config.Context) *Parser {
	return &Parser{ctx: ctx}
}

// Parse -
func (p *Parser) Parse(operation *operation.Operation, symLink string, result *parsers.Result) error {
	if !operation.IsOrigination() {
		return errors.Errorf("invalid operation kind in computeContractMetrics: %s", operation.Kind)
	}

	contract := contract.Contract{
		Network:    operation.Network,
		Level:      operation.Level,
		Timestamp:  operation.Timestamp,
		Manager:    operation.Source,
		Account:    operation.Destination,
		Delegate:   operation.Delegate,
		LastAction: operation.Timestamp,
	}

	if err := p.computeMetrics(operation, symLink, &contract); err != nil {
		return err
	}
	result.Contracts = append(result.Contracts, &contract)
	return nil
}

func (p *Parser) computeMetrics(operation *operation.Operation, symLink string, c *contract.Contract) error {
	script, err := astContract.NewParser(operation.Script)
	if err != nil {
		return errors.Wrap(err, "astContract.NewParser")
	}
	operation.AST = script.Code

	contractScript, err := p.ctx.Scripts.ByHash(script.Hash)
	if err != nil {
		if !p.ctx.Storage.IsRecordNotFound(err) {
			return err
		}
		var s bcd.RawScript
		if err := json.Unmarshal(script.CodeRaw, &s); err != nil {
			return err
		}

		constants, err := script.FindConstants()
		if err != nil {
			return errors.Wrap(err, "script.FindConstants")
		}

		if len(constants) > 0 {
			globalConstants, err := p.ctx.GlobalConstants.All(c.Network, constants...)
			if err != nil {
				return err
			}
			contractScript.Constants = globalConstants
			p.replaceConstants(&contractScript, operation)

			script, err = astContract.NewParser(operation.Script)
			if err != nil {
				return errors.Wrap(err, "astContract.NewParser")
			}
			operation.AST = script.Code
		}

		if err := script.Parse(); err != nil {
			return err
		}

		params, err := script.Code.Parameter.ToTypedAST()
		if err != nil {
			return err
		}

		operation.Script = script.CodeRaw
		contractScript = contract.Script{
			Hash:                 script.Hash,
			Code:                 s.Code,
			Parameter:            s.Parameter,
			Storage:              s.Storage,
			Views:                s.Views,
			FingerprintParameter: script.Fingerprint.Parameter,
			FingerprintCode:      script.Fingerprint.Code,
			FingerprintStorage:   script.Fingerprint.Storage,
			FailStrings:          script.FailStrings.Values(),
			Annotations:          script.Annotations.Values(),
			Tags:                 types.NewTags(script.Tags.Values()),
			Hardcoded:            script.HardcodedAddresses.Values(),
			Entrypoints:          params.GetEntrypoints(),
		}

		switch symLink {
		case bcd.SymLinkAlpha:
			c.Alpha = contractScript
		case bcd.SymLinkBabylon:
			c.Babylon = contractScript
		}
	} else {
		switch symLink {
		case bcd.SymLinkAlpha:
			c.AlphaID = contractScript.ID
			c.Alpha = contractScript
		case bcd.SymLinkBabylon:
			c.BabylonID = contractScript.ID
			c.Babylon = contractScript
		}
	}

	c.Tags = contractScript.Tags

	return nil
}

func (p *Parser) replaceConstants(c *contract.Script, operation *operation.Operation) {
	pattern := `{"prim":"constant","args":[{"string":"%s"}]}`
	for i := range c.Constants {
		operation.Script = bytes.ReplaceAll(
			operation.Script,
			[]byte(fmt.Sprintf(pattern, c.Constants[i].Address)),
			c.Constants[i].Value,
		)
	}
}
