import React, { useState } from 'react';
import { Link, useLocation } from 'react-router-dom';
import { IconType } from 'react-icons';
import { 
  FiBox, 
  FiGrid, 
  FiBarChart2, 
  FiSettings, 
  FiUsers,
  FiChevronDown,
  FiChevronUp
} from 'react-icons/fi';
import Logo from './Logo';

interface NavItemProps {
  icon: IconType;
  label: string;
  to?: string;
  items?: { label: string; to: string }[];
}

const NavItem: React.FC<NavItemProps> = ({ icon: Icon, label, to, items }) => {
  const [isOpen, setIsOpen] = useState(false);
  const location = useLocation();
  const isActive = to ? location.pathname === to : 
    items?.some(item => location.pathname === item.to);

  if (items) {
    return (
      <div className="mb-2">
        <button
          onClick={() => setIsOpen(!isOpen)}
          className={`w-full flex items-center px-4 py-2 text-sm font-medium rounded-lg
            ${isActive ? 'text-primary bg-base-200' : 'text-base-content hover:bg-base-200'}`}
        >
          <Icon className="h-5 w-5 mr-3" />
          <span className="flex-1 text-left">{label}</span>
          {isOpen ? <FiChevronUp className="h-4 w-4" /> : <FiChevronDown className="h-4 w-4" />}
        </button>
        {isOpen && (
          <div className="ml-9 mt-2 space-y-1">
            {items.map((item, index) => (
              <Link
                key={index}
                to={item.to}
                className={`block px-4 py-2 text-sm rounded-lg
                  ${location.pathname === item.to 
                    ? 'text-primary bg-base-200' 
                    : 'text-base-content hover:bg-base-200'}`}
              >
                {item.label}
              </Link>
            ))}
          </div>
        )}
      </div>
    );
  }

  return (
    <Link
      to={to || '#'}
      className={`mb-2 flex items-center px-4 py-2 text-sm font-medium rounded-lg
        ${isActive ? 'text-primary bg-base-200' : 'text-base-content hover:bg-base-200'}`}
    >
      <Icon className="h-5 w-5 mr-3" />
      <span>{label}</span>
    </Link>
  );
};

const Navigation: React.FC = () => {
  return (
    <nav className="w-64 bg-base-100 border-r border-base-200 h-screen p-4">
      <div className="mb-8 w-48 mx-auto">
        <Logo className="h-12" />
      </div>

      <div className="space-y-4">
        <div>
          <h3 className="px-4 text-xs font-semibold text-base-content/70 uppercase tracking-wider mb-2">
            PORTFOLIO MANAGEMENT
          </h3>
          <NavItem
            icon={FiBox}
            label="Software"
            items={[
              { label: 'Overview', to: '/software' },
              { label: 'Add software', to: '/software/new' }
            ]}
          />
          <NavItem
            icon={FiGrid}
            label="Entities"
            items={[
              { label: 'Overview', to: '/entities' },
              { label: 'Add entity', to: '/entities/new' }
            ]}
          />
        </div>

        <div>
          <h3 className="px-4 text-xs font-semibold text-base-content/70 uppercase tracking-wider mb-2">
            INSIGHT
          </h3>
          <NavItem
            icon={FiBarChart2}
            label="Reports"
            items={[
              { label: 'Overview', to: '/reports' }
            ]}
          />
        </div>

        <div>
          <h3 className="px-4 text-xs font-semibold text-base-content/70 uppercase tracking-wider mb-2">
            SETTINGS
          </h3>
          <NavItem
            icon={FiSettings}
            label="Portfolio settings"
            items={[
              { label: 'General', to: '/settings/general' },
              { label: 'Custom fields', to: '/settings/custom-fields' },
              { label: 'Bulk import', to: '/settings/bulk-import' }
            ]}
          />
          <NavItem
            icon={FiUsers}
            label="User Management"
            items={[
              { label: 'Overview', to: '/users' }
            ]}
          />
        </div>
      </div>
    </nav>
  );
};

export default Navigation; 