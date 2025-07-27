import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "./pages/Home";
import URLDetail from "./pages/URLDetail";

function App() {
    return (
        <Router>
            <Routes>
                {/* Home Page */}
                <Route path="/" element={<Home />} />

                {/* URL Detail Page */}
                <Route path="/url/:id" element={<URLDetail />} />
            </Routes>
        </Router>
    );
}

export default App;
