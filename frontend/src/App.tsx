import React, { useState, useEffect } from 'react';

function App() {
  const [theme, setTheme] = useState(localStorage.getItem('theme') || 'light');

  useEffect(() => {
    document.documentElement.setAttribute('data-theme', theme);
    localStorage.setItem('theme', theme);
  }, [theme]);

  const toggleTheme = () => {
    setTheme(theme === 'light' ? 'dark' : 'light');
  };

  return (
    <div className="min-h-screen bg-base-200 transition-colors duration-300">
      {/* Navbar */}
      <div className="navbar bg-base-100 shadow-md">
        <div className="navbar-start">
          <div className="dropdown">
            <div tabIndex={0} role="button" className="btn btn-ghost lg:hidden">
              <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M4 6h16M4 12h8m-8 6h16" /></svg>
            </div>
            <ul tabIndex={0} className="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52">
              <li><a>Homepage</a></li>
              <li><a>About</a></li>
              <li><a>Services</a></li>
              <li><a>Contact</a></li>
            </ul>
          </div>
          <a className="btn btn-ghost normal-case text-xl">My Awesome Site</a>
        </div>
        <div className="navbar-center hidden lg:flex">
          <ul className="menu menu-horizontal px-1">
            <li><a>Homepage</a></li>
            <li><a>About</a></li>
            <li><a>Services</a></li>
            <li><a>Contact</a></li>
          </ul>
        </div>
        <div className="navbar-end">
          <button className="btn" onClick={toggleTheme}>
            {theme === 'light' ? (
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="lucide lucide-moon-star"><path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z"/><path d="m17 13 1.5 1.5M5 16l1.5 1.5"/></svg>
            ) : (
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="lucide lucide-sun"><circle cx="12" cy="12" r="4"/><path d="M12 2v2M12 20v2M4.93 4.93l1.41 1.41M17.66 17.66l1.41 1.41M2 12h2M20 12h2M4.93 19.07l1.41-1.41M17.66 6.34l1.41-1.41"/></svg>
            )}
          </button>
          <button className="btn btn-primary ml-2">Get Started</button>
        </div>
      </div>

      {/* Page Content */}
      <div className="container mx-auto p-8">
        <div className="hero bg-base-100 rounded-lg shadow-md p-8">
          <div className="hero-content text-center">
            <div className="max-w-md">
              <h1 className="text-5xl font-bold text-primary mb-4">Welcome!</h1>
              <p className="py-6 text-lg text-gray-700">
                This is a sample static page with a beautiful navbar and a dark/light mode toggle built using DaisyUI and styled with Tailwind CSS.
              </p>
              <button className="btn btn-secondary">Explore More</button>
            </div>
          </div>
        </div>

        <div className="py-8 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div className="card bg-base-100 shadow-md">
            <div className="card-body">
              <h2 className="card-title text-secondary">Feature One</h2>
              <p className="text-gray-600">Description of feature one goes here.</p>
              <div className="card-actions justify-end">
                <button className="btn btn-sm">Learn More</button>
              </div>
            </div>
          </div>
          <div className="card bg-base-100 shadow-md">
            <div className="card-body">
              <h2 className="card-title text-secondary">Feature Two</h2>
              <p className="text-gray-600">Description of feature two goes here.</p>
              <div className="card-actions justify-end">
                <button className="btn btn-sm">Learn More</button>
              </div>
            </div>
          </div>
          <div className="card bg-base-100 shadow-md">
            <div className="card-body">
              <h2 className="card-title text-secondary">Feature Three</h2>
              <p className="text-gray-600">Description of feature three goes here.</p>
              <div className="card-actions justify-end">
                <button className="btn btn-sm">Learn More</button>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Footer */}
      <footer className="bg-base-300 text-base-content py-6 text-center">
        <p>Copyright © 2025 - My Awesome Site</p>
      </footer>
    </div>
  );
}

export default App;