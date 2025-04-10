
import '../App.css'
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import { GetUsers } from '../components/GetUser'
import { AddUser } from '../components/AddUser';
import { DropBox } from "../components/Dropbox";
import { GetFiles } from "../components/GetFIles";

export function Home() {
  return(
    <div >
      <div >
    <GetUsers></GetUsers>
    <DropBox />
    <GetFiles />
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
