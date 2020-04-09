package migrations

import (
	"log"
	"time"

	"github.com/baking-bad/bcdhub/internal/metrics"
)

// SetBMDStrings - migration that set key and value strings array at big map diff
type SetBMDStrings struct{}

// Do - migrate function
func (m *SetBMDStrings) Do(ctx *Context) error {
	log.Print("Start SetBMDStrings migration...")
	start := time.Now()
	allBigMapDiffs, err := ctx.ES.GetAllBigMapDiff()
	if err != nil {
		return err
	}
	log.Printf("Found %d big map diff", len(allBigMapDiffs))

	opIDs := make(map[string]struct{})
	for i := range allBigMapDiffs {
		opIDs[allBigMapDiffs[i].OperationID] = struct{}{}
	}

	log.Printf("Found %d unique operations with big map diff", len(opIDs))

	h := metrics.New(ctx.ES, ctx.DB)
	for id := range opIDs {
		log.Printf("Compute for operation with id: %s", id)
		if err := h.SetBigMapDiffsStrings(id); err != nil {
			return err
		}
	}

	log.Printf("Time spent: %v", time.Since(start))

	return nil
}
