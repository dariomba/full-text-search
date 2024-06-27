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

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gorilla/mux"
)

const (
	csvFile   = "movies.csv"
	indexName = "movies"
)

var es *elasticsearch.Client

func init() {
	var err error
	es, err = elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	var buf map[string]interface{}
	err := json.Unmarshal([]byte(fmt.Sprintf(`{"query": {"match": {"name": "%s"}}}`, query)), &buf)
	if err != nil {
		http.Error(w, "Error parsing the query", http.StatusInternalServerError)
		return
	}

	var b []byte
	b, err = json.Marshal(buf)
	if err != nil {
		http.Error(w, "Error marshaling the query", http.StatusInternalServerError)
		return
	}

	req := esapi.SearchRequest{
		Index: []string{"students"},
		Body:  bytes.NewReader(b),
	}

	res, err := req.Do(r.Context(), es)
	if err != nil {
		http.Error(w, "Error getting response from Elasticsearch", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		http.Error(w, res.String(), http.StatusInternalServerError)
		return
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		http.Error(w, "Error parsing the response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func createIndex(indexName string) {
	req := esapi.IndicesCreateRequest{
		Index: indexName,
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Cannot create index: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		log.Fatalf("Cannot create index: %s", res.String())
	}
}

func indexRecords(indexName string, records []map[string]interface{}) {
	for _, record := range records {
		data, err := json.Marshal(record)
		if err != nil {
			log.Fatalf("Cannot encode record: %s", err)
		}

		req := esapi.IndexRequest{
			Index:   indexName,
			Body:    bytes.NewReader(data),
			Refresh: "true",
		}

		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Printf("[%s] Error indexing document ID=%s", res.Status(), record["id"])
		} else {
			log.Printf("[%s] Document indexed successfully", res.Status())
		}
	}
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
	records, err := readCSV("datasets/" + csvFile)
	if err != nil {
		log.Fatalf("Error reading CSV file: %s", err)
	}

	log.Println(records)

	createIndex(indexName)
	indexRecords(indexName, records)

	log.Println("Data imported successfully")

	r := mux.NewRouter()
	r.HandleFunc("/search", searchHandler).Methods("GET")

	http.Handle("/", r)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
