import React, { useEffect, useState, useRef } from "react";
import "ol/ol.css";
import "./App.css";
import MapComponent from "./components/MapComponent.js";
import { Manager } from "socket.io-client";

const IP = "192.168.6.107:8080";

const route_data = [
  {
    bus_route: "1",
    start_stop: "30",
    stop_order: "1",
    stop_name: "DCIS Circle",
    lat: "23.129273",
    long: "72.5819126",
  },
  {
    bus_route: "1",
    start_stop: "31",
    stop_order: "2",
    stop_name: "Sarthi Bunglows",
    lat: "23.1229731",
    long: "72.5778669",
  },
];

/* Initially get all stations information then no need to bother api about this. */
function getStationInfoById(stationId, stationsMap) {
  // see if id present then give name
  if (stationsMap.has(Number(stationId))) {
    return stationsMap.get(Number(stationId));
  }

  return {
    id: Number(stationId),
    name: "not available",
  };
}

const MapApp = () => {
  // all stations information
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

  // initially set it false
  const [expanded, setExpanded] = useState(false);

  // Event handler for source station select
  const handleSourceChange = (newValue) => {
    setSourceStation(newValue);
  };

  // Event handler for destination station select
  const handleDestinationChange = (newValue) => {
    setDestinationStation(newValue);
  };
  const showAllBuses = () => {
    socketConn.emit("sourceSelected", [1,2,3,4,5,6,7,8,9,10,11,12,13,14])
  };
  const selectBus = (busInfo) => {
    console.log(busInfo)
    console.log("SOCKET EMIT EVENT")
    setCurrentBuses([])
    socketConn.emit("busSelected", busInfo.bus_id)
  };

  const resetSelections = () => {
    console.log("reset");
    setSourceStation(null);
    setDestinationStation(null);

    socketConn.emit("sourceSelected", [])

    // //

    // sourceStationRef.current = null
    // destinationStationRef.current = null

    document.getElementById("source").selectedIndex = 0;
    document.getElementById("destination").selectedIndex = 0;

    // jo data uski vajah se aaya he use bhi nikal do
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

    let newArray = currentStationRoutes
    // as this data changes see if there is any update or not if there then update cuurrent routes
    currentBuses.forEach(busData => {
      newArray  = newArray.map(routeInfo => routeInfo.bus_id == busData.bus_id ? {...routeInfo, lat: busData.lat, long: busData.long, last_station_order: busData.last_station_order}: routeInfo)
    })

    setCurrentStationRoutes(newArray)
  }, [currentBuses])
  

  // fetch route infomation when sourcestation adn destination station changes
  useEffect(() => {
    // if(sourceStation == null){

    // }

    // if(destinationStation == null){

    // }

    if (sourceStation != null) {
      let data = {};
      data["source"] = Number(sourceStation);
      if (destinationStation != null) {
        data["destination"] = Number(destinationStation);
      }

      // Reset Station Routes
      setCurrentStationRoutes([]);
      setCurrentStationRoutesError(null);

      fetch(`http://${IP}/api/schedule/GetUpcomingBus`, {
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
            let updatedData = routeStations.map((info) => {
              return {
                ...info,
                departure_time: info.departure_time.substr(11, 5),
                sourceName: getStationInfoById(info.source, stationsMap).name,
                destinationName: getStationInfoById(
                  info.destination,
                  stationsMap
                ).name,
                last_updated: new Date()
              };
            });

            // if it has lat long then add this bus to current buses
            // ig let all the buses have lat long
            setCurrentBuses(updatedData)

            setCurrentStationRoutes(updatedData);

            let routes = new Set();
            updatedData.forEach((element) => {
              routes.add(element.route_id);
            });

            console.log(routes);
            console.log(updatedData);

            let routesArray = [];
            routes.forEach((route) => routesArray.push(route));

            // Now set data to socket
            console.log("SOCKET EMITTED");
            socketConn.emit("sourceSelected", routesArray);
          }
        })
        .catch((error) => {
          console.error("Error fetching stations data:", error);
        });
    }
  }, [sourceStation, destinationStation, stations]);

  useEffect(() => {
    fetch(`http://${IP}/api/station/`)
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        let stations = data.Data;
        setStations(stations);

        // Put stations data in map for later use.
        let mp = new Map();

        stations.forEach((station) => {
          mp.set(station.id, station);
        });

        setStationsMap(mp);
      })
      .catch((error) => {
        console.error("Error fetching stations data:", error);
      });

    // connect with socket connections
    const manager = new Manager("ws://192.168.6.107:8080", {
      reconnectionDelayMax: 100000,
      transports: ["polling"],
    });

    const socket = manager.socket("/", { transports: ["polling", "websocket"] });

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

    socket.on("update", (busInfo) => {
      console.log("UPDATE", busInfo);

      // check already array me he ki nahi
      let isAlreadyBusDisplayed = false;
      currentBuses.forEach((bus) => {
        if (bus.bus_id == busInfo.bus_id) {
          isAlreadyBusDisplayed = true;
        }
      });

      if (isAlreadyBusDisplayed) {
        // update this information
        let newArray = currentBuses.map((bus) =>
          bus.bus_id == busInfo.bus_id
            ? {
                ...busInfo,
              }
            : bus
        );
        setCurrentBuses(newArray);

        // agar ye bus current routes me he then updated,
      } else {
        currentBuses.push(busInfo);
      }

      console.log(currentBuses);
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
          routeStations={route_data}
          currentBuses={currentBuses}
          sourceStationRef={sourceStationRef}
          destinationStationRef={destinationStationRef}
        />
      )}
      <div
        className={`side-bar ${expanded ? "expanded-view" : "shrink-view"}  `}
      >
        <button className="expand-shrink-btn" onClick={expandShrinkView}>
          {expanded ? "SHRINK VIEW >>>>>" : " <<<<< EXPAND VIEW"}
        </button>

        <h1 className="heading">Find Bus</h1>
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
            <option key={null} value={null}>NOT SELECTED</option>
            {stations.map((station) => (
              <option key={station.id} value={station.id}>
                {station.name}
              </option>
            ))}
          </select>
        </div>
        <button className="reset-selection-btn" onClick={resetSelections}>
          RESET SELECTIONS
        </button>
        <button className="reset-selection-btn" onClick={showAllBuses}>
          SHOW ALL BUSES
        </button>

        {currentStationRoutes && currentStationRoutes.length > 0 && (
          <div>
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
                <th>LAST STATION</th>
                <th>LAST UPDATED LOCATION</th>
              </tr>
            </thead>
            <tbody>
              {!currentStationRoutes && <div>Stations Data Not Loaded</div>}
              {currentStationRoutes &&
                currentStationRoutes.map((currentStationRoute, index) => (
                  <tr key={index} onClick={()=>selectBus(currentStationRoute)}>
                    {/* <td>{currentStationRoute.id}</td> */}
                    <td>{currentStationRoute.route_name}</td>
                    <td>{currentStationRoute.sourceName}</td>
                    <td>{currentStationRoute.destinationName}</td>
                    <td>{currentStationRoute.departure_time}</td>
                    <td>{getStationInfoById(currentStationRoute.last_station_order, stationsMap).name}</td>
                    <td>{currentStationRoute.lat},{currentStationRoute.long}</td>
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
  );
};

export default MapApp;
