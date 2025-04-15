import React, { useState } from 'react';
import { Link, useNavigate, useSearchParams } from 'react-router-dom';
import Logo from '../../components/Logo';

const ResetPassword: React.FC = () => {
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const token = searchParams.get('token');

  const [formData, setFormData] = useState({
    password: '',
    confirmPassword: ''
  });
  const [error, setError] = useState('');
  const [passwordChecks, setPasswordChecks] = useState({
    length: false,
    capital: false,
    number: false,
    special: false
  });

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));

    if (name === 'password') {
      setPasswordChecks({
        length: value.length >= 12,
        capital: /[A-Z]/.test(value),
        number: /[0-9]/.test(value),
        special: /[!@#$%^&*]/.test(value)
      });
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    if (formData.password !== formData.confirmPassword) {
      setError('Passwords do not match');
      return;
    }

    if (!token) {
      setError('Invalid reset token');
      return;
    }

    try {
      const response = await fetch('/api/v1/auth/reset-password', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          token,
          password: formData.password
        }),
      });

      if (!response.ok) {
        throw new Error('Failed to reset password');
      }

      // Redirect to login page
      navigate('/login', { 
        state: { message: 'Password reset successful. Please login with your new password.' }
      });
    } catch (err) {
      setError('Failed to reset password. Please try again.');
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-base-200">
      <div className="max-w-xl w-full bg-base-100 shadow-xl rounded-lg p-8">
        <div className="mb-8">
          <Logo className="h-12 w-12 mb-4" />
          <h1 className="text-2xl font-bold mb-2">Reset your password</h1>
          <p className="text-gray-600">Fill in your new password below</p>
        </div>

        <form onSubmit={handleSubmit} className="space-y-4">
          {error && (
            <div className="alert alert-error">
              <span>{error}</span>
            </div>
          )}

          <div className="form-control">
            <input
              type="password"
              name="password"
              placeholder="Password"
              className="input input-bordered w-full"
              value={formData.password}
              onChange={handleInputChange}
              required
            />
            <div className="mt-2 space-y-1">
              <div className={`text-sm ${passwordChecks.length ? 'text-success' : 'text-base-content'}`}>
                ✓ At least 12 characters long
              </div>
              <div className={`text-sm ${passwordChecks.capital ? 'text-success' : 'text-base-content'}`}>
                ✓ At least 1 capital
              </div>
              <div className={`text-sm ${passwordChecks.number ? 'text-success' : 'text-base-content'}`}>
                ✓ At least 1 number
              </div>
              <div className={`text-sm ${passwordChecks.special ? 'text-success' : 'text-base-content'}`}>
                ✓ At least 1 special character (!@#$%^&*)
              </div>
            </div>
          </div>

          <div className="form-control">
            <input
              type="password"
              name="confirmPassword"
              placeholder="Confirm Password"
              className="input input-bordered w-full"
              value={formData.confirmPassword}
              onChange={handleInputChange}
              required
            />
          </div>

          <button type="submit" className="btn btn-primary w-full">
            Reset password
          </button>

          <Link to="/login" className="text-sm text-primary hover:underline block text-center">
            Back to Login page
          </Link>
        </form>
      </div>
    </div>
  );
};

export default ResetPassword; 