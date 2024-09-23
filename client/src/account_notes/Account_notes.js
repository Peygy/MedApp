import React from 'react';
import { gql, useQuery } from '@apollo/client';
import Cookies from 'js-cookie';
import { Link } from 'react-router-dom';

const GET_USER_VISIT_RECORDS_QUERY = gql`
  query GetUserVisitRecords($input: UserAccountData!) {
    getUserVisitRecords(input: $input) {
      recordNumber
      doctorName
      specialization
      visitDate
    }
  }
`;

function VisitRecords() {
  const userId = Cookies.get('userId');
  const { data, loading, error } = useQuery(GET_USER_VISIT_RECORDS_QUERY, {
    variables: { input: { userId } },
    context: {
      uri: 'http://localhost:4000/graphql/account/notes',
      credentials: 'include',
    },
  });

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  return (
    <div>
      <h1>Visit Records</h1>
      {data.getUserVisitRecords.length === 0 ? (
        <p>No visit records found.</p>
      ) : (
        <table>
          <thead>
            <tr>
              <th>Record Number</th>
              <th>Doctor Name</th>
              <th>Specialization</th>
              <th>Visit Date</th>
            </tr>
          </thead>
          <tbody>
            {data.getUserVisitRecords.map((record) => (
              <tr key={record.recordNumber}>
                <td>{record.recordNumber}</td>
                <td>{record.doctorName}</td>
                <td>{record.specialization}</td>
                <td>{record.visitDate}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
      <Link to="/account">Back to Account</Link>
    </div>
  );
}

export default VisitRecords;