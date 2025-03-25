import React, { useState, useEffect } from 'react';

export function GetUsers() {
  // let endpoint = "/api/users/"
  const [posts, setPosts] = useState([]);
   useEffect(() => {
      fetch("http://localhost:8080/api/emails_to_users_test")
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
            <div className="container" key={post.Email}>
                <h2 className="name">{post.Email}</h2>
                <div className="button">
                <div className="delete-btn">{post._id}</div>
                </div>
             </div>
          );
       })}
    </div>
    );
}

