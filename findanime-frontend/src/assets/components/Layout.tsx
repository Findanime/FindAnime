// components/Layout.tsx
import { Outlet } from 'react-router-dom';
import Navbar from './Navbar';

export default function Layout() {
  return (
    <div>
      <Navbar />
      <main>
        <Outlet /> {/* renders the page component here */}
      </main>
    </div>
  );
}