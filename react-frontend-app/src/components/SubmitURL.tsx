import React, { useState } from "react";
import axios from "axios";
import { SubmitURLProps } from "../pages/types";

const SubmitURL: React.FC<SubmitURLProps> = ({ onClose }) => {
    const [url, setUrl] = useState<string>("");
    const [message, setMessage] = useState<string>("");

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            // Submit the URL
            const response = await axios.post("http://localhost:8080/urls", {
                url,
            });

            const urlId = response.data.id;
            setMessage(`URL submitted successfully! ID: ${urlId}`);

            // Trigger the crawl
            await axios.post("http://localhost:8080/crawl", {
                url_id: urlId,
            });

            setMessage(
                `URL submitted and crawling started successfully! ID: ${urlId}`
            );
            setUrl("");
        } catch (error) {
            console.error("Error during submission or crawling:", error);
            setMessage(
                "Failed to submit URL or start crawling. Please try again."
            );
        }
    };

    return (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center">
            <div className="bg-white p-6 rounded-lg shadow-lg w-full max-w-md space-y-4">
                <h2 className="text-xl font-semibold text-gray-800">
                    Submit a URL
                </h2>
                <form onSubmit={handleSubmit} className="space-y-4">
                    <div>
                        <input
                            type="text"
                            placeholder="Enter URL"
                            value={url}
                            onChange={(e) => setUrl(e.target.value)}
                            required
                            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                        />
                    </div>
                    <button
                        type="submit"
                        className="w-full bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 transition"
                    >
                        Submit
                    </button>
                </form>
                {message && (
                    <p
                        className={`text-sm font-medium ${
                            message.includes("successfully")
                                ? "text-green-600"
                                : "text-red-600"
                        }`}
                    >
                        {message}
                    </p>
                )}
                <button
                    onClick={onClose}
                    className="w-full bg-gray-300 text-gray-800 py-2 px-4 rounded-lg hover:bg-gray-400 transition"
                >
                    Close
                </button>
            </div>
        </div>
    );
};

export default SubmitURL;
