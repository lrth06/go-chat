export default function Profile() {
  function logout () {
    localStorage.removeItem('token');
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
