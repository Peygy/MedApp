import React, { useState } from 'react';
import { gql, useQuery, useMutation } from '@apollo/client';
import Cookies from 'js-cookie';

const GET_USER_INFO_QUERY = gql`
  query GetUserInfo($input: UserAccountData!) {
    getUserInfo(input: $input) {
      username
      age
      height
      weight
      pulse
      pressure
      dailyWater
      bodyMassIndex
    }
  }
`;

const UPDATE_USER_HEALTH_DATA_MUTATION = gql`
  mutation UpdateUserHealthData($input: UpdateUserHealthDataInput!) {
    updateUserHealthData(input: $input) {
      username
      age
      height
      weight
      pulse
      pressure
      dailyWater
      bodyMassIndex
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

  const [updateUserHealthData] = useMutation(UPDATE_USER_HEALTH_DATA_MUTATION, {
    context: {
      uri: 'http://localhost:4000/graphql/account/update',
      credentials: 'include',
    },
  });

  const [formState, setFormState] = useState({});

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormState({
      ...formState,
      [name]: value ? parseFloat(value) : '',
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    // Значения, которые были получены с сервера
    const defaultData = data.getUserInfo;

    // Заполняем пропущенные данные из основного блока
    const input = {
      userId,
      age: formState.age !== undefined ? formState.age : defaultData.age,
      height: formState.height !== undefined ? formState.height : defaultData.height,
      weight: formState.weight !== undefined ? formState.weight : defaultData.weight,
      pulse: formState.pulse !== undefined ? formState.pulse : defaultData.pulse,
      pressure: formState.pressure !== undefined ? formState.pressure : defaultData.pressure,
    };

    console.log('Submitting:', input);
    updateUserHealthData({ variables: { input } }).then(() => {
      window.location.reload();
    })
    .catch((err) => {
      console.error('Ошибка обновления данных:', err);
    });
  };

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  const {
    username,
    age,
    height,
    weight,
    pulse,
    pressure,
    dailyWater,
    bodyMassIndex,
  } = data.getUserInfo;

  const displayValue = (value, defaultText = 'Данные не установлены') =>
    value === 0 || value === '' || value === 0.0 ? defaultText : value;

  return (
    <div>
      <h1>Account</h1>
      <p>Username: {username}</p>
      <p>Age: {displayValue(age)}</p>
      <p>Height: {displayValue(height)}</p>
      <p>Weight: {displayValue(weight)}</p>
      <p>Pulse: {displayValue(pulse)}</p>
      <p>Pressure: {displayValue(pressure)}</p>
      <p>
        Daily Water: {weight ? displayValue(dailyWater, 'Требуется вес') : 'Требуется вес'}
      </p>
      <p>
        Body Mass Index:{' '}
        {weight && height
          ? displayValue(bodyMassIndex, 'Требуются вес и рост')
          : 'Требуются вес и рост'}
      </p>

      <h2>Update Health Data</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Age:
          <input
            type="number"
            name="age"
            value={formState.age || ''}
            onChange={handleChange}
          />
        </label>
        <label>
          Height:
          <input
            type="number"
            name="height"
            step="0.01"
            value={formState.height || ''}
            onChange={handleChange}
          />
        </label>
        <label>
          Weight:
          <input
            type="number"
            name="weight"
            step="0.01"
            value={formState.weight || ''}
            onChange={handleChange}
          />
        </label>
        <label>
          Pulse:
          <input
            type="number"
            name="pulse"
            value={formState.pulse || ''}
            onChange={handleChange}
          />
        </label>
        <label>
          Pressure:
          <input
            type="text"
            name="pressure"
            value={formState.pressure || ''}
            onChange={handleChange}
          />
        </label>
        <button type="submit">Update</button>
      </form>
    </div>
  );
}

export default Account;
