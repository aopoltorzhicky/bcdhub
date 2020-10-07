package elastic

import (
	"strconv"
	"strings"

	"github.com/baking-bad/bcdhub/internal/models"
	"github.com/pkg/errors"
)

// GetTokens -
func (e *Elastic) GetTokens(network, tokenInterface string, lastAction, size int64) ([]models.Contract, int64, error) {
	tags := []string{"fa12", "fa1", "fa2"}
	if tokenInterface == "fa12" || tokenInterface == "fa1" || tokenInterface == "fa2" {
		tags = []string{tokenInterface}
	}

	query := newQuery().Query(
		boolQ(
			filter(
				matchQ("network", network),
				in("tags", tags),
			),
		),
	).Sort("last_action", "desc").Size(size)

	if lastAction != 0 {
		query = query.SearchAfter([]interface{}{lastAction * 1000})
	}

	result, err := e.query([]string{DocContracts}, query)
	if err != nil {
		return nil, 0, err
	}

	contracts := make([]models.Contract, 0)
	for _, hit := range result.Get("hits.hits").Array() {
		var contract models.Contract
		contract.ParseElasticJSON(hit)
		contracts = append(contracts, contract)
	}
	return contracts, result.Get("hits.total.value").Int(), nil
}

// GetTokensStats -
func (e *Elastic) GetTokensStats(network string, addresses, entrypoints []string) (map[string]TokenUsageStats, error) {
	addressFilters := make([]qItem, len(addresses))
	for i := range addresses {
		addressFilters[i] = matchPhrase("destination", addresses[i])
	}

	entrypointFilters := make([]qItem, len(entrypoints))
	for i := range entrypoints {
		entrypointFilters[i] = matchPhrase("entrypoint", entrypoints[i])
	}

	query := newQuery().Query(
		boolQ(
			must(
				matchQ("network", network),
				boolQ(
					should(addressFilters...),
					minimumShouldMatch(1),
				),
				boolQ(
					should(entrypointFilters...),
					minimumShouldMatch(1),
				),
			),
		),
	).Add(
		aggs(
			aggItem{
				"body",
				composite(
					maxQuerySize,
					aggItem{
						"destination", termsAgg("destination.keyword", 0),
					},
					aggItem{
						"entrypoint", termsAgg("entrypoint.keyword", 0),
					},
				).Extend(
					aggs(
						aggItem{
							"average_consumed_gas", avg("result.consumed_gas"),
						},
					),
				),
			},
		),
	).Zero()

	result, err := e.query([]string{DocOperations}, query)
	if err != nil {
		return nil, err
	}

	response := make(map[string]TokenUsageStats)
	buckets := result.Get("aggregations.body.buckets").Array()
	for _, bucket := range buckets {
		address := bucket.Get("key.destination").String()
		method := bucket.Get("key.entrypoint").String()
		usage := TokenMethodUsageStats{
			Count:       bucket.Get("doc_count").Int(),
			ConsumedGas: bucket.Get("average_consumed_gas.value").Int(),
		}

		if _, ok := response[address]; !ok {
			response[address] = make(TokenUsageStats)
		}
		response[address][method] = usage
	}

	return response, nil
}

// GetTokenVolumeSeries -
func (e *Elastic) GetTokenVolumeSeries(network, period string, contracts []string, initiators []string, tokenID uint) ([][]int64, error) {
	hist := qItem{
		"date_histogram": qItem{
			"field":             "timestamp",
			"calendar_interval": period,
		},
	}

	hist.Append("aggs", qItem{
		"result": qItem{
			"sum": qItem{
				"field": "amount",
			},
		},
	})

	matches := []qItem{
		matchQ("network", network),
		matchQ("status", "applied"),
		term("token_id", tokenID),
	}
	if len(contracts) > 0 {
		addresses := make([]qItem, len(contracts))
		for i := range contracts {
			addresses[i] = matchPhrase("contract", contracts[i])
		}
		matches = append(matches, boolQ(
			should(addresses...),
			minimumShouldMatch(1),
		))
	}

	if len(initiators) > 0 {
		addresses := make([]qItem, len(initiators))
		for i := range initiators {
			addresses[i] = matchPhrase("initiator", initiators[i])
		}
		matches = append(matches, boolQ(
			should(addresses...),
			minimumShouldMatch(1),
		))
	}

	query := newQuery().Query(
		boolQ(
			filter(
				matches...,
			),
		),
	).Add(
		aggs(aggItem{"hist", hist}),
	).Zero()

	response, err := e.query([]string{DocTransfers}, query)
	if err != nil {
		return nil, err
	}

	data := response.Get("aggregations.hist.buckets").Array()
	histogram := make([][]int64, 0)
	for _, hit := range data {
		item := []int64{
			hit.Get("key").Int(),
			hit.Get("result.value").Int(),
		}
		histogram = append(histogram, item)
	}
	return histogram, nil
}

