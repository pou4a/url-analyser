import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "axios";
import { CrawlResult } from "./types";

const URLDetail: React.FC = () => {
    const { id } = useParams<{ id: string }>();
    const [results, setResults] = useState<CrawlResult | null>(null);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchResults = async () => {
            try {
                const response = await axios.get(
                    `http://localhost:8080/crawl-results?url_id=${id}`
                );
                setResults(response.data);
            } catch (error) {
                console.error("Failed to fetch crawl results:", error);
            }
        };

        fetchResults();
    }, [id]);

    return (
        <div className="min-h-screen bg-gray-100 flex flex-col items-center p-6">
            <header className="mb-8">
                <h1 className="text-4xl font-bold text-blue-600">URL Detail</h1>
                <button
                    onClick={() => navigate("/")}
                    className="bg-gray-300 text-gray-800 px-4 py-2 rounded-lg hover:bg-gray-400 transition"
                >
                    Back to Home
                </button>
            </header>

            <main className="w-full max-w-4xl space-y-8">
                {results ? (
                    <div className="p-6 bg-white shadow-md rounded-lg border border-gray-200 space-y-4">
                        <h2 className="text-xl font-semibold text-gray-800">
                            Results for URL ID: {id}
                        </h2>
                        <p className="text-sm text-gray-600">
                            <span className="font-semibold">Title:</span>{" "}
                            {results.title}
                        </p>
                        <p className="text-sm text-gray-600">
                            <span className="font-semibold">HTML Version:</span>{" "}
                            {results.html_version}
                        </p>
                        <p className="text-sm text-gray-600">
                            <span className="font-semibold">
                                Internal Links:
                            </span>{" "}
                            {results.internal_links}
                        </p>
                        <p className="text-sm text-gray-600">
                            <span className="font-semibold">
                                External Links:
                            </span>{" "}
                            {results.external_links}
                        </p>
                        <p className="text-sm text-gray-600">
                            <span className="font-semibold">
                                Has Login Form:
                            </span>{" "}
                            {results.has_login_form ? "Yes" : "No"}
                        </p>
                        <p className="text-sm text-gray-600">
                            <span className="font-semibold">Analyzed At:</span>{" "}
                            {new Date(results.analyzed_at).toLocaleString()}
                        </p>
                    </div>
                ) : (
                    <p className="text-gray-600">
                        No crawl results found for URL ID: {id}
                    </p>
                )}
            </main>
        </div>
    );
};

export default URLDetail;
