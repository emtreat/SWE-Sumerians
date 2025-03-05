import './App.css'

function FileRow({ file }) {
  const name = <span style={{ color: 'white' }}>
      {file.name}
    </span>;

  return (
    <tr>
      <td>{name}</td>
    </tr>
  );
}

function ProductTable({ files }) {
  const rows = [];

  files.forEach((file) => {
    rows.push(
      <FileRow
      file={file}
        key={file.name} />
    );
  });

  return (
    <table>
      <thead>
      </thead>
      <tbody>{rows}</tbody>
    </table>
  );
}

function SearchBar() {
  return (
    <form>
      <input type="text" placeholder="Search..." />
    </form>
  );
}

function FilterableProductTable({ files }) {
  return (
    <div>
      <SearchBar />
      <ProductTable files={files} />
    </div>
  );
}

const FILES = [
  {name: "file_1.jpeg"},
  {name: "file_2.jpeg"},
  ];

function App() {
  const handleClick = () => {
    alert('File Added!');
  };

  return (
    <div className="app">
      <div className="name">User's Name</div>
      <div className="container">
        <button onClick={handleClick}>Add File</button>
      </div>
      <div className="container">
        <FilterableProductTable files={FILES} />
      </div>
    </div>
  );
}

export default App
