import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';

test('renders home page test', () => {
  render(<App />);
  const linkElement = screen.getByText(/Home page/i);
  expect(linkElement).toBeInTheDocument();
});
