export default async function Logout(token) {
  const res = await fetch('/api/v1/auth/logout', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`
    }
  });
  let data = await res.json();

  // set token to invalidated token in case it is not removed from local storage
  localStorage.setItem('token', data.token);
  return data;
}
