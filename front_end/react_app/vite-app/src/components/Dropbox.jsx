import "../App.css";
import React, { useCallback, useState } from "react";
import { useDropzone } from "react-dropzone";

export function DropBox({ props_className }) {
  const [files, setFiles] = useState([]);
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
          Object.assign(file, { preview: URL.createObjectURL(file) })
        ),
      ]);
    }
  }, []);
  function removefiles(name) {
    setFiles((files) =>
      files.filter((file) => {
        file.name !== name;
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
            <li key={file.name}>
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
    </div>
  );
}
