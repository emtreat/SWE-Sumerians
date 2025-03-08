import './App.css'
import React, { useState, useEffect } from 'react';

const GetUsers = () => {
  const [posts, setPosts] = useState([]);
   useEffect(() => {
      fetch("http://localhost:8080/api/users")
         .then((response) => response.json())
         .then((data) => {
            console.log(data);
            setPosts(data);
         })
         .catch((err) => {
            console.log(err.message);
         });
   }, []);

   return (
    <div className="container">
       {posts.map((post) => {
          return (
            <div className="container" key={post._id}>
              console.log(key);
                <h2 className="name">{post.name}</h2>
                console.log(name);
                <div className="button">
                <div className="delete-btn">Delete</div>
                </div>
             </div>
          );
       })}
    </div>
    );
}

// function FileRow({ file }) {
//   const name = <span style={{ color: 'white' }}>
//       {file.name}
//     </span>;

//   return (
//     <tr>
//       <td>{name}</td>
//     </tr>
//   );
// }

// function ProductTable({ files }) {
//   const rows = [];

//   files.forEach((file) => {
//     rows.push(
//       <FileRow
//       file={file}
//         key={file.name} />
//     );
//   });

//   return (
//     <table>
//       <thead>
//       </thead>
//       <tbody>{rows}</tbody>
//     </table>
//   );
// }

// function SearchBar() {
//   return (
//     <form>
//       <input type="text" placeholder="Search..." />
//     </form>
//   );
// }

// function FilterableProductTable({ files }) {
//   return (
//     <div>
//       <SearchBar />
//       <ProductTable files={files} />
//     </div>
//   );
// }

// const FILES = [
//   {name: "file_1.jpeg"},
//   {name: "file_2.jpeg"},
//   ];

function App() {
  // const handleClick = () => {
  //   alert('File Added!');
  // };

  return (
    <div className="app">
    {/* //   <div className="name">User's Name</div>
    //   <div className="container">
    //     <button onClick={handleClick}>Add File</button>
      // </div>
      <div> */}
      <h1>Did it run?</h1>
        <GetUsers />
      </div>
    // </div>
  );
}

export default App
