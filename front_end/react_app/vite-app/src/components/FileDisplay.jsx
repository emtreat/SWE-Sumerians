import "../App.css";
import axios from "axios";
import React, { useState } from "react";
import { useParams, useNavigate } from "react-router-dom";

export function FileDisplay({ file }) {
  const { email } = useParams();
  //   const navigate = useNavigate();
  // const [files, setFiles] = useState([]);
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState("");
  const [error, setError] = useState("");

  console.log("here");

  //Download Button
  const downloadFile = async () => {
    setError("");
    setMessage("");

    setLoading(true);
    try {
      const url = window.URL.createObjectURL(new Blob([file.file_blob]));
      const link = document.createElement("a");
      link.href = url;
      link.setAttribute("download", file.file_name);
      document.body.appendChild(link);
      link.click();
      link.remove();
      window.URL.revokeObjectURL(url);

      setMessage("File Downloaded successfully! Redirecting...");
      // setTimeout(() => navigate(`/Home/${email}`), 1500);
    } catch (err) {
      setError("Download failed. Please try again. Error: ", err);
    } finally {
      setLoading(false);
    }
  };

  function fileType(bytes) {
    const hex = bytes
      .slice(0, 8)
      .map((b) => b.toString(16).padStart(2, "0"))
      .join(" ")
      .toUpperCase();

    switch (hex) {
      case hex.startsWith("89 50 4E 47"):
        return "image/png";
      case hex.startsWith("FF D8 FF"):
        return "image/jpeg";
      case hex.startsWith("25 50 44 46"):
        return "application/pdf";
      case hex.startsWith("47 49 46 38"):
        return "image/gif";
      case hex.startsWith("50 48 03 04"):
        return "application/zip";
      default:
        "unkown";
    }
  }

  function base64Byte(base64) {
    const binaryStr = atob(base64);
    const bytes = new Uint8Array(binaryStr.length);
    for (let i = 0; i < binaryStr.length; i++) {
      bytes[i] = binaryStr.charCodeAt(i);
    }
    return bytes;
  }

  const file_convert_from_base64_to_bytes = base64Byte(file.file_blob);
  const typefromfile = fileType(file_convert_from_base64_to_bytes);

  console.log(file);
  console.log(file_convert_from_base64_to_bytes);
  console.log(typefromfile);

  const url = window.URL.createObjectURL(new Blob([file.file_blob]));
  console.log(url);

  // end of send blob section

  if (loading) return <div className="container">File Loading</div>;

  return (
    <div>
      <h2>{file.file_name}</h2>
      <button onClick={downloadFile}>Dowload</button>
      <img src={url}></img>
    </div>
  );
}
