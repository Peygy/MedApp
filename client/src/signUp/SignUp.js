import React, { useState } from 'react';
import { gql, useMutation } from '@apollo/client';
import { useNavigate } from 'react-router-dom';
import Cookies from 'js-cookie';

const SIGN_UP_MUTATION = gql`
  mutation SignUp($input: UserInput!) {
    signUp(input: $input) {
      userId
      role
    }
  }
`;

function SignUp() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [signUp, { loading, error }] = useMutation(SIGN_UP_MUTATION);
  const navigate = useNavigate();  // Хук для редиректа

  const handleSubmit = (e) => {
    e.preventDefault();
    signUp({ variables: { input: { username, password } } })
      .then((response) => {
        const { userId, role } = response.data.signUp;

        // Сохраняем в куки
        Cookies.set('userId', userId, { expires: 7 });
        Cookies.set('role', role, { expires: 7 });

        // Редирект на страницу входа
        navigate('/');
      })
      .catch((err) => {
        console.error('Error during sign up:', err);
      });
  };

  return (
    <div>
      <h1>Sign Up</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={username}
          placeholder="Username"
          onChange={(e) => setUsername(e.target.value)}
        />
        <input
          type="password"
          value={password}
          placeholder="Password"
          onChange={(e) => setPassword(e.target.value)}
        />
        <button type="submit" disabled={loading}>
          {loading ? 'Signing Up...' : 'Sign Up'}
        </button>
      </form>
      {error && <p>Error: {error.message}</p>}
    </div>
  );
}

export default SignUp;
