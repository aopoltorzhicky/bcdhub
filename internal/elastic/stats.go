package elastic

import "github.com/pkg/errors"

// GetNetworkCountStats -
func (e *Elastic) GetNetworkCountStats(network string) (stats NetworkCountStats, err error) {
	query := newQuery().Query(
		boolQ(
			filter(
				matchQ("network", network),
			),
			should(
				exists("entrypoint"),
				exists("fingerprint"),
			),
			minimumShouldMatch(1),
		),
	).Add(
		aggs("by_index", qItem{
			"terms": qItem{
				"field": "_index",
				"size":  maxQuerySize,
			},
		}),
	).Zero()

	response, err := e.query([]string{DocContracts, DocOperations}, query)
	if err != nil {
		return
	}

	data := response.Get("aggregations.by_index.buckets").Array()
	for _, item := range data {
		key := item.Get("key").String()
		count := item.Get("doc_count").Int()
		switch key {
		case DocContracts:
			stats.Contracts = count
		case DocOperations:
			stats.Operations = count
		default:
			return stats, errors.Errorf("Unknwon index: %s", key)
		}
	}

	return
}

// GetCallsCountByNetwork -
func (e *Elastic) GetCallsCountByNetwork() (map[string]int64, error) {
	query := newQuery().Query(exists("entrypoint")).Add(
		aggs("network", qItem{
			"terms": qItem{
				"field": "network.keyword",
			},
		}),
	).Zero()

	response, err := e.query([]string{DocOperations}, query)
	if err != nil {
		return nil, err
	}

	data := response.Get("aggregations.network.buckets").Array()
	counts := make(map[string]int64)
	for _, item := range data {
		key := item.Get("key").String()
		count := item.Get("doc_count").Int()
		counts[key] = count
	}
	return counts, nil
}

// GetContractStatsByNetwork -
func (e *Elastic) GetContractStatsByNetwork() (map[string]ContractCountStats, error) {
	query := newQuery().Add(
		aggs("network", qItem{
			"terms": qItem{
				"field": "network.keyword",
			},
			"aggs": qItem{
				"same": qItem{
					"cardinality": qItem{
						"script": "doc['fingerprint.parameter'].value + '|' + doc['fingerprint.storage'].value + '|' + doc['fingerprint.code'].value",
					},
				},
				"balance":         sum("balance"),
				"total_withdrawn": sum("total_withdrawn"),
			},
		}),
	).Zero()
	response, err := e.query([]string{DocContracts}, query)
	if err != nil {
		return nil, err
	}

	data := response.Get("aggregations.network.buckets").Array()
	counts := make(map[string]ContractCountStats)
	for _, item := range data {
		key := item.Get("key").String()
		counts[key] = ContractCountStats{
			Total:          item.Get("doc_count").Int(),
			SameCount:      item.Get("same.value").Int(),
			Balance:        item.Get("balance.value").Int(),
			TotalWithdrawn: item.Get("total_withdrawn.value").Int(),
		}
	}
	return counts, nil
}

// GetFACountByNetwork -
func (e *Elastic) GetFACountByNetwork() (map[string]int64, error) {
	query := newQuery().Query(
		in("tags", []string{
			"fa1",
			"fa12",
		}),
	).Add(
		aggs("network", qItem{
			"terms": qItem{
				"field": "network.keyword",
			},
		}),
	).Zero()

	response, err := e.query([]string{DocContracts}, query)
	if err != nil {
		return nil, err
	}

	data := response.Get("aggregations.network.buckets").Array()
	counts := make(map[string]int64)
	for _, item := range data {
		key := item.Get("key").String()
		count := item.Get("doc_count").Int()
		counts[key] = count
	}
	return counts, nil
}

// GetLanguagesForNetwork -
func (e *Elastic) GetLanguagesForNetwork(network string) (map[string]int64, error) {
	query := newQuery().Query(
		boolQ(
			filter(
				matchQ("network", network),
			),
		),
	).Add(
		aggs("languages", qItem{
			"terms": qItem{
				"field": "language.keyword",
			},
		}),
	).Zero()

	response, err := e.query([]string{DocContracts}, query)
	if err != nil {
		return nil, err
	}

	data := response.Get("aggregations.languages.buckets").Array()
	counts := make(map[string]int64)
	for _, item := range data {
		key := item.Get("key").String()
		count := item.Get("doc_count").Int()
		counts[key] = count
	}
	return counts, nil
}
