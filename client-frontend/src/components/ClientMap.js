import React, { useEffect, useState, useRef } from "react";
import "ol/ol.css";
import "../App.js";
import MapComponent from "./MapComponent.js";
import resetIcon from "../images/reset-svgrepo-com.svg";
import backIcon from "../images/backward-3-svgrepo-com.svg";
import IP from "../IP.js";
import { IndirectRoutes } from "./sub-components/indirect_route_component.js";
import { DirectRoutes } from "./sub-components/direct_route_component.js";
import {
  FetchStations,
  fetchAllRoutes,
} from "./sub-components/fetch_routes_api.js";

/* Initially get all stations information then no need to bother api about this. */
function getStationInfoById(stationId, stationsMap) {
  // See If ID present then give name
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
  const [specialStationRoutes, setSpecialStationRoutes] = useState([]);

  const [currentBuses, setCurrentBuses] = useState([]);
  // References to source and destination states to Re-Render Map
  const sourceStationRef = useRef();
  const destinationStationRef = useRef();

  // Initially set it false
  const [expanded, setExpanded] = useState(false);
  var socket = window._DEFAULT_DATA;


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
    socket.emit("busSelected", busInfo.bus_id);
  };

  const resetSelections = () => {
    let btn = document.querySelector(".reset-selection-btn");
    btn.style.rotate = "360deg";

    setTimeout(() => {
      btn.style.rotate = "0deg";
    }, 1000);

    setSourceStation(null);
    setDestinationStation(null);

    socket.emit("sourceSelected", []);

    document.getElementById("source").selectedIndex = 0;
    document.getElementById("destination").selectedIndex = 0;

    setCurrentBuses([]);
    setCurrentStationRoutes([]);
    setSpecialStationRoutes([]);
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
    // As this data changes see if there is any update or not if there then update cuurrent routes
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
      console.log("UPDATE CAME", busInfo);
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
      } else {
        // Append this information
        setCurrentBuses([...currentBuses, busInfo]);
      }
    });

    setCurrentStationRoutes(newArray);
  }, [currentBuses]);

  // fetch route infomation when sourcestation adn destination station changes
  useEffect(() => {
    if (sourceStation != null) {
      let data = {};
      data["source"] = Number(sourceStation);
      if (destinationStation != null) {
        data["destination"] = Number(destinationStation);
      }
      // Reset Station Routes
      setCurrentStationRoutes([]);
      setCurrentStationRoutesError(null);

      fetchAllRoutes(
        data,
        destinationStation,
        getStationInfoById,
        stationsMap,
        setCurrentStationRoutes,
        setSpecialStationRoutes,
        setCurrentStationRoutesError,
        socket
      );
    }
    setCurrentBuses([]);
  }, [sourceStation, destinationStation, stations]);

  useEffect(() => {

    FetchStations(setStations, setStationsMap,socket);
    socket.on("roomJoined", (data) => {
      // console.log("Room joined", data);
    });
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
            {expanded ? (
              <img className="translate-img" src={backIcon}></img>
            ) : (
              <img src={backIcon}></img>
            )}
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
              value={sourceStation != null ? sourceStation : ""}
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
              value={destinationStation != null ? destinationStation : ""}
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

          {currentStationRoutes.length != 0 ? (
            <DirectRoutes
              currentStationRoutes={currentStationRoutes}
              selectBus={selectBus}
              currentStationRoutesError={currentStationRoutesError}
            />
          ) : (
            <IndirectRoutes
              specialStationRoutes={specialStationRoutes}
              getStationInfoById={getStationInfoById}
              sourceStation={sourceStation}
              stationsMap={stationsMap}
              destinationStation={destinationStation}
            />
          )}
        </div>
      </div>
    </div>
  );
};

export default ClientMap;
