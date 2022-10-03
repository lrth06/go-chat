export function SuccessAlert(message) {
  console.log(message);
  return (
    <div className="bg-green-100 border border-green-400 text-green-700 px4 py3 rounded relative" role="alert">
    <strong className="font-bold">Success!</strong>
    <br />
    <strong>{message}</strong>
    </div>
  );
}
