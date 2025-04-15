import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import Logo from '../../components/Logo';

const ForgotPassword: React.FC = () => {
  const [email, setEmail] = useState('');
  const [error, setError] = useState('');
  const [success, setSuccess] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setSuccess(false);

    try {
      const response = await fetch('/api/v1/auth/forgot-password', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email }),
      });

      if (!response.ok) {
        throw new Error('Failed to send reset email');
      }

      setSuccess(true);
    } catch (err) {
      setError('Failed to send reset email. Please try again.');
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-base-200">
      <div className="max-w-xl w-full bg-base-100 shadow-xl rounded-lg p-8">
        <div className="mb-8">
          <Logo className="h-12 w-12 mb-4" />
          <h1 className="text-2xl font-bold mb-2">Forgot your password?</h1>
          <p className="text-gray-600">Enter your email below to reset your password</p>
        </div>

        {success ? (
          <div className="alert alert-success">
            <span>Reset password link has been sent to your email.</span>
          </div>
        ) : (
          <form onSubmit={handleSubmit} className="space-y-4">
            {error && (
              <div className="alert alert-error">
                <span>{error}</span>
              </div>
            )}

            <div className="form-control">
              <input
                type="email"
                placeholder="Email address"
                className="input input-bordered w-full"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
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
        )}
      </div>
    </div>
  );
};

export default ForgotPassword; 