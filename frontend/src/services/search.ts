export interface SearchResponse {
  movies: Movie[];
}

export interface Movie {
  ID: string;
  Genre: string;
  Original_Language: string;
  Overview: string;
  Popularity: string;
  Poster_Url: string;
  Release_Date: string;
  Title: string;
  Vote_Average: string;
  Vote_Count: string;
}

export const searchAPI = async (query: string) => {
  const res = await fetch(`${import.meta.env.VITE_API_URL}/search?q=${query}`);
  const moviesResponse = (await res.json()) as SearchResponse;
  return moviesResponse.movies;
};
