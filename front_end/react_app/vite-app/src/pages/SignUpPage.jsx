/* eslint-disable no-unused-vars */

import { Link, useNavigate } from 'react-router-dom';
import React, { useState } from 'react';
import axios from 'axios';
import { CancelButton } from '../components/CancelButton';

export function AccountCreation() {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    firstName: '',
    lastName: '',
    email: ''
  });
  const [response, setResponse] = useState(null);
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const [success, setSuccess] = useState(false);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
  };

  const postData = async (formData) => {
    setIsLoading(true);
    setError(null);
    setResponse(null);
    setSuccess(false);
    
    // Structure the data according to your backend requirements
    const data = {
      // "_id": Math.random().toString(36).substring(2, 9), // Generate a random ID
      "Email": formData.email,
      "FirstName": formData.firstName,
      "LastName": formData.lastName,
      "Files": []
    };

    console.log(data)

    try {
      const result = await axios.post('http://localhost:8080/api/users', data, {
        headers: {
          'Content-Type': 'application/json',
        }
      });
      setResponse(result.data);
      setSuccess(true); // Show success message
      
      // Redirect to /Home/email after 1.5 seconds
      setTimeout(() => {
        navigate(`/Home/${formData.email}`, { 
          state: { 
            user: result.data,
            successMessage: 'Account created successfully!'
          } 
        });
      }, 1500);
      
      return result.data;
    } catch (err) {
      setError(err.response?.data?.message || err.message || 'An error occurred');
      throw err;
    } finally {
      setIsLoading(false);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    await postData(formData);
  };

  return (
    <div style={{
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center',
      height: '100vh'
    }}>
      <Link to={`/`}>
      <CancelButton/>
      </Link>
      
      <form onSubmit={handleSubmit} style={{
        display: 'flex',
        flexDirection: 'column',
        gap: '15px',
        width: '450px',
        padding: '20px',
        boxShadow: '0 0 10px rgba(0,0,0,0.1)',
        borderRadius: '8px',
        position: 'relative'
      }}>
        <h2 style={{ textAlign: 'center' }}>Create Account</h2>
        
        {error && (
          <div style={{
            color: 'red',
            padding: '10px',
            backgroundColor: '#ffebee',
            borderRadius: '4px',
            textAlign: 'center'
          }}>
            {error}
          </div>
        )}
        
        {success ? (
          <div style={{
            color: 'green',
            padding: '10px',
            backgroundColor: '#e8f5e9',
            borderRadius: '4px',
            textAlign: 'center'
          }}>
            Account created successfully! Redirecting...
          </div>
        ) : (
          <>
            <input
              type="text"
              name="firstName"
              placeholder="First Name"
              value={formData.firstName}
              onChange={handleChange}
              required
            />
            
            <input
              type="text"
              name="lastName"
              placeholder="Last Name"
              value={formData.lastName}
              onChange={handleChange}
              required
            />
            
            <input
              type="email"
              name="email"
              placeholder="Email"
              value={formData.email}
              onChange={handleChange}
              required
            />
            
            <button 
              type="submit"
              disabled={isLoading}
            >
              {isLoading ? 'Creating Account...' : 'Create Account'}
            </button>
          </>
        )}
      </form>
    </div>
  );
}