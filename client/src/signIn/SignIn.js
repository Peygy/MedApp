import React, { useState } from 'react';
import { gql, useMutation } from '@apollo/client';
import { useNavigate } from 'react-router-dom';
import Cookies from 'js-cookie';

const SIGN_IN_MUTATION = gql`
  mutation SignIn($input: AuthData!) {
    signIn(input: $input) {
      userId
      role
    }
  }
`;

function SignIn() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [signIn, { loading, error }] = useMutation(SIGN_IN_MUTATION, {
    context: {
      uri: 'http://localhost:4000/graphql/signin'
    }
  });
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();
    signIn({ variables: { input: { username, password } } })
      .then((response) => {
        const { userId, role } = response.data.signIn;

        // Сохраняем в куки
        Cookies.set('userId', userId, { expires: 7 });
        Cookies.set('role', role, { expires: 7 });

        // Перенаправление после успешного входа
        navigate('/account');
      })
      .catch((err) => {
        console.error('Error during sign in:', err);
      });
  };

  return (
    <div>
      <h1>Login</h1>
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
          {loading ? 'Logging In...' : 'Login'}
        </button>
      </form>
      {error && <p>Error: {error.message}</p>}
    </div>
  );
}

export default SignIn;
