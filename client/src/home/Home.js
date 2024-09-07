import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import Cookies from 'js-cookie';

function Home() {
  const navigate = useNavigate();
  const userId = Cookies.get('userId');

  const handleLogout = () => {
    Cookies.remove('userId');
    Cookies.remove('role');
    navigate('/');
  };

  return (
    <div>
      <h1>Welcome to the App</h1>
      {userId ? (
        <div>
          <Link to="/account">
            <button>Account</button>
          </Link>
          <button onClick={handleLogout}>Logout</button>
        </div>
      ) : (
        <div>
          <Link to="/signup">
            <button>Sign Up</button>
          </Link>
          <Link to="/signin">
            <button>Login</button>
          </Link>
        </div>
      )}
    </div>
  );
}

export default Home;