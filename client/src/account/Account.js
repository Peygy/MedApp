import React from 'react';
import { gql, useQuery } from '@apollo/client';
import Cookies from 'js-cookie';

const GET_USER_INFO_QUERY = gql`
  query GetUserInfo($input: UserAccountData!) {
    getUserInfo(input: $input) {
      username
    }
  }
`;

function Account() {
  const userId = Cookies.get('userId');

  const { data, loading, error } = useQuery(GET_USER_INFO_QUERY, {
    variables: { input: { userId } },
    context: {
      uri: 'http://localhost:4000/graphql/account',
      credentials: 'include',
    },
  });

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  return (
    <div>
      <h1>Account</h1>
      {data && <p>Username: {data.getUserInfo.username}</p>}
    </div>
  );
}

export default Account;