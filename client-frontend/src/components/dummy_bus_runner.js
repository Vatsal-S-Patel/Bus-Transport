import { useEffect, useState } from "react";
import coords from "./latlongJSONArray";

const BusHome = () => {
  // Usestate to handle the form data
  const [formData, setFormData] = useState({
    last_station_order: 0,
    last_updated: "",
    traffic: 0,
    status: 0,
  });

  const [routeId, setRouteId] = useState("");
  const [routeName, setRouteName] = useState("");
  const [busId, setBusId] = useState("");

  var socket = window._DEFAULT_DATA;

  // The function to change state as per the change in the form data
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  // Func to handle the submit and emit information to socket
  const handleSubmit = (e) => {
    e.preventDefault();

    var i = 0;
    console.log("dfgds");
    // Set the interval to send data to socket after a specific time
    var clear = setInterval(() => {
      // Add data to send in the socket
      formData.last_updated = new Date().toLocaleTimeString().substring(0, 5);
      formData.status = parseInt(formData.status);
      formData.traffic = parseInt(formData.traffic);

      formData.last_station_order = 1;
      formData.last_station_name = coords[i % 41].location;
      formData.bus_id = parseInt(busId);

      formData.route_name = routeName;
      formData.lat = coords[i % 41].latitude;
      formData.long = coords[i % 41].longitude;

      i++;

      if (i == 41) {
        formData.status = 0;
      }
      console.log(coords[i % 41].latitude);
      // Emit the request on update with json Data and RouteId as params work as Rooms to join for client
      console.log(formData);
      socket.emit("update", JSON.stringify(formData), parseInt(routeId));

      if (i == 41) {
        clearInterval(clear);
      }
    }, 5000);
  };

  return (
    <div className="max-w-md mx-auto mt-8 p-6 bg-gray-100 rounded-md shadow-md">
      <h2 className="text-lg font-semibold mb-4">
        Enter Bus Route Information
      </h2>
      <form>
        <div className="mb-4">
          <label
            htmlFor="last_station_order"
            className="block text-sm font-medium text-gray-700"
          >
            Bus Id
          </label>
          <input
            type="number"
            id="last_station_order"
            name="last_station_order"
            value={busId}
            onChange={(e) => {
              setBusId(e.target.value);
            }}
            className="mt-1 p-2 w-full border rounded-md focus:outline-none focus:ring focus:border-blue-300"
            required
          />
        </div>
        <div className="mb-4">
          <label
            htmlFor="last_station_order"
            className="block text-sm font-medium text-gray-700"
          >
            Route Id
          </label>
          <input
            type="number"
            id="last_station_order"
            name="last_station_order"
            value={routeId}
            onChange={(e) => {
              setRouteId(e.target.value);
            }}
            className="mt-1 p-2 w-full border rounded-md focus:outline-none focus:ring focus:border-blue-300"
            required
          />
        </div>
        <div className="mb-4">
          <label
            htmlFor="last_station_order"
            className="block text-sm font-medium text-gray-700"
          >
            Route Name
          </label>
          <input
            type="text"
            id="last_station_order"
            name="last_station_order"
            value={routeName}
            onChange={(e) => setRouteName(e.target.value)}
            className="mt-1 p-2 w-full border rounded-md focus:outline-none focus:ring focus:border-blue-300"
            required
          />
        </div>
        <div className="mb-4">
          <label
            htmlFor="traffic"
            className="block text-sm font-medium text-gray-700"
          >
            Traffic
          </label>
          <input
            type="number"
            id="traffic"
            name="traffic"
            value={formData.traffic}
            onChange={handleChange}
            className="mt-1 p-2 w-full border rounded-md focus:outline-none focus:ring focus:border-blue-300"
            required
          />
        </div>
        <div className="mb-4">
          <label
            htmlFor="status"
            className="block text-sm font-medium text-gray-700"
          >
            Status
          </label>
          <select
            id="status"
            name="status"
            value={formData.status}
            onChange={handleChange}
            className="mt-1 p-2 w-full border rounded-md focus:outline-none focus:ring focus:border-blue-300"
            required
          >
            <option value="">Select Status</option>
            <option value="1">Active</option>
            <option value="0">Inactive</option>
          </select>
        </div>
        <button
          type="submit"
          onClick={handleSubmit}
          id="submitBtn"
          className="w-full bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:bg-blue-600"
        >
          Submit
        </button>
      </form>
    </div>
  );
};

export default BusHome;
