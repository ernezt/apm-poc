import React from 'react';
import Navigation from './Navigation';

interface AppLayoutProps {
  children: React.ReactNode;
}

const AppLayout: React.FC<AppLayoutProps> = ({ children }) => {
  return (
    <div className="flex min-h-screen">
      <Navigation />
      <main className="flex-1 p-8 bg-base-200">
        {children}
      </main>
    </div>
  );
};

export default AppLayout; 