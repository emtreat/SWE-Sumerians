import '../App.css'
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { GetUsers } from '../components/GetUser'


export function Home() {
  return(
    <GetUsers></GetUsers>
  )
}