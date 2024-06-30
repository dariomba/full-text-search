import { Movie } from '../../services/search';
import { MovieCard } from '../MovieCard/MovieCard';

interface MoviesProps {
  movies: Movie[];
}

export const Movies: React.FC<MoviesProps> = ({ movies }) => {
  return movies.length > 0 ? (
    <ul className="flex flex-col md:grid md:grid-cols-3 md:items-start p-4 gap-10 justify-center items-center">
      {movies.map((movie) => (
        <MovieCard key={movie.ID} movie={movie} />
      ))}
    </ul>
  ) : null;
};
