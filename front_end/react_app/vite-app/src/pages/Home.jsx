
import '../App.css'
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { AddFile } from "../components/AddFile";
import { GetUsers } from "../components/GetUser";
import { DropBox } from "../components/Dropbox";



export function Home() {
  return (
    <div className="app">
      <AddFile />
      <GetUsers />
      <DropBox />
    </div>
  );
}
