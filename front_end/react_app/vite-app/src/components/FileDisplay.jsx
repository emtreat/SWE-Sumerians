import "../App.css";
import axios from "axios";
import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";

export function FileDisplay({ file }) {
  const { email } = useParams();
  //   const navigate = useNavigate();
  // const [files, setFiles] = useState([]);
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState("");
  const [error, setError] = useState("");
  const [fileURL, setFileURL] = useState(null);
  const [fileTypes, setFileTypes] = useState("application/octet-stream");

  console.log("here");

  useEffect(() => {
    if (file && file.file_blob) {
      try {
        const binStr = atob(file.file_blob);
        const bytes = new Uint8Array(binStr.length);
        for (let i = 0; i < binStr.length; i++) {
          bytes[i] = binStr.charCodeAt(i);
        }
        const fileExtension = fileType(file.file_name);
        setFileTypes(fileExtension);
        const blob = new Blob([bytes]);
        const url = URL.createObjectURL(blob);
        setFileURL(url);
      } catch (err) {
        console.error("Error with file:", err);
        setError("failed to load file");
      }
    }
    return () => {
      if (fileURL) {
        URL.revokeObjectURL(fileURL);
      }
    };
  }, [file]);

  //Download Button
  const downloadFile = async () => {
    if (!fileURL) {
      setError("");
    }
    try {
      setLoading(true);
      const link = document.createElement("a");
      link.href = fileURL;
      link.setAttribute("download", file.file_name);
      document.body.appendChild(link);
      link.click();
      link.remove();

      setMessage("File Downloaded successfully!");
      // setTimeout(() => navigate(`/Home/${email}`), 1500);
    } catch (err) {
      setError("Download failed. Please try again. Error: ", err);
    } finally {
      setLoading(false);
    }
  };

  function fileType(fileName) {
    const fileExtension = fileName.split(".").pop().toLowerCase();
    console.log(fileExtension);
    const mapFileType = {
      //mime types
      jpg: "image/jpeg",
      jpeg: "image/jpeg",
      png: "image/png",
      gif: "image/gif",
      pdf: "application/pdf",
      doc: "application/msword",
      docx: "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
      xls: "application/vnd.ms-excel",
      xlsx: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
      zip: "application/zip",
      mp3: "audio/mpeg",
      mp4: "video/mp4",
      txt: "text/plain",
      csv: "text/csv",
    };
    return mapFileType[fileExtension];
  }

  console.log(file);
  console.log(fileType(file.file_name));

  if (loading) return <div className="container">File Loading</div>;
  if (fileTypes == "application/pdf") {
    return (
      <div>
        <h2 style={{ color: "white", backgroundColor: "rgba(0,0,0,0.9)" }}>
          {file.file_name}
        </h2>
        <object
          data={fileURL}
          type="application/pdf"
          width={200}
          height={200}
        ></object>
      </div>
    );
  } else if (
    fileTypes == "image/jpeg" ||
    fileTypes == "image/png" ||
    fileTypes == "image/gif"
  ) {
    return (
      <div
        style={{
          position: "fixed",
          top: "0",
          left: "0",
          width: "100vw",
          height: "100vh",
          backgroundColor: "rgba(0,0,0,0.9)",
          zOndex: "1000", // in the front
        }}
      >
        <h2 style={{ color: "white", backgroundColor: "rgba(0,0,0,0.9)" }}>
          {file.file_name}
        </h2>
        <img src={fileURL}></img>
      </div>
    );
  }

  return (
    <div>
      {error && <div>{error}</div>}
      {message && <div>{message}</div>}
      <h2>{file.file_name}</h2>
      <button onClick={downloadFile} disabled={!fileURL}>
        Dowload
      </button>
      {fileURL && <iframe src={fileURL}></iframe>}
    </div>
  );
}
