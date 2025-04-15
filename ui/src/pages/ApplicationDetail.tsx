import { useState } from 'react';
import { useParams, Link } from 'react-router-dom';

interface Stakeholder {
  name: string;
  role: string;
}

interface Application {
  id: string;
  display_name: string;
  description: string;
  implementation_status: string;
  lifecycle_status: string;
  type: string;
  available_since: string;
  rating: number;
  portfolio: string;
  functionalities: string[];
  groups: string[];
  install_type: string;
  manufacturer: string;
  vendor: string;
  documentation_url: string;
  stakeholders: Stakeholder[];
}

const ApplicationDetail = () => {
  const { id } = useParams<{ id: string }>();
  const [activeTab, setActiveTab] = useState('latest-news');
  const [application] = useState<Application>({
    id: 'A-10AD148',
    display_name: 'Application 15',
    description: 'Praeterea, ex culpa non invenies unum aut non accusatis unum. Et nihil molutam. Nemo nocere tibi erit, et non inimicos, et ne illa laederentur.',
    implementation_status: 'In use',
    lifecycle_status: 'Maintenance',
    type: 'COTS',
    available_since: '2025-01-02',
    rating: 4,
    portfolio: 'Microsoft B.V.',
    functionalities: ['Functionality 2', 'Functionality 1572'],
    groups: ['Functionality 1'],
    install_type: 'Software as a Service (SaaS)',
    manufacturer: 'Microsoft B.V.',
    vendor: 'One',
    documentation_url: 'https://example.com/docs',
    stakeholders: [
      { name: 'Henk de Vries', role: 'Owner' },
      { name: 'Jaap van Vliet', role: 'Contract manager' },
      { name: 'Pien Schelten', role: 'License manager' },
    ],
  });

  return (
    <div className="min-h-screen bg-base-200">
      {/* Breadcrumb and actions */}
      <div className="bg-base-100 border-b">
        <div className="container mx-auto px-6 py-4">
          <div className="flex justify-between items-center">
            <div className="flex items-center space-x-2 text-sm breadcrumbs">
              <ul>
                <li><Link to="/applications">Overview</Link></li>
                <li>{application.display_name}</li>
              </ul>
            </div>
            <div className="flex space-x-2">
              <button className="btn btn-outline btn-sm">
                <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" viewBox="0 0 20 20" fill="currentColor">
                  <path d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z" />
                </svg>
                Export
              </button>
              <Link to={`/applications/${id}/edit`} className="btn btn-outline btn-sm">
                <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" viewBox="0 0 20 20" fill="currentColor">
                  <path d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z" />
                </svg>
                Edit
              </Link>
              <button className="btn btn-outline btn-error btn-sm">
                <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" viewBox="0 0 20 20" fill="currentColor">
                  <path fillRule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clipRule="evenodd" />
                </svg>
                Remove
              </button>
            </div>
          </div>
        </div>
      </div>

      <div className="container mx-auto px-6 py-8">
        <div className="grid grid-cols-3 gap-6">
          {/* Main content */}
          <div className="col-span-2">
            {/* Header */}
            <div className="flex items-start gap-6 mb-8">
              <div className="avatar placeholder">
                <div className="bg-neutral text-neutral-content rounded-lg w-24">
                  <span className="text-3xl">{application.display_name.charAt(0)}</span>
                </div>
              </div>
              <div className="flex-1">
                <h1 className="text-3xl font-bold mb-2">{application.display_name}</h1>
                <div className="rating rating-sm mb-4">
                  {[1, 2, 3, 4, 5].map((star) => (
                    <input
                      key={star}
                      type="radio"
                      name="rating-2"
                      className="mask mask-star-2 bg-orange-400"
                      checked={star === application.rating}
                      readOnly
                    />
                  ))}
                </div>
                <div className="flex flex-wrap gap-2">
                  <span className="badge badge-lg badge-success">{application.implementation_status}</span>
                  <span className="badge badge-lg">{application.lifecycle_status}</span>
                  <span className="badge badge-lg badge-ghost">{application.type}</span>
                  <span className="badge badge-lg badge-outline">Available since {new Date(application.available_since).toLocaleDateString()}</span>
                </div>
              </div>
            </div>

            {/* Screenshots */}
            <div className="mb-8">
              <h2 className="text-xl font-semibold mb-4">Screenshots</h2>
              <div className="grid grid-cols-3 gap-4">
                {[1, 2, 3].map((n) => (
                  <div key={n} className="aspect-video bg-base-300 rounded-lg flex items-center justify-center">
                    <svg xmlns="http://www.w3.org/2000/svg" className="h-12 w-12 opacity-30" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                    </svg>
                  </div>
                ))}
              </div>
            </div>

            {/* Description */}
            <div className="mb-8">
              <h2 className="text-xl font-semibold mb-4">Description</h2>
              <div className="prose max-w-none">
                <p>{application.description}</p>
              </div>
            </div>

            {/* Tabs */}
            <div className="tabs tabs-bordered mb-6">
              <a 
                className={`tab ${activeTab === 'latest-news' ? 'tab-active' : ''}`}
                onClick={() => setActiveTab('latest-news')}
              >
                Latest news
              </a>
              <a 
                className={`tab ${activeTab === 'competitors' ? 'tab-active' : ''}`}
                onClick={() => setActiveTab('competitors')}
              >
                Competitors
              </a>
              <a 
                className={`tab ${activeTab === 'comments' ? 'tab-active' : ''}`}
                onClick={() => setActiveTab('comments')}
              >
                Comments
              </a>
              <a 
                className={`tab ${activeTab === 'history' ? 'tab-active' : ''}`}
                onClick={() => setActiveTab('history')}
              >
                History
              </a>
            </div>

            {/* Tab content */}
            <div className="bg-base-100 rounded-lg p-6">
              {activeTab === 'latest-news' && (
                <div className="grid grid-cols-3 gap-4">
                  {[1, 2, 3].map((n) => (
                    <div key={n} className="card bg-base-200">
                      <div className="card-body">
                        <h3 className="card-title text-sm">Lorem ipsum dolor sit</h3>
                        <div className="text-xs opacity-70">
                          Lorem ipsum dolor sit amet, consectetur adipiscing elit...
                        </div>
                        <div className="card-actions justify-end mt-2">
                          <div className="badge badge-ghost">Application name</div>
                          <div className="badge badge-ghost">Company name</div>
                        </div>
                      </div>
                    </div>
                  ))}
                </div>
              )}
            </div>
          </div>

          {/* Sidebar */}
          <div className="col-span-1 space-y-6">
            {/* Portfolio */}
            <div className="bg-base-100 rounded-lg p-6">
              <h3 className="font-semibold mb-4">Portfolio</h3>
              <p className="text-sm">{application.portfolio}</p>
            </div>

            {/* Functionalities */}
            <div className="bg-base-100 rounded-lg p-6">
              <h3 className="font-semibold mb-4">Functionalities</h3>
              <div className="flex flex-wrap gap-2">
                {application.functionalities.map((func, index) => (
                  <span key={index} className="badge badge-ghost">{func}</span>
                ))}
              </div>
            </div>

            {/* Groups */}
            <div className="bg-base-100 rounded-lg p-6">
              <h3 className="font-semibold mb-4">Groups</h3>
              <div className="flex flex-wrap gap-2">
                {application.groups.map((group, index) => (
                  <span key={index} className="badge">{group}</span>
                ))}
              </div>
            </div>

            {/* Install type */}
            <div className="bg-base-100 rounded-lg p-6">
              <h3 className="font-semibold mb-4">Install type</h3>
              <p className="text-sm">{application.install_type}</p>
            </div>

            {/* Manufacturer */}
            <div className="bg-base-100 rounded-lg p-6">
              <h3 className="font-semibold mb-4">Manufacturer</h3>
              <p className="text-sm">{application.manufacturer}</p>
            </div>

            {/* Product documentation */}
            <div className="bg-base-100 rounded-lg p-6">
              <h3 className="font-semibold mb-4">Product documentation</h3>
              <a href={application.documentation_url} className="link link-primary text-sm">
                Application 15 documentation
              </a>
            </div>

            {/* Vendor */}
            <div className="bg-base-100 rounded-lg p-6">
              <h3 className="font-semibold mb-4">Vendor</h3>
              <p className="text-sm">{application.vendor}</p>
            </div>

            {/* Stakeholders */}
            <div className="bg-base-100 rounded-lg p-6">
              <h3 className="font-semibold mb-4">Stakeholders</h3>
              <div className="space-y-4">
                {application.stakeholders.map((stakeholder, index) => (
                  <div key={index} className="flex justify-between items-center">
                    <span className="text-sm link-hover">{stakeholder.name}</span>
                    <span className="text-xs opacity-70">{stakeholder.role}</span>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ApplicationDetail; 