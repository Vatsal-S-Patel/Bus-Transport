import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const AdminLogin = () => {
  // States to Manage Form
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  
  // Navigator to Navigate to different Routes
  var navigator = useNavigate();

  // Handle Login Function
  const handleLogin = () => {
    if (username === 'admin' && password === 'password') {
      console.log('Login successful');
      localStorage.setItem('user', "admin");
      navigator("/admin/bus")
    } else {
      setError('Invalid username or password');
    }
  };

  return (
    <div className="flex h-screen justify-center items-center bg-gray-100">
      <div className="bg-white p-8 rounded shadow-md">
        <h2 className="text-2xl font-bold mb-4">Admin Login</h2>
        {error && <p className="text-red-500 mb-4">{error}</p>}
        <input
          type="text"
          placeholder="Username"
          className="w-full border border-gray-300 p-2 mb-4 rounded"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <input
          type="password"
          placeholder="Password"
          className="w-full border border-gray-300 p-2 mb-4 rounded"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button
          className="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600"
          onClick={handleLogin}
        >
          Login
        </button>
      </div>
    </div>
  );
};

export default AdminLogin;
