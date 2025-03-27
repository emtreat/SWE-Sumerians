import './App.css'
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Home } from './pages/Home'
import  { Login } from './pages/Login'

function App() {
    return (
      <Router>
        <Routes>
        <Route path="/" element={<Login/>} />
        <Route path="/Home" element={<Home />} />
      </Routes>
      </Router>
      );
}


export default App
