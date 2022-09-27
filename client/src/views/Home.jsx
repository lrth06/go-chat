import { useNavigate } from 'react-router-dom';
export default function Home() {
  const navigate = useNavigate();
  return (
    <div className="container mx-auto min-h-full p-5">
      <section className="body-font text-gray-600">
        <div className="container mx-auto flex flex-col items-center px-5 py-24 md:flex-row">
          <div className="mb-16 flex flex-col items-center text-center md:mb-0 md:w-1/2 md:items-start md:pr-16 md:text-left lg:flex-grow lg:pr-24">
            <h1 className="title-font mb-4 text-3xl font-medium text-gray-900 sm:text-4xl">
              Network and Communicate
              <br />
              Instantly and Securely
            </h1>
            <p className="mb-8 leading-relaxed">
              Using the latest security protocols, your data is kept secure and
              confidential, with tools like password protection on chat spaces,
              as well as Admin and Moderator roles for each space.
            </p>
            <div className="flex justify-center">
              <button
                onClick={
                  //scroll to bottom of page
                  () => {
                    window.scrollTo(0, document.body.scrollHeight);
                  }
                }
                className="inline-flex rounded border-0 bg-blue-500 py-2 px-6 text-lg text-white hover:bg-blue-600 focus:outline-none"
              >
                Features
              </button>
              <button
                onClick={() => {
                  navigate('/pricing');
                }}
                className="ml-4 inline-flex rounded border-0 bg-gray-100 py-2 px-6 text-lg text-gray-700 hover:bg-gray-200 focus:outline-none"
              >
                Pricing
              </button>
            </div>
          </div>
          <div className="w-5/6 md:w-1/2 lg:w-full lg:max-w-lg">
            <img
              className="rounded object-cover object-center"
              alt="hero"
              src="/images/support-team.svg"
            />
          </div>
        </div>
      </section>
      <section className="body-font text-gray-600">
        <div className="container mx-auto px-5 py-24">
          <div className="mb-20 text-center">
            <h1 className="title-font mb-4 text-2xl font-medium text-gray-900 sm:text-3xl">
              A Community for Growth
            </h1>
            <p className="text-gray-500s mx-auto text-base leading-relaxed lg:w-3/4 xl:w-2/4">
              Meet and interact with others you know in controlled spaces, or
              join a shared space to expand your network.
            </p>
            <div className="mt-6 flex justify-center">
              <div className="inline-flex h-1 w-16 rounded-full bg-blue-500"></div>
            </div>
          </div>
          <div className="-mx-4 -mb-10 -mt-4 flex flex-wrap space-y-6 sm:-m-4 md:space-y-0">
            <div className="flex flex-col items-center p-4 text-center md:w-1/3">
              <div className="mb-5 inline-flex h-20 w-20 flex-shrink-0 items-center justify-center rounded-full bg-blue-100 text-blue-500">
                <svg
                  fill="none"
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  className="h-10 w-10"
                  viewBox="0 0 24 24"
                >
                  <path d="M22 12h-4l-3 9L9 3l-3 9H2"></path>
                </svg>
              </div>
              <div className="flex-grow">
                <h2 className="title-font mb-3 text-lg font-medium text-gray-900">
                  General
                </h2>
                <p className="text-base leading-relaxed">
                  A place for all users to meet. Users capped at 100.
                </p>
                <a
                  href="/room/general"
                  className="mt-3 inline-flex items-center text-blue-500"
                >
                  Join Room
                  <svg
                    fill="none"
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="2"
                    className="ml-2 h-4 w-4"
                    viewBox="0 0 24 24"
                  >
                    <path d="M5 12h14M12 5l7 7-7 7"></path>
                  </svg>
                </a>
              </div>
            </div>
            <div className="flex flex-col items-center p-4 text-center md:w-1/3">
              <div className="mb-5 inline-flex h-20 w-20 flex-shrink-0 items-center justify-center rounded-full bg-blue-100 text-blue-500">
                <svg
                  fill="none"
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  className="h-10 w-10"
                  viewBox="0 0 24 24"
                >
                  <circle cx="6" cy="6" r="3"></circle>
                  <circle cx="6" cy="18" r="3"></circle>
                  <path d="M20 4L8.12 15.88M14.47 14.48L20 20M8.12 8.12L12 12"></path>
                </svg>
              </div>
              <div className="flex-grow">
                <h2 className="title-font mb-3 text-lg font-medium text-gray-900">
                  Open Collaboration
                </h2>
                <p className="text-base leading-relaxed">
                  Come here to meet new professionals wanting to collab with
                  others.
                  <br /> Community heavily moderated.
                </p>
                <a
                  href="/room/colab"
                  className="mt-3 inline-flex items-center text-blue-500"
                >
                  Join Room
                  <svg
                    fill="none"
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="2"
                    className="ml-2 h-4 w-4"
                    viewBox="0 0 24 24"
                  >
                    <path d="M5 12h14M12 5l7 7-7 7"></path>
                  </svg>
                </a>
              </div>
            </div>
            <div className="flex flex-col items-center p-4 text-center md:w-1/3">
              <div className="mb-5 inline-flex h-20 w-20 flex-shrink-0 items-center justify-center rounded-full bg-blue-100 text-blue-500">
                <svg
                  fill="none"
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  className="h-10 w-10"
                  viewBox="0 0 24 24"
                >
                  <path d="M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2"></path>
                  <circle cx="12" cy="7" r="4"></circle>
                </svg>
              </div>
              <div className="flex-grow">
                <h2 className="title-font mb-3 text-lg font-medium text-gray-900">
                  Rants and Raves
                </h2>
                <p className="text-base leading-relaxed">
                  Need to get something off your chest? Community members are
                  encouraged to share their complaints here.
                  <br /> Community heavily moderated.
                </p>
                <a
                  href="/room/rants"
                  className="mt-3 inline-flex items-center text-blue-500"
                >
                  Join Room
                  <svg
                    fill="none"
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="2"
                    className="ml-2 h-4 w-4"
                    viewBox="0 0 24 24"
                  >
                    <path d="M5 12h14M12 5l7 7-7 7"></path>
                  </svg>
                </a>
              </div>
            </div>
          </div>
          {/* <button className="flex mx-auto mt-16 text-white bg-blue-500 border-0 py-2 px-8 focus:outline-none hover:bg-blue-600 rounded text-lg">Join Random Room</button> */}
        </div>
      </section>

      <section className="body-font text-gray-600">
        <div className="container mx-auto px-5 py-24">
          <div className="mx-auto w-full text-center lg:w-3/4 xl:w-1/2">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              className="mb-8 inline-block h-8 w-8 text-gray-400"
              viewBox="0 0 975.036 975.036"
            >
              <path d="M925.036 57.197h-304c-27.6 0-50 22.4-50 50v304c0 27.601 22.4 50 50 50h145.5c-1.9 79.601-20.4 143.3-55.4 191.2-27.6 37.8-69.399 69.1-125.3 93.8-25.7 11.3-36.8 41.7-24.8 67.101l36 76c11.6 24.399 40.3 35.1 65.1 24.399 66.2-28.6 122.101-64.8 167.7-108.8 55.601-53.7 93.7-114.3 114.3-181.9 20.601-67.6 30.9-159.8 30.9-276.8v-239c0-27.599-22.401-50-50-50zM106.036 913.497c65.4-28.5 121-64.699 166.9-108.6 56.1-53.7 94.4-114.1 115-181.2 20.6-67.1 30.899-159.6 30.899-277.5v-239c0-27.6-22.399-50-50-50h-304c-27.6 0-50 22.4-50 50v304c0 27.601 22.4 50 50 50h145.5c-1.9 79.601-20.4 143.3-55.4 191.2-27.6 37.8-69.4 69.1-125.3 93.8-25.7 11.3-36.8 41.7-24.8 67.101l35.9 75.8c11.601 24.399 40.501 35.2 65.301 24.399z"></path>
            </svg>
            <p className="text-lg leading-relaxed">
              This is a personal project gone rogue. A full blown community with
              membership tiers and enterprise level support. The underlying code
              is largely open source, and collaboration on the project is
              welcomed. To contribute, please submit an issue, rfc, or pull
              request on <a href="https://github.com/lrth06/go-chat">Github</a>
            </p>
            <span className="mt-8 mb-6 inline-block h-1 w-10 rounded bg-blue-500"></span>
            <h2 className="title-font text-sm font-medium tracking-wider text-gray-900">
              TOBY HAGAN
            </h2>
            <p className="text-gray-500">Project Architect and Maintainer</p>
          </div>
        </div>
      </section>
    </div>
  );
}
