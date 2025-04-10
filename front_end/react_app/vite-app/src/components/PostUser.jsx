import React, { useState } from 'react';
import axios from 'axios'; // You'll need to install axios: npm install axios

export function PostUser() {
  const [response, setResponse] = useState(null);
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(false);

  const postData = async () => {
    setIsLoading(true);
    setError(null);
    setResponse(null);
    
    const data = {
      "_id": "67d10adf1365ac824b9c79d5",
      "Email": "jedwards@udallas.edu",
      "Files": []
    };

    try {
      const result = await axios.post('http://localhost:8080/api/users', data, {
        headers: {
          'Content-Type': 'application/json',
          // Add any auth headers if needed:
          // 'Authorization': 'Bearer YOUR_TOKEN_HERE'
        }
      });
      setResponse(result.data);
    } catch (err) {
      setError(err.message || 'An error occurred');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div>
      <h1>Post Test</h1>
      <button 
        onClick={postData}
        disabled={isLoading}
      >
        {isLoading ? 'Posting...' : 'Post Data'}
      </button>

      {isLoading && <p>Loading...</p>}
      
      {error && (
        <div style={{ 
          marginTop: '20px', 
          padding: '10px', 
          border: '1px solid #ff4444',
          borderRadius: '4px',
          backgroundColor: '#ffebee'
        }}>
          <strong>Error:</strong> {error}
        </div>
      )}

      {response && (
        <div style={{ 
          marginTop: '20px', 
          padding: '10px', 
          border: '1px solid #ddd',
          borderRadius: '4px'
        }}>
          <strong>Success!</strong>
          <pre>{JSON.stringify(response, null, 2)}</pre>
        </div>
      )}
    </div>
  );
};
