import '../App.css'
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import { GetUsers } from '../components/GetUser'
import { AddUser } from '../components/AddUser';


export function Home() {
  return(
    <div >
      <div >
    <GetUsers></GetUsers>
    </div>
    <div >
      <Link to={`UploadFile/`}>
            <AddUser>
              Upload New File
            </AddUser>
          </Link>
    </div>
    </div>
  
    
  
  )
}