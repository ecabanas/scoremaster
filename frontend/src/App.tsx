import HomePage from './pages/Home';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

function App() {

  return (
    <Router>
      <div className="min-h-screen bg-base-200">
        <div className="flex-grow">
          <Routes>
            <Route path="/" element={<HomePage />} />
          </Routes>
        </div>
        <footer className="bg-base-300 text-base-content py-6 text-center">
          <p>Copyright © 2025 - My Awesome Site!</p>
        </footer>
      </div>
    </Router>
  );
}

export default App;