import { useForm } from '../../hooks/useForm';
import parseJwt from '../../utils/parseJwt';
export default function MessageBar({ ws, token }) {
  const [values, handleChange] = useForm({
    message: '',
  });

  const handleSubmit = (e) => {
    e.preventDefault();
    ws.current.send(
      JSON.stringify({
        type: 'message',
        data: {
          text: values.message,
          time: new Date().getTime(),
          user: parseJwt(token),
        },
      })
    );
    values.message = '';
  };
  return (
    <form className="bottom-0 flex items-center py-5" onSubmit={handleSubmit}>
      <input
        className="h-10 w-full rounded-lg bg-gray-200 p-2 outline-none"
        type="text"
        required={true}
        value={values.message}
        onChange={handleChange}
        autoComplete="off"
        name="message"
        placeholder="Type a message..."
      />
      <button
        className="m-1 h-10 w-10 rounded-lg bg-gray-200 p-1"
        type="submit"
      >
        <svg
          className="h-full w-full"
          viewBox="0 0 24 24"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path d="M2.01 21L23 12 2.01 3 2 10l15 2-15 2z" fill="currentColor" />
        </svg>
      </button>
    </form>
  );
}
