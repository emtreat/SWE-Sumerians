import './App.css'
import { GetUsers } from './components/GetUser'

export function Home() {
    return (
        <div className="app">
          <h1>User ID</h1>
            <GetUsers />
          </div>
      );
}