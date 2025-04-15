import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import Logo from '../../components/Logo';

interface RegisterForm {
  firstname: string;
  lastname: string;
  email: string;
  organisation: string;
  password: string;
  confirmPassword: string;
}

const Register: React.FC = () => {
  const navigate = useNavigate();
  const [formData, setFormData] = useState<RegisterForm>({
    firstname: '',
    lastname: '',
    email: '',
    organisation: '',
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

    try {
      const response = await fetch('/api/v1/auth/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          firstname: formData.firstname,
          lastname: formData.lastname,
          email: formData.email,
          organisation: formData.organisation,
          password: formData.password
        }),
      });

      if (!response.ok) {
        throw new Error('Registration failed');
      }

      // Navigate to success page
      navigate('/register/success', { 
        state: { firstname: formData.firstname }
      });
    } catch (err) {
      setError('Registration failed. Please try again.');
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-base-200">
      <div className="max-w-4xl w-full bg-base-100 shadow-xl rounded-lg flex">
        {/* Left side - Register form */}
        <div className="flex-1 p-8">
          <div className="mb-8">
            <Logo className="h-12 w-12 mb-4" />
            <h1 className="text-2xl font-bold mb-2">Register</h1>
          </div>

          {/* Social login buttons */}
          <div className="flex gap-4 mb-6">
            <button className="btn btn-outline flex-1">
              <svg className="w-5 h-5 mr-2" viewBox="0 0 24 24">
                <path fill="currentColor" d="M12.545,10.239v3.821h5.445c-0.712,2.315-2.647,3.972-5.445,3.972c-3.332,0-6.033-2.701-6.033-6.032s2.701-6.032,6.033-6.032c1.498,0,2.866,0.549,3.921,1.453l2.814-2.814C17.503,2.988,15.139,2,12.545,2C7.021,2,2.543,6.477,2.543,12s4.478,10,10.002,10c8.396,0,10.249-7.85,9.426-11.748L12.545,10.239z"/>
              </svg>
              Google
            </button>
            <button className="btn btn-outline flex-1">
              <svg className="w-5 h-5 mr-2" viewBox="0 0 24 24">
                <path fill="currentColor" d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
              </svg>
              Microsoft
            </button>
            <button className="btn btn-outline flex-1">
              <svg className="w-5 h-5 mr-2" viewBox="0 0 24 24">
                <path fill="currentColor" d="M22.675 0h-21.35c-.732 0-1.325.593-1.325 1.325v21.351c0 .731.593 1.324 1.325 1.324h11.495v-9.294h-3.128v-3.622h3.128v-2.671c0-3.1 1.893-4.788 4.659-4.788 1.325 0 2.463.099 2.795.143v3.24l-1.918.001c-1.504 0-1.795.715-1.795 1.763v2.313h3.587l-.467 3.622h-3.12v9.293h6.116c.73 0 1.323-.593 1.323-1.325v-21.35c0-.732-.593-1.325-1.325-1.325z"/>
              </svg>
              LinkedIn
            </button>
          </div>

          <div className="divider">Or</div>

          {/* Registration form */}
          <form onSubmit={handleSubmit} className="space-y-4">
            {error && (
              <div className="alert alert-error">
                <span>{error}</span>
              </div>
            )}

            <div className="grid grid-cols-2 gap-4">
              <div className="form-control">
                <input
                  type="text"
                  name="lastname"
                  placeholder="Lastname"
                  className="input input-bordered w-full"
                  value={formData.lastname}
                  onChange={handleInputChange}
                  required
                />
              </div>
              <div className="form-control">
                <input
                  type="text"
                  name="firstname"
                  placeholder="Firstname"
                  className="input input-bordered w-full"
                  value={formData.firstname}
                  onChange={handleInputChange}
                  required
                />
              </div>
            </div>

            <div className="form-control">
              <input
                type="email"
                name="email"
                placeholder="Email address"
                className="input input-bordered w-full"
                value={formData.email}
                onChange={handleInputChange}
                required
              />
            </div>

            <div className="form-control">
              <input
                type="text"
                name="organisation"
                placeholder="Organisation"
                className="input input-bordered w-full"
                value={formData.organisation}
                onChange={handleInputChange}
                required
              />
            </div>

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
              Register
            </button>
          </form>
        </div>

        {/* Right side - Login CTA */}
        <div className="flex-1 bg-base-200 p-8 flex flex-col items-center justify-center">
          <h2 className="text-2xl font-bold mb-4">Already have an account?</h2>
          <p className="text-center mb-8">
            Then go back to the previous page and log in
          </p>
          <Link to="/login" className="btn btn-outline">
            Login
          </Link>
        </div>
      </div>
    </div>
  );
};

export default Register; 