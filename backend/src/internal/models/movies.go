package models

type MoviesResponse struct {
	Movies []Movie `json:"movies"`
}

type Movie struct {
	ID               string `json:"ID"`
	Genre            string `json:"Genre"`
	OriginalLanguage string `json:"Original_Language"`
	Overview         string `json:"Overview"`
	Popularity       string `json:"Popularity"`
	PosterURL        string `json:"Poster_Url"`
	ReleaseDate      string `json:"Release_Date"`
	Title            string `json:"Title"`
	VoteAverage      string `json:"Vote_Average"`
	VoteCount        string `json:"Vote_Count"`
}
