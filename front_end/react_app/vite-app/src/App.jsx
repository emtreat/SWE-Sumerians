import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Home } from './pages/Home'
import  { Login } from './pages/Login'
import { Post } from './pages/AddPage';

function App() {

    return (
      <Router>
        <Routes>
        <Route path="/" element={<Login/>} />
        <Route path="/Home/:email" element={<Home/>} />
        <Route path="/AddPage" element ={<Post/>}/>
      </Routes>
    </Router>
  );
}

export default App;
