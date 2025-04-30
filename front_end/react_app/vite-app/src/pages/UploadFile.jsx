import { DropBox } from "../components/Dropbox";

export function AddFile() {
  return (
    <div
      style={{
        flexDirection: "column",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "100vh", // Make sure the container takes up the full viewport height
      }}
    >
      <div className="add-file-page">
        <h2>Upload New File</h2>
      </div>
      <DropBox />
    </div>
  );
}
