import { render, screen } from '@testing-library/react';
import App from './App';

test('Renders Home Page', () => {
  render(<App />);
  const linkElement = screen.getByText(/Enter a Room Name:/i);
  expect(linkElement).toBeInTheDocument();
});
