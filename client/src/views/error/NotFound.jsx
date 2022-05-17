import { useLocation } from 'react-router-dom';
export default function NotFoundPage() {
  const location = useLocation();
  return (
    <div>
      <h1>404</h1>
      <h3>{location} not found. How did you get here?</h3>
    </div>
  );
}
