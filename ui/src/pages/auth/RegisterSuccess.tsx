import React from 'react';
import { useLocation } from 'react-router-dom';
import Logo from '../../components/Logo';

const RegisterSuccess: React.FC = () => {
  const location = useLocation();
  const firstname = location.state?.firstname || '';

  return (
    <div className="min-h-screen flex items-center justify-center bg-base-200">
      <div className="max-w-xl w-full bg-base-100 shadow-xl rounded-lg p-8 text-center">
        <div className="mb-8">
          <Logo className="h-12 w-12 mx-auto mb-4" />
        </div>
        
        <h1 className="text-2xl font-bold mb-4">Welcome {firstname}!</h1>
        <p className="text-gray-600 mb-4">
          You're account & workspace is currently being prepared. You will receive a confirmation and activation email momentarily.
        </p>
      </div>
    </div>
  );
};

export default RegisterSuccess; 