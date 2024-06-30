import { Movie } from '../../services/search';

interface MovieProps {
  movie: Movie;
}

export const MovieCard: React.FC<MovieProps> = ({ movie }) => {
  return (
    <li className="text-white">
      <article className="max-w-80 border rounded-lg shadow bg-gray-800 border-gray-700">
        <img
          className="rounded-t-lg min-w-full max-h-[345px]"
          src={movie.Poster_Url}
          loading="lazy"
          alt={`Poster image of ${movie.Title}`}
        />
        <div className="p-5">
          <h5 className="mb-2 text-2xl font-bold tracking-tight text-white">{movie.Title}</h5>
          <p className="mb-3 font-normal text-gray-400 text-pretty">{movie.Overview}</p>
        </div>
      </article>
    </li>
  );
};
