import "../App.css";
import React, { useState, useEffect } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import axios from 'axios';
import { useParams } from 'react-router-dom';
import { AddFile } from '../pages/UploadFile';
import { AddUser } from './AddUser';

export function GetUsers() {
  const { email } = useParams();
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await axios.get(
          `http://localhost:8080/api/users/${email}`
        );
        setUser(response.data);
      } catch (err) {
        setError(err.response?.data?.error || "Failed to fetch user data");
      } finally {
        setLoading(false);
      }
    };
    fetchUserData();
  }, [email]);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;
  if (!user) return <div>No user data found</div>;

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        minHeight: "100vh",
        padding: "20px",
      }}
    >
      <div style={{ textAlign: "center", margin: "40px 0" }}>
        <h1 style={{ fontSize: "2rem", margin: 0 }}>Welcome, {email}</h1>
      </div>

      {/* Files container */}
      <div
        style={{
          flex: 1,
          backgroundColor: "white",
          borderRadius: "8px",
          padding: "20px",
          margin: "0 auto 30px auto",
          width: "80%",
          maxWidth: "800px",
          boxShadow: "0 2px 4px rgba(0,0,0,0.1)",
        }}
      >
        <h2 style={{ marginTop: 0 }}>Your Files:</h2>

        {user.Files.length > 0 ? (
          <ul
            style={{
              listStyle: "none",
              padding: 0,
              maxHeight: "400px",
              overflowY: "auto",
            }}
          >
            {user.Files.map((file, index) => (
              <li key={index} style={{ 
                display: 'flex', 
                justifyContent: 'space-between',
                borderBottom: '1px solid #eee'
              }}> <button style={{display: 'flex', 
                justifyContent: 'space-between',}}>
                <span>{file.file_name}</span>
                <span style={{ fontWeight: 'bold' }}>{file.file_size} KB</span>
              </button>
              </li>
            ))}
          </ul>
        ) : (
          <p style={{ textAlign: "center", color: "#666" }}>
            No files uploaded yet
          </p>
        )}
      </div>
    </div>
  );
}
