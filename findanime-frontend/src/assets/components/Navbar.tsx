// components/Navbar.tsx
import { Link } from 'react-router-dom';

export default function Navbar() {
  return (
    <nav className="text-white px-6 py-4 flex justify-between items-center">
      <Link to="/" className="text-xl font-bold">F<span className="text-pink-500">!</span>ndAnime</Link>
      <div className="space-x-4">
        <Link to="/">Home</Link>
        <Link to="/about">About</Link>
        <Link to="/search">Search</Link>
      </div>
    </nav>
  );
}