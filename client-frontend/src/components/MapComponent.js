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
import "./MapComponent.css";
import { getBusStyle, getCurrentLocationStyle, getHighlightedStyle, getNormalStyle } from "./MapStyles";


const MapComponent = ({
  stations,
  sourceStation,
  setSourceStation,
  destinationStation,
  setDestinationStation,
  currentBuses,
  sourceStationRef,
  destinationStationRef,
}) => {

  const [latitude, setLatitude] = useState(null);
  const [longitude, setLongitude] = useState(null);
  const [error, setError] = useState(null);
  const [map, setMap] = useState();
  const mapElement = useRef();
  const mapRef = useRef();
  const busLayerRef = useRef();
  const markerSourceRef = useRef();

  const [center, setCenter] = useState(
    fromLonLat([72.64360178036434, 23.058133475294028])
  );

  const stationsFeatureSource = useRef();
  const isSourceSelected = useRef();
  const isDestinationSelected = useRef();

  /* UseEffects */
  useEffect(() => {

    if (sourceStation == null) {
      document.getElementById("source").style.border = "none";

      // if source station ref here then remove its style
      if (sourceStationRef.current) {
        let stationId = sourceStationRef.current.get("id");
        let stationName = sourceStationRef.current.get("name");
        sourceStationRef.current.setStyle(
          getNormalStyle({ stationId, stationName, showText: false })
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
          getNormalStyle({ stationId, stationName, showText: false })
        );
      }
    }
    // if someone selects sourceStation then update it on map
    if (sourceStation != null) {
      // ye ya fir destination change hua he
      document.getElementById("source").style.border = "1px solid blue";
    }
    if (destinationStation != null) {
      // get feature from features list and highlight it.
      document.getElementById("destination").style.border = "1px solid blue";
    }
  }, [sourceStation, destinationStation]);

  /* UseEffects */
  useEffect(() => {

    // current bus array updated
    console.log("current bus array updated");
    /* Bus features */
    const busFeatures = currentBuses.map((busInfo) => {
      console.log(busInfo)
      let latitude = busInfo.lat;
      let longitude = busInfo.long;

      const busFeature = new Feature({
        geometry: new Point(fromLonLat([longitude, latitude])),
      });

      busFeature.setId(busInfo.bus_id);
      busFeature.set("lat", latitude);
      busFeature.set("long", longitude);
      busFeature.set("type", "bus");
      busFeature.setStyle(getBusStyle(busInfo.route_name,busInfo.status,busInfo.last_station_name));
      return busFeature;
    });

    console.log("bus features now are", busFeatures)

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
      if (mapRef.current) {
        mapRef.current.addLayer(busLayer);
        mapRef.current.render();
      }
    }
    // update live bus features
  }, [currentBuses]);

  /* UseEffects */
  useEffect(() => {
    isSourceSelected.current = false;
    isDestinationSelected.current = false;

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
      view: mapView,
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

    markerSourceRef.current = markerSource;
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

    const markerLayer = new VectorLayer({
      map: map,
      className: "vector-layer",
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
      busFeature.setStyle(getBusStyle(busInfo.route_name,busInfo.status));
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
        } else if (feature.get("type") == "bus") {
          console.log("bus clicked");
        } else {
          const stationName = feature.get("name");
          const stationId = feature.get("id");

          console.log(stationId, stationName);

          if (isSourceSelected.current == false) {

            console.log("first selected")

            setSourceStation(stationId);
            isSourceSelected.current = true;
            sourceStationRef.current = feature;
            feature.set("isSelected", true);
            feature.setStyle(
              getHighlightedStyle({ stationName, stationId, showText: true })
            );
          } else if (isDestinationSelected.current == false) {
            
            console.log("second selected")

            setDestinationStation(stationId);
            feature.set("isSelected", true);
            isDestinationSelected.current = true;
            destinationStationRef.current = feature;
            feature.setStyle(
              getHighlightedStyle({ stationName, stationId, showText: true })
            );
          } else {
            var source = markerLayer.getSource();
            console.log("third selected")

            // source and destination dono features ki style normal kar do
            // let sourceStationFeature = sourceStationRef.current
            // let destinationStationFeature = destinationStation.current
            // sourceStationFeature.setStyle(getHighlightedStyle({stationId: sourceStationFeature.get("id"), stationName: sourceStationFeature.get("name")}))

            // TODO very important method to work with features.
            source.getFeatures().forEach(function (feature) {
              if (
                feature.getId() == "userLocation" ||
                feature.get("type") == "bus"
              ) {
                console.log("user location or bus")
              } else {
                feature.set("isSelected", false);
                feature.setStyle(
                  getNormalStyle({
                    stationName: feature.get("name"),
                    stationId: feature.get("id"),
                    showText: false,
                  })
                );
              }
            });

            isDestinationSelected.current = false;

            setSourceStation(stationId);
            setDestinationStation(null);
            feature.set("isSelected", true);
            sourceStationRef.current = feature;
            destinationStationRef.current = null;
            feature.setStyle(
              getHighlightedStyle({ stationName, stationId, showText: true })
            );
          }
        }
      });
    });

    mapRef.current = initialMap;

    mapView.on("change:resolution", function (event) {
      // Make all the features size small and big according to it.
      let features = markerSource.getFeatures();
      // Now add normal style to it with specified width and all stuff and highlighted style for speciified
      let showText = true;

      if (event.oldValue > 15) {
        showText = false;
      }

      features.forEach((feature) => {
        if(feature.getId() == "userLocation" || feature.getId() == "bus"){
          
        } else {
          let isSelected = feature.get("isSelected");

        if (isSelected) {

        } else {
          feature.setStyle(
            getNormalStyle({
              stationId: feature.get("id"),
              stationName: feature.get("name"),
              showText: showText,
            })
          );
        }
        }
        
      });
    });

    return () => {
      initialMap.setTarget = null;
    };
  }, []);

  return (
    <div>
      <div id="map" ref={mapElement} className="map-container"></div>
    </div>
  );
};

export default MapComponent;