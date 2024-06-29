import { useDebounce } from '@uidotdev/usehooks';
import { useEffect, useState } from 'react';
import { Movie, searchAPI } from '../../services/search';

const DEBOUNCE_DELAY = 50;

export const Search: React.FC = () => {
  const [search, setSearch] = useState(() => {
    const searchParams = new URLSearchParams(window.location.search);
    return searchParams.get('q') || '';
  });
  const [movies, setMovies] = useState<Movie[]>([]);
  const debounceSearch = useDebounce(search, DEBOUNCE_DELAY);

  const handleSearch = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearch(event.target.value);
  };

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
      return;
    }
    setMovies([]);
  }, [debounceSearch]);

  async function searchMovies(searchQuery: string) {
    try {
      const moviesFiltered = await searchAPI(searchQuery);
      setMovies(moviesFiltered);
    } catch (e) {
      console.error(e);
    }
  }

  return (
    <>
      <form>
        <div className="relative">
          <div className="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
            <svg
              className="w-4 h-4 text-gray-400"
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
          </div>
          <input
            type="search"
            id="default-search"
            onChange={handleSearch}
            defaultValue={search}
            className="block w-full p-4 ps-10 text-sm border rounded-2xl bg-gray-700 border-gray-600 placeholder-gray-400 text-white focus:border-gray-500"
            placeholder="Avatar, The Matrix, Inception..."
          />
        </div>
      </form>
      <ul>
        {movies.map((movie) => (
          <li key={movie.ID} className="text-white">
            <article>
              <p key={`movie-title-${movie.ID}`}>{movie.Title}</p>
            </article>
          </li>
        ))}
      </ul>
    </>
  );
};
