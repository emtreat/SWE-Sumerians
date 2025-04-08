import React, { useState, useEffect } from 'react';
import { GetFiles } from './GetFIles';
import '../App.css'

// export function GetUsers() {
//   // let endpoint = "/api/users/"
//   const [posts, setPosts] = useState([]);
//    useEffect(() => {
//       fetch("http://localhost:8080/api/emails_to_users_test")
//          .then((response) => response.json())
//          .then((data) => {
//             console.log(data);
//             setPosts(data);
//          })
//          .catch((err) => {
//             console.log(err.message);
//          });
//    }, []);

//    return (
//     <div className="container">
//        {posts.map((post) => {
//           return (
//             <div className="container" key={post.Email}>
//                 <h2 className="name">{post.Email}</h2>
//                 <div className="button">
//                 <div className="delete-btn">{post._id}</div>
//                 </div>
//              </div>
//           );
//        })}
//     </div>
//     );
// }

//place holder function while I work on a dynamic GET request
export function GetUsers() {
  // eslint-disable-next-line no-unused-vars
  const [posts, setPosts] = useState([]);
  const [firstPost, setFirstPost] = useState(null); // New state for first post

  useEffect(() => {
    fetch("http://localhost:8080/api/users")
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        setPosts(data);
        if (data.length > 0) {
          setFirstPost(data[2]); // Store the first post
        }
      })
      .catch((err) => {
        console.log(err.message);
      });
  }, []);

  // const items = firstPost.map(item => <li>{item}</li>)

  return (
    <div className="container">
      {firstPost ? ( // Only render if firstPost exists
        <div className="container" key={firstPost.Email}>
          <h2 className="name">{firstPost.Email}</h2>
          <GetFiles/>
        </div>
      ) : (
        <p>No data available...</p> // Fallback content
      )}
    </div>
  );
}