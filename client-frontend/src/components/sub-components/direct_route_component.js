
import gone from "../../images/gone.png";
import live from "../../images/live.png";
import notGone from "../../images/notGone.png";

export function DirectRoutes({currentStationRoutes, selectBus,currentStationRoutesError}) {
  return (
    <div id="route-info-wrapper">
      <center>
        <h1 className="text-purple-600 text-xl text-center rounded-sm">
          Bus Routes
        </h1>
        <hr className="my-2 bg-purple-200"></hr>
        <table id="route-info">
          <tbody>
            {!currentStationRoutes && <div>Stations Data Not Loaded</div>}
            {currentStationRoutes &&
              currentStationRoutes.map((currentStationRoute, index) => (
                <tr key={index} onClick={() => selectBus(currentStationRoute)}>
                  <td>
                    {currentStationRoute.status == 2 ? (
                      <img src={live} width="20px" height="20px" />
                    ) : !currentStationRoute.status == 1 ? (
                      <img src={gone} width="20px" height="20px" />
                    ) : (
                      <img src={notGone} width="20px" height="20px" />
                    )}
                  </td>
                  <td className="text-lg w-4/12">
                    {currentStationRoute.sourceName}
                  </td>
                  <td>
                    <center>
                      <div className="w-6 h-6 bg-purple-800 text-white p-1 my-1 rounded-md text-xs">
                        {currentStationRoute.route_name}
                      </div>
                      <div className="text-purple-800 text-xs">
                        Schedule At {currentStationRoute.departure_time}
                      </div>
                    </center>
                  </td>
                  <td className="text-lg w-4/12">
                    {currentStationRoute.destinationName}
                  </td>
                </tr>
              ))}
          </tbody>
        </table>
        <div className="current-station-error">{currentStationRoutesError}</div>
      </center>
    </div>
  );
}
