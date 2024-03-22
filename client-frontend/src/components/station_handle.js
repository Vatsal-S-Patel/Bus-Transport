import React, { useEffect, useState } from "react";
import IP from "../IP";
import { useNavigate } from "react-router-dom";
import { CustomizeTable } from "./sub-components/table_for_data";

const StationHandle = () => {
  const fetchData = async () => {
    try {
      const res = await fetch("http://" + IP + ":8080/api/station/");
      const jsonres = await res.json();

      setStations(jsonres.Data);
      setStaticStation(jsonres.Data);
    } catch (err) {
      alert(err);
    }
  };

  const [stationDetails, setStationDetails] = useState({
    id: null,
    name: "",
    lat: null,
    long: null,
  });
  const [stations, setStations] = useState([]);
  const [staticStation, setStaticStation] = useState([]);
  var navigator = useNavigate();

  useEffect(() => {
    localStorage.getItem("user") == undefined
      ? navigator("/admin/login")
      : navigator();

    fetchData();
  }, []);

  const handleChange = (e) => {
    const { name, value } = e.target;
    if (name == "name") {
      setStationDetails({ ...stationDetails, [name]: value });
    } else {
      setStationDetails({ ...stationDetails, [name]: parseFloat(value) });
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    fetch("http://" + IP + ":8080/api/station/", {
      method: "POST",
      type: "application/json",
      body: JSON.stringify(stationDetails),
    }).then((res) => {
      if (res.status === 200) {
        alert("Success");
        setStations([...stations, stationDetails]);
        setStaticStation([...stations, stationDetails]);
      } else {
        alert("Error in Inserting Data");
      }
    }).catch((err) => {alert(err) })
    

    setStationDetails({ id: "", name: "", lat: "", long: "" });
  };

  const handleDelete = (id, index) => {
    const updatedStations = [...stations];
    updatedStations.splice(index, 1);

    fetch("http://" + IP + ":8080/api/station/" + id, {
      method: "POST",
    })
      .then((res) => {
        if (res.status == 200) {
          alert("Success");
          setStations(updatedStations);
        } else {
          alert("Cannot be deleted");
        }
      })
      .catch((err) => {
        alert(err);
      });
  };

  const handleFilteredItems = (e) => {
    setStations(
      staticStation.filter((item) =>
        item.name.toLowerCase().includes(e.target.value.toLowerCase())
      )
    );
  };

  return (
    <>
      <center>
        <div className="container mx-auto mt-8">
          <h2 className="text-2xl font-bold mb-4">Add Station</h2>
          <form onSubmit={handleSubmit} className="max-w-md">
            <div className="grid grid-cols-2 gap-4">
              <div className="mb-4">
                <input
                  type="number"
                  id="id"
                  name="id"
                  value={stationDetails.id}
                  onChange={handleChange}
                  placeholder="Enter Station ID"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
                <input
                  type="text"
                  id="name"
                  name="name"
                  value={stationDetails.name}
                  onChange={handleChange}
                  placeholder="Enter Station Name"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
                <input
                  type="number"
                  step="0.0000000001"
                  id="lat"
                  name="lat"
                  value={stationDetails.lat}
                  onChange={handleChange}
                  placeholder="Enter Latitude"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
                <input
                  type="number"
                  step="0.00000000001"
                  id="long"
                  name="long"
                  value={stationDetails.long}
                  onChange={handleChange}
                  placeholder="Enter Longitude"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
            </div>
            <div className="flex justify-center my-2">
              <button
                type="submit"
                onChange={handleSubmit}
                className="bg-blue-500 p-1 w-full text-white  rounded hover:bg-blue-600 focus:outline-none focus:ring focus:border-blue-300"
              >
                Submit
              </button>
            </div>
          </form>
        </div>
        <input
          type="text"
          placeholder="Search..."
          onChange={handleFilteredItems}
          className="px-2 py-2 mt-10 border w-4/12 border-gray-300 rounded-md text-base focus:outline-none focus:border-blue-500 transition duration-300"
        />
      </center>

      <CustomizeTable details={stationDetails} handleDelete={handleDelete} items={stations}/>
    </>
  );
};

export default StationHandle;
