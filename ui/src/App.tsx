import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Login from './pages/auth/Login';
import Register from './pages/auth/Register';
import RegisterSuccess from './pages/auth/RegisterSuccess';
import ForgotPassword from './pages/auth/ForgotPassword';
import ResetPassword from './pages/auth/ResetPassword';
import ApplicationsOverview from './pages/ApplicationsOverview';
import ApplicationDetail from './pages/ApplicationDetail';
import ApplicationEdit from './pages/ApplicationEdit';

function App() {
  return (
    <Router>
      <div className="min-h-screen bg-base-200">
        <Routes>
          {/* Auth routes */}
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/register/success" element={<RegisterSuccess />} />
          <Route path="/forgot-password" element={<ForgotPassword />} />
          <Route path="/reset-password" element={<ResetPassword />} />

          {/* Application routes */}
          <Route path="/applications" element={<ApplicationsOverview />} />
          <Route path="/applications/:id" element={<ApplicationDetail />} />
          <Route path="/applications/:id/edit" element={<ApplicationEdit />} />
          <Route path="/" element={<Navigate to="/applications" replace />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
