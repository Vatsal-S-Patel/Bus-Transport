import React, { useEffect, useState, useRef } from "react";
import "ol/ol.css";
import "./App.css";
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import AdminLogin from './components/admin_login';
import AdminNavbar from './components/admin_navbar';
import BusHandle from './components/bus_handle';
import BusRouteHandle from './components/bus_route_handle';
import ScheduleHandle from './components/schedule_handle';
import DriverHandle from './components/driver_handle';
import StationHandle from './components/station_handle';
import BusHome from './components/dummy_bus_runner';
import ClientMap from "./components/ClientMap";
import BusLive from "./components/live_location_bus";


const MapApp = () => {
  return (
    <div id="app">
    <Router>
      <Routes>
        {/* All Routes for Admin */}
        <Route exact path="/" element={<ClientMap/>} />
        <Route exact path="/admin/login" element={<AdminLogin/>} />
        <Route exact path="/admin/bus" element={<><AdminNavbar/><BusHandle/></>} />
        <Route exact path="/admin/schedule" element={<><AdminNavbar/><ScheduleHandle/></>} />
        <Route exact path="/admin/route" element={<><AdminNavbar/><BusRouteHandle/></>} />
        <Route exact path="/admin/driver" element={<><AdminNavbar/><DriverHandle/></>} />
        <Route exact path="/admin/station" element={<><AdminNavbar/><StationHandle/></>} />
        <Route exact path="/bus/home" element={<BusHome/>} />
        <Route exact path="/bus/live" element={<BusLive/>} />

      </Routes>
    </Router>
    </div>
  );
};

export default MapApp;
