import "../App.css";
import React, { useCallback, useState } from "react";
import { useDropzone } from "react-dropzone";
import { v4 as uuidv4 } from "uuid";
import { PostFileBlob } from "./AddBlob";

export function DropBox({ props_className }) {
  const [files, setFiles] = useState([]);
  const [loading, setLoading] = useState(false);
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
    setFiles((files) =>
      files.filter((file) => {
        file.id !== id;
      })
    );
  }
  const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop });

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
        <ul>
          {files.map((file) => (
            <li key={file.id}>
              <p>File {file.name}</p>
              {/* <Image
                src={file.preview}
                width={100}
                height={100}
                onload={() => {
                  URL.revokeObjectURL(file.preview);
                }}
              /> */}
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
          onClick={PostFileBlob()}
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
          onClick={() => setFiles([])}
          className="btn secondary"
        >
          Cancel
        </button>
      </div>
    </div>
  );
}
