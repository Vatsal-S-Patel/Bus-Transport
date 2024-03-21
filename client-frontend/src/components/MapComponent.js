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
import {
  getBusStyle,
  getCurrentLocationStyle,
  getHighlightedStyle,
  getNormalStyle,
} from "./sub-components/MapStyles";
import { SetupForMap } from "./sub-components/settingup_inital_map";

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
    function selectSourceDestinationMarker(isSource) {
      if (markerSourceRef.current) {
        markerSourceRef.current.getFeatures().forEach((feature) => {
          let stationId = feature.get("id");
          const stationName = feature.get("name");

          if(feature.getId() == "userLocation"){
              // console.log("user lcoation")
          } else {
            if (
              isSource
                ? stationId === Number(sourceStation)
                : stationId === Number(destinationStation)
            ) {
              if (isSource) {
                setSourceStation(stationId);
                isSourceSelected.current = true;
                sourceStationRef.current = feature;
              } else {
                setDestinationStation(stationId);
                isDestinationSelected.current = true;
                destinationStationRef.current = feature;
              }
  
              // Feature related styles
              feature.set("isSelected", true);
              feature.setStyle(
                getHighlightedStyle({
                  stationId: stationId,
                  stationName: stationName,
                  showText: true,
                })
              );
            } else if (
              isSource
                ? destinationStation && stationId == Number(destinationStation)
                : sourceStation && stationId == Number(sourceStation)
            ) {
              // Fo nothing
            } else {
              feature.set("isSelected", true);
              feature.setStyle(
                getNormalStyle({
                  stationId: stationId,
                  stationName: stationName,
                  showText: false,
                })
              );
            }
          }
        });

        if (isSource) {
          document.getElementById("source").style.border = "1px solid blue";
        } else {
          document.getElementById("destination").style.border =
            "1px solid blue";
        }
      }
    }

    if (sourceStation == null) {
      document.getElementById("source").style.border = "none";

      // If source station ref here then remove its style
      if (sourceStationRef.current) {
        console.log("style changed");
        let stationId = sourceStationRef.current.get("id");
        let stationName = sourceStationRef.current.get("name");
        sourceStationRef.current.setStyle(
          getNormalStyle({ stationId, stationName, showText: false })
        );
      }
    }

    if (destinationStation == null) {
      document.getElementById("destination").style.border = "none";
      // If destination station ref here then remove its style
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
      selectSourceDestinationMarker(true);
    }
    if (destinationStation != null) {
      selectSourceDestinationMarker(false);
      // get feature from features list and highlight it.
    }
  }, [sourceStation, destinationStation]);

  /* UseEffects */
  useEffect(() => {
    // current bus array updated
    // console.log("current bus array updated");
    /* Bus features */
    const busFeatures = currentBuses.map((busInfo) => {
      console.log(busInfo);
      let latitude = busInfo.lat;
      let longitude = busInfo.long;

      const busFeature = new Feature({
        geometry: new Point(fromLonLat([longitude, latitude])),
      });

      busFeature.setId(busInfo.bus_id);
      busFeature.set("lat", latitude);
      busFeature.set("long", longitude);
      busFeature.set("type", "bus");
      busFeature.setStyle(
        getBusStyle(
          busInfo.route_name,
          busInfo.status,
          busInfo.last_station_name
        )
      );
      return busFeature;
    });

    const busSource = new VectorSource({
      features: busFeatures,
    });

    // create bus layer if not exist
    if (busLayerRef.current) {
      // Already layer he then just update map reference and its source
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

      // console.log(typeof stationId)

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

      const busFeature = new Feature({
        geometry: new Point(fromLonLat([longitude, latitude])),
      });

      busFeature.setId(busInfo.bus_id);
      busFeature.set("lat", latitude);
      busFeature.set("long", longitude);
      busFeature.set("type", "bus");
      busFeature.setStyle(getBusStyle(busInfo.route_name, busInfo.status));
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

    SetupForMap(initialMap,mapRef,mapView,markerSource,transform,isSourceSelected,setSourceStation,sourceStationRef,isDestinationSelected,setDestinationStation,destinationStationRef,markerLayer)

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
