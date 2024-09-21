import React from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import Home from './../home/Home';
import SignUp from '../signUp/SignUp';
import SignIn from '../signIn/SignIn';
import Account from '../account/Account';
import Notes from '../note/note';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/signup" element={<SignUp />} />
        <Route path="/signin" element={<SignIn />} />
        <Route path="/account" element={<Account />} />
        <Route path="/logout" element={<Account />} />
        <Route path="/notes" element={<Notes />} />
        <Route path="*" element={<Navigate to="/" />} /> {/* Редирект на главную, если путь неверный */}
      </Routes>
    </Router>
  );
}

export default App;