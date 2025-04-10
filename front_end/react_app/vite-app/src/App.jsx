import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Home } from "./pages/Home";
import { Login } from "./pages/Login";
import { DropBox } from "./components/Dropbox";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/Home" element={<Home />} />
        {/* <Route path="/DropBox" element={<DropBox />} /> */}
      </Routes>
    </Router>
  );
}

export default App;
