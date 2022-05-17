export default function UserList(users) {
  return (
    <div className="-space-x-4">
      {users.map((user) => (
        <div key={user.id} className="flex items-center space-x-2">
          <div className="flex-shrink-0">
            <img
              src={user.avatar}
              className="h-8 w-8 rounded-full"
              alt={user.name}
            />
          </div>
          <div className="flex-1">
            <div className="text-sm font-medium leading-5 text-gray-900">
              {user.name}
            </div>
            <div className="text-sm leading-5 text-gray-500">{user.email}</div>
          </div>
        </div>
      ))}
    </div>
  );
}
