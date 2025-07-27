import { useEffect, useState } from "react";
import axios from "axios";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { useNavigate } from "react-router-dom";
import { URLType } from "../pages/types";

const ListURLs = () => {
    const [urls, setUrls] = useState<URLType[]>([]);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchURLs = async () => {
            try {
                const response = await axios.get("http://localhost:8080/urls");
                setUrls(response.data);
            } catch (error) {
                console.error("Failed to fetch URLs:", error);
            }
        };

        fetchURLs();
    }, []);

    // Define columns for the DataGrid
    const columns: GridColDef[] = [
        { field: "url", headerName: "URL", width: 300 },
        { field: "status", headerName: "Status", width: 150 },
        {
            field: "created_at",
            headerName: "Created At",
            width: 200,
        },
        {
            field: "actions",
            headerName: "Actions",
            width: 150,

            renderCell: (params) => (
                <button
                    onClick={() => navigate(`/url/${params.row.id}`)}
                    className="text-blue-800 text-md font-bold rounded inline-flex items-center"
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        className="h-5 w-5 mr-2"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                    >
                        <path d="M10 2a8 8 0 100 16 8 8 0 000-16zm1 11H9v-2h2v2zm0-4H9V7h2v2z" />
                    </svg>
                    <span>Show info</span>
                </button>
            ),
        },
    ];

    return (
        <div className="space-y-4">
            <div style={{ height: 400, width: "100%" }}>
                <DataGrid
                    rows={urls}
                    columns={columns}
                    pagination
                    rowSelection={false}
                />
            </div>
        </div>
    );
};

export default ListURLs;
