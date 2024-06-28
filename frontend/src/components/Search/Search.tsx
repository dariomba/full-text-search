import { useEffect, useState } from 'react';

export const Search: React.FC = () => {
  const [search, setSearch] = useState('');

  const handleSearch = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearch(event.target.value);
  };

  useEffect(() => {
    if (search == '') {
      window.history.pushState({}, '', window.location.pathname);
      return;
    }

    window.history.pushState({}, '', `?q=${search}`);
  }, [search]);

  return (
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
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"
            />
          </svg>
        </div>
        <input
          type="search"
          id="default-search"
          onChange={handleSearch}
          className="block w-full p-4 ps-10 text-sm border rounded-2xl bg-gray-700 border-gray-600 placeholder-gray-400 text-white focus:border-gray-500"
          placeholder="Avatar, The Matrix, Inception..."
        />
      </div>
    </form>
  );
};
