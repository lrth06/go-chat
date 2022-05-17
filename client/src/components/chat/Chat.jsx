import { useEffect } from 'react';
import Message from './Message';
export default function Chat({ messages }) {
  useEffect(() => {
    const messagesEnd = document.getElementById('end');
    messagesEnd.scrollIntoView({ behavior: 'smooth' });
  }, [messages.length]);
  return (
    <div
      style={{ height: 'calc( 100vh - 17vh )' }}
      className="max-h-screen w-full bg-gray-100 p-2"
    >
      <div
        style={{ scrollSnapAlign: 'end', scrollSnapType: 'y proximity' }}
        className="h-fit max-h-full overflow-y-scroll"
      >
        <div
          id="messages"
          className="mx-auto mb-auto grid h-full w-full grid-cols-1 grid-rows-2 space-y-5"
        >
          <>
            {messages.map((message, index) => (
              <Message key={index} message={JSON.stringify(message)} />
            ))}
            <div id="end" />
          </>
        </div>
      </div>
    </div>
  );
}
