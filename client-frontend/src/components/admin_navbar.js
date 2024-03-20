import React from 'react';
import { Link } from 'react-router-dom';
import bus from '../bus.png'

const AdminNavbar = () => {
  return (
    <nav className="bg-gray-800 p-4">
      <div className="container mx-auto flex justify-between items-center">
        <div className='flex justify-around items-center'>
          <img src={bus} width="40px" />
          <h1 className="text-white ml-4 text-xl hover:text-gray-300" >BusBuddy</h1>
        </div>
        <div>
          <Link to="/admin/bus" className="text-white mr-4 hover:text-gray-300">Buses</Link>
          <Link to="/admin/route" className="text-white mr-4 hover:text-gray-300">Routes</Link>
          <Link to="/admin/schedule" className="text-white mr-4 hover:text-gray-300">Schedules</Link>
          <Link to="/admin/driver" className="text-white  mr-4 hover:text-gray-300">Drivers</Link>
          <Link to="/admin/station" className="text-white hover:text-gray-300">Station</Link>
        </div>
        {/* You can add admin profile/logout button here if needed */}
      </div>
    </nav>
  );
};

export default AdminNavbar;
