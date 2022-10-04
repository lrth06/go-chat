export function SuccessAlert(message) {
  console.log(message);
  return (
    <div
      className="px4 py3 relative rounded border border-green-400 bg-green-100 text-green-700"
      role="alert"
    >
      <strong className="font-bold">Success!</strong>
      <br />
      <strong>{message}</strong>
    </div>
  );
}
