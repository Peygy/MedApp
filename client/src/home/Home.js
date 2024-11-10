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
      <h1>Добро пожаловать на MedApp!</h1>
      {userId ? (
        <div>
          <Link to="/account">
            <button>Ваш личный кабинет</button>
          </Link>
          <Link to="/notes">
            <button>Запись к врачу</button>
          </Link>
          <button onClick={handleLogout}>Выход из аккаунта</button>
        </div>
      ) : (
        <div>
          <Link to="/signup">
            <button>Регистрация аккаунта</button>
          </Link>
          <Link to="/signin">
            <button>Вход в аккаунт</button>
          </Link>
        </div>
      )}
    </div>
  );
}

export default Home;