// TokenBalance -
type TokenBalance struct {
	Address string
	TokenID int64
}

// GetBalances -
func (e *Elastic) GetBalances(network, contract string, level int64, addresses ...TokenBalance) (map[TokenBalance]int64, error) {
	filters := []qItem{
		matchQ("network", network),
	}

	if contract != "" {
		filters = append(filters, matchPhrase("contract", contract))
	}

	if level > 0 {
		filters = append(filters, rangeQ("level", qItem{
			"lt": level,
		}))
	}

	b := boolQ(
		filter(filters...),
	)

	if len(addresses) > 0 {
		addressFilters := make([]qItem, 0)

		for _, a := range addresses {
			addressFilters = append(addressFilters, boolQ(
				filter(
					matchPhrase("from", a.Address),
					term("token_id", a.TokenID),
				),
			))
		}

		b.Get("bool").Extend(
			should(addressFilters...),
		)
		b.Get("bool").Extend(minimumShouldMatch(1))
	}

	query := newQuery().Query(b).Add(
		qItem{
			"aggs": qItem{
				"balances": qItem{
					"scripted_metric": qItem{
						"init_script": "state.balances = [:]",
						"map_script": `
						if (!state.balances.containsKey(doc['from.keyword'].value)) {
							state.balances[doc['from.keyword'].value + '_' + doc['token_id'].value] = doc['amount'].value;
						} else {
							state.balances[doc['from.keyword'].value + '_' + doc['token_id'].value] = state.balances[doc['from.keyword'].value + '_' + doc['token_id'].value] - doc['amount'].value;
						}
						
						if (!state.balances.containsKey(doc['to.keyword'].value)) {
							state.balances[doc['to.keyword'].value + '_' + doc['token_id'].value] = doc['amount'].value;
						} else {
							state.balances[doc['to.keyword'].value + '_' + doc['token_id'].value] = state.balances[doc['to.keyword'].value + '_' + doc['token_id'].value] + doc['amount'].value;
						}
						`,
						"combine_script": `
						Map balances = [:]; 
						for (entry in state.balances.entrySet()) { 
							if (!balances.containsKey(entry.getKey())) {
								balances[entry.getKey()] = entry.getValue();
							} else {
								balances[entry.getKey()] = balances[entry.getKey()] + entry.getValue();
							}
						} 
						return balances;
						`,
						"reduce_script": `
						Map balances = [:]; 
						for (state in states) { 
							for (entry in state.entrySet()) {
								if (!balances.containsKey(entry.getKey())) {
									balances[entry.getKey()] = entry.getValue();
								} else {
									balances[entry.getKey()] = balances[entry.getKey()] + entry.getValue();
								}
							}
						} 
						return balances;
						`,
					},
				},
			},
		},
	).Zero()
	response, err := e.query([]string{DocTransfers}, query)
	if err != nil {
		return nil, err
	}

	balances := make(map[TokenBalance]int64)
	for address, balance := range response.Get("aggregations.balances.value").Map() {
		parts := strings.Split(address, "_")
		if len(parts) != 2 {
			return nil, errors.Errorf("Invalid addressToken key split size: %d", len(parts))
		}
		tokenID, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return nil, err
		}
		balances[TokenBalance{
			Address: parts[0],
			TokenID: tokenID,
		}] = balance.Int()
	}
	return balances, nil
}

