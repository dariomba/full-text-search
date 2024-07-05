package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dariomba/full-text-search/src/internal/constants"
	"github.com/dariomba/full-text-search/src/internal/handlers"
	csv "github.com/dariomba/full-text-search/src/internal/services/csv"
	elastic "github.com/dariomba/full-text-search/src/internal/services/elastic"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{os.Getenv("ELASTIC_ADDRESS")},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
		return
	}

	elasticService := elastic.NewElasticService(es)

	records, err := csv.ReadCSV("datasets/" + constants.CsvFilename)
	if err != nil {
		log.Fatalf("Error reading CSV file: %s", err)
		return
	}

	elasticService.PopulateElastic(constants.IndexName, records)

	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("ALLOWED_HOSTS_URLS"), ","),
		AllowCredentials: true,
	})

	handlers.NewSearchHandler(r, es, elasticService)

	handler := c.Handler(r)

	http.Handle("/", handler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
