import { UserContext } from '../context/UserContext';
import { useEffect, useRef, useState, useContext } from 'react';
import UserDropdown from './chat/UserDropdown';
import {useNavigate} from 'react-router-dom';
export function Header() {
  const navigate = useNavigate();
  const [hidden, setHidden] = useState(true);
  const menuRef = useRef(null);
  const { user, setUser } = useContext(UserContext);
  const handleToggle = () => {
    setHidden(!hidden);
  };
  return (
    <nav className="h-15 bg-gray-100">
      <div className="mx-auto px-4">
        <div className="flex justify-between">
          <div className="flex space-x-4">
            <div>
              <a
                href="/"
                className="flex items-center py-5 px-2 text-gray-700 hover:text-gray-900"
              >
                <img src="/images/chat.svg" className="h-fit px-1" />
                <span className="font-bold">Go-Chat</span>
              </a>
            </div>

            <div className="hidden items-center space-x-1 md:flex">
              <a
                href="/room"
                className="py-5 px-3 text-gray-700 hover:text-gray-900"
              >
                Rooms
              </a>
              <a
                href="/blog"
                className="py-5 px-3 text-gray-700 hover:text-gray-900"
              >
                Blog
              </a>
            </div>
          </div>

          {user ? (
            <UserDropdown user={user} />
          ) : (
            <div className="hidden items-center space-x-1 md:flex">
              <a href="/auth/login" className="py-5 px-3">
                Login
              </a>
              <a
                href="/auth/register"
                className="hover:text-white-800 rounded bg-green-400 py-2 px-3 text-gray-100 transition duration-300 hover:bg-green-300"
              >
                Signup
              </a>
            </div>
          )}

          <div
            onClick={(e) => handleToggle(e)}
            className="flex items-center md:hidden"
          >
            <button className="mobile-menu-button">
              <svg
                className="h-6 w-6"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M4 6h16M4 12h16M4 18h16"
                />
              </svg>
            </button>
          </div>
        </div>
      </div>
      <div
        ref={menuRef}
        className={`mobile-menu md:hidden ${hidden == true && 'hidden'}`}
      >
        <a onClick={()=>navigate("/rooms")} className="block py-2 px-4 text-sm hover:bg-gray-200">
          Rooms
        </a>
        <a onClick={()=>navigate("/blog")} className="block py-2 px-4 text-sm hover:bg-gray-200">
          Blog
        </a>
      </div>
    </nav>
  );
}

export function Footer() {
  return (
    <footer className="body-font text-gray-600">
      <div className="container mx-auto flex flex-col items-center px-5 py-8 sm:flex-row">
        <a className="title-font flex items-center justify-center font-medium text-gray-900 md:justify-start">
          <img src="/images/chat.svg" alt="" />
          <span className="ml-3 text-xl">Go-Chat</span>
        </a>
        <p className="mt-4 text-sm text-gray-500 sm:ml-4 sm:mt-0 sm:border-l-2 sm:border-gray-200 sm:py-2 sm:pl-4">
          © {new Date().getFullYear()} Go-Chat —
          <a
            href="https://github.com/lrth06"
            className="ml-1 text-gray-600"
            rel="noopener noreferrer"
            target="_blank"
          >
            @Lrth06
          </a>
        </p>
        <span className="mt-4 inline-flex justify-center sm:ml-auto sm:mt-0 sm:justify-start">
          <a
            href="https://github.com/lrth06/go-chat"
            rel="noopener noreferrer"
            target="_blank"
            className="text-gray-500"
          >
            <svg
              className="h-5 w-5"
              fill="currentColor"
              xmlns="http://www.w3.org/2000/svg"
              width="1024"
              height="1024"
              viewBox="0 0 1024 1024"
            >
              <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M8 0C3.58 0 0 3.58 0 8C0 11.54 2.29 14.53 5.47 15.59C5.87 15.66 6.02 15.42 6.02 15.21C6.02 15.02 6.01 14.39 6.01 13.72C4 14.09 3.48 13.23 3.32 12.78C3.23 12.55 2.84 11.84 2.5 11.65C2.22 11.5 1.82 11.13 2.49 11.12C3.12 11.11 3.57 11.7 3.72 11.94C4.44 13.15 5.59 12.81 6.05 12.6C6.12 12.08 6.33 11.73 6.56 11.53C4.78 11.33 2.92 10.64 2.92 7.58C2.92 6.71 3.23 5.99 3.74 5.43C3.66 5.23 3.38 4.41 3.82 3.31C3.82 3.31 4.49 3.1 6.02 4.13C6.66 3.95 7.34 3.86 8.02 3.86C8.7 3.86 9.38 3.95 10.02 4.13C11.55 3.09 12.22 3.31 12.22 3.31C12.66 4.41 12.38 5.23 12.3 5.43C12.81 5.99 13.12 6.7 13.12 7.58C13.12 10.65 11.25 11.33 9.47 11.53C9.76 11.78 10.01 12.26 10.01 13.01C10.01 14.08 10 14.94 10 15.21C10 15.42 10.15 15.67 10.55 15.59C13.71 14.53 16 11.53 16 8C16 3.58 12.42 0 8 0Z"
                transform="scale(64)"
                fill="#1B1F23"
              />
            </svg>
          </a>
        </span>
      </div>
    </footer>
  );
}
