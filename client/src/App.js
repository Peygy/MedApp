import React, { useState } from 'react';
import { gql, useMutation } from '@apollo/client';
import Cookies from 'js-cookie';  // Пакет для работы с куки (можно установить через npm install js-cookie)

const SIGN_UP_MUTATION = gql`
  mutation SignUp($input: SignUpInput!) {
    signUp(input: $input) {
      userId
      role
    }
  }
`;

function App() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [signUp, { data, loading, error }] = useMutation(SIGN_UP_MUTATION);

  const handleSubmit = (e) => {
    e.preventDefault();
    signUp({ variables: { input: { username, password } } })
      .then((response) => {
        const { userId, role } = response.data.signUp;
        
        // Сохраняем userId и role в куки
        Cookies.set('userId', userId, { expires: 7 });  // Куки сохраняется на 7 дней
        Cookies.set('role', role, { expires: 7 });

        console.log('User ID and Role saved to cookies:', { userId, role });
      })
      .catch((err) => {
        console.error('Error during sign up:', err);
      });
  };

  return (
    <div>
      <h1>GraphQL Client</h1>
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
        <button type="submit">Submit</button>
      </form>
      {loading && <p>Loading...</p>}
      {error && <p>Error: {error.message}</p>}
      {data && (
        <div>
          <p>User ID: {data.signUp.userId}</p>
          <p>Role: {data.signUp.role}</p>
        </div>
      )}
    </div>
  );
}

export default App;