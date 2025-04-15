import React, { useState, useEffect } from 'react';
import './ApplicationOverview.css'; // We will create this CSS file next

// Define the structure of a Software object based on models.SoftwareResponse
interface Software {
  id: string;
  foreign_key?: string;
  display_name: string;
  description?: string;
  software_type: string; // Use string for simplicity, can refine with enum later
  software_subtype?: string;
  vendor?: string;
  manufacturer?: string;
  install_type?: string;
  product_type?: string;
  context?: string;
  lifecycle_status?: string;
  implementation_status?: string;
  created_at: string; // Assuming ISO string format from backend
  updated_at: string;
  rating?: number; // Added for star rating
}

// Define structure for the new software form data based on models.CreateSoftwareRequest
interface NewSoftwareData {
  display_name: string;
  description: string;
  software_type: string; // Default to 'web' or another suitable default
  // Add other optional fields as needed
  vendor: string;
  manufacturer: string;
  // ... add other fields from CreateSoftwareRequest if needed in the form
}

const ApplicationOverview: React.FC = () => {
  // State for the list of software
  const [softwareList, setSoftwareList] = useState<Software[]>([]);
  // State to control modal visibility
  const [isModalOpen, setIsModalOpen] = useState(false);
  // State for the new software form data
  const [newSoftwareData, setNewSoftwareData] = useState<NewSoftwareData>({
    display_name: '',
    description: '',
    software_type: 'web', // Default type
    vendor: '',
    manufacturer: '',
  });
  const [isLoading, setIsLoading] = useState(false); // Add loading state
  const [error, setError] = useState<string | null>(null); // Add error state
  const [selectedSoftware, setSelectedSoftware] = useState<Software | null>(null);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [editSoftwareData, setEditSoftwareData] = useState<NewSoftwareData | null>(null);

  // Fetch software from the backend when the component mounts
  useEffect(() => {
    fetchSoftware();
  }, []);

  const fetchSoftware = async () => {
    setIsLoading(true);
    setError(null);
    console.log("Fetching software...");
    try {
      const response = await fetch('/api/v1/software'); // Use proxied path
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      // Ensure data.data is an array before setting state
      setSoftwareList(Array.isArray(data.data) ? data.data : []); 
    } catch (e) {
      console.error("Failed to fetch software:", e);
      setError(e instanceof Error ? e.message : 'An unknown error occurred');
      // TODO: Display error more prominently to the user
    } finally {
      setIsLoading(false);
    }
  };

  // Handler to open the modal and reset form
  const handleOpenModal = () => {
    setNewSoftwareData({ display_name: '', description: '', software_type: 'web', vendor: '', manufacturer: '' });
    setIsModalOpen(true);
  };

  // Handler to close the modal
  const handleCloseModal = () => {
    setIsModalOpen(false);
  };

  // Handler for input changes in the form
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setNewSoftwareData(prevData => ({
      ...prevData,
      [name]: value,
    }));
  };

  // Handler for form submission
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true); // Set loading true during submission
    setError(null);
    console.log('Submitting:', newSoftwareData);
    try {
      const response = await fetch('/api/v1/software', { // Use proxied path
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          // TODO: Add Authorization header if required
          // 'Authorization': `Bearer ${your_jwt_token}`
        },
        body: JSON.stringify(newSoftwareData),
      });

      if (!response.ok) {
         // Attempt to read error details from backend response
        const errorData = await response.json().catch(() => null); // Avoid crashing if body is not JSON
        console.error('Backend error:', errorData);
        throw new Error(`HTTP error! status: ${response.status} - ${errorData?.error || 'Failed to create software'}`);
      }

      const createdSoftware: Software = await response.json();
      // Add the new software to the list
      setSoftwareList(prevList => [...prevList, createdSoftware]);
      handleCloseModal();
    } catch (e) {
      console.error("Failed to add software:", e);
      setError(e instanceof Error ? e.message : 'An unknown error occurred');
      // Keep modal open on error so user can retry or fix data
      // TODO: Show specific error message within the modal
    } finally {
      setIsLoading(false); // Set loading false after submission attempt
    }
  };

  // Handler for viewing software details
  const handleViewDetails = (software: Software) => {
    setSelectedSoftware(software);
    setIsModalOpen(true);
  };

  // Handler for editing software
  const handleEdit = (software: Software) => {
    setSelectedSoftware(software);
    setEditSoftwareData({
      display_name: software.display_name,
      description: software.description || '',
      software_type: software.software_type,
      vendor: software.vendor || '',
      manufacturer: software.manufacturer || '',
    });
    setIsEditModalOpen(true);
  };

  // Handler for deleting software
  const handleDelete = async (softwareId: string) => {
    if (!window.confirm('Are you sure you want to delete this software?')) {
      return;
    }

    setIsLoading(true);
    try {
      const response = await fetch(`/api/v1/software/${softwareId}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      // Remove the deleted software from the list
      setSoftwareList(prevList => prevList.filter(sw => sw.id !== softwareId));
    } catch (e) {
      console.error("Failed to delete software:", e);
      setError(e instanceof Error ? e.message : 'An unknown error occurred');
    } finally {
      setIsLoading(false);
    }
  };

  // Handler for updating software
  const handleUpdate = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!selectedSoftware || !editSoftwareData) return;

    setIsLoading(true);
    try {
      const response = await fetch(`/api/v1/software/${selectedSoftware.id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(editSoftwareData),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      // For 204 No Content, we don't need to parse JSON
      if (response.status === 204) {
        // Update the local state with the edited data
        setSoftwareList(prevList => 
          prevList.map(sw => sw.id === selectedSoftware.id ? {
            ...sw,
            ...editSoftwareData,
            updated_at: new Date().toISOString()
          } : sw)
        );
        setIsEditModalOpen(false);
        setSelectedSoftware(null);
        setEditSoftwareData(null);
        return;
      }

      // If we get here, there was an unexpected response
      throw new Error('Unexpected response from server');
    } catch (e) {
      console.error("Failed to update software:", e);
      setError(e instanceof Error ? e.message : 'An unknown error occurred');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-base-200">
      <div className="container mx-auto p-4">
        {/* Header with breadcrumb */}
        <div className="flex items-center justify-between mb-6">
          <div className="flex items-center space-x-2 text-sm breadcrumbs">
            <ul>
              <li><a href="/">Overview</a></li>
              <li>Application {selectedSoftware?.display_name || ''}</li>
            </ul>
          </div>
          <div className="flex space-x-2">
            <button className="btn btn-outline btn-sm">
              <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" viewBox="0 0 20 20" fill="currentColor">
                <path d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z" />
              </svg>
              Export
            </button>
            <button className="btn btn-outline btn-sm" onClick={() => handleEdit(selectedSoftware!)}>
              <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" viewBox="0 0 20 20" fill="currentColor">
                <path d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z" />
              </svg>
              Edit
            </button>
            <button className="btn btn-outline btn-error btn-sm" onClick={() => handleDelete(selectedSoftware!.id)}>
              <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" viewBox="0 0 20 20" fill="currentColor">
                <path fillRule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clipRule="evenodd" />
              </svg>
              Remove
            </button>
          </div>
        </div>

        {/* Main content grid */}
        <div className="grid grid-cols-3 gap-6">
          {/* Main content area - 2 columns */}
          <div className="col-span-2">
            {/* Software list with better styling */}
            <div className="bg-base-100 rounded-lg shadow-lg">
              <table className="table w-full">
                <thead>
                  <tr>
                    <th>Display Name</th>
                    <th>Type</th>
                    <th>Status</th>
                    <th>Last Updated</th>
                    <th>Actions</th>
                  </tr>
                </thead>
                <tbody>
                  {isLoading && !softwareList.length && (
                    <tr>
                      <td colSpan={5} className="text-center">
                        <span className="loading loading-spinner loading-md"></span>
                      </td>
                    </tr>
                  )}
                  {!isLoading && error && !softwareList.length && (
                    <tr>
                      <td colSpan={5} className="text-center text-error">{error}</td>
                    </tr>
                  )}
                  {!isLoading && !error && softwareList.length === 0 && (
                    <tr>
                      <td colSpan={5} className="text-center">No software found.</td>
                    </tr>
                  )}
                  {softwareList.map((sw: Software) => (
                    <tr key={sw.id} className="hover">
                      <td>
                        <div className="flex items-center space-x-3">
                          <div className="avatar placeholder">
                            <div className="bg-neutral text-neutral-content rounded-lg w-12">
                              <span className="text-xl">{sw.display_name.charAt(0)}</span>
                            </div>
                          </div>
                          <div>
                            <div className="font-bold">{sw.display_name}</div>
                            <div className="text-sm opacity-50">{sw.vendor || 'No vendor'}</div>
                          </div>
                        </div>
                      </td>
                      <td>
                        <span className="badge badge-ghost">{sw.software_type}</span>
                      </td>
                      <td>
                        <div className="flex flex-col">
                          <span className={`badge ${sw.implementation_status === 'IN_USE' ? 'badge-success' : 'badge-warning'}`}>
                            {sw.implementation_status || 'Unknown'}
                          </span>
                          <span className="text-xs opacity-50">{sw.lifecycle_status}</span>
                        </div>
                      </td>
                      <td>
                        <div className="flex flex-col">
                          <span>{new Date(sw.updated_at).toLocaleDateString()}</span>
                          <span className="text-xs opacity-50">
                            {new Date(sw.updated_at).toLocaleTimeString()}
                          </span>
                        </div>
                      </td>
                      <td>
                        <div className="flex space-x-1">
                          <button 
                            className="btn btn-ghost btn-xs"
                            onClick={() => handleViewDetails(sw)}
                          >
                            View
                          </button>
                          <button 
                            className="btn btn-ghost btn-xs"
                            onClick={() => handleEdit(sw)}
                          >
                            Edit
                          </button>
                          <button 
                            className="btn btn-ghost btn-error btn-xs"
                            onClick={() => handleDelete(sw.id)}
                          >
                            Delete
                          </button>
                        </div>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>

          {/* Right sidebar - 1 column */}
          <div className="col-span-1">
            <div className="bg-base-100 rounded-lg shadow-lg p-4 space-y-4">
              <h3 className="text-lg font-semibold">Quick Stats</h3>
              <div className="stats stats-vertical shadow">
                <div className="stat">
                  <div className="stat-title">Total Applications</div>
                  <div className="stat-value">{softwareList.length}</div>
                </div>
                <div className="stat">
                  <div className="stat-title">Active</div>
                  <div className="stat-value text-success">
                    {softwareList.filter(sw => sw.implementation_status === 'IN_USE').length}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        {/* View Details Modal with improved styling */}
        {isModalOpen && selectedSoftware && !isEditModalOpen && (
          <div className="modal modal-open">
            <div className="modal-box w-11/12 max-w-5xl">
              <div className="flex justify-between items-start mb-6">
                <div>
                  <h2 className="text-2xl font-bold">{selectedSoftware.display_name}</h2>
                  <div className="rating rating-sm">
                    {[1, 2, 3, 4, 5].map((star) => (
                      <input
                        key={star}
                        type="radio"
                        name="rating-2"
                        className="mask mask-star-2 bg-orange-400"
                        checked={star === (selectedSoftware.rating || 0)}
                        readOnly
                      />
                    ))}
                  </div>
                </div>
                <button 
                  className="btn btn-sm btn-circle btn-ghost"
                  onClick={() => {
                    setIsModalOpen(false);
                    setSelectedSoftware(null);
                  }}
                >
                  ✕
                </button>
              </div>

              <div className="grid grid-cols-2 gap-6">
                <div className="space-y-4">
                  <div className="flex flex-wrap gap-2 mb-4">
                    <div className="badge badge-lg">{selectedSoftware.implementation_status}</div>
                    <div className="badge badge-lg badge-outline">{selectedSoftware.lifecycle_status}</div>
                    <div className="badge badge-lg badge-ghost">{selectedSoftware.software_type}</div>
                  </div>

                  <div className="prose">
                    <h3>Description</h3>
                    <p>{selectedSoftware.description || 'No description available.'}</p>
                  </div>

                  <div className="divider"></div>

                  <div className="grid grid-cols-2 gap-4">
                    <div>
                      <h4 className="font-semibold mb-2">Details</h4>
                      <ul className="space-y-2">
                        <li><span className="opacity-70">Vendor:</span> {selectedSoftware.vendor || 'N/A'}</li>
                        <li><span className="opacity-70">Manufacturer:</span> {selectedSoftware.manufacturer || 'N/A'}</li>
                        <li><span className="opacity-70">Install Type:</span> {selectedSoftware.install_type || 'N/A'}</li>
                        <li><span className="opacity-70">Product Type:</span> {selectedSoftware.product_type || 'N/A'}</li>
                      </ul>
                    </div>
                    <div>
                      <h4 className="font-semibold mb-2">Dates</h4>
                      <ul className="space-y-2">
                        <li>
                          <span className="opacity-70">Created:</span>
                          {new Date(selectedSoftware.created_at).toLocaleDateString()}
                        </li>
                        <li>
                          <span className="opacity-70">Last Updated:</span>
                          {new Date(selectedSoftware.updated_at).toLocaleDateString()}
                        </li>
                      </ul>
                    </div>
                  </div>
                </div>

                <div className="border-l pl-6">
                  <h3 className="font-semibold mb-4">Screenshots</h3>
                  <div className="grid grid-cols-2 gap-4">
                    {[1, 2, 3, 4].map((n) => (
                      <div key={n} className="aspect-video bg-base-200 rounded-lg flex items-center justify-center">
                        <svg xmlns="http://www.w3.org/2000/svg" className="h-12 w-12 opacity-30" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                        </svg>
                      </div>
                    ))}
                  </div>
                </div>
              </div>
            </div>
          </div>
        )}

        {/* Edit Modal with improved styling */}
        {isEditModalOpen && selectedSoftware && editSoftwareData && (
          <div className="modal modal-open">
            <div className="modal-box w-11/12 max-w-5xl">
              <h2 className="text-2xl font-bold mb-6">Edit Software</h2>
              {error && <div className="alert alert-error mb-4">{error}</div>}
              <form onSubmit={handleUpdate} className="space-y-4">
                <div className="grid grid-cols-2 gap-6">
                  <div className="form-control">
                    <label className="label">
                      <span className="label-text">Display Name *</span>
                    </label>
                    <input
                      type="text"
                      className="input input-bordered"
                      value={editSoftwareData.display_name}
                      onChange={(e) => setEditSoftwareData({...editSoftwareData, display_name: e.target.value})}
                      required
                    />
                  </div>

                  <div className="form-control">
                    <label className="label">
                      <span className="label-text">Software Type *</span>
                    </label>
                    <select
                      className="select select-bordered"
                      value={editSoftwareData.software_type}
                      onChange={(e) => setEditSoftwareData({...editSoftwareData, software_type: e.target.value})}
                      required
                    >
                      <option value="api">API</option>
                      <option value="web">Web</option>
                      <option value="mobile">Mobile</option>
                      <option value="desktop">Desktop</option>
                      <option value="embedded">Embedded</option>
                      <option value="middleware">Middleware</option>
                      <option value="library">Library</option>
                    </select>
                  </div>

                  <div className="form-control col-span-2">
                    <label className="label">
                      <span className="label-text">Description</span>
                    </label>
                    <textarea
                      className="textarea textarea-bordered h-24"
                      value={editSoftwareData.description}
                      onChange={(e) => setEditSoftwareData({...editSoftwareData, description: e.target.value})}
                    />
                  </div>

                  <div className="form-control">
                    <label className="label">
                      <span className="label-text">Vendor</span>
                    </label>
                    <input
                      type="text"
                      className="input input-bordered"
                      value={editSoftwareData.vendor}
                      onChange={(e) => setEditSoftwareData({...editSoftwareData, vendor: e.target.value})}
                    />
                  </div>

                  <div className="form-control">
                    <label className="label">
                      <span className="label-text">Manufacturer</span>
                    </label>
                    <input
                      type="text"
                      className="input input-bordered"
                      value={editSoftwareData.manufacturer}
                      onChange={(e) => setEditSoftwareData({...editSoftwareData, manufacturer: e.target.value})}
                    />
                  </div>
                </div>

                <div className="modal-action">
                  <button type="submit" className="btn btn-primary" disabled={isLoading}>
                    {isLoading ? <span className="loading loading-spinner"></span> : 'Update Software'}
                  </button>
                  <button 
                    type="button"
                    className="btn btn-ghost"
                    onClick={() => {
                      setIsEditModalOpen(false);
                      setSelectedSoftware(null);
                      setEditSoftwareData(null);
                    }}
                    disabled={isLoading}
                  >
                    Cancel
                  </button>
                </div>
              </form>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default ApplicationOverview; 