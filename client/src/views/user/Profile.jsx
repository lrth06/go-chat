import Logout from '../../utils/logout';
export default function Profile() {

  function logout () {
    // TODO: #7 Use the logout function to invalidate the token
   return Logout(localStorage.getItem('token')).then(() => {
      localStorage.removeItem('token');
      window.location.href = '/';
  })};

  //redirect to home if page is accessed without a token
  if (!localStorage.getItem('token')) {
    window.location.href = '/';
  }

  return (
    <div className="container mx-auto min-h-full p-5">
      <h1>Profile</h1>
      {/* logout button */}
      <button onClick={()=>{logout()}} className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded">
        Logout
      </button>
    </div>
  );
}
