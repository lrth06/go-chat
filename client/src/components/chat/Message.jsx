import TimeAgo from '../TimeAgo';
import UserDropdown from './UserDropdown';
export default function Message({ message }) {
  let msg = JSON.parse(message);
  if (msg.isServer) {
    return (
      <div className="h-7 rounded bg-gray-200 text-center">
        <span className="max-w-4 w-fit p-5 text-xs">
          <b>Server: </b>
          {msg.text}{' '}
          <span className="p-1 text-xs text-gray-400">
            <TimeAgo date={msg.time} />
          </span>
        </span>
      </div>
    );
  }
  return (
    <div className={msg.isSelf ? 'place-self-start' : 'place-self-end'}>
      <div className="flex items-center">
        {msg.isSelf ? (
          <>
            <UserDropdown user={msg.user} />

            <div className="max-w-prose rounded-lg rounded-l-none bg-white p-3 text-sm">
              <b>You: </b> {msg.text}
            </div>
            <span className="place-self-center p-1 text-xs text-gray-400">
              <TimeAgo date={msg.time} />
            </span>
          </>
        ) : (
          <>
            <span className="place-self-center p-1 text-xs text-gray-400">
              <TimeAgo date={msg.time} />
            </span>
            <div className="max-w-prose rounded-lg rounded-r-none bg-blue-100 p-3  text-sm">
              <b>{msg.user.name}: </b> {msg.text}
            </div>
            <UserDropdown user={msg.user} />
          </>
        )}
      </div>
    </div>
  );
}
