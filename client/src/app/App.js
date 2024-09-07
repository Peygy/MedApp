import React from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import Home from './../home/Home';
import SignIn from '../signIn/SignIn';
import SignUp from '../signup/SignUp';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/signup" element={<SignUp />} />
        <Route path="/signin" element={<SignIn />} />
        <Route path="*" element={<Navigate to="/" />} /> {/* Редирект на главную, если путь неверный */}
      </Routes>
    </Router>
  );
}

export default App;