import '../App.css'
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

