import { useDebounce } from '@uidotdev/usehooks';
import { useCallback, useEffect, useRef, useState } from 'react';
import { Movie, searchAPI } from '../../services/search';
import { Movies } from '../Movies/Movies';

const DEBOUNCE_DELAY = 20;

export const Search: React.FC = () => {
  const [search, setSearch] = useState(() => {
    const searchParams = new URLSearchParams(window.location.search);
    return searchParams.get('q') || '';
  });
  const [showMovies, setShowMovies] = useState(true);
  const [movies, setMovies] = useState<Movie[]>([]);
  const [selectedSuggestion, setSelectedSuggestion] = useState<number | null>(null);
  const debounceSearch = useDebounce(search, DEBOUNCE_DELAY);
  const suggestionsRef = useRef<HTMLUListElement | null>(null);

  const handleKeyDown = useCallback(
    (event: KeyboardEvent) => {
      if (movies.length === 0) return;

      if (event.key === 'ArrowDown') {
        event.preventDefault();
        setSelectedSuggestion((prev) => (prev === null || prev === movies.length - 1 ? 0 : prev + 1));
      } else if (event.key === 'ArrowUp') {
        event.preventDefault();
        setSelectedSuggestion((prev) => (prev === null || prev === 0 ? movies.length - 1 : prev - 1));
      } else if (event.key === 'Enter' && selectedSuggestion !== null) {
        event.preventDefault();
        const selectedMovie = movies[selectedSuggestion];
        if (selectedMovie) {
          setSearch(selectedMovie.Title);
          setShowMovies(true);
        }
      }
    },
    [movies, selectedSuggestion],
  );

  useEffect(() => {
    if (debounceSearch === '') {
      window.history.pushState({}, '', window.location.pathname);
      return;
    }

    window.history.pushState({}, '', `?q=${debounceSearch}`);
  }, [debounceSearch]);

  useEffect(() => {
    if (debounceSearch !== '') {
      searchMovies(debounceSearch);
      setSelectedSuggestion(null);
      return;
    }
    setMovies([]);
    setSelectedSuggestion(null);
  }, [debounceSearch]);

  useEffect(() => {
    document.addEventListener('keydown', handleKeyDown);
    return () => {
      document.removeEventListener('keydown', handleKeyDown);
    };
  }, [handleKeyDown]);

  async function searchMovies(searchQuery: string) {
    try {
      const moviesFiltered = await searchAPI(searchQuery);
      setMovies(moviesFiltered);
    } catch (e) {
      console.error(e);
    }
  }

  const handleSearch = (event: React.ChangeEvent<HTMLInputElement>) => {
    setShowMovies(false);
    setSearch(event.target.value);
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setShowMovies(true);
  };

  const handleClickSuggestion = (event: React.MouseEvent<HTMLLIElement, MouseEvent>) => {
    const clickedElement = event.currentTarget.textContent;
    if (clickedElement) {
      setSearch(clickedElement);
      setShowMovies(true);
    }
  };

  return (
    <>
      <form onSubmit={handleSubmit} autoComplete="off">
        <div className="flex">
          <div className="relative w-full">
            <input
              type="search"
              id="default-search"
              onChange={handleSearch}
              value={search}
              className="flex w-full p-4 text-sm border rounded-2xl bg-gray-700 border-gray-600 placeholder-gray-400 text-white focus:border-gray-500 focus:bg-gray-700"
              placeholder="Avatar, The Matrix, Inception..."
            />
            <button
              type="submit"
              className="absolute top-0 end-0 p-2.5 text-sm font-medium h-full text-white rounded-e-lg border border-blue-700  focus:ring-4 focus:outline-none  bg-blue-600 hover:bg-blue-700 focus:border-gray-500"
            >
              <svg
                className="w-4 h-4"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 20 20"
              >
                <path
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"
                />
              </svg>
              <span className="sr-only">Search</span>
            </button>
            {!showMovies && movies.length > 0 && (
              <ul
                id="autocomplete-results"
                ref={suggestionsRef}
                className="absolute z-10 w-1/2 bg-gray-700 border border-gray-600 rounded-xl mt-1 p-4"
              >
                {movies.map((movie, index) => (
                  <li
                    key={`suggestion-movie-${movie.ID}`}
                    className={`text-white cursor-pointer hover:bg-gray-800 hover:rounded ${selectedSuggestion === index ? 'bg-gray-800' : ''}`}
                    onClick={handleClickSuggestion}
                  >
                    {movie.Title}
                  </li>
                ))}
              </ul>
            )}
          </div>
        </div>
      </form>

      {showMovies && <Movies movies={movies} />}

      {search === '' && <p className="text-white text-center p-6">Search a movie and check how fast it is! 👀</p>}

      {showMovies && search !== '' && movies.length === 0 && (
        <p className="text-white text-center p-6">No results found 😞, please try again with another movie</p>
      )}
    </>
  );
};
