import { getHighlightedStyle, getNormalStyle } from "./MapStyles";

export function SetupForMap(initialMap,mapRef,mapView,markerSource,transform,isSourceSelected,setSourceStation,sourceStationRef,isDestinationSelected,setDestinationStation,destinationStationRef,markerLayer){

    initialMap.on("click", function (event) {
        console.log(
          "coordinates clicked ",
          transform(event.coordinate, "EPSG:3857", "EPSG:4326")
        );
  
        initialMap.forEachFeatureAtPixel(event.pixel, function (feature) {
          if (feature.getId() == "userLocation") {
          } else if (feature.get("type") == "bus") {
          } else {
            const stationName = feature.get("name");
            const stationId = feature.get("id");
  
            if (isSourceSelected.current == false) {
              setSourceStation(stationId);
              isSourceSelected.current = true;
              sourceStationRef.current = feature;
              feature.set("isSelected", true);
              feature.setStyle(
                getHighlightedStyle({ stationName, stationId, showText: true })
              );
            } else if (isDestinationSelected.current == false) {
              setDestinationStation(stationId);
              feature.set("isSelected", true);
              isDestinationSelected.current = true;
              destinationStationRef.current = feature;
              feature.setStyle(
                getHighlightedStyle({ stationName, stationId, showText: true })
              );
            } else {
              var source = markerLayer.getSource();
              console.log("third selected");
  
              source.getFeatures().forEach(function (feature) {
                if (
                  feature.getId() == "userLocation" ||
                  feature.get("type") == "bus"
                ) {
                  console.log("user location or bus");
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
          if (feature.getId() == "userLocation" || feature.getId() == "bus") {
          } else {
            let isSelected = feature.get("isSelected");
  
            if (!isSelected) {
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
}