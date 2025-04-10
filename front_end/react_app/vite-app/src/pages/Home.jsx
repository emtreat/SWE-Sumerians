import "../App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { GetFiles } from "../components/GetFIles";
import { GetUsers } from "../components/GetUser";
import { DropBox } from "../components/Dropbox";

export function Home() {
  return (
    <div className="app">
      <GetFiles />
      <GetUsers />
      <DropBox />
    </div>
  );
}
