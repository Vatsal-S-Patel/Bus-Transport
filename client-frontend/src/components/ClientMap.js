import React, { useEffect, useState, useRef } from "react";
import "ol/ol.css";
import "../App.js";
import MapComponent from "./MapComponent.js";
import { Manager } from "socket.io-client";
import resetIcon from "../images/reset-svgrepo-com.svg"
import backIcon from "../images/backward-3-svgrepo-com.svg"
import IP from "../IP.js";


/* Initially get all stations information then no need to bother api about this. */
export function getStationInfoById(stationId, stationsMap) {
  // see if id present then give name
  if (stationsMap.has(Number(stationId))) {
    return stationsMap.get(Number(stationId));
  }
  return {
    id: Number(stationId),
    name: "not available",
  };
}

const ClientMap = () => {
  // All stations information
  const [stations, setStations] = useState([]);
  const [stationsMap, setStationsMap] = useState();

  // user selected sourceStation and destinationStation
  const [sourceStation, setSourceStation] = useState(null);
  const [destinationStation, setDestinationStation] = useState(null);

  // After selecting source and destination we will show user below information
  const [currentStationRoutes, setCurrentStationRoutes] = useState([]);
  const [currentStationRoutesError, setCurrentStationRoutesError] =
    useState(null);

  const [socketConn, setSocketConn] = useState(null);
  const [currentBuses, setCurrentBuses] = useState([]);
  

  const sourceStationRef = useRef();
  const destinationStationRef = useRef();

  // Initially set it false
  const [expanded, setExpanded] = useState(false);

  // Event handler for source station select
  const handleSourceChange = (newValue) => {
    setSourceStation(newValue);
  };

  // Event handler for destination station select
  const handleDestinationChange = (newValue) => {
    setDestinationStation(newValue);
  };

  const selectBus = (busInfo) => {
    console.log("SOCKET EMIT EVENT");
    setCurrentBuses([]);
    socketConn.emit("busSelected", busInfo.bus_id);
  };

  const resetSelections = () => {
    let btn = document.querySelector(".reset-selection-btn")
    btn.style.rotate = "360deg"

    setTimeout(() => {
     btn.style.rotate = "0deg"
    }, 1000);

    setSourceStation(null);
    setDestinationStation(null);

    socketConn.emit("sourceSelected", []);

    document.getElementById("source").selectedIndex = 0;
    document.getElementById("destination").selectedIndex = 0;

    // Jo data uski vajah se aaya he use bhi nikal do
    setCurrentBuses([]);
    setCurrentStationRoutes([]);
    setCurrentStationRoutesError("Not Selected");
  };

  const expandShrinkView = (event) => {
    if (expanded) {
      setExpanded(false);
    } else {
      setExpanded(true);
    }
  };

  useEffect(() => {
    let newArray = currentStationRoutes;
    // as this data changes see if there is any update or not if there then update cuurrent routes
    currentBuses.forEach((busData) => {
      newArray = newArray.map((routeInfo) =>
        routeInfo.bus_id == busData.bus_id
          ? {
              ...routeInfo,
              lat: busData.lat,
              long: busData.long,
              last_station_order: busData.last_station_order,
            }
          : routeInfo
      );
    });

    socket.on("update", (busInfo) => {
      console.log("UPDATE CAME", busInfo)
      // Check Already array me he ki nahi
      let isAlreadyBusDisplayed = false;
      
      currentBuses.map((bus) => {
        if (bus.bus_id == busInfo.bus_id) {
          isAlreadyBusDisplayed = true;
        }
      });
     
      if (isAlreadyBusDisplayed) {
        // Update this information
        let newArray = currentBuses.map((bus) =>
          bus.bus_id == busInfo.bus_id
            ? {
                ...busInfo,
              }
            : bus
        );
        setCurrentBuses(newArray);
        // Agar ye bus current routes me he then updated,
      } else {
        setCurrentBuses([...currentBuses , busInfo])
      }
    });

    setCurrentStationRoutes(newArray);
  }, [currentBuses]);

  // fetch route infomation when sourcestation adn destination station changes
  useEffect(() => {
    setCurrentBuses([]);
    
    if (sourceStation != null) {
      let data = {};
      data["source"] = Number(sourceStation);
      if (destinationStation != null) {
        data["destination"] = Number(destinationStation);
      }
      // Reset Station Routes
      setCurrentStationRoutes([]);
      setCurrentStationRoutesError(null);

      fetch(`http://${IP}:8080/api/schedule/GetUpcomingBus`, {
        method: "POST",
        body: JSON.stringify(data),
        type: "application/json",
      })
        .then((response) => response.json())
        .then((data) => {
          let routeStations = data.Data;

          if (data.Code == 500 || data.Code == 404) {
            setCurrentStationRoutesError(data.Message);

            console.log("SOCKET EMITTED");
            socketConn.emit("sourceSelected", []);
          } else {
            let time = new Date()
            let updatedData = routeStations.map((info) => {
              return {
                ...info,
                departure_time: info.departure_time.substr(11, 5),
                sourceName: getStationInfoById(info.source, stationsMap).name,
                destinationName: getStationInfoById(
                  info.destination,
                  stationsMap
                ).name,
                last_updated: time.getHours() + ":"+time.getMinutes(),
              };
            });

            // If it has lat long then add this bus to current buses
            // it let all the buses have lat long
            
            setCurrentStationRoutes(updatedData);

            let routes = new Set();
            updatedData.forEach((element) => {
              routes.add(element.route_id);
            });

            let routesArray = [];
            routes.forEach((route) => routesArray.push(route));

            socketConn.emit("sourceSelected", routesArray);
          }
        })
        .catch((error) => {
          console.error("Error fetching stations data:", error);
        });
    }
  }, [sourceStation, destinationStation, stations]);

  useEffect(() => {
    fetch(`http://${IP}:8080/api/station/`)
      .then((response) => response.json())
      .then((data) => {
        let stations = data.Data;
        setStations(stations);

        let mp = new Map();
        stations.forEach((station) => {
          mp.set(station.id, station);
        });

        setStationsMap(mp);
      })
      .catch((error) => {
        console.error("Error fetching stations data:", error);
      });

    setSocketConn(socket);

    socket.on("connect", () => {
      console.log("connected");
    });

    socket.on("disconnect", () => {
      setCurrentBuses([]);
      console.log("disconnected");
    });

    socket.on("error", (error) => {
      setCurrentBuses([]);
      console.log("Error", error);
    });

    socket.on("roomJoined", (data) => {
      console.log("Room joined", data);
    });

    return () => {
      socket.close();
    };
  }, []);

  return (
    <div style={{ height: "100vh", width: "100%" }}>
      {stations.length > 0 && (
        <MapComponent
          stations={stations}
          sourceStation={sourceStation}
          setSourceStation={setSourceStation}
          destinationStation={destinationStation}
          setDestinationStation={setDestinationStation}
          currentBuses={currentBuses}
          sourceStationRef={sourceStationRef}
          destinationStationRef={destinationStationRef}
        />
      )}
      <div
        className={`side-bar ${expanded ? "expanded-view" : "shrink-view"}  `}
      >
       <div className="top-bar">
       <button className="expand-shrink-btn" onClick={expandShrinkView}>
          {expanded ? 
          <img className="translate-img" src={backIcon}></img>
          : 
          <img src={backIcon}></img>}
        </button>

        <button className={`reset-selection-btn`} onClick={resetSelections}>
          <img src={resetIcon}></img>
        </button>
       </div>

        <div className="side-bar-content">
        <div className="input-div">
          <label>Source</label>
          <select
            id="source"
            value={sourceStation}
            onChange={(e) => handleSourceChange(e.target.value)}
          >
            <option value={null}>NOT SELECTED</option>
            {stations.map((station) => (
              <option key={station.id} value={station.id}>
                {station.name}
              </option>
            ))}
          </select>
        </div>

        <div className="input-div">
          <label>Destination</label>
          <select
            id="destination"
            value={destinationStation}
            onChange={(e) => handleDestinationChange(e.target.value)}
          >
            <option key={null} value={null}>
              NOT SELECTED
            </option>
            {stations.map((station) => (
              <option key={station.id} value={station.id}>
                {station.name}
              </option>
            ))}
          </select>
        </div>
       
        {currentStationRoutes && currentStationRoutes.length > 0 && (
          <div className="about-bus-info">
            Click on particular bus info to get live location of it to be
            displayed on map.
          </div>
        )}
        <div id="route-info-wrapper">
          {/* <h2>Routes Available</h2> */}
          <table id="route-info">
            <thead>
              <tr>
                <th>ROUTE</th>
                <th>SOURCE</th>
                <th>DESTINATION</th>
                <th>DEPARTURE TIME</th>
              </tr>
            </thead>
            <tbody>
              {!currentStationRoutes && <div>Stations Data Not Loaded</div>}
              {currentStationRoutes &&
                currentStationRoutes.map((currentStationRoute, index) => (
                  <tr
                    key={index}
                    onClick={() => selectBus(currentStationRoute)}
                  >
                    <td>{currentStationRoute.route_name}</td>
                    <td>{currentStationRoute.sourceName}</td>
                    <td>{currentStationRoute.destinationName}</td>
                    <td>{currentStationRoute.departure_time}</td>
                  </tr>
                ))}
            </tbody>
          </table>
          <div className="current-station-error">
            {currentStationRoutesError}
          </div>
        </div>
        </div>
      </div>
    </div>
  );
};

export default ClientMap;
