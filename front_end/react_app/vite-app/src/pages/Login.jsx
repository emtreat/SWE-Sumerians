import '../App.css'
// import { BrowserRouter as Routes, Route } from 'react-router-dom';
import { Button } from '../components/LoginButton';

export function Login() {
  return (
            <div style={{ // position button in the middle of the page
          display: 'flex', 
          justifyContent: 'center', 
          alignItems: 'center',
          height: '100vh' // Make sure the container takes up the full viewport height
        }}>
          <Button></Button>
        </div>
          );
};
