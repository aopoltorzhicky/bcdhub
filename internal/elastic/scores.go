package elastic

import (
	"fmt"

	"github.com/baking-bad/bcdhub/internal/models"
)

var indexToModel = map[string]models.Scorable{
	DocContracts:  models.Contract{},
	DocOperations: models.Operation{},
	DocBigMapDiff: models.BigMapDiff{},
}

// GetSearchScores -
func GetSearchScores(search string, indices []string) ([]string, error) {
	if len(indices) == 0 {
		return nil, nil
	}

	result := make([]string, 0)
	for i := range indices {
		model, ok := indexToModel[indices[i]]
		if !ok {
			return nil, fmt.Errorf("[GetSearchScores] Unknown scorable model: %s", indices[i])
		}
		modelScores := model.GetScores(search)
		result = append(result, modelScores...)
	}

	return result, nil
}
