import { Feature } from "ol";
import { useEffect, useRef, useState } from "react";
import Map from "ol/Map";
import View from "ol/View";
import TileLayer from "ol/layer/Tile";
import OSM from "ol/source/OSM";
import VectorLayer from "ol/layer/Vector";
import VectorSource from "ol/source/Vector";
import Point from "ol/geom/Point";
import { fromLonLat, transform } from "ol/proj";
import Style from "ol/style/Style";
import Icon from "ol/style/Icon";
import Text from "ol/style/Text";
import Fill from "ol/style/Fill";
import stationIcon from "../images/bus_image2.png";
import busIcon from "../images/image.png";
import currentLocationIcon from "../current_location.svg";
import "./MapComponent.css";
import { Circle, LineString } from "ol/geom";
import Stroke from "ol/style/Stroke";

// mapView.on("change:resolution", function(event){
//   // Make all the features size small and big according to it.

//   // stations
//   let features = markerSource.getFeatures()
  
//   // Now add normal style to it with specified width and all stuff and highlighted style for speciified

//   let showText = true

//   if(event.oldValue > 15){
//     showText = false
//   }
//   features.forEach(feature => {
//     let busId = feature.get("id")
//     console.log(busId, sourceStation)
//     if(Number(busId) == Number(sourceStation) || busId == destinationStation){
//     feature.setStyle(getHighlightedStyle({ stationId: feature.get("id"),stationName:feature.get("name"), showText: showText}))

//     } else
//     feature.setStyle(getNormalStyle({ stationId: feature.get("id"),stationName:feature.get("name"), showText: showText}))
//   })
// })

const buses = [
  {
    id: 1,
  },
  {
    id: 2,
  },
];

const route7d = [
  {
    name: "DCIS Circle",
    id: 30,
    lat: 23.129273,
    long: 72.5819126,
  },
  {
    name: "DCIS Circle Cont 1",
    id: 30,
    lat: 23.126927661766928,
    long: 72.58367768034088,
  },
  {
    name: "DCIS Circle Cont 2",
    id: 30,
    lat: 23.124164927878155,
    long: 72.58297058247635,
  },
  {
    name: "Sarthi Bunglows",
    id: 31,
    lat: 23.1229731,
    long: 72.5778669,
  },
  {
    name: "Sarthi Bunglows Cont 1",
    id: 31,
    lat: 23.1193977234677,
    long: 72.58196886181285,
  },
  {
    name: "Chandkheda Gam",
    id: 32,
    lat: 23.1165397,
    long: 72.5819121,
  },
  {
    name: "Shiv Shaktinagar",
    id: 33,
    lat: 23.1125755,
    long: 72.5838263,
  },
];

