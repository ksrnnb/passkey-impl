import React from 'react';
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import SignIn from './pages/SignIn';
import { useAuth } from './hooks/Auth';
import AuthContext from './context/AuthContext';

function App() {
  const {isLoading, token, setToken, user} = useAuth();

  if (isLoading) {
    return <></>;
  }

  return (
    <AuthContext.Provider value={{user, setToken}}>
      <BrowserRouter>
        <Routes>
          {token ?
            <>
              <Route path="/signin" element={<Navigate to="/" />} />
              <Route path="/" element={<Home />} />
            </>
          :
            <>
              <Route path="/signin" element={<SignIn />} />
              <Route path="*" element={<Navigate to="/signin" />} />
            </>
          }
        </Routes>
      </BrowserRouter>
    </AuthContext.Provider>
  );
}

export default App;
