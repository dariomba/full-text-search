interface SearchInputProps {
  search: string;
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
  onSubmit: (event: React.FormEvent<HTMLFormElement>) => void;
}

export const SearchInput: React.FC<SearchInputProps> = ({ search, onChange, onSubmit }) => {
  return (
    <form onSubmit={onSubmit} autoComplete="off">
      <div className="flex">
        <div className="relative w-full">
          <input
            type="search"
            id="default-search"
            onChange={onChange}
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
        </div>
      </div>
    </form>
  );
};
