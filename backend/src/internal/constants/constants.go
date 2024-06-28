package constants

const (
	IndexName         = "movies"
	CsvFilename       = "movies.csv"
	MappingIndexQuery = `{
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
)
