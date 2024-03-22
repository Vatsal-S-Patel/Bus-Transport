import Style from "ol/style/Style";
import Icon from "ol/style/Icon";
import Text from "ol/style/Text";
import Fill from "ol/style/Fill";
import stationIcon from "../../images/bus_image2.png";
import busIcon from "../../images/bus.png";
import currentLocationIcon from "../../current_location.svg";


export function getHighlightedStyle({ stationId, stationName, showText }) {
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

export function getNormalStyle({ stationId, stationName, showText }) {
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

export function getCurrentLocationStyle(user) {
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

 export function getBusStyle(route,status,last) {
  // console.log(route,status,last);
  
  if (last.length > 15 ){
    last = last.substring(0,12)+"..."
  }
    return new Style({
      image: new Icon({
        crossOrigin: "anonymous",
        src: busIcon,
        width: 40,
        height: 40,
      }),
      text: new Text({
        text: `${route}`,
        offsetY: -20,
        fill: new Fill({
          color: "#000",
        }),
        scale: 1.2,
        backgroundFill: new Fill({ color: "white" }),
      }),
    });
  }
