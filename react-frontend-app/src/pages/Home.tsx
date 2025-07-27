import React, { useState } from "react";
import ListURLs from "../components/ListURLs";
import SubmitURL from "../components/SubmitURL";

const Home: React.FC = () => {
    const [showSubmitModal, setShowSubmitModal] = useState<boolean>(false);

    const handleSubmitModal = () => {
        setShowSubmitModal(true); // Open the submit URL modal
    };

    return (
        <div className="min-h-screen bg-gray-100 flex flex-col items-center p-6">
            <header className="mb-8">
                <h1 className="text-4xl font-bold text-blue-600">
                    Url scraper and analyzer
                </h1>
                <p className="text-gray-600 mt-2 text-center">
                    Manage URLs and view crawl results.
                </p>
            </header>

            <main className="w-full max-w-4xl space-y-8">
                <section className="bg-white shadow-md rounded-lg p-6">
                    <div className="flex justify-between items-center mb-4">
                        <h2 className="text-2xl font-semibold text-gray-800">
                            List of URLs
                        </h2>
                        <button
                            onClick={handleSubmitModal}
                            className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
                        >
                            Submit URL
                        </button>
                    </div>
                    <ListURLs />
                </section>
            </main>

            {/* Submit URL Modal */}
            {showSubmitModal && (
                <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center">
                    <div className="bg-white p-6 rounded-lg shadow-lg w-full max-w-md">
                        <SubmitURL onClose={() => setShowSubmitModal(false)} />
                    </div>
                </div>
            )}
        </div>
    );
};

export default Home;
