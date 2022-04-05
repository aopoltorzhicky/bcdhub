package elastic

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/baking-bad/bcdhub/internal/search"
	"github.com/cenkalti/backoff"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Elastic -
type Elastic struct {
	*elasticsearch.Client

	bulkIndexer esutil.BulkIndexer
}

// New -
func New(addresses []string) (*Elastic, error) {
	retryBackoff := backoff.NewExponentialBackOff()
	elasticConfig := elasticsearch.Config{
		Addresses:     addresses,
		RetryOnStatus: []int{502, 503, 504, 429},
		RetryBackoff: func(i int) time.Duration {
			if i == 1 {
				retryBackoff.Reset()
			}
			return retryBackoff.NextBackOff()
		},
		MaxRetries: 5,
	}
	es, err := elasticsearch.NewClient(elasticConfig)
	if err != nil {
		return nil, err
	}

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client:        es,               // The Elasticsearch client
		NumWorkers:    2,                // The number of worker goroutines
		FlushBytes:    1024 * 512,       // The flush threshold in bytes
		FlushInterval: 10 * time.Second, // The periodic flush interval
	})
	if err != nil {
		return nil, err
	}

	e := &Elastic{es, bi}
	return e, e.ping()
}

// WaitNew -
func WaitNew(addresses []string, timeout int) *Elastic {
	var es *Elastic
	var err error

	for es == nil {
		if es, err = New(addresses); err != nil {
			logger.Warning().Msgf("Waiting elastic up %d seconds...", timeout)
			time.Sleep(time.Second * time.Duration(timeout))
		}
	}

	return es
}

func (e *Elastic) getResponse(resp *esapi.Response, result interface{}) error {
	if resp.IsError() {
		if resp.StatusCode == 404 {
			return NewRecordNotFoundErrorFromResponse(resp)
		}
		return errors.Errorf(resp.String())
	}

	if result == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(result)
}

func (e *Elastic) query(indices []string, query map[string]interface{}, response interface{}) (err error) {
	buf := bytes.NewBuffer([]byte{})
	if err = json.NewEncoder(buf).Encode(query); err != nil {
		return
	}

	// logger.InterfaceToJSON(query)
	// logger.InterfaceToJSON(indices)

	var resp *esapi.Response
	options := []func(*esapi.SearchRequest){
		e.Search.WithContext(context.Background()),
		e.Search.WithIndex(indices...),
		e.Search.WithBody(buf),
	}

	resp, err = e.Search(
		options...,
	)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return e.getResponse(resp, response)
}

func (e *Elastic) ping() (err error) {
	res, err := e.Info()
	if err != nil {
		return
	}
	defer res.Body.Close()

	var result TestConnectionResponse
	return e.getResponse(res, &result)
}

func (e *Elastic) createIndexIfNotExists(index string) error {
	req := esapi.IndicesExistsRequest{
		Index: []string{index},
	}
	res, err := req.Do(context.Background(), e)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if !res.IsError() {
		return nil
	}

	jsonFile, err := os.Open(fmt.Sprintf("mappings/%s.json", index))
	if err != nil {
		res, err = e.Indices.Create(index)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.IsError() {
			return errors.Errorf("%s", res)
		}
		return nil
	}
	defer jsonFile.Close()

	res, err = e.Indices.Create(index, e.Indices.Create.WithBody(jsonFile))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return errors.Errorf("%s", res)
	}
	return nil
}

// CreateIndexes -
func (e *Elastic) CreateIndexes() error {
	for _, index := range search.Indices {
		if err := e.createIndexIfNotExists(index); err != nil {
			return err
		}
	}
	return nil
}

func (e *Elastic) deleteWithQuery(indices []string, query map[string]interface{}) (result *DeleteByQueryResponse, err error) {
	var buf bytes.Buffer
	if err = json.NewEncoder(&buf).Encode(query); err != nil {
		return
	}

	// logger.InterfaceToJSON(query)
	// logger.InterfaceToJSON(indices)

	options := []func(*esapi.DeleteByQueryRequest){
		e.DeleteByQuery.WithContext(context.Background()),
		e.DeleteByQuery.WithConflicts("proceed"),
		e.DeleteByQuery.WithWaitForCompletion(true),
		e.DeleteByQuery.WithRefresh(true),
	}
	resp, err := e.DeleteByQuery(
		indices,
		bytes.NewReader(buf.Bytes()),
		options...,
	)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = e.getResponse(resp, &result)
	return
}

type countResponse struct {
	Count int64 `json:"count"`
}

func (e *Elastic) count(indices []string, query map[string]interface{}) (int64, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return 0, err
	}

	// logger.InterfaceToJSON(query)
	// logger.InterfaceToJSON(indices)

	var resp *esapi.Response
	options := []func(*esapi.CountRequest){
		e.Count.WithContext(context.Background()),
		e.Count.WithIndex(indices...),
		e.Count.WithBody(&buf),
	}

	resp, err := e.Count(
		options...,
	)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var response countResponse
	if err := e.getResponse(resp, &response); err != nil {
		return 0, err
	}
	return response.Count, nil
}
