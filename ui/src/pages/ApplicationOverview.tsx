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
                {/* Add relevant actions for software */}
                <button>View Details</button>
                <button>Edit</button>
                <button>Delete</button> {/* Add Delete action */} 
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
    </div>
  );
};

export default ApplicationOverview; 