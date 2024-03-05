package database

import (
	"busproject/model"
	"database/sql"
	"log"
)

func InsertRouteStation(db *sql.DB, routeStation model.RouteStation) error {
	sqlStatement := `INSERT INTO transport.routestations (route_id, station_id, station_order) VALUES ($1, $2, $3)`

	_, err := db.Exec(sqlStatement, routeStation.RouteId, routeStation.StationId, routeStation.StationOrder)
	if err != nil {
		return err
	}

	log.Println("RouteStation inserted successfully")
	return nil
}

func InsertAllRouteStation(db *sql.DB, sqlStatement string) error {

	log.Println(sqlStatement)
	_, err := db.Exec(sqlStatement[:len(sqlStatement)-1] + ";")
	if err != nil {
		return err
	}

	log.Println("All RouteStation inserted successfully")
	return nil
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
