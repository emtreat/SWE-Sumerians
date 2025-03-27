import '../App.css'
import { AddFile } from '../components/AddFile';
import { GetUsers } from '../components/GetUser'

export function Home() {
  return (
    <div className="app">
     <AddFile/>
      <GetUsers />
    </div>
  );
}