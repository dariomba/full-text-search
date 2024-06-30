import { useDebounce } from '@uidotdev/usehooks';
import { useCallback, useEffect, useState } from 'react';
import { Movie, searchAPI } from '../../services/search';
import { Movies } from '../Movies/Movies';
import { SearchInput } from '../SearchInput/SearchInput';
import { Suggestions } from '../Suggestions/Suggestions';

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
      <SearchInput onChange={handleSearch} onSubmit={handleSubmit} search={search} />

      {!showMovies && movies.length > 0 && (
        <Suggestions
          movies={movies}
          onClickSuggestion={handleClickSuggestion}
          selectedSuggestion={selectedSuggestion}
        />
      )}

      {showMovies && <Movies movies={movies} />}

      {search === '' && <p className="text-white text-center p-6">Search a movie and check how fast it is! 👀</p>}

      {showMovies && search !== '' && movies.length === 0 && (
        <p className="text-white text-center p-6">No results found 😞, please try again with another movie</p>
      )}
    </>
  );
};
