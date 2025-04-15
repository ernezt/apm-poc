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
    <div className="application-overview-container">
       {/* Optional: Display loading indicator */} 
       {isLoading && <div className="loading-indicator">Loading...</div>}
       {/* Optional: Display error message */} 
       {error && <div className="error-message">Error: {error}</div>} 

      <div className="header">
        <h1>Software Portfolio</h1>
        <button className="add-button" onClick={handleOpenModal} disabled={isLoading}>Add New Software</button>
      </div>
      <table>
        <thead>
          <tr>
            {/* Updated table headers based on Software model */}
            <th>Display Name</th>
            <th>Description</th>
            <th>Type</th>
            <th>Vendor</th>
            <th>Created At</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {/* Display loading/error directly in the table body if preferred */} 
          {isLoading && !softwareList.length && (
             <tr><td colSpan={6} style={{ textAlign: 'center' }}>Loading software...</td></tr>
          )}
          {!isLoading && error && !softwareList.length && (
            <tr><td colSpan={6} style={{ textAlign: 'center' }} className="error-message">Failed to load software: {error}</td></tr>
          )}
          {!isLoading && !error && softwareList.length === 0 && (
            <tr><td colSpan={6} style={{ textAlign: 'center' }}>No software found.</td></tr>
          )}
          {softwareList.map((sw: Software) => (
            <tr key={sw.id}>
              <td>{sw.display_name}</td>
              {/* Truncate description for display? */}
              <td>{sw.description?.substring(0, 50)}{sw.description && sw.description.length > 50 ? '...' : ''}</td>
              <td>{sw.software_type}</td>
              <td>{sw.vendor}</td>
              <td>{new Date(sw.created_at).toLocaleDateString()}</td>
              <td>
                <button onClick={() => handleViewDetails(sw)}>View Details</button>
                <button onClick={() => handleEdit(sw)}>Edit</button>
                <button onClick={() => handleDelete(sw.id)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>

      {/* Modal for Adding Software */}
      {isModalOpen && (
        <div className="modal-backdrop">
          <div className="modal-content">
            <h2>Add New Software</h2>
             {/* Display error inside modal */} 
             {error && <div className="error-message" style={{ marginBottom: '15px' }}>Error: {error}</div>} 
            <form onSubmit={handleSubmit}>
              {/* Updated form fields for CreateSoftwareRequest */} 
              <div className="form-group">
                <label htmlFor="display_name">Display Name *</label>
                <input
                  type="text"
                  id="display_name"
                  name="display_name"
                  value={newSoftwareData.display_name}
                  onChange={handleInputChange}
                  required
                />
              </div>
              <div className="form-group">
                <label htmlFor="description">Description</label>
                <textarea
                  id="description"
                  name="description"
                  value={newSoftwareData.description}
                  onChange={handleInputChange} // Need to cast e.target type correctly if using single handler
                  rows={3}
                />
              </div>
              <div className="form-group">
                <label htmlFor="software_type">Software Type *</label>
                <select
                  id="software_type"
                  name="software_type"
                  value={newSoftwareData.software_type}
                  onChange={handleInputChange}
                  required
                >
                  {/* Options based on SoftwareType enum */}
                  <option value="api">API</option>
                  <option value="web">Web</option>
                  <option value="mobile">Mobile</option>
                  <option value="desktop">Desktop</option>
                  <option value="embedded">Embedded</option>
                  <option value="middleware">Middleware</option>
                  <option value="library">Library</option>
                </select>
              </div>
               <div className="form-group">
                <label htmlFor="vendor">Vendor</label>
                <input
                  type="text"
                  id="vendor"
                  name="vendor"
                  value={newSoftwareData.vendor}
                  onChange={handleInputChange}
                />
              </div>
               <div className="form-group">
                <label htmlFor="manufacturer">Manufacturer</label>
                <input
                  type="text"
                  id="manufacturer"
                  name="manufacturer"
                  value={newSoftwareData.manufacturer}
                  onChange={handleInputChange}
                />
              </div>
              {/* Add more form groups for other fields if needed */}
              <div className="modal-actions">
                 {/* Disable button while loading */}
                <button type="submit" className="submit-button" disabled={isLoading}> {isLoading ? 'Adding...' : 'Add Software'}</button>
                <button type="button" onClick={handleCloseModal} className="cancel-button" disabled={isLoading}>Cancel</button>
              </div>
            </form>
          </div>
        </div>
      )}

      {/* Add Edit Modal */}
      {isEditModalOpen && selectedSoftware && editSoftwareData && (
        <div className="modal-backdrop">
          <div className="modal-content">
            <h2>Edit Software</h2>
            {error && <div className="error-message" style={{ marginBottom: '15px' }}>Error: {error}</div>}
            <form onSubmit={handleUpdate}>
              <div className="form-group">
                <label htmlFor="edit_display_name">Display Name *</label>
                <input
                  type="text"
                  id="edit_display_name"
                  name="display_name"
                  value={editSoftwareData.display_name}
                  onChange={(e) => setEditSoftwareData({...editSoftwareData, display_name: e.target.value})}
                  required
                />
              </div>
              <div className="form-group">
                <label htmlFor="edit_description">Description</label>
                <textarea
                  id="edit_description"
                  name="description"
                  value={editSoftwareData.description}
                  onChange={(e) => setEditSoftwareData({...editSoftwareData, description: e.target.value})}
                  rows={3}
                />
              </div>
              <div className="form-group">
                <label htmlFor="edit_software_type">Software Type *</label>
                <select
                  id="edit_software_type"
                  name="software_type"
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
              <div className="form-group">
                <label htmlFor="edit_vendor">Vendor</label>
                <input
                  type="text"
                  id="edit_vendor"
                  name="vendor"
                  value={editSoftwareData.vendor}
                  onChange={(e) => setEditSoftwareData({...editSoftwareData, vendor: e.target.value})}
                />
              </div>
              <div className="form-group">
                <label htmlFor="edit_manufacturer">Manufacturer</label>
                <input
                  type="text"
                  id="edit_manufacturer"
                  name="manufacturer"
                  value={editSoftwareData.manufacturer}
                  onChange={(e) => setEditSoftwareData({...editSoftwareData, manufacturer: e.target.value})}
                />
              </div>
              <div className="modal-actions">
                <button type="submit" className="submit-button" disabled={isLoading}>
                  {isLoading ? 'Updating...' : 'Update Software'}
                </button>
                <button 
                  type="button" 
                  onClick={() => {
                    setIsEditModalOpen(false);
                    setSelectedSoftware(null);
                    setEditSoftwareData(null);
                  }} 
                  className="cancel-button" 
                  disabled={isLoading}
                >
                  Cancel
                </button>
              </div>
            </form>
          </div>
        </div>
      )}

      {/* Add View Details Modal */}
      {isModalOpen && selectedSoftware && !isEditModalOpen && (
        <div className="modal-backdrop">
          <div className="modal-content">
            <h2>Software Details</h2>
            <div className="software-details">
              <p><strong>Display Name:</strong> {selectedSoftware.display_name}</p>
              <p><strong>Description:</strong> {selectedSoftware.description || 'N/A'}</p>
              <p><strong>Type:</strong> {selectedSoftware.software_type}</p>
              <p><strong>Subtype:</strong> {selectedSoftware.software_subtype || 'N/A'}</p>
              <p><strong>Vendor:</strong> {selectedSoftware.vendor || 'N/A'}</p>
              <p><strong>Manufacturer:</strong> {selectedSoftware.manufacturer || 'N/A'}</p>
              <p><strong>Install Type:</strong> {selectedSoftware.install_type || 'N/A'}</p>
              <p><strong>Product Type:</strong> {selectedSoftware.product_type || 'N/A'}</p>
              <p><strong>Context:</strong> {selectedSoftware.context || 'N/A'}</p>
              <p><strong>Lifecycle Status:</strong> {selectedSoftware.lifecycle_status || 'N/A'}</p>
              <p><strong>Implementation Status:</strong> {selectedSoftware.implementation_status || 'N/A'}</p>
              <p><strong>Created At:</strong> {new Date(selectedSoftware.created_at).toLocaleString()}</p>
              <p><strong>Updated At:</strong> {new Date(selectedSoftware.updated_at).toLocaleString()}</p>
            </div>
            <div className="modal-actions">
              <button 
                type="button" 
                onClick={() => {
                  setIsModalOpen(false);
                  setSelectedSoftware(null);
                }} 
                className="cancel-button"
              >
                Close
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default ApplicationOverview; 