package migrations

import (
	"log"
	"time"

	"github.com/baking-bad/bcdhub/internal/elastic"
	"github.com/baking-bad/bcdhub/internal/metrics"
)

// SetOperationStrings - migration that set storage and parameter strings array at operation
type SetOperationStrings struct{}

// Description -
func (m *SetOperationStrings) Description() string {
	return "parse parameter & storage strings"
}

// Do - migrate function
func (m *SetOperationStrings) Do(ctx *Context) error {
	log.Print("Start SetOperationStrings migration...")
	start := time.Now()

	for _, network := range ctx.Config.Migrations.Networks {
		operations, err := ctx.ES.GetAllOperations(network)
		if err != nil {
			return err
		}
		log.Printf("Found %d operations", len(operations))

		var lastIdx int
		h := metrics.New(ctx.ES, ctx.DB)
		for i := range operations {
			log.Printf("Compute for operation with id: %s", operations[i].ID)
			h.SetOperationStrings(&operations[i])

			if (i%1000 == 0 || i == len(operations)-1) && i > 0 {
				log.Printf("Saving updated data from %d to %d...", lastIdx, i)
				updates := make([]elastic.Identifiable, len(operations[lastIdx:i]))
				for j := range operations[lastIdx:i] {
					updates[j] = operations[lastIdx:i][j]
				}
				if err := ctx.ES.BulkUpdate("operation", updates); err != nil {
					return err
				}
				lastIdx = i
			}
		}
	}

	log.Printf("Time spent: %v", time.Since(start))

	return nil
}
