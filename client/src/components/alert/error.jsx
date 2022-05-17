export function errorAlert({ message }) {
  return (
    <div className="w-full bg-red-100 p-5 sm:w-1/2">
      <div className="flex space-x-3">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          className="h-4 w-4 flex-none fill-current text-red-500"
        >
          <path d="M12 0c-6.627 0-12 5.373-12 12s5.373 12 12 12 12-5.373 12-12-5.373-12-12-12zm4.597 17.954l-4.591-4.55-4.555 4.596-1.405-1.405 4.547-4.592-4.593-4.552 1.405-1.405 4.588 4.543 4.545-4.589 1.416 1.403-4.546 4.587 4.592 4.548-1.403 1.416z" />
        </svg>
        <div className="flex flex-col space-y-2 leading-tight">
          <div className="text-sm font-medium text-red-700">{message.err}</div>
          <div className="flex-1 text-sm leading-snug text-red-600">
            {message.msg}
          </div>
        </div>
      </div>
    </div>
  );
}
