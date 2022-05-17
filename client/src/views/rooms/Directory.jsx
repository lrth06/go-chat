import { useForm } from '../../hooks/useForm';
import { useNavigate } from 'react-router-dom';
export default function RoomIndex() {
  const [values, handleChange] = useForm({
    room: '',
  });
  const token = localStorage.getItem('token') || '12345';
  const headers = new Headers({ Authorization: 'Bearer ' + token });

  const navigate = useNavigate();
  function handleSubmit(e) {
    e.preventDefault();
    goToRoom(values.room);
  }
  function goToRoom(room) {
    if (!room) {
      fetch('http://localhost:3000/api/v1/random', {
        method: 'GET',
        headers,
      })
        .then((res) => res.json())
        .then((data) => {
          return navigate(`/room/${data.id}`);
        });
    }
    navigate(`/room/${room}`);
  }

  return (
    <div className="container mx-auto min-h-full p-5 font-primary">
      <section className="body-font text-gray-600">
        <div className="container mx-auto px-5 py-24">
          <div className="mb-20 flex w-full flex-col flex-wrap items-center text-center">
            <h1 className="title-font mb-2 text-2xl font-medium text-gray-900 sm:text-3xl">
              Directory
            </h1>
            <p className="w-full font-secondary leading-relaxed text-gray-500 lg:w-1/2">
              Browse public rooms or create your own.
            </p>
          </div>
          <div className="-m-4 flex flex-wrap">
            <div className="p-4 md:w-1/2 xl:w-1/3">
              <div className="rounded-lg border border-gray-200 p-6">
                <div className="mb-4 inline-flex h-10 w-10 items-center justify-center rounded-full bg-blue-100 text-blue-500">
                  <svg
                    fill="none"
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="2"
                    className="h-6 w-6"
                    viewBox="0 0 24 24"
                  >
                    <path d="M22 12h-4l-3 9L9 3l-3 9H2"></path>
                  </svg>
                </div>
                <h2 className="title-font mb-2 text-lg font-medium text-gray-900">
                  <a href="/room/general" className="text-blue-500">
                    General
                  </a>
                </h2>
                <p className="font-primary text-sm leading-relaxed">
                  A place for all. Users capped at 100.
                </p>
              </div>
            </div>
          </div>
          <form className="p-12 md:p-24" onSubmit={handleSubmit}>
            <input
              className="w-full bg-gray-200 py-2 pl-12 focus:outline-none md:py-4"
              required={true}
              type="text"
              name="room"
              placeholder="Room Name"
              value={values.room}
              onChange={handleChange}
            />
            <button
              className="mx-auto mt-16 flex rounded border-0 bg-blue-500 py-2 px-8 text-lg text-white hover:bg-blue-600 focus:outline-none"
              type="submit"
            >
              Send
            </button>
          </form>
          <button
            name="random"
            value="random"
            onClick={() => goToRoom()}
            className="mx-auto mt-16 flex rounded border-0 bg-green-500 py-2 px-8 text-lg text-white hover:bg-green-600 focus:outline-none"
          >
            Random Room
          </button>
        </div>
      </section>
    </div>
  );
}
