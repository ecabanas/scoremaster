import Navbar from './components/Navbar';
import HomePage from './pages/Home';
import AboutPage from './pages/About';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import PopupForm from './components/popup-component';

function App() {
  return (
    <Router>
      <div className="min-h-screen bg-base-200">
        <Navbar />
        <div className="flex-grow">
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/about" element={<AboutPage />} />
          </Routes>
        </div>
        <div className = "flex justify-center my-6">
          <PopupForm />
        </div>
        <footer className="bg-base-300 text-base-content py-6 text-center">
          <p>Copyright © 2025 - My Awesome Site!</p>
        </footer>
      </div>
    </Router>
  );
}

export default App;