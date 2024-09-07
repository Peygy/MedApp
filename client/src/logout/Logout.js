import React, { useEffect } from 'react';
import { useHistory } from 'react-router-dom';
import Cookies from 'js-cookie';

const Logout = () => {
  const history = useHistory();

  useEffect(() => {
    Cookies.remove('userId');
    Cookies.remove('role');

    history.push('/');
  }, [history]);

  return null;
};

export default Logout;