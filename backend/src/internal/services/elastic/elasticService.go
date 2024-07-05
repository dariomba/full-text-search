package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dariomba/full-text-search/src/internal/constants"
	"github.com/dariomba/full-text-search/src/internal/models"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ElasticService struct {
	client *elasticsearch.Client
}

func NewElasticService(client *elasticsearch.Client) ElasticService {
	return ElasticService{
		client: client,
	}
}

func (e ElasticService) SearchMovies(ctx context.Context, query string) ([]models.Movie, error) {
	esQuery := fmt.Sprintf(`{
		"query": {
			"match_phrase_prefix": {
				"Title": "%s"
			}
		}
	}`, query)

	req := esapi.SearchRequest{
		Index: []string{constants.IndexName},
		Body:  bytes.NewReader([]byte(esQuery)),
	}

	res, err := req.Do(ctx, e.client)
	if err != nil {
		return []models.Movie{}, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return []models.Movie{}, err
	}

	var elasticResponse models.ElasticResponse
	if err := json.NewDecoder(res.Body).Decode(&elasticResponse); err != nil {
		return []models.Movie{}, err
	}

	movies := []models.Movie{}

	for _, hit := range elasticResponse.Hits.Hits {
		movies = append(movies, hit.NewMovie())
	}

	return movies, nil
}

func (e ElasticService) PopulateElastic(indexName string, records []map[string]interface{}) {
	indexExistsReq := esapi.IndicesGetRequest{
		Index: []string{indexName},
	}

	resIndex, err := indexExistsReq.Do(context.Background(), e.client)
	if err != nil {
		log.Fatalf("error trying to find index: %s", err)
	}

	if resIndex.StatusCode == http.StatusNotFound {
		log.Println("index not found, creating index...")
		e.createIndex(indexName)

		log.Println("indexing records...")
		e.indexRecords(indexName, records)

		log.Println("data imported successfully")
	}
}

func (e ElasticService) indexRecords(indexName string, records []map[string]interface{}) {
	var buf bytes.Buffer

	for _, record := range records {
		meta := []byte(fmt.Sprintf(`{ "index" : { "_index" : "%s" } }%s`, indexName, "\n"))
		data, err := json.Marshal(record)
		if err != nil {
			log.Fatalf("cannot encode record %v: %s", record, err)
		}

		data = append(data, "\n"...)

		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)
	}

	req := esapi.BulkRequest{
		Body:    bytes.NewReader(buf.Bytes()),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), e.client)
	if err != nil {
		log.Fatalf("failed to perform bulk request: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("bulk request failed: %s", res.String())
	}

	log.Println("bulk indexing completed successfully")
}

func (e ElasticService) createIndex(indexName string) {
	mapping := constants.MappingIndexQuery

	req := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  bytes.NewReader([]byte(mapping)),
	}
	res, err := req.Do(context.Background(), e.client)
	if err != nil {
		log.Fatalf("cannot create index: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		log.Fatalf("cannot create index: %s", res.String())
	}
}
