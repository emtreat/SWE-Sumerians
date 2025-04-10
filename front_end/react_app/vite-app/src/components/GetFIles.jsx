import React, { useState, useEffect } from 'react';
import '../App.css';

export function GetFiles() {
  const [files, setFiles] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/api/users")
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        setFiles(data[2]);
      })
      .catch((err) => {
        console.log(err.message);
      });
  }, []);

  return (
    <div className="main-container">
      <h2>Files</h2>
      <div className="files-container">
        {files.length > 0 ? (
          files.map((file) => (
            <><button
                  key={file._id} // Use a unique identifier
                  className="file-button"
                  onClick={() => {
                      // Handle file button click
                      console.log('Selected file:', file);
                  } }
              >
                  {file.name || file.files || 'Unnamed File'} {/* Fallback text */}
              </button><br></br></>
          ))
        ) : (
          <p>No files available...</p>
        )}
      </div>
    </div>
  );
}