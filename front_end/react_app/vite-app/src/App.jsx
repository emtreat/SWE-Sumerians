import './App.css'
import { HashRouter as Router, Routes, Route} from 'react-router-dom'
import { Home } from './pages/Home'
import { Login } from './pages/Login'

function App() {

  return (
    <div className="app">
      <h1>User ID</h1>
        <GetUsers />
      </div>
  );

  // return (
  //   <Router>
  //     <Routes>
  //       <Route path="/" element={<Login/>}/> 
  //       <Route path="/Home" element={<Home/>}/> 
  //     </Routes>
  //   </Router>
  // );
}

export default App
