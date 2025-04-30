// import { Route, Routes, useNavigate } from "react-router-dom"
import '../App.css'
import { useNavigate } from 'react-router-dom';

export function SignUp() {
    const navigate = useNavigate();
    
    return (
    <div>
        <button
        onClick={() => navigate("/SignUpPage")}
        style={{
            backgroundColor: "grey",
            color: "white",
            border: "none",
            cursor: "pointer",
        }}
        >
        Sign Up
        </button>
    </div>
    ); 
};

