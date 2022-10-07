import { useContext,useEffect, useState } from 'react';
import { UserContext } from "../../context/UserContext";
import Logout from '../../utils/logout';
import parseJwt from '../../utils/parseJwt';
import axios from 'axios';
import { SuccessAlert } from '../../components/alert/success';
import {useNavigate} from 'react-router-dom';
export default function Profile() {
  const navigate = useNavigate();
  const { user, setUser } = useContext(UserContext);
  function logout() {
    return Logout(localStorage.getItem('token')).then(() => {
      localStorage.removeItem('token');
      window.location.href = '/';
    });
  }
  function deleteUser(){
    try{
      let res = axios.delete(`/api/v1/user/${user.id}`, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
      });
      console.log(res);
      localStorage.removeItem('token');
      SuccessAlert('Account deleted');
      setTimeout(() => {
        window.location.href = '/';
      }, 5000);
    }catch(e){
      console.log(e);
    }
  }



  //redirect to home if page is accessed without a token
  if (!localStorage.getItem('token')) {
    window.location.href = '/';
  }

  return (
    <div className="container  mx-auto min-h-full p-5">
      <div className="flex flex-col items-center justify-center">
        <img src={user?.avatar} alt="avatar" className="w-32 h-32 rounded-full" />
        <h1 className="text-2xl font-bold">{user?.name}</h1>
        <p className="text-gray-500">{user?.email}</p>
        <button
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mt-5"
          onClick={logout}
        >
          Logout
        </button>
        <button
          onClick={() => navigate(`/users/${user?.id}/edit`)}
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mt-5">
            Edit Profile
            </button>
          {/*delete account button with confirmation  */}
          <button
            className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded mt-5"
            onClick={() => {
              if (window.confirm('Are you sure you wish to delete your account?')) {
                deleteUser();
              }
            }}
          >
            Delete Account
          </button>

      </div>
    </div>
  );
}
