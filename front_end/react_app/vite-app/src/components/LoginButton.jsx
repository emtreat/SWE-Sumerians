// import { Route, Routes, useNavigate } from "react-router-dom"
import '../App.css'
import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Routes, useNavigate } from 'react-router-dom';

export function Button() {
    const [email, setEmail] = useState('');
    const [error, setError] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const navigate = useNavigate();
  
    const handleSubmit = async (e) => {
      e.preventDefault();
      setError('');
      setIsLoading(true);
  
      try {
        navigate(`/Home/${(email)}`);
      } catch (err) {
        setError('Error during navigation');
        console.error(err);
      } finally {
        setIsLoading(false);
      }
    };
    return (
      <div>
      <h1>Login</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="Enter your email"
          required
        />
        <button type="submit" disabled={isLoading}>
          {isLoading ? 'Loading...' : 'Login'}
        </button>
        {error && <p className="error">{error}</p>}
      </form>
    </div>
    ); 
  
};

