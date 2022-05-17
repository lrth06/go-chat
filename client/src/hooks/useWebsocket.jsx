// useEffect(() => {
//   if (!id || !token || id === 'undefined') {
//     return;
//   }
//   ws.current = new WebSocket(
//     `ws://localhost:3000/ws/${id}?v=1.0&token=${token}`
//   );
// }, [id, token]);

// useEffect(() => {
//   if (!ws.current) {
//     return;
//   }
//   ws.current.onopen = (e) => {
//     console.log('connected');
//     e.target.send(JSON.stringify({ type: 'join', data: token }));
//   };
//   ws.current.onmessage = (e) => {
//     const data = JSON.parse(e.data);
//     setInfo(data);
//     if (data.type === 'ping') {
//       let msg = {
//         isServer: true,
//         text: data.data,
//       };
//       setMessages((prev) => [...prev, msg]);
//     }
//     if (data.type === 'message') {
//       let msg = {
//         isServer: false,
//         text: data.data,
//         isSelf: false,
//         user:user
//       };
//       console.log(msg);
//       setMessages((prev) => [...prev, msg]);
//     }
//   };
//   ws.current.onclose = () => {
//     console.log('disconnected');
//   };
//   return () => {
//     ws.current.close();
//   };
// }, [token]);

// useEffect(() => {
//   const interval = setInterval(() => {
//     if (run) {
//       ws.current.send(
//         JSON.stringify({
//           type: 'ping',
//           data: new Date().toLocaleTimeString(),
//         })
//       );
//     }
//   }, 250);
//   return () => clearInterval(interval);
// }, [run]);

export function useWebsocket(ref, id) {
  const ws = useRef(null);
  ws.current = new WebSocket(
    `ws://localhost:3000/ws/${id}?v=1.0&token=${localStorage.getItem('token')}`
  );
  useEffect(() => {
    if (!id || !localStorage.getItem('token') || id === 'undefined') {
      return;
    }
    ws.current.onopen = (e) => {
      console.log('connected');
      e.target.send(
        JSON.stringify({ type: 'join', data: localStorage.getItem('token') })
      );
    };
    ws.current.onclose = () => {
      console.log('disconnected');
    };
  }, [id, localStorage.getItem('token')]);
  return ws.current;
}
