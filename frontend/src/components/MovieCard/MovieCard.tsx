import { Movie } from '../../services/search';

interface MovieProps {
  movie: Movie;
}

export const MovieCard: React.FC<MovieProps> = ({ movie }) => {
  return (
    <li className="text-white">
      <article>
        <p>{movie.Title}</p>
      </article>
    </li>
  );
};
