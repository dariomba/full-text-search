package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dariomba/full-text-search/src/internal/models"
	services "github.com/dariomba/full-text-search/src/internal/services/elastic"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gorilla/mux"
)

type SearchHandler struct {
	elasticClient  *elasticsearch.Client
	elasticService services.ElasticService
}

func NewSearchHandler(
	app *mux.Router,
	es *elasticsearch.Client,
	elasticService services.ElasticService,
) {
	searchHandler := SearchHandler{
		elasticClient:  es,
		elasticService: elasticService,
	}

	app.HandleFunc("/search", searchHandler.search).Methods("GET")
}

func (s SearchHandler) search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	movies, err := s.elasticService.SearchMovies(r.Context(), query)
	if err != nil {
		log.Println("an error has ocurred searching movies", err)
		http.Error(w, "Error searching movies", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		models.MoviesResponse{
			Movies: movies,
		})
}
