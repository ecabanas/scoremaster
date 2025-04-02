
function HomePage() {
  return (
    <div className="container mx-auto p-8">
      <div className="hero bg-base-100 rounded-lg shadow-md p-8">
        <div className="hero-content text-center">
          <div className="max-w-md">
            <h1 className="text-5xl font-bold text-primary mb-4">Welcome Home!</h1>
            <p className="py-6 text-lg text-gray-700">
              This is the homepage of our awesome website. Feel free to explore!
            </p>
            <button className="btn btn-secondary">Learn More</button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default HomePage;