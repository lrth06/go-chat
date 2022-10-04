import { useEffect, useState } from 'react';
import Logout from '../../utils/logout';
import parseJwt from '../../utils/parseJwt';
export default function Profile() {
  const [user, setUser] = useState();
  function logout() {
    // TODO: #7 Use the logout function to invalidate the token
    return Logout(localStorage.getItem('token')).then(() => {
      localStorage.removeItem('token');
      window.location.href = '/';
    });
  }

  //redirect to home if page is accessed without a token
  if (!localStorage.getItem('token')) {
    window.location.href = '/';
  }

  useEffect(() => {
    //get user data from token
    const token = parseJwt(localStorage.getItem('token'));
    setUser(token);
  }, []);
  return (
    <div className="container mx-auto min-h-full p-5">

      <div className="p-16">
        <div className="mt-24 bg-white p-8 shadow">
          {' '}
          <div className="grid grid-cols-1 md:grid-cols-3">
            {' '}
            <div className="order-last mt-20 grid grid-cols-3 text-center md:order-first md:mt-0">
              {' '}
              <div>
                {' '}
                <p className="text-xl font-bold text-gray-700">22</p>{' '}
                <p className="text-gray-400">Friends</p>{' '}
              </div>{' '}
              <div>
                {' '}
                <p className="text-xl font-bold text-gray-700">10</p>{' '}
                <p className="text-gray-400">Photos</p>{' '}
              </div>{' '}
              <div>
                {' '}
                <p className="text-xl font-bold text-gray-700">89</p>{' '}
                <p className="text-gray-400">Comments</p>{' '}
              </div>{' '}
            </div>{' '}
            <div className="relative">
              {' '}
              <div className="absolute inset-x-0 top-0 mx-auto -mt-24 flex h-48 w-48 items-center justify-center rounded-full bg-indigo-100 text-indigo-500 shadow-2xl">
                <img
                  className="h-24 w-24 rounded-full"
                  src={user?.avatar}
                  alt="profile"
                />
              </div>{' '}
            </div>{' '}
            <div className="mt-32 flex justify-between space-x-8 md:mt-0 md:justify-center">
              <button className="transform rounded bg-blue-400 py-2 px-4 font-medium uppercase text-white shadow transition hover:-translate-y-0.5 hover:bg-blue-500 hover:shadow-lg">
                {' '}
                Edit Profile
              </button>{' '}
              <button onClick={()=>{
                logout();
              }} className="transform rounded bg-red-700 py-2 px-4 font-medium uppercase text-white shadow transition hover:-translate-y-0.5 hover:bg-red-800 hover:shadow-lg">
                {' '}
                Logout
              </button>{' '}
            </div>{' '}
          </div>{' '}
          <div className="mt-20 border-b pb-12 text-center">
            {' '}
            <h1 className="text-4xl font-medium text-gray-700">
              {user?.name}
              <span className="text-2xl font-light text-gray-500">{" "}({user?.email})</span>
            </h1>{' '}
            {/* <p className="mt-3 font-light text-gray-600">Bucharest, Romania</p>{' '} */}
            {/* <p className="mt-8 text-gray-500">
              Solution Manager - Creative Tim Officer
            </p>{' '}
            <p className="mt-2 text-gray-500">University of Computer Science</p>{' '} */}
          </div>{' '}
          <div className="mt-12 flex flex-col justify-center">
            {' '}
            <p className="text-center font-light text-gray-600 lg:px-16">
              {/* An artist of considerable range, Ryan — the name taken by
              Melbourne-raised, Brooklyn-based Nick Murphy — writes, performs
              and records all of his own music, giving it a warm, intimate feel
              with a solid groove structure. An artist of considerable range. */}
            </p>{' '}
            {/* <button className="mt-4 py-2 px-4  font-medium text-indigo-500">
              {' '}
              Show more
            </button>{' '} */}
          </div>
        </div>
      </div>
    </div>
  );
}
