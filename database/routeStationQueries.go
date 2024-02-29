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
