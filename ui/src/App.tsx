import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Login from './pages/Login';
import ApplicationsOverview from './pages/ApplicationsOverview';
import ApplicationDetail from './pages/ApplicationDetail';

function App() {
  return (
    <Router>
      <div className="min-h-screen bg-base-200">
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/applications" element={<ApplicationsOverview />} />
          <Route path="/applications/:id" element={<ApplicationDetail />} />
          <Route path="/" element={<Navigate to="/applications" replace />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
