import React from 'react';
import { useParams, Link } from 'react-router-dom';
import { ChevronRightIcon, HomeIcon } from '@heroicons/react/24/outline';

const ApplicationEdit: React.FC = () => {
  const { id } = useParams();

  return (
    <div className="container mx-auto px-4 py-8">
      {/* Breadcrumb */}
      <nav className="flex mb-8" aria-label="Breadcrumb">
        <ol className="inline-flex items-center space-x-1 md:space-x-3">
          <li className="inline-flex items-center">
            <Link to="/" className="inline-flex items-center text-sm font-medium text-gray-700 hover:text-blue-600">
              <HomeIcon className="w-4 h-4 mr-2" />
              Home
            </Link>
          </li>
          <li>
            <div className="flex items-center">
              <ChevronRightIcon className="w-4 h-4 text-gray-400" />
              <Link to="/applications" className="ml-1 text-sm font-medium text-gray-700 hover:text-blue-600 md:ml-2">Applications</Link>
            </div>
          </li>
          <li>
            <div className="flex items-center">
              <ChevronRightIcon className="w-4 h-4 text-gray-400" />
              <Link to={`/applications/${id}`} className="ml-1 text-sm font-medium text-gray-700 hover:text-blue-600 md:ml-2">Application Details</Link>
            </div>
          </li>
          <li aria-current="page">
            <div className="flex items-center">
              <ChevronRightIcon className="w-4 h-4 text-gray-400" />
              <span className="ml-1 text-sm font-medium text-gray-500 md:ml-2">Edit</span>
            </div>
          </li>
        </ol>
      </nav>

      <div className="flex gap-8">
        {/* Main Content */}
        <div className="flex-grow">
          {/* General Information */}
          <div className="bg-white rounded-lg shadow-sm p-6 mb-6">
            <h2 className="text-xl font-semibold mb-4">General Information</h2>
            <div className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Application Name</label>
                <input type="text" className="input input-bordered w-full" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Application Type</label>
                <select className="select select-bordered w-full">
                  <option>Web Application</option>
                  <option>Mobile Application</option>
                  <option>Desktop Application</option>
                </select>
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Status</label>
                <select className="select select-bordered w-full">
                  <option>Active</option>
                  <option>Inactive</option>
                  <option>Maintenance</option>
                </select>
              </div>
            </div>
          </div>

          {/* Description */}
          <div className="bg-white rounded-lg shadow-sm p-6 mb-6">
            <h2 className="text-xl font-semibold mb-4">Description</h2>
            <textarea className="textarea textarea-bordered w-full h-32" placeholder="Enter application description..."></textarea>
          </div>

          {/* Screenshots */}
          <div className="bg-white rounded-lg shadow-sm p-6 mb-6">
            <h2 className="text-xl font-semibold mb-4">Screenshots</h2>
            <div className="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center">
              <p className="text-gray-600">Drag and drop screenshots here or click to upload</p>
              <button className="btn btn-primary mt-4">Upload Screenshots</button>
            </div>
          </div>

          {/* Stakeholders */}
          <div className="bg-white rounded-lg shadow-sm p-6">
            <h2 className="text-xl font-semibold mb-4">Stakeholders</h2>
            <div className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Product Owner</label>
                <input type="text" className="input input-bordered w-full" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Technical Lead</label>
                <input type="text" className="input input-bordered w-full" />
              </div>
            </div>
          </div>
        </div>

        {/* Sidebar */}
        <div className="w-80">
          {/* Actions */}
          <div className="bg-white rounded-lg shadow-sm p-6 mb-6">
            <h2 className="text-xl font-semibold mb-4">Actions</h2>
            <div className="space-y-3">
              <button className="btn btn-primary w-full">Save Changes</button>
              <Link to={`/applications/${id}`} className="btn btn-outline w-full">Cancel</Link>
            </div>
          </div>

          {/* Portal Settings */}
          <div className="bg-white rounded-lg shadow-sm p-6">
            <h2 className="text-xl font-semibold mb-4">Portal Settings</h2>
            <div className="space-y-4">
              <div className="form-control">
                <label className="label cursor-pointer justify-start gap-3">
                  <input type="checkbox" className="checkbox" />
                  <span className="label-text">Show in Portal</span>
                </label>
              </div>
              <div className="form-control">
                <label className="label cursor-pointer justify-start gap-3">
                  <input type="checkbox" className="checkbox" />
                  <span className="label-text">Featured Application</span>
                </label>
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Portal Category</label>
                <select className="select select-bordered w-full">
                  <option>Business</option>
                  <option>Development</option>
                  <option>Infrastructure</option>
                </select>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ApplicationEdit; 