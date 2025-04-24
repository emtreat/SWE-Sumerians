import "../App.css";
import axios from "axios";
import React, { useCallback, useState } from "react";
import { useDropzone } from "react-dropzone";
import { useParams, useNavigate } from "react-router-dom";
import { v4 as uuidv4 } from "uuid";

export function DropBox({ props_className }) {
  const { email } = useParams();
  const navigate = useNavigate();
  const [files, setFiles] = useState([]);
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState("");
  const [error, setError] = useState("");

  const onDrop = useCallback((acceptedFiles) => {
    console.log(acceptedFiles);
    acceptedFiles.forEach((file) => {
      const reader = new FileReader();
      reader.onabort = () => console.log("file reading was aborted");
      reader.oneerror = () => console.log("file reading has failed");
      reader.onload = () => {
        const binaryStr = reader.result;
        console.log(binaryStr);
      };
      reader.readAsArrayBuffer(file);
    });

    if (acceptedFiles?.length) {
      setFiles((previousFiles) => [
        ...previousFiles,
        ...acceptedFiles.map((file) =>
          Object.assign(file, {
            id: uuidv4(),
            preview: URL.createObjectURL(file),
          })
        ),
      ]);
    }
  }, []);
  function removefiles(id) {
    setFiles((files) => files.filter((file) => file.id !== id));
  }
  const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop });

  //send blob section
  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    setMessage("");

    if (files.length <= 0) {
      setError("Upload at least one file");
      return;
    }
    setLoading(true);
    try {
      const readFileAsBase64 = (file) => {
        return new Promise((resolve, reject) => {
          const reader = new FileReader();
          reader.onabort = () => reject("file reading was aborted");
          reader.onerror = () => reject("file reading has failed");
          reader.onload = () => resolve(reader.result);
          reader.readAsDataURL(file);
          // reader.readAsArrayBuffer(file);
        });
      };
      //loop through files and post them
      for (const file of files) {
        const binaryStr = await readFileAsBase64(file);
        const base64data = binaryStr.split(",")[1];
        const fileSize = Blob([binaryStr]).size;
        console.log(binaryStr);
        console.log(base64data);
        console.log(fileSize);
        const response = await axios.post(
          `http://localhost:8080/api/users/${email}/files`,
          {
            file_name: file.name,
            file_data: binaryStr,
            file_size: fileSize,
            //file_blob: Array.from(new Uint8Array(binaryStr)),
          },
          {
            headers: {
              "Content-Type": "application/json",
            },
          }
        );
        console.log("here");

        console.log("response recieved", response);
      }

      setMessage("File uploaded successfully! Redirecting...");
      setTimeout(() => navigate(`/Home/${email}`), 1500);
    } catch (err) {
      setError(err.response?.data?.error || "Upload failed. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  // end of send blob section

  return (
    // <div className="DropBoxClass container">
    <div className="DropBoxClass">
      <form className="dropboxtext">
        <div
          {...getRootProps({
            className: "container dropboxtext",
          })}
        >
          <input {...getInputProps()} />
          {isDragActive ? <p> Upload ...</p> : <p> Drag and Drop files</p>}
        </div>

        <h2> Accepted Files</h2>
        {message && <div className="alert success">{message}</div>}

        {error && <div className="alert error">{error}</div>}
        <ul>
          {files.map((file) => (
            <li key={file.id}>
              <p>File {file.name}</p>
              <img
                src={file.preview}
                width={100}
                height={100}
                onLoad={() => {
                  URL.revokeObjectURL(file.preview);
                }}
              />
            </li>

            // console.log(file.name, file.preview);
            // console.log(files.length);
            // console.log(file.name);
          ))}
        </ul>
      </form>
      <div className="button-group">
        <button
          type="submit"
          disabled={loading}
          className="btn primary"
          onClick={handleSubmit}
        >
          {loading ? (
            <>
              <span className="spinner"></span> Uploading...
            </>
          ) : (
            "Upload File"
          )}
        </button>

        <button
          type="button"
          onClick={() => {
            setFiles([]), setError("");
          }}
          className="btn secondary"
        >
          Cancel
        </button>
      </div>
    </div>
  );
}
