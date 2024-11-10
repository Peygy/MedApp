import React from 'react';
import { useQuery, useMutation, gql } from '@apollo/client';
import Cookies from 'js-cookie';

const GET_DOCTORS = gql`
  query GetDoctors {
    getDoctors {
      doctorId
      doctorName
      specialization
      experienceYears
    }
  }
`;

const ADD_VISIT_RECORD = gql`
  mutation AddVisitRecord($input: AddVisitRecordInput!) {
    addVisitRecord(input: $input)
  }
`;

const DoctorsList = () => {
  const userId = Cookies.get('userId'); // Получение userId из cookie

  const { loading, error, data } = useQuery(GET_DOCTORS, {
    context: {
      uri: 'http://localhost:4000/graphql/notes',
      credentials: 'include',
    },
  });

  const [addVisitRecord] = useMutation(ADD_VISIT_RECORD, {
    context: {
      uri: 'http://localhost:4000/graphql/notes/add',
      credentials: 'include',
    },
  });

  if (!userId) {
    return <p>Не удалось получить идентификатор пользователя из cookie.</p>;
  }

  if (loading) return <p>Загрузка...</p>;
  if (error) return <p>Ошибка: {error.message}</p>;

  const handleRecordVisit = (doctorName, specialization) => {
    let visitDate = new Date();
    visitDate.setDate(visitDate.getDate() + 3);
  
    const formattedVisitDate = visitDate.toLocaleDateString('ru-RU', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit'
    });

    addVisitRecord({
      variables: {
        input: {
          userId,
          doctorName,
          specialization,
          visitDate: formattedVisitDate,
        },
      },
    })
      .then(() => {
        alert('Запись на прием успешно добавлена!');
      })
      .catch((err) => {
        console.error('Ошибка при добавлении записи:', err);
        alert('Не удалось добавить запись на прием.');
      });
  };

  return (
    <div>
      <h1>Список врачей</h1>
      <ul>
        {data?.getDoctors.map((doctor) => (
          <li key={doctor.doctorId}>
            <p>Фио врача: {doctor.doctorName}</p>
            <p>Специализация: {doctor.specialization}</p>
            <p>Опыт: {doctor.experienceYears} лет</p>
            <button onClick={() => handleRecordVisit(doctor.doctorName, doctor.specialization)}>
              Записаться на прием
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default DoctorsList;
