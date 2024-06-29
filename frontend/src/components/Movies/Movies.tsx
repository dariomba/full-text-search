import { Movie } from '../../services/search';
import { MovieCard } from '../MovieCard/MovieCard';

interface MoviesProps {
  movies: Movie[];
}

export const Movies: React.FC<MoviesProps> = ({ movies }) => {
  return (
    <ul>
      {movies.map((movie) => (
        <MovieCard key={movie.ID} movie={movie} />
      ))}
    </ul>
  );
};
