// src/App.tsx
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import GitHubAuthButton from './components/GitHubAuthButton';
import AuthCallback from './components/AuthCallback';
import Profile from './components/Profile';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<GitHubAuthButton />} />
        <Route path="/auth/callback" element={<AuthCallback />} />
        <Route path="/profile" element={<Profile />} /> {/* Новый маршрут */}
      </Routes>
    </Router>
  );
};

export default App;