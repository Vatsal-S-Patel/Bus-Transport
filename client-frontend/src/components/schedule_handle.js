import React, { useEffect, useState } from "react";
import IP from "../IP";
import { useNavigate } from "react-router-dom";
import { CustomizeTable } from "./sub-components/table_for_data";

const ScheduleHandle = () => {
  // States to manage Form Data and Schedule Data from Server
  const [scheduleDetails, setScheduleDetails] = useState({
    id: "",
    bus_id: "",
    route_id: "",
    dep: "",
  });
  const [schedules, setSchedules] = useState([]);
  const [allSchedule, setAllSchedule] = useState([]);
  // Used to Navigate through the website
  var navigator = useNavigate();

  // function to fetch the data from server
  const fetchData = async () => {
    try {
      const res = await fetch("http://" + IP + ":8080/api/schedule/");
      const jsonres = await res.json();

      // set Data in both states
      setSchedules(jsonres.Data);
      setAllSchedule(jsonres.Data);
    } catch (err) {
      alert(err);
    }
  };

  // The Function to run before mounting the components on website
  useEffect(() => {
    localStorage.getItem("user") == undefined
      ? navigator("/admin/login")
      : navigator();
    fetchData();
  }, []);

  // Handle the form value and convert them to int if needed
  const handleChange = (e) => {
    const { name, value } = e.target;
    if (name == "id" || name == "bus_id" || name == "route_id") {
      setScheduleDetails({ ...scheduleDetails, [name]: parseInt(value) });
    } else {
      setScheduleDetails({ ...scheduleDetails, [name]: value });
    }
  };

  // Handle the form submission for schedule and change data in states
  const handleSubmit = async (e) => {
    e.preventDefault();

    // POST request to server for data entry
    try {
      const res = await fetch("http://" + IP + ":8080/api/schedule/", {
        method: "POST",
        body: JSON.stringify(scheduleDetails),
      });
      // checking the Status Code for error handling
      if (res.status === 200) {
        alert("Success");
        setSchedules([...schedules, scheduleDetails]);
        setAllSchedule([...schedules, scheduleDetails]);
      }
    } catch (err) {
      alert(err);
    }

    // Clear the values of the Form
    setScheduleDetails({ id: "", bus_id: "", route_id: "", dep: "" });
  };
  // Handle the Delete functionality for schedule and change data in states
  const handleDelete = (id, index) => {
    const updatedSchedules = [...schedules];
    updatedSchedules.splice(index, 1);

    // POST request for deleting schedules and handling the error by status code
    fetch("http://" + IP + ":8080/api/schedule/" + id, {
      method: "POST",
    })
      .then((res) => {
        if (res.status == 200) {
          // Update the states for user
          alert("Success");
          setSchedules(updatedSchedules);
          setAllSchedule(updatedSchedules);
        } else {
          alert("Cannot be deleted");
        }
      })
      .catch((err) => {
        alert(err);
      });
  };

  const handleFilteredItems = (e) => {
    console.log(allSchedule);
    setSchedules(
      allSchedule.filter(
        (item) =>
          String(item.route_id)
            .toLowerCase()
            .includes(e.target.value.toLowerCase()) ||
          item.dep.toLowerCase().includes(e.target.value.toLowerCase())
      )
    );
  };

  return (
    <>
      <center>
        <div className="container mx-auto mt-8">
          <h2 className="text-2xl font-bold mb-4">Add Schedule</h2>
          <form onSubmit={handleSubmit} className="max-w-md">
            <div className="grid grid-cols-2 gap-4">
              <div className="mb-4">
                <input
                  type="number"
                  id="id"
                  name="id"
                  value={scheduleDetails.id}
                  onChange={handleChange}
                  placeholder="Enter Schedule Id"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
                <input
                  type="number"
                  id="bus_id"
                  name="bus_id"
                  value={scheduleDetails.bus_id}
                  onChange={handleChange}
                  placeholder="Enter Bus Id"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
                <input
                  type="number"
                  id="route_id"
                  name="route_id"
                  value={scheduleDetails.route_id}
                  onChange={handleChange}
                  placeholder="Enter Route Id"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
              <div className="mb-4">
                <input
                  type="text"
                  id="dep"
                  name="dep"
                  value={scheduleDetails.dep}
                  onChange={handleChange}
                  placeholder="Enter HH:MM"
                  className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
                />
              </div>
            </div>
            <div className="flex justify-center my-2">
              <button
                type="submit"
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

      <CustomizeTable details={scheduleDetails} handleDelete={handleDelete} items={schedules}/>
    </>
  );
};

export default ScheduleHandle;
