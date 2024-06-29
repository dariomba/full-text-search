package models

type ElasticResponse struct {
	Shards   Shards `json:"_shards"`
	Hits     Hits   `json:"hits"`
	TimedOut bool   `json:"timed_out"`
	Took     int    `json:"took"`
}

type Shards struct {
	Failed     int `json:"failed"`
	Skipped    int `json:"skipped"`
	Successful int `json:"successful"`
	Total      int `json:"total"`
}

type Hits struct {
	Hits     []Hit   `json:"hits"`
	MaxScore float64 `json:"max_score"`
	Total    Total   `json:"total"`
}

type Hit struct {
	ID    string  `json:"_id"`
	Index string  `json:"_index"`
	Score float64 `json:"_score"`
	Movie Movie   `json:"_source"`
}

type Total struct {
	Relation string `json:"relation"`
	Value    int    `json:"value"`
}
