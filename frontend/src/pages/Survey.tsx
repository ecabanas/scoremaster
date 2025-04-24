import { useParams } from 'react-router-dom';

function SurveyPage() {
    const { pageNumber } = useParams();

    return (
        <div className="container mx-auto p-8">
            <div className="hero bg-base-100 rounded-lg shadow-md p-8">
                <div className="hero-content text-center">
                    <div className="max-w-md">
                        <h1 className="text-5xl font-bold text-primary mb-4">
                            Survey Page {pageNumber || '1'}
                        </h1>
                        <p className="py-6 text-lg text-gray-700">
                            This is page {pageNumber || '1'} of your survey.
                        </p>
                        {/* Your survey content here */}
                    </div>
                </div>
            </div>
        </div>
    );
}

export default SurveyPage;