package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/baking-bad/bcdhub/internal/models"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
)

// BulkInsert -
func (e *Elastic) BulkInsert(index string, buf *bytes.Buffer) error {
	req := esapi.BulkRequest{
		Index:   index,
		Body:    bytes.NewReader(buf.Bytes()),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), e)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = e.getResponse(res)
	return err
}

// BulkInsertOperations -
func (e *Elastic) BulkInsertOperations(v []models.Operation) error {
	bulk := bytes.NewBuffer([]byte{})
	for i := range v {
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id": "%s"} }%s`, v[i].ID, "\n"))
		data, err := json.Marshal(v[i])
		if err != nil {
			return err
		}
		data = append(data, "\n"...)

		bulk.Grow(len(meta) + len(data))
		bulk.Write(meta)
		bulk.Write(data)
	}
	return e.BulkInsert(DocOperations, bulk)
}

// BulkInsertContracts -
func (e *Elastic) BulkInsertContracts(v []models.Contract) error {
	bulk := bytes.NewBuffer([]byte{})
	for i := range v {
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id": "%s"} }%s`, v[i].ID, "\n"))
		data, err := json.Marshal(v[i])
		if err != nil {
			return err
		}
		data = append(data, "\n"...)

		bulk.Grow(len(meta) + len(data))
		bulk.Write(meta)
		bulk.Write(data)
	}
	return e.BulkInsert(DocContracts, bulk)
}

// BulkSaveBigMapDiffs -
func (e *Elastic) BulkSaveBigMapDiffs(diffs []models.BigMapDiff) error {
	bulk := bytes.NewBuffer([]byte{})
	for i := range diffs {
		id := uuid.New().String()
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id": "%s"} }%s`, id, "\n"))
		data, err := json.Marshal(diffs[i])
		if err != nil {
			log.Println(err)
			continue
		}
		data = append(data, "\n"...)

		bulk.Grow(len(meta) + len(data))
		bulk.Write(meta)
		bulk.Write(data)
	}
	return e.BulkInsert(DocBigMapDiff, bulk)
}

// BulkInsertMigrations -
func (e *Elastic) BulkInsertMigrations(v []models.Migration) error {
	bulk := bytes.NewBuffer([]byte{})
	for i := range v {
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id": "%s"} }%s`, v[i].ID, "\n"))
		data, err := json.Marshal(v[i])
		if err != nil {
			return err
		}
		data = append(data, "\n"...)

		bulk.Grow(len(meta) + len(data))
		bulk.Write(meta)
		bulk.Write(data)
	}
	return e.BulkInsert(DocMigrations, bulk)
}
