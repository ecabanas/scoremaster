
function AboutPage() {
  return (
    <div className="container mx-auto p-8">
      <div className="bg-base-100 rounded-lg shadow-md p-8">
        <h2 className="text-3xl font-semibold text-info mb-4">About Us</h2>
        <p className="py-4 text-gray-700">
          This is the about page. Here you can learn more about our company and our mission.
        </p>
        <ul className="list-disc pl-6 text-gray-600">
          <li>Our history</li>
          <li>Our team</li>
          <li>Our values</li>
        </ul>
      </div>
    </div>
  );
}

export default AboutPage;