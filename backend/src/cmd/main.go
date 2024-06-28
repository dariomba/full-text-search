package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dariomba/full-text-search/src/internal/constants"
	"github.com/dariomba/full-text-search/src/internal/handlers"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gorilla/mux"
)

var es *elasticsearch.Client

func init() {
	var err error
	cfg := elasticsearch.Config{
		Addresses: []string{"http://127.0.0.1:9200"},
	}
	es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
}

func populateElastic(indexName string, records []map[string]interface{}) {
	indexExistsReq := esapi.IndicesGetRequest{
		Index: []string{indexName},
	}

	resIndex, err := indexExistsReq.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("error trying to find index: %s", err)
	}

	if resIndex.StatusCode == http.StatusNotFound {
		log.Println("index not found, creating index...")
		createIndex(indexName)

		log.Println("indexing records...")
		indexRecords(indexName, records)

		log.Println("data imported successfully")
	}
}

func createIndex(indexName string) {
	mapping := `{
		"settings": {
			"number_of_shards": 1,
			"number_of_replicas": 1,
			"analysis": {
				"analyzer": {
					"autocomplete": {
						"type": "custom",
						"tokenizer": "standard",
						"filter": ["lowercase", "autocomplete_filter"]
					}
				},
				"filter": {
					"autocomplete_filter": {
						"type": "edge_ngram",
						"min_gram": 1,
						"max_gram": 20
					}
				}
			}
		},
		"mappings": {
			"properties": {
				"Release_Date": {
					"type": "date",
					"format": "yyyy-MM-dd"
				},
				"Title": {
					"type": "text",
					"analyzer": "autocomplete",
					"search_analyzer": "standard"
				},
				"Overview": {
					"type": "text",
					"analyzer": "standard"
				},
				"Popularity": {
					"type": "float"
				},
				"Vote_Count": {
					"type": "integer"
				},
				"Vote_Average": {
					"type": "float"
				},
				"Original_Language": {
					"type": "keyword"
				},
				"Genre": {
					"type": "keyword"
				},
				"Poster_Url": {
					"type": "keyword",
					"index": false
				}
			}
		}
	}`

	req := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  bytes.NewReader([]byte(mapping)),
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("cannot create index: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		log.Fatalf("cannot create index: %s", res.String())
	}
}

func indexRecords(indexName string, records []map[string]interface{}) {
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

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("failed to perform bulk request: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("bulk request failed: %s", res.String())
	}

	log.Println("bulk indexing completed successfully")
}

func readCSV(filename string) ([]map[string]interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("could not read CSV headers: %v", err)
	}

	var records []map[string]interface{}
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("could not read CSV line: %v", err)
		}

		record := make(map[string]interface{})
		for i, header := range headers {
			record[header] = line[i]
		}
		records = append(records, record)
	}
	return records, nil
}

func main() {
	records, err := readCSV("datasets/" + constants.CsvFilename)
	if err != nil {
		log.Fatalf("Error reading CSV file: %s", err)
		return
	}

	populateElastic(constants.IndexName, records)

	r := mux.NewRouter()

	handlers.NewSearchHandler(r, es)

	http.Handle("/", r)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
