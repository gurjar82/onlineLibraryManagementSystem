// Home.js
import React from 'react';
import { Link } from 'react-router-dom';

function Home() {
  return (
    <div>
      <h1>Welcome to My React App</h1>
      <Link to="/login">
        <button>Login</button>
      </Link>
      <Link to="/signup">
        <button>Signup</button>
      </Link>
    </div>
  );
}

export default Home;