const MapComponent = ({
  stations,
  sourceStation,
  setSourceStation,
  destinationStation,
  setDestinationStation,
  routeStations,
  currentBuses,
  sourceStationRef,
  destinationStationRef
}) => {
  const [latitude, setLatitude] = useState(null);
  const [longitude, setLongitude] = useState(null);
  const [error, setError] = useState(null);
  const [map, setMap] = useState();
  const mapElement = useRef();
  const mapRef = useRef();
  const busLayerRef = useRef();
  const markerSourceRef = useRef();

  const [center, setCenter] = useState(fromLonLat([72.64360178036434, 23.058133475294028]));

  const stationsFeatureSource = useRef();
  const isSourceSelected = useRef();
  const isDestinationSelected = useRef();

  // Function to generate route features
  function generateRouteFeatures() {
    const routeFeatures = [];
    for (let i = 0; i < routeStations.length - 1; i++) {
      const startPoint = fromLonLat([
        routeStations[i].long,
        routeStations[i].lat,
      ]);
      const endPoint = fromLonLat([
        routeStations[i + 1].long,
        routeStations[i + 1].lat,
      ]);
      const routeGeometry = new LineString([startPoint, endPoint]);
      const routeFeature = new Feature({
        geometry: routeGeometry,
      });
      routeFeatures.push(routeFeature);
    }
    return routeFeatures;
  }

  // Generate a random integer between min and max (inclusive)
  function getRandomInteger(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
  }

  function getHighlightedStyle({ stationId, stationName, showText}) {
    return new Style({
      image: new Icon({
        crossOrigin: "anonymous",
        src: stationIcon,
        width: 40,
        height: 40,
      }),
      text: new Text({
        text: showText ? stationId + " " + stationName : "",
        offsetY: -30,
        fill: new Fill({
          color: "green",
        }),
        scale: 1.2,
      }),
    });
  }

  function getNormalStyle({ stationId, stationName, showText}) {
    return new Style({
      image: new Icon({
        crossOrigin: "anonymous",
        src: stationIcon,
        width: 22,
        height: 22,
      }),
      text: new Text({
        text: showText ? stationId + " " + stationName : "",
        offsetY: -20,
        fill: new Fill({
          color: "black",
        }),
        scale: 1.1,
      }),
    });
  }

  function getCurrentLocationStyle(user) {
    return new Style({
      image: new Icon({
        crossOrigin: "anonymous",
        src: currentLocationIcon,
        width: 25,
        height: 25,
      }),
      text: new Text({
        text: user,
        offsetY: 20,
        offsetX: 0,
        fill: new Fill({
          color: "#123",
        }),
        scale: 1.1,
      }),
    });
  }

  function getBusStyle(lastUpdated) {
    return new Style({
      image: new Icon({
        // color: "#8959A8",
        crossOrigin: "anonymous",
        src: busIcon,
        width: 40,
        height: 40,
      }),
      text: new Text({
        text: `Last Updated ${lastUpdated}`,
        offsetY: -20,
        // fill: new Fill({
        //   color: "#000",
        // }),
        scale: 1.2,
      }),
    });
  }

  function getData() {
    console.log(mapElement.current);
  }

  const updateOneBusLocation = (busIndex, busFeatures, lat, long) => {
    console.log("update bus loc ", busFeatures, lat, long);
    let busFeature = busFeatures[busIndex];
    const newCoordinates = fromLonLat([lat, long]);

    busFeature.getGeometry().setCoordinates(newCoordinates);

    // trigger map redraw ha ha// IMPORTANT
    console.log(mapRef.current);
    if (mapRef.current) {
      mapRef.current.render();
    }
  };

  useEffect(() => {
    // as source station or destination station change change the feature
    console.log("source and destination changed", sourceStation, destinationStation);
    console.log("references are", sourceStationRef.current, destinationStationRef.current)

    if (sourceStation == null) {
      document.getElementById("source").style.border = "none";

      // if source station ref here then remove its style
      if (sourceStationRef.current) {
        let stationId = sourceStationRef.current.get("id");
        let stationName = sourceStationRef.current.get("name");
        sourceStationRef.current.setStyle(
          getNormalStyle({ stationId, stationName })
        );
      }
    }

    if (destinationStation == null) {
      document.getElementById("destination").style.border = "none";

      // if destination station ref here then remove its style
      if (destinationStationRef.current) {
        let stationId = destinationStationRef.current.get("id");
        let stationName = destinationStationRef.current.get("name");
        destinationStationRef.current.setStyle(
          getNormalStyle({ stationId, stationName })
        );
      }
    }

    // if someone selects sourceStation then update it on map
    if (sourceStation != null) {
      // ye ya fir destination change hua he
      // also highlight that particular feature
    // markerSource.getFeatureById(1).setst
      // if(markerSourceRef.current){
      //   markerSourceRef.current.getFeatureById(String(sourceStation)).setStyle(getNormalStyle({stationId: sourceStation, stationName: "great"}))
      // }
      document.getElementById("source").style.border = "1px solid blue";
    }
    if (destinationStation != null) {
      // get feature from features list and highlight it.
      document.getElementById("destination").style.border = "1px solid blue";
    }
  }, [sourceStation, destinationStation]);

  useEffect(() => {
    // current bus array updated
    console.log("current bus array updated");

    /* Bus features */
    const busFeatures = currentBuses.map((busInfo) => {
      let latitude = busInfo.lat;
      let longitude = busInfo.long;

      console.log(latitude, longitude);

      const busFeature = new Feature({
        geometry: new Point(fromLonLat([longitude, latitude])),
      });

      busFeature.setId(busInfo.bus_id);
      busFeature.set("lat", latitude);
      busFeature.set("long", longitude);
      busFeature.setStyle(getBusStyle(busInfo.last_updated));
      return busFeature;
    });

    const busSource = new VectorSource({
      features: busFeatures,
    });

    // create bus layer if not exist

    if (busLayerRef.current) {
      // already layer he then just update map reference and its source
      busLayerRef.current.setSource(busSource);

      // then render it.
      if (mapRef.current) {
        mapRef.current.render();
      }
    } else {
      const busLayer = new VectorLayer({
        source: busSource,
        className: "vector-layer",
      });

      busLayerRef.current = busLayer;

      // busLayer.getSource().getFeatureById(1)

      if (mapRef.current) {
        mapRef.current.addLayer(busLayer);
        mapRef.current.render();
      }
    }

    // update live bus features
  }, [currentBuses]);

  useEffect(() => {
    const mapView = new View({
      center: center,
      zoom: 12.5,
    });
    const initialMap = new Map({
      target: mapElement.current,
      layers: [
        new TileLayer({
          source: new OSM(),
          className: "tile-layer",
        }),
      ],
      view: mapView
    });

    const markerFeatures = stations.map((coord) => {
      const feature = new Feature({
        geometry: new Point(fromLonLat([coord.long, coord.lat])),
      });

      let stationName = coord.name;
      let stationId = coord.id;

      feature.set("id", stationId);
      feature.set("name", stationName);
      feature.set("isSelected", false);
      initialMap.set("isDestination", false);
      initialMap.set("isSource", false);

      feature.setStyle(getNormalStyle({ stationId, stationName }));

      return feature;
    });

    const markerSource = new VectorSource({
      features: markerFeatures,
    });

    markerSourceRef.current = markerSource

    // markerSource.getFeatureById(1).setst

    // get current location of user
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          console.log(position);
          setLatitude(position.coords.latitude);
          setLongitude(position.coords.longitude);

          const userLocationFeature = new Feature({
            geometry: new Point(
              fromLonLat([position.coords.longitude, position.coords.latitude])
            ),
          });

          userLocationFeature.setId("userLocation");
          userLocationFeature.setStyle(getCurrentLocationStyle("You"));

          markerSource.addFeature(userLocationFeature);
        },
        (error) => {
          setError(error.message);
        }
      );
    } else {
      alert("give me geolocation");
      setError("Geolocation is not supported by this browser.");
    }

    // const routeFeatures = generateRouteFeatures(stations);
    // const routeSource = new VectorSource({
    //   features: routeFeatures,
    // });
    // const routeLayer = new VectorLayer({
    //   source: routeSource,
    //   style: new Style({
    //     stroke: new Stroke({
    //       color: "#FF5733",
    //       width: 3,
    //     }),
    //   }),
    // });

    // initialMap.addLayer(routeLayer);

  

    const markerLayer = new VectorLayer({
      map: map,
      className: "vector-layer"
    });



    initialMap.addLayer(markerLayer);

    markerLayer.setSource(markerSource);

    /* Bus features */
    const busFeatures = currentBuses.map((busInfo) => {
      let latitude = busInfo.lat;
      let longitude = busInfo.long;

      console.log(latitude, longitude);

      const busFeature = new Feature({
        geometry: new Point(fromLonLat([longitude, latitude])),
      });

      busFeature.setId(busInfo.bus_id);
      busFeature.set("lat", latitude);
      busFeature.set("long", longitude);
      busFeature.set("type", "bus");
      busFeature.setStyle(getBusStyle("bus"));
      return busFeature;
    });

    const busSource = new VectorSource({
      features: busFeatures,
    });

    stationsFeatureSource.current = busSource;

    // // create new layer and add it
    const busLayer = new VectorLayer({
      source: busSource,
      className: "vector-layer",
    });

    busLayerRef.current = busLayer;

    initialMap.addLayer(busLayer);

    setMap(initialMap);

    // DummyBus(0 ,busFeatures)
    // DummyBus(1 ,busFeatures)

    initialMap.on("click", function (event) {
      console.log(
        "coordinates clicked ",
        transform(event.coordinate, "EPSG:3857", "EPSG:4326")
      );

      initialMap.forEachFeatureAtPixel(event.pixel, function (feature) {
        if (feature.getId() == "userLocation") {
          console.log("got user location clicked");
        } else if(feature.get("type") == "bus"){
          console.log("bus clicked")
        } else {
          const stationName = feature.get("name");
          const stationId = feature.get("id");
          const isSelected = feature.get("isSelected");
          // const isDestination = initialMap.get("isDestination");
          // const isSource = initialMap.get("isSource");

          console.log(stationId, stationName);

          if (isSourceSelected.current == false) {
            setSourceStation(stationId);
            isSourceSelected.current = true
            sourceStationRef.current = feature;
            feature.set("isSelected",true)
            feature.setStyle(getHighlightedStyle(stationName));
          } else if (isDestinationSelected.current == false) {
            setDestinationStation(stationId);
            feature.set("isSelected",true)
            isDestinationSelected.current = true
            destinationStationRef.current = feature;
            feature.setStyle(getHighlightedStyle(stationName));
          } else {
            var source = markerLayer.getSource();


            // source and destination dono features ki style normal kar do
            // let sourceStationFeature = sourceStationRef.current
            // let destinationStationFeature = destinationStation.current
            // sourceStationFeature.setStyle(getHighlightedStyle({stationId: sourceStationFeature.get("id"), stationName: sourceStationFeature.get("name")}))

            // TODO very important method to work with features.
            source.getFeatures().forEach(function (feature) {
              if(feature.get("id") != "userLocation" && feature.get("type") != "bus"){
                feature.setStyle(
                  getNormalStyle({
                    stationName: feature.get("name"),
                    stationId: feature.get("id"),
                  })
                );
              }
            });

            isDestinationSelected.current = false

            // because of bellow one
            setSourceStation(stationId);
            setDestinationStation(null);

            sourceStationRef.current = feature;
            destinationStationRef.current = null;
            feature.setStyle(getHighlightedStyle(stationName));
          }
        }
      });
    }) 
    
    mapRef.current = initialMap;

    return () => {
      initialMap.setTarget = null;
    };
  }, []);

  return (
    <div>
      {console.log(sourceStation, destinationStation)}
      {/* <button onClick={getData}>{sourceStation}</button> */}

      <div id="map" ref={mapElement} className="map-container"></div>
    </div>
  );
};

export default MapComponent;
