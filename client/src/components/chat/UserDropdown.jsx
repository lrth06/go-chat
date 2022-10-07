export default function UserDropdown({ user }) {
  return (
    <>
      <div className="hidden items-center space-x-1 md:flex">
        <a
          href={`/users/${user.id}`}
          // target="_blank"
          // rel="noopener noreferrer"
          className="py-5 px-3 text-gray-700 hover:text-gray-900"
        >
          <img
            id="avatar"
            type="button"
            data-dropdown-toggle="userDropdown"
            data-dropdown-placement="bottom-start"
            className="h-7 w-7 cursor-pointer rounded-full bg-slate-400"
            src={user.avatar}
            alt="User Dropdown"
          />
        </a>
      </div>
    </>
  );
}
