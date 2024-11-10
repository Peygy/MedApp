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

  if (loading) return <p>Загрузка...</p>;
  if (error) return <p>Ошибка: {error.message}</p>;

  return (
    <div>
      <h1>Ваши записи к врачам</h1>
      {data.getUserVisitRecords.length === 0 ? (
        <p>No visit records found.</p>
      ) : (
        <table>
          <thead>
            <tr>
              <th>Номер записи</th>
              <th>ФИО врача</th>
              <th>Специализация врача</th>
              <th>Дата приема</th>
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
      <Link to="/account">Вернуться в личный кабинет</Link>
    </div>
  );
}

export default VisitRecords;