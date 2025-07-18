// App.tsx
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { Home } from './assets/pages/home/Home';
import Layout from './assets/components/Layout';
export default function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
        </Route>
      </Routes>
    </Router>
  );
}
