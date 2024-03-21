
import arrow from "../../images/right-arrow.png";

export function IndirectRoutes({specialStationRoutes,getStationInfoById,sourceStation,stationsMap,destinationStation}) {
  return (
    <div id="route-info-wrapper">
      <center>
        <h1 className="text-purple-600 text-xl text-center rounded-sm">
          No Direct Routes Found
        </h1>
        <hr className="my-2 bg-purple-200"></hr>
        <table id="route-info">
          <tbody>
            {specialStationRoutes &&
              specialStationRoutes.map((currentStationRoute, index) => (
                <tr key={index}>
                  <td>
                    <div className="w-10 h-10 bg-purple-800 text-white p-2 m-2 rounded-lg">
                      {currentStationRoute.source_route_name}
                    </div>
                  </td>
                  <td>
                    <div className="text-purple-800 my-2 text-xs">
                      {`${getStationInfoById(sourceStation, stationsMap).name}`}
                    </div>
                  </td>
                  <td>
                    <div className="my-2 mx-1 w-5">
                      <img src={arrow} width="30px" />
                    </div>
                  </td>
                  <td>
                    <div className="text-purple-800 my-2 text-xs font-bold">
                      {`${currentStationRoute.junction_name}`}
                    </div>
                  </td>
                  <td>
                    <div className="my-2 mx-1 w-5">
                      <img src={arrow} width="30px" />
                    </div>
                  </td>
                  <td>
                    <div className="text-purple-800 my-2 text-xs">
                      {`${
                        getStationInfoById(destinationStation, stationsMap).name
                      }`}
                    </div>
                  </td>
                  <td>
                    <div className="w-10 h-10 bg-purple-800 text-white p-2 my-2 rounded-lg">
                      {currentStationRoute.destination_route_name}
                    </div>
                  </td>
                </tr>
              ))}
          </tbody>
        </table>
      </center>
    </div>
  );
}
