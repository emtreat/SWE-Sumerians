
import '../App.css'
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import { GetFiles } from '../components/GetUserData'
import { AddFile } from '../components/AddFile';
import { LogoutButton } from '../components/LogoutButton';

export function Home() {
  return(
    <div style={{ position: 'relative' }}>
      <div>
        <Link to={`/`}>
        <LogoutButton />
        </Link>
      </div>
      <div >
    <GetFiles />
    </div>
    <div >
      <Link to={`UploadFile/`}>
            <AddFile>
              Upload New File
            </AddFile>
          </Link>
    </div>
    </div>
  
    
  
  )
}
