import React, { useEffect, useState } from "react";
import IP from "../IP";
import { useNavigate } from "react-router-dom";
import { CustomizeTable } from "./sub-components/table_for_data";

const BusHandle = () => {
  // States to save the Form Data and the Record from
  const [busDetails, setBusDetails] = useState({
    id: 0,
    capacity: 0,
    model: "",
    registration_number: "",
  });
  
  // TWO STATES ONE FOR STORING AND ONE FOR FILTERING , SO THE DATA IN STATE CAN NOT BE LOST
  const [filteredBuses, setFilteredBuses] = useState([]);
  const [allBuses, setAllBuses] = useState([]);
  var navigator = useNavigate();

  // FUNCTION TO FETCH THE DATA IN STATEES
  const fetchData = async () => {
    try {
      const res = await fetch("http://" + IP + ":8080/api/bus/");
      const jsonres = await res.json();
      setFilteredBuses(jsonres.Data);
      setAllBuses(jsonres.Data);
    }catch(err){
      alert(err)
    }
  };
  // FUNCTION TO RUN BEFORE COMPONENT MOUNT
  useEffect(() => {
    localStorage.getItem("user") == undefined
      ? navigator("/admin/login")
      : navigator();
    fetchData();
    ;
  }, []);

  // CHANGE THE STATE DETAILS FOR FORM ON CHANGE
  const handleChange = (e) => {
    const { name, value } = e.target;
    if (e.target.type == "number") {
      setBusDetails({ ...busDetails, [name]: parseInt(value) });
    } else {
      setBusDetails({ ...busDetails, [name]: value });
    }
  };

  // HANDLE THE SUBMISSION OF FORM DATA TO BACKEND AND HANDLE ERROR
  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const res = await fetch("http://" + IP + ":8080/api/bus/", {
        method: "POST",
        body: JSON.stringify(busDetails),
      });
      if (res.status == 200) {
        alert("Success");
        setFilteredBuses([...filteredBuses, busDetails]);
        setAllBuses([...allBuses, busDetails]);
      }else{
        res.json().then((d)=>{
          alert(d)
          console.log(d)
        })
      }
    } catch (err) {
      alert(err);
    }

    setBusDetails({
      registration_number: "",
      model: "",
      capacity: "",
      model: "",
      id: "",
    });
  };

  // FUNCTION TO DELETE THE RECORD FROM DATABASE AS WELL AS STATE
  const handleDelete = (id, index) => {
    const updatedBuses = [...allBuses];
    updatedBuses.splice(index, 1);

    fetch("http://" + IP + ":8080/api/bus/" + id, {
      method: "POST",
    })
      .then((res) => {
        if (res.status == 200) {
          alert("Success");
          setFilteredBuses(updatedBuses);
          setAllBuses(updatedBuses);
        } else {
          alert("Cannot be deleted");
        }
      })
      .catch((err) => {
        alert(err);
      });
  };

  // HANDLE THE FILTER OF THE DATA RECORD ON BASIS OF REGISSTRATION NUMBER OR MODEL OF BUS
  const handleFilteredItems = (e) => {
    setFilteredBuses(
      allBuses.filter(
        (item) =>
          item.registration_number
            .toLowerCase()
            .includes(e.target.value.toLowerCase()) ||
          item.model.toLowerCase().includes(e.target.value.toLowerCase())
      )
    );
  };

  return (
    <>
      <center>
        <div className="container place-items-center mt-8">
          <h2 className="text-2xl font-bold mb-4">Add Bus</h2>
          <form onSubmit={handleSubmit} className="max-w-md">
            <div className="grid grid-cols-2 gap-4">
              <div className="mb-4">
                <label
                  htmlFor="id"
                  className="block text-sm font-medium text-gray-700"
                >
                  ID:
                </label>
                <input
                  type="number"
                  id="id"
                  name="id"
                  value={busDetails.id}
                  onChange={handleChange}
                  placeholder="Enter ID"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
                <label
                  htmlFor="capacity"
                  className="block text-sm font-medium text-gray-700"
                >
                  Capacity:
                </label>
                <input
                  type="number"
                  id="capacity"
                  name="capacity"
                  value={busDetails.capacity}
                  onChange={handleChange}
                  placeholder="Enter Capacity"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
                <label
                  htmlFor="model"
                  className="block text-sm font-medium text-gray-700"
                >
                  Model:
                </label>
                <input
                  type="text"
                  id="model"
                  name="model"
                  value={busDetails.model}
                  onChange={handleChange}
                  placeholder="Eicher"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
                <label
                  htmlFor="numberPlate"
                  className="block text-sm font-medium text-gray-700"
                >
                  Number Plate:
                </label>
                <input
                  type="text"
                  id="registration_number"
                  name="registration_number"
                  value={busDetails.registration_number}
                  onChange={handleChange}
                  placeholder="GJ01XB1231"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
            </div>
            <div className="flex justify-end mt-4">
              <button
                onClick={handleSubmit}
                className="bg-blue-500 text-white p-2 w-full rounded hover:bg-blue-600 focus:outline-none focus:ring focus:border-blue-300"
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

      <CustomizeTable details={busDetails} handleDelete={handleDelete} items={filteredBuses}/>
    </>
  );
};

export default BusHandle;
