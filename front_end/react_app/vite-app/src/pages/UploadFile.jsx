import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate, useParams } from 'react-router-dom';
import { DropBox } from '../components/Dropbox';

export function AddFile() {
  const { email } = useParams();
  const [fileName, setFileName] = useState('');
  const [fileSize, setFileSize] = useState('');
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  // Optional: Verify email exists on component mount
  useEffect(() => {
    if (!email) {
      navigate('/'); // Redirect to login if no email
    }
  }, [email, navigate]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setMessage('');

    if (!fileName || !fileSize) {
      setError('File name and size are required');
      return;
    }
console.log(email)
    setLoading(true);
    try {
      const response = await axios.post(
        `http://localhost:8080/api/users/${email}/files`,
        {
          file_name: fileName,
          file_size: parseInt(fileSize)
        },
        {
          headers: {
            'Content-Type': 'application/json'
          }
        }
      );
      
      console.log(response)

      setMessage('File uploaded successfully! Redirecting...');
      setTimeout(() => navigate(`/Home/${email}`), 1500);
    } catch (err) {
      setError(err.response?.data?.error || 'Upload failed. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ // position button in the middle of the page
      display: 'flex', 
      justifyContent: 'center', 
      alignItems: 'center',
      height: '100vh' // Make sure the container takes up the full viewport height
    }}>
      <div className="add-file-page">
      <h2>Upload New File</h2>
      {/* <p className="user-email">Uploading as: {email}</p> */}

      {message && (
        <div className="alert success">
          {message}
        </div>
      )}

      {error && (
        <div className="alert error">
          {error}
        </div>
      )}

      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="fileName">File Name</label>
          <input
            id="fileName"
            type="text"
            value={fileName}
            onChange={(e) => setFileName(e.target.value)}
            placeholder="report.pdf"
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="fileSize">File Size (bytes)</label>
          <input
            id="fileSize"
            type="number"
            value={fileSize}
            onChange={(e) => setFileSize(e.target.value)}
            min="1"
            placeholder="1024"
            required
          />
        </div>

        <div className="button-group">
          <button
            type="submit"
            disabled={loading}
            className="btn primary"
          >
            {loading ? (
              <>
                <span className="spinner"></span> Uploading...
              </>
            ) : (
              'Upload File'
            )}
          </button>

          <button
            type="button"
            onClick={() => navigate(`/Home/${email}`)}
            className="btn secondary"
          >
            Cancel
          </button>
        </div>
      </form>
    </div>
    <DropBox />
    </div>
    
  );
};
