package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dariomba/full-text-search/src/internal/constants"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gorilla/mux"
)

type SearchHandler struct {
	elasticClient *elasticsearch.Client
}

func NewSearchHandler(
	app *mux.Router,
	es *elasticsearch.Client,
) {
	searchHandler := SearchHandler{
		elasticClient: es,
	}

	app.HandleFunc("/search", searchHandler.search).Methods("GET")
}

func (s SearchHandler) search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

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

	res, err := req.Do(r.Context(), s.elasticClient)
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
