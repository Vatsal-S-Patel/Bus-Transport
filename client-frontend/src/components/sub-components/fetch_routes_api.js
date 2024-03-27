import IP from "../../IP";

export function FetchStations(setStations,setStationsMap,socket) {
  // Function to fetch all the stations and set them to states
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
}

export function fetchAllRoutes(
  data,
  destinationStation,
  getStationInfoById,
  stationsMap,
  setCurrentStationRoutes,
  setSpecialStationRoutes,
  setCurrentStationRoutesError,
  socket
) {
  fetch(`http://${IP}:8080/api/schedule/GetUpcomingBus`, {
    method: "POST",
    body: JSON.stringify(data),
    type: "application/json",
  })
    .then((response) => response.json())
    .then((d) => {
      let routeStations = d.Data;
      if (d.Code == 404) {
        if (destinationStation != null) {
          console.log(data);
          fetch(`http://${IP}:8080/api/schedule/GetUpcomingSpecialBus`, {
            method: "POST",
            body: JSON.stringify(data),
            type: "application/json",
          })
            .then((response) => response.json())
            .then((data) => {
              if (data.Data == null) {
                return;
              }

              const uniqueElements = new Set(
                data.Data.map((obj) => obj.source_name)
              );
              const uniqueArray = Array.from(uniqueElements).map((id) =>
                data.Data.find((obj) => obj.source_name === id)
              );

              setSpecialStationRoutes(data.Data);
            });
        } else {
          setCurrentStationRoutesError("No bus found");
        }
      } else if (data.Code == 500) {
        setCurrentStationRoutesError(data.Message);

        console.log("SOCKET EMITTED");
        socket.emit("sourceSelected", []);
      } else {
        let time = new Date();
        let updatedData = routeStations.map((info) => {
          return {
            ...info,
            departure_time: info.departure_time.substr(11, 5),
            sourceName: getStationInfoById(info.source, stationsMap).name,
            destinationName: getStationInfoById(info.destination, stationsMap)
              .name,
            last_updated: time.getHours() + ":" + time.getMinutes(),
          };
        });

        const now = new Date().getHours();
        var upcomingBusSchedule = [];
        var scheduledBuses = [];

        for (let i = 0; i < updatedData.length; i++) {
          const [hours1, minutes] = updatedData[i].departure_time.split(":");
          var a = Number(hours1);

          if (updatedData[i].status == 1) {
            updatedData[i].status = 2;
            upcomingBusSchedule.push(updatedData[i]);
            continue;
          }
          if (a > Number(now)) {
            updatedData[i].status = 1;
            upcomingBusSchedule.push(updatedData[i]);
          } else {
            updatedData[i].status = 0;
            scheduledBuses.push(updatedData[i]);
          }
        }

        setCurrentStationRoutes([...upcomingBusSchedule, ...scheduledBuses]);

        let routes = new Set();
        updatedData.forEach((element) => {
          routes.add(element.route_id);
        });

        let routesArray = [];
        routes.forEach((route) => routesArray.push(route));

        socket.emit("sourceSelected", routesArray);
      }
    })
    .catch((error) => {
      console.error("Error fetching stations data:", error);
    });
}
