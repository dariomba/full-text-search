import { Movie } from '../../services/search';

interface SuggestionsProps {
  movies: Movie[];
  selectedSuggestion: number | null;
  onClickSuggestion: (event: React.MouseEvent<HTMLLIElement, MouseEvent>) => void;
}

export const Suggestions: React.FC<SuggestionsProps> = ({ movies, selectedSuggestion, onClickSuggestion }) => {
  return (
    <ul
      id="autocomplete-results"
      className="absolute z-10 w-1/2 bg-gray-700 border border-gray-600 rounded-xl mt-1 p-4"
    >
      {movies.map((movie, index) => (
        <li
          key={`suggestion-movie-${movie.ID}`}
          className={`text-white cursor-pointer hover:bg-gray-800 hover:rounded ${selectedSuggestion === index ? 'bg-gray-800' : ''}`}
          onClick={onClickSuggestion}
        >
          {movie.Title}
        </li>
      ))}
    </ul>
  );
};
