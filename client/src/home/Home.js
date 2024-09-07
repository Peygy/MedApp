import React from 'react';
import { Link } from 'react-router-dom';

function Home() {
  return (
    <div>
      <h1>Welcome to the App</h1>
      <Link to="/signup">
        <button>Sign Up</button>
      </Link>
      <Link to="/signin">
        <button>Login</button>
      </Link>
    </div>
  );
}

export default Home;