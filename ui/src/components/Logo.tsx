import React from 'react';

interface LogoProps {
  className?: string;
}

const Logo: React.FC<LogoProps> = ({ className = "h-12 w-12" }) => {
  return (
    <div className={className}>
      <svg
        viewBox="0 0 100 100"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        className="w-full h-full"
      >
        {/* Background circle */}
        <circle cx="50" cy="50" r="45" fill="#4F46E5" />
        
        {/* APM letters - stylized and modern */}
        <text
          x="50"
          y="65"
          fontSize="36"
          fontWeight="bold"
          fill="white"
          textAnchor="middle"
          fontFamily="system-ui, -apple-system, sans-serif"
        >
          APM
        </text>
        
        {/* Decorative line */}
        <path
          d="M25 70 L75 70"
          stroke="white"
          strokeWidth="2"
          strokeLinecap="round"
          opacity="0.6"
        />
      </svg>
    </div>
  );
};

export default Logo; 