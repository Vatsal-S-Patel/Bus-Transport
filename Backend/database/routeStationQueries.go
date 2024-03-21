package database

import (
	"busproject/model"
	"database/sql"
	"log"
)

func InsertRouteStation(db *sql.DB, routeStation model.RouteStation) error {
	sqlStatement,err :=db.Prepare(`INSERT INTO transport.routestations (route_id, station_id, station_order) VALUES ($1, $2, $3)`)
	if err != nil {
		return err
	}
	defer sqlStatement.Close()

	_, err = sqlStatement.Exec(routeStation.RouteId, routeStation.StationId, routeStation.StationOrder)
	return err
}

func InsertAllRouteStation(db *sql.DB, sqlStatement string) error {

	_, err := db.Exec(sqlStatement[:len(sqlStatement)-1] + ";")
	return err
}

func GetAllRouteStation(db *sql.DB) ([]model.RouteStation, error) {
	sqlStatement := `SELECT * FROM transport.routestations;`

	res, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var routeStations []model.RouteStation
	for res.Next() {
		var routeStation model.RouteStation

		err := res.Scan(&routeStation.RouteId, &routeStation.StationId, &routeStation.StationOrder)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		
		routeStations = append(routeStations, routeStation)
	}

	return routeStations, nil
}
