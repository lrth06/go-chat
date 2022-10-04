import { useState, useRef, useContext, useEffect } from 'react';
import Chat from '../../components/chat/Chat';
import MessageBar from '../../components/chat/MessageBar';
import { UserContext } from '../../context/UserContext';
import { useParams } from 'react-router-dom';
import parseJwt from '../../utils/parseJwt';
import UserList from '../../components/chat/UserList';
export default function Room() {
  const { user } = useContext(UserContext);

  const [messages, setMessages] = useState([]);
  const [users, setUsers] = useState([]);
  const { id } = useParams();
  const [token] = useState(localStorage.getItem('token'));
  const ws = useRef(null);

  //redirect to login if not logged in

  useEffect(() => {
    if (!token) {
      window.location.href = '/auth/login';
    }
  }, [token]);

  useEffect(() => {
    if (!user) return;
    ws.current = new WebSocket(
      `ws://localhost:3000/ws/room/${id}?v=1.0&token=${token}`
    );
    ws.current.onopen = (e) => {
      console.log('connected');
      e.target.send(JSON.stringify({ type: 'join', data: token }));
    };
    ws.current.onmessage = (e) => {
      const data = JSON.parse(e.data);
      console.log(e);
      let msg = {};
      if (data.type === 'join') {
        const sender = parseJwt(data.data);
        if (sender.id !== user.id) {
          console.log(sender.id, user.id);
          let payload = {
            type: 'existingUser',
            data: token,
          };
        }
        msg = {
          isServer: true,
          text: `${sender.name} has joined the room`,
          time: new Date(),
        };
      }
      if (data.type === 'leave') {
        const sender = parseJwt(data.data);
        msg = {
          isServer: true,
          text: `${sender.name} has left the room`,
          time: new Date(),
        };

        console.log(users);
      }
      if (data.type === 'message') {
        console.log(data.data.user.name);
        msg = {
          isSelf: data.data.user.name === user.name,
          text: data.data.text,
          time: data.data.time,
          user: data.data.user,
        };
      }
      setMessages((prev) => [...prev, msg]);
    };
    ws.current.onclose = () => {
      console.log('disconnected');
    };
    return () => {
      ws.current.close();
    };
  }, [user]);



  useEffect(() => {
    window.onbeforeunload = function (e) {
      ws.current.send(JSON.stringify({ type: 'leave', data: token }));
    };
  }, []);

  return (
    <div className="container mx-auto min-h-full p-5">
      <Chat messages={messages} />
      <MessageBar ws={ws} token={token} />
    </div>
  );
}
