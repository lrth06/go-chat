import axios from "axios";
export default async function Logout(token) {
  try {
  const res = await axios.post(
    "/api/v1/auth/logout",
    {},
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );


  // set token to invalidated token in case it is not removed from local storage
  localStorage.setItem('token', res.data.token);
  return data;
  } catch (e) {
    console.log(e);
  }
}
