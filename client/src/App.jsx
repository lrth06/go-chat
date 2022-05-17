import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { useContext, useEffect, useMemo, useState } from 'react';
import Home from './views/Home';
import { Header, Footer } from './components/Layout';
import RoomIndex from './views/rooms/Directory';
import Room from './views/rooms/Room';
import Profile from './views/user/Profile';
import Login from './views/auth/Login';
import Register from './views/auth/Register';
import RoomEditor from './views/rooms/Editor';
import './css/styles.css';
import { ThemeContext } from './context/ThemeContext';
import { UserContext } from './context/UserContext';
import PasswordReset from './views/auth/PasswordReset';

function App() {
  const [dark, setDark] = useState(null);
  const themeValue = useMemo(() => ({ dark, setDark }), [dark, setDark]);
  const [token, setToken] = useState(localStorage.getItem('token'));
  const [user, setUser] = useState(null);
  const userValue = useMemo(() => ({ user, setUser }), [user, setUser]);
  useEffect(() => {
    if (!token) return;
    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace('-', '+').replace('_', '/');
    const parsedToken = JSON.parse(window.atob(base64));
    setUser(parsedToken);
  }, []);

  return (
    <BrowserRouter>
      <ThemeContext.Provider value={themeValue}>
        <UserContext.Provider value={userValue}>
          <Header />
          <Routes>
            <Route path="*" element={<Navigate to={'/'} />} />
            <Route path="/room/undefined" element={<Navigate to={'/room'} />} />
            <Route path="/" element={<Home />} />
            <Route path="/auth">
              <Route path="login" element={<Login />} />
              <Route path="register" element={<Register />} />
              <Route path="recover" element={<PasswordReset />} />
            </Route>
            <Route path="/users">
              <Route path=":id" element={<Profile />} />
            </Route>
            <Route path="/room">
              <Route index element={<RoomIndex />} />
              <Route path="new" element={<RoomEditor />} />
              <Route path=":id" element={<Room />}>
                <Route path="edit" element={<RoomEditor />} />
              </Route>
            </Route>
          </Routes>
        </UserContext.Provider>
      </ThemeContext.Provider>
    </BrowserRouter>
  );
}

export default App;
