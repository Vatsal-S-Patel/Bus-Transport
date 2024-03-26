import React, { useEffect, useState } from "react";
import IP from "../IP";
import { useNavigate } from "react-router-dom";
import { CustomizeTable, Table } from "./sub-components/table_for_data";

const DriverHandle = () => {
  // States to handle form data and data from server
  const [driverDetails, setDriverDetails] = useState({
    id: null,
    name: "",
    phone: "",
    gender: null,
    dob: "",
  });
  const [drivers, setDrivers] = useState([]);
  const [allDrivers, setallDrivers] = useState([]);
  var navigator = useNavigate();

  // Fetch Data from Driver Database
  const fetchData = async () => {
    try {
      const res = await fetch("http://" + IP + ":8080/api/driver/");
      const jsonres = await res.json();

      setDrivers(jsonres.Data);
      setallDrivers(jsonres.Data);
    } catch (err) {
      alert(err);
    }
  };
  // Set the Driver data into states
  useEffect(() => {
    localStorage.getItem("user") == undefined
      ? navigator("/admin/login")
      : navigator();
    fetchData();
  }, []);

  // Handle the form data in integer and string
  const handleChange = (e) => {
    const { name, value } = e.target;
    if (name == "id" || name == "gender") {
      setDriverDetails({ ...driverDetails, [name]: parseInt(value) });
    } else {
      setDriverDetails({ ...driverDetails, [name]: value });
    }
  };

  // Handle the post request for driver database
  const handleSubmit = async (e) => {
    e.preventDefault();
    // Post request to Driver and check the Error
    try {
      const res = await fetch("http://" + IP + ":8080/api/driver/", {
        method: "POST",
        body: JSON.stringify(driverDetails),
      });

      if (res.status === 200) {
        alert("Success");
        setDrivers([...drivers, driverDetails]);
        setallDrivers([...allDrivers, driverDetails]);
      }else{
        alert(res.status)
      }
    } catch (err) {
      alert(err);
    }

    // Clear the driver state
    setDriverDetails({ id: "", name: "", phone: "", gender: "", dob: "" });
  };

  // Handle delete in database and in states
  const handleDelete = (id, index) => {
    const updatedDrivers = [...drivers];
    updatedDrivers.splice(index, 1);
    // POST request for deleting data accroding to id and handle error
    fetch("http://" + IP + ":8080/api/driver/" + id, {
      method: "POST",
    })
      .then((res) => {
        if (res.status == 200) {
          alert("Success");
          setDrivers(updatedDrivers);
          setallDrivers(updatedDrivers);
        } else {
          alert("Cannot be deleted");
        }
      })
      .catch((err) => {
        alert(err);
      });
  };

  // Filter the data according to name and phone number of driver
  const handleFilteredItems = (e) => {
    setDrivers(
      allDrivers.filter(
        (item) =>
          item.name.toLowerCase().includes(e.target.value.toLowerCase()) ||
          item.phone.toLowerCase().includes(e.target.value.toLowerCase())
      )
    );
  };

  return (
    <>
      <center>
        <div className="container mx-auto mt-8">
          <h2 className="text-2xl font-bold mb-4">Add Driver</h2>
          <form onSubmit={handleSubmit} className="max-w-md">
            <div className="grid grid-cols-2 gap-4">
              <div className="mb-4">
                <input
                  type="number"
                  id="id"
                  name="id"
                  value={driverDetails.id}
                  onChange={handleChange}
                  placeholder="Enter Driver ID"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
                <input
                  type="text"
                  id="name"
                  name="name"
                  value={driverDetails.name}
                  onChange={handleChange}
                  placeholder="Enter Name"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
                <input
                  type="text"
                  id="phone"
                  name="phone"
                  value={driverDetails.phone}
                  onChange={handleChange}
                  placeholder="+91 75676 75676"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
                <input
                  type="number"
                  id="gender"
                  name="gender"
                  value={driverDetails.gender}
                  onChange={handleChange}
                  placeholder="Enter Gender"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
                <input
                  type="text"
                  id="dob"
                  name="dob"
                  value={driverDetails.dob}
                  onChange={handleChange}
                  placeholder="DD-MM-YYYY"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="flex justify-center my-3 mx-2">
                <button
                  onClick={handleSubmit}
                  className="bg-blue-500 text-white w-full  rounded hover:bg-blue-600 focus:outline-none focus:ring focus:border-blue-300"
                >
                  Submit
                </button>
              </div>
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
      
      <CustomizeTable details={driverDetails} handleDelete={handleDelete} items={drivers}/>
    </>
  );
};

export default DriverHandle;
