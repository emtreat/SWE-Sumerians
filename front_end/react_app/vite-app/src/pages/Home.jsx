import '../App.css'
import { AddFile } from '../components/AddUser';
import { GetUsers } from '../components/GetUser'

export function Home() {
  return (
    <div className="app">
     <AddUser/>
      <GetUsers />
    </div>
  );
}