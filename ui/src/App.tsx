import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Login from './pages/auth/Login';
import Register from './pages/auth/Register';
import RegisterSuccess from './pages/auth/RegisterSuccess';
import ForgotPassword from './pages/auth/ForgotPassword';
import ResetPassword from './pages/auth/ResetPassword';
import AppLayout from './components/AppLayout';
import ApplicationsOverview from './pages/ApplicationsOverview';
import ApplicationDetail from './pages/ApplicationDetail';
import ApplicationEdit from './pages/ApplicationEdit';
import ApplicationOverview from './pages/ApplicationOverview';

// Placeholder components for other routes
const Entities = () => <h1>Entities Overview</h1>;
const AddEntity = () => <h1>Add Entity</h1>;
const Reports = () => <h1>Reports Overview</h1>;
const Settings = () => <h1>Settings</h1>;
const CustomFields = () => <h1>Custom Fields</h1>;
const BulkImport = () => <h1>Bulk Import</h1>;
const Users = () => <h1>User Management</h1>;

function App() {
  return (
    <Router>
      <Routes>
        {/* Auth routes without layout */}
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/register/success" element={<RegisterSuccess />} />
        <Route path="/forgot-password" element={<ForgotPassword />} />
        <Route path="/reset-password" element={<ResetPassword />} />

        {/* Application routes with layout */}
        <Route path="/software" element={<AppLayout><ApplicationsOverview /></AppLayout>} />
        <Route path="/software/:id" element={<AppLayout><ApplicationDetail /></AppLayout>} />
        <Route path="/software/:id/edit" element={<AppLayout><ApplicationEdit /></AppLayout>} />
        <Route path="/software/overview" element={<AppLayout><ApplicationOverview /></AppLayout>} />

        {/* Other routes with layout */}
        <Route path="/entities" element={<AppLayout><Entities /></AppLayout>} />
        <Route path="/entities/new" element={<AppLayout><AddEntity /></AppLayout>} />
        <Route path="/reports" element={<AppLayout><Reports /></AppLayout>} />
        <Route path="/settings/general" element={<AppLayout><Settings /></AppLayout>} />
        <Route path="/settings/custom-fields" element={<AppLayout><CustomFields /></AppLayout>} />
        <Route path="/settings/bulk-import" element={<AppLayout><BulkImport /></AppLayout>} />
        <Route path="/users" element={<AppLayout><Users /></AppLayout>} />

        {/* Redirect root to software overview */}
        <Route path="/" element={<Navigate to="/software" replace />} />
      </Routes>
    </Router>
  );
}

export default App;
