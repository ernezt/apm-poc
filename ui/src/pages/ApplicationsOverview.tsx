import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

interface Application {
  id: string;
  display_name: string;
  type: string;
  status: string;
  user_count: number;
  created_at: string;
  implementation_status: string;
}

const ApplicationsOverview = () => {
  const navigate = useNavigate();
  const [selectedApplications, setSelectedApplications] = useState<string[]>([]);
  const [applications, setApplications] = useState<Application[]>([
    // Sample data
    { id: 'A-10AD134', display_name: 'Application 1', type: 'COTS', status: 'In use', user_count: 5, created_at: '2025-01-01', implementation_status: 'In use' },
    { id: 'A-10AD135', display_name: 'Application 2', type: 'COTS', status: 'In use', user_count: 5, created_at: '2025-01-01', implementation_status: 'In use' },
    // Add more sample data as needed
  ]);
  const [rowsPerPage, setRowsPerPage] = useState(25);
  const [currentPage, setCurrentPage] = useState(1);

  const handleCheckboxChange = (id: string) => {
    setSelectedApplications(prev => 
      prev.includes(id) 
        ? prev.filter(appId => appId !== id)
        : [...prev, id]
    );
  };

  const handleSelectAll = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.checked) {
      setSelectedApplications(applications.map(app => app.id));
    } else {
      setSelectedApplications([]);
    }
  };

  const handleBulkAction = (action: string) => {
    // Implement bulk actions (export, update, delete)
    console.log('Bulk action:', action, 'on:', selectedApplications);
  };

  return (
    <div className="p-6">
      {/* Header */}
      <div className="flex justify-between items-center mb-6">
        <div>
          <h1 className="text-2xl font-bold mb-2">Applications</h1>
          <div className="stats shadow">
            <div className="stat">
              <div className="stat-title">Active applications</div>
              <div className="stat-value">431</div>
            </div>
          </div>
        </div>
        <div className="flex gap-2">
          <div className="dropdown">
            <button className="btn btn-primary">Add Application</button>
          </div>
          <button className="btn btn-outline" onClick={() => navigate('/applications/import')}>
            Import
          </button>
        </div>
      </div>

      {/* Actions Bar */}
      <div className="flex justify-between items-center mb-4">
        <div className="flex gap-2">
          <select 
            className="select select-bordered w-48"
            onChange={(e) => handleBulkAction(e.target.value)}
            value=""
          >
            <option value="" disabled>Select bulk action...</option>
            <option value="export">Export</option>
            <option value="update">Update</option>
            <option value="delete">Delete</option>
          </select>
          <button className="btn btn-primary btn-sm">Bulk action</button>
        </div>

        <div className="flex gap-2">
          <div className="dropdown dropdown-end">
            <label tabIndex={0} className="btn btn-ghost btn-sm">
              Column Options
              <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
              </svg>
            </label>
            <ul tabIndex={0} className="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-52">
              <li><a>Edit columns</a></li>
              <li><a>Filter</a></li>
              <li><a>Hide column</a></li>
            </ul>
          </div>
          <button className="btn btn-outline btn-sm">
            Advanced search
          </button>
        </div>
      </div>

      {/* Table */}
      <div className="bg-base-100 rounded-lg shadow">
        <table className="table w-full">
          <thead>
            <tr>
              <th>
                <input
                  type="checkbox"
                  className="checkbox"
                  onChange={handleSelectAll}
                  checked={selectedApplications.length === applications.length}
                />
              </th>
              <th>ID</th>
              <th>Display name</th>
              <th>Type</th>
              <th>Status</th>
              <th>User count</th>
              <th>Created</th>
            </tr>
          </thead>
          <tbody>
            {applications.map(app => (
              <tr 
                key={app.id}
                className="hover cursor-pointer"
                onClick={() => navigate(`/applications/${app.id}`)}
              >
                <td onClick={e => e.stopPropagation()}>
                  <input
                    type="checkbox"
                    className="checkbox"
                    checked={selectedApplications.includes(app.id)}
                    onChange={() => handleCheckboxChange(app.id)}
                  />
                </td>
                <td>{app.id}</td>
                <td>{app.display_name}</td>
                <td>
                  <span className="badge badge-ghost">{app.type}</span>
                </td>
                <td>
                  <span className={`badge ${
                    app.implementation_status === 'In use' ? 'badge-success' : 
                    app.implementation_status === 'Retired' ? 'badge-warning' :
                    'badge-ghost'
                  }`}>
                    {app.status}
                  </span>
                </td>
                <td>{app.user_count}</td>
                <td>{new Date(app.created_at).toLocaleDateString()}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {/* Pagination */}
      <div className="flex justify-between items-center mt-4">
        <div className="flex items-center gap-2">
          <span className="text-sm">Rows per page:</span>
          <select 
            className="select select-bordered select-sm w-20"
            value={rowsPerPage}
            onChange={(e) => setRowsPerPage(Number(e.target.value))}
          >
            <option value={25}>25</option>
            <option value={50}>50</option>
            <option value={100}>100</option>
          </select>
        </div>
        <div className="flex items-center gap-2">
          <span className="text-sm">1-25 of 521</span>
          <div className="join">
            <button className="join-item btn btn-sm" onClick={() => setCurrentPage(prev => Math.max(1, prev - 1))}>«</button>
            <button className="join-item btn btn-sm" onClick={() => setCurrentPage(prev => prev + 1)}>»</button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ApplicationsOverview; 