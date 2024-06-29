import './App.css';
import { Search } from './components/Search/Search';

function App() {
  return (
    <div className="p-4">
      <h1 className="text-3xl font-bold mb-6">Search your favorite movie!</h1>
      <Search />
    </div>
  );
}

export default App;