// GetAccountBalances -
func (e *Elastic) GetAccountBalances(network, address string) (map[TokenBalance]int64, error) {
	filters := []qItem{
		matchQ("network", network),
		matchQ("status", "applied"),
	}

	if address != "" {
		filters = append(filters, boolQ(
			should(
				matchPhrase("from", address),
				matchPhrase("to", address),
			),
		))
	}

	b := boolQ(
		filter(filters...),
	)

	query := newQuery().Query(b).Add(
		qItem{
			"aggs": qItem{
				"balances": qItem{
					"scripted_metric": qItem{
						"params": qItem{
							"address": address,
						},
						"init_script": "state.balances = [:]",
						"map_script": `
						def amount = doc['amount'].value;
						if (doc['from.keyword'].value == params.address) {
							amount = -amount;
						}

						def key = doc['contract.keyword'].value + '_' + doc['token_id'].value;
						if (!state.balances.containsKey(key)) {
							state.balances[key] = amount;
						} else {
							state.balances[key] = state.balances[key] + amount;
						}						
						`,
						"combine_script": `
						Map balances = [:]; 
						for (entry in state.balances.entrySet()) { 
							if (!balances.containsKey(entry.getKey())) {
								balances[entry.getKey()] = entry.getValue();
							} else {
								balances[entry.getKey()] = balances[entry.getKey()] + entry.getValue();
							}
						} 
						return balances;
						`,
						"reduce_script": `
						Map balances = [:]; 
						for (state in states) { 
							for (entry in state.entrySet()) {
								if (!balances.containsKey(entry.getKey())) {
									balances[entry.getKey()] = entry.getValue();
								} else {
									balances[entry.getKey()] = balances[entry.getKey()] + entry.getValue();
								}
							}
						} 
						return balances;
						`,
					},
				},
			},
		},
	).Zero()
	response, err := e.query([]string{DocTransfers}, query)
	if err != nil {
		return nil, err
	}

	balances := make(map[TokenBalance]int64)
	for address, balance := range response.Get("aggregations.balances.value").Map() {
		parts := strings.Split(address, "_")
		if len(parts) != 2 {
			return nil, errors.Errorf("Invalid addressToken key split size: %d", len(parts))
		}
		tokenID, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return nil, err
		}
		balances[TokenBalance{
			Address: parts[0],
			TokenID: tokenID,
		}] = balance.Int()
	}
	return balances, nil
}

// TokenSupply -
type TokenSupply struct {
	Supply     float64 `json:"supply"`
	Transfered float64 `json:"transfered"`
}

// GetTokenSupply -
func (e *Elastic) GetTokenSupply(network, address string, tokenID int64) (result TokenSupply, err error) {
	query := newQuery().Query(
		boolQ(
			filter(
				matchQ("network", network),
				matchPhrase("contract", address),
				term("token_id", tokenID),
				matchQ("status", "applied"),
			),
		),
	).Add(
		qItem{
			"aggs": qItem{
				"result": qItem{
					"scripted_metric": qItem{
						"init_script": `state.result = ["supply":0, "transfered":0]`,
						"map_script": `
							if (doc['from.keyword'].value == "") {
								state.result["supply"] = state.result["supply"] + doc["amount"].value;
							} else if (doc['to.keyword'].value == "") {
								state.result["supply"] = state.result["supply"] - doc["amount"].value;
							} else {							
								state.result["transfered"] = state.result["transfered"] + doc["amount"].value;
						}`,
						"combine_script": `return state.result`,
						"reduce_script": `
							Map result = ["supply":0, "transfered":0]; 
							for (state in states) { 
								result["transfered"] = result["transfered"] + state["transfered"];
								result["supply"] = result["supply"] + state["supply"];
							} 
							return result;
						`,
					},
				},
			},
		},
	).Zero()
	response, err := e.query([]string{DocTransfers}, query)
	if err != nil {
		return
	}

	result.Supply = response.Get("aggregations.result.value.supply").Float()
	result.Transfered = response.Get("aggregations.result.value.transfered").Float()

	return
}
