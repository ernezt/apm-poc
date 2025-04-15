import React from 'react';

interface LogoProps {
  className?: string;
}

const Logo: React.FC<LogoProps> = ({ className = "h-12" }) => {
  return (
    <div className={`flex justify-center ${className}`}>
      <img 
        src="/images/unicorn-logo.png" 
        alt="Unicorn Logo" 
        className="h-full w-auto object-contain"
        style={{ maxWidth: 'none' }}
      />
    </div>
  );
};

export default Logo; 