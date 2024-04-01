import React, { useEffect, useState } from "react";
import { json, useNavigate } from "react-router-dom";
import IP from "../IP";
import { CustomizeTable } from "./sub-components/table_for_data";

const BusRouteHandle = () => {
  // STATES TO MANAGE THE FORM DATA AND DATA FROM DATABASE
  const [routeDetails, setRouteDetails] = useState({
    id: 0,
    name: "",
    source: 0,
    destination: 0,
    status: 0,
    Mapping: [],
  });
  const [routes, setRoutes] = useState([]);
  const [allRoutes, setAllRoutes] = useState([]);
  const [station, setStation] = useState([]);
  const [isSelected, setIsSelected] = useState([]);
  const [orderStations, setOrderStations] = useState([]);
  // TO NAVIGATE TO DIFFERENT PAGES
  var navigator = useNavigate();

  // FUNCTION TO FETCH DATA OF STATIONS FROM DATABASE THROUGH SERVER
  const fetchDataStation = async () => {
    try {
      const res = await fetch("http://" + IP + ":8080/api/station/");
      const jsonres = await res.json();
      var len = jsonres.Data.map((station) => {
        return false;
      });

      setIsSelected(len);
      setStation(jsonres.Data);
    } catch (err) {
      alert(err);
    }
  };

  // FUNCTION TO FETCH THE ROUTE FROM DATABASE THROUGH SERVER
  const fetchDataRoute = async () => {
    try {
      const res = await fetch("http://" + IP + ":8080/api/route/");
      const jsonres = await res.json();

      console.log(jsonres.Data)
      setRoutes(jsonres.Data);
      setAllRoutes(jsonres.Data);
    } catch (err) {
      alert(err);
    }
  };
  // RUN FUNCTION BEFORE COMPONENT MOUNT TO CHECK ADMIN IS VALID AND FETCH THE DATA'S
  useEffect(() => {
    localStorage.getItem("user") == undefined
      ? navigator("/admin/login")
      : navigator();

    fetchDataStation();
    fetchDataRoute();
  }, []);
  // HANDLE THE CHANGE FOR THE FORM
  const handleChange = (e) => {
    const { name, value } = e.target;
    if (name == "name") {
      setRouteDetails({ ...routeDetails, [name]: value });
    } else {
      setRouteDetails({ ...routeDetails, [name]: parseInt(value) });
    }
  };

  // HANDLE THE SUBMIT FOR FORM FOR ROUTE AS WELL AS ORDER OF ROUTE
  const handleSubmit = (e) => {
    e.preventDefault();
    // CREATE ORDER STATION MAPPING JSON ARRAY OBJECT
    var OrderOfStationsInRoute = orderStations.map((station, index) => {
      return { station_order: index + 1, station_id: station.order };
    });
    // ADD DATA WITH THE ROUTE DATA OF FORM
    var dataToSend = { ...routeDetails, ["Mapping"]: OrderOfStationsInRoute };

    // POST REQUEST FOR TH EROUTES AND MAPPING , CHECK THE ERROR
    try {
      fetch("http://" + IP + ":8080/api/route/", {
        method: "POST",
        type: "application/json",
        body: JSON.stringify(dataToSend),
      })
        .then((res) => {
          if (res.status === 200) {
            alert("Success");
            setRoutes([...routes, routeDetails]);
            setAllRoutes([...allRoutes, routeDetails]);
          } else {
            alert("Error in Inserting Data");
          }
        })
        .catch((err) => {
          alert(err);
        });
    } catch (err) {
      alert(err);
    }
  };
  // HANDLE THE DELETE OF ROUTE AS WELL AS MAPPING
  const handleDelete = (id, index) => {
    const updatedRoutes = [...routes];
    updatedRoutes.splice(index, 1);
    // Delete data from Route Database
    fetch("http://" + IP + ":8080/api/route/" + id, {
      method: "POST",
    })
      .then((res) => {
        if (res.status == 200) {
          alert("Success");
          setRoutes(updatedRoutes);
          setAllRoutes(updatedRoutes);
        } else {
          alert("Cannot be deleted");
        }
      })
      .catch((err) => {
        alert(err);
      });
  };

  // Handle the order of the Station while Entering and deleting
  const handleStationsInRoute = (id, name) => {
    // If the selected Station is not in Route Add it Otherwise not
    if (
      isSelected == undefined ||
      isSelected[id] == false ||
      isSelected[id] == null
    ) {
      var len = isSelected;
      len[id] = true;
      setIsSelected(len);
      setOrderStations([...orderStations, { order: id, station: name }]);
    }
  };

  // Delete the station from the order
  const handleDeleteStation = (id) => {
    // The Array with the station remaining without id record
    var finalRemain = orderStations.filter((station) => {
      if (station.order != id) return station;
    });

    // Alter the new Selected map according to the delete station
    var newSelected = isSelected.map((ele, index) => {
      if (ele == true) {
        return index == id ? false : true;
      } else {
        return false;
      }
    });
    // Set Both OrderStation and Selected
    setIsSelected(newSelected);
    setOrderStations(finalRemain);
  };

  // Filter function to see routes according to name, source or destination
  const handleFilteredItems = (e) => {
    setRoutes(
      allRoutes.filter(
        (item) =>
          item.name.toLowerCase().includes(e.target.value.toLowerCase()) ||
          String(item.source)
            .toLowerCase()
            .includes(e.target.value.toLowerCase()) ||
          String(item.destination)
            .toLowerCase()
            .includes(e.target.value.toLowerCase())
      )
    );
  };

  return (
    <div>
      <center>
        <div className="container mx-auto mt-8">
          <h2 className="text-2xl font-bold mb-4">Add Route</h2>
          <form onSubmit={handleSubmit} className="max-w-md">
            <div className="grid grid-cols-2 gap-4">
              <div className="mb-4">
                <label>Route Id</label>
                <input
                  type="number"
                  id="id"
                  name="id"
                  value={routeDetails.id}
                  onChange={handleChange}
                  placeholder="Enter Route Id"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
              <label>Route Name</label>
                <input
                  type="text"
                  id="name"
                  name="name"
                  value={routeDetails.name}
                  onChange={handleChange}
                  placeholder="Enter Name"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
                <label>Source Id</label>
                <input
                  type="number"
                  id="source"
                  name="source"
                  value={routeDetails.source}
                  onChange={handleChange}
                  placeholder="Enter Source"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
              <label>Destination Id</label>
                <input
                  type="number"
                  id="destination"
                  name="destination"
                  value={routeDetails.destination}
                  onChange={handleChange}
                  placeholder="Enter Destination"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
              <label>Route Status</label>
                <input
                  type="number"
                  id="status"
                  name="status"
                  value={routeDetails.status}
                  onChange={handleChange}
                  placeholder="Enter Status"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="flex justify-center my-3 mx-2">
                <button
                  type="submit"
                  className="bg-blue-500 text-white w-full  rounded hover:bg-blue-600 focus:outline-none focus:ring focus:border-blue-300"
                >
                  Submit
                </button>
              </div>
            </div>
          </form>
        </div>
        <hr></hr>
        <div className="grid grid-cols-2 ">
          <div className="station display-flex">
            {station!= undefined ? station.map((s, index) => {
              return (
                <button
                  key={index}
                  type="button"
                  onClick={() => {
                    handleStationsInRoute(s.id, s.name);
                  }}
                  className="text-white font-semibold p-1 px-1 m-2 text-sm text-blue-500 border-2 rounded hover:bg-blue-500 hover:text-white focus:outline-none focus:ring focus:border-blue-300"
                >
                  {s.name} ({s.id})
                </button>
              );
            }):""}
          </div>
          <div className="tableOrder">
            <table className="w-5/12">
              <thead>
                <tr>
                  <th>Order</th>
                  <th>Station</th>
                  <th>Delete </th>
                </tr>
              </thead>
              <tbody>
                {orderStations != undefined ? orderStations.map((orderStation, index) => (
                  <>
                    <tr className="p-3 " key={index}>
                      <td className="text-sm m-5">{index + 1}</td>
                      <td className="text-sm m-5">{orderStation.station}</td>
                      <td className="m-5">
                        <button
                          className="bg-red-500 text-xs text-white w-11/12 p-1 rounded hover:bg-red-600 focus:outline-none focus:ring focus:border-red-300"
                          onClick={() => {
                            handleDeleteStation(orderStation.order);
                          }}
                        >
                          Delete
                        </button>
                      </td>
                    </tr>
                    <tr>
                      <td></td>
                      <td></td>
                      <td></td>
                    </tr>
                  </>
                )):""}
              </tbody>
            </table>
          </div>
        </div>
        <input
          type="text"
          placeholder="Search..."
          onChange={handleFilteredItems}
          className="px-2 py-2 mt-10 border w-4/12 border-gray-300 rounded-md text-base focus:outline-none focus:border-blue-500 transition duration-300"
        />
      </center>

      <CustomizeTable details={routeDetails} handleDelete={handleDelete} items={routes} isRoute={true}/>
    </div>
  );
};

export default BusRouteHandle;
