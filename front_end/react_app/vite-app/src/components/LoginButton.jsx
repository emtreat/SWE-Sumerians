// import { Route, Routes, useNavigate } from "react-router-dom"
import '../App.css'

export function Button() {
//   const navigate = useNavigate();

//   const handleClick = () => {
//     navigate("/");
// };
    return (
      <button
    //   onClick={handleClick}
      style={{
        padding: "10px 20px",
        fontSize: "16px",
        cursor: "pointer",
        outline: "none",
        border: "2px solid green",
        borderRadius: "4px",
        marginTop: "16px",
        marginBottom: "16px",
      }}
    >
      LOGIN
    </button>
    ); 
  
};

