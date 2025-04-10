import '../App.css'
import { BrowserRouter as Routes, Route } from 'react-router-dom';
import { Link } from 'react-router-dom';
import { PostUser } from '../components/PostUser';
import { AddFile } from './UploadFile';

export function Post() {
    return (
        <div style={{ // position button in the middle of the page
      display: 'flex', 
      justifyContent: 'center', 
      alignItems: 'center',
      height: '100vh' // Make sure the container takes up the full viewport height
    }}>
        <AddFile></AddFile>
    </div>
      );
}

