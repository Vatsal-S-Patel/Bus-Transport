import { useEffect, useState } from "react";

const BusLive = () => {
  // Usestate to handle the form data
  const [formData, setFormData] = useState({
    last_station_order: 0,
    last_updated: "",
    traffic: 0,
    status: 0,
    bus_id: 0,
  });

  const [routeId, setRouteId] = useState("");
  const [routeName, setRouteName] = useState("");
  var socket = window._DEFAULT_DATA;

  // The function to change state as per the change in the form data
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  useEffect(() => {
    setInterval(() => {
      document.getElementById("submitBtn").click();
    }, 4000);
  }, []);

  // Func to handle the submit and emit information to socket
  const handleSubmit = (e) => {
    e.preventDefault();

    // Set the interval to send data to socket after a specific time

    // Add data to send in the socket
    formData.last_updated = new Date().toLocaleTimeString().substring(0, 5);
    formData.status = parseInt(formData.status);
    formData.traffic = parseInt(formData.traffic);

    formData.route_name = routeName;
    if (formData.bus_id === "") {
      formData.bus_id = 0;
    }
    formData.bus_id = parseInt(formData.bus_id);
    // Error message for Fetching live location of Bus
    function errorCallback(error) {
      console.error("Error getting current location:", error.message);
    }

    // Fetch the current location
    if (navigator.geolocation) {
      navigator.geolocation.watchPosition(
        (position) => {
          formData.lat = position.coords.latitude;
          formData.long = position.coords.longitude;
        },
        errorCallback,
        { enableHighAccuracy: true }
      );
    }

    // Emit the request on update with json Data and RouteId as params work as Rooms to join for client
    console.log(formData);
    socket.emit("update", JSON.stringify(formData), parseInt(routeId));
    socket.emit("bus", formData.bus_id);
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
            id="bus_id"
            name="bus_id"
            value={formData.bus_id}
            onChange={(e) => {
              handleChange(e);
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
            id="route_id"
            name="route_id"
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
            id="route_name"
            name="route_name"
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
            onChange={(e) => handleChange(e)}
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
            onChange={(e) => {
              handleChange(e);
            }}
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

export default BusLive;
