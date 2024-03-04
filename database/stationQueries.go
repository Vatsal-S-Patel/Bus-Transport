package database

import (
	"busproject/model"
	"database/sql"
	"log"
	"strconv"
)

func InsertStation(db *sql.DB, station model.Station) error {
	sqlStatement := `INSERT INTO transport.station (id, name, lat, long) VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(sqlStatement, station.Id, station.Name, station.Lat, station.Long)
	if err != nil {
		return err
	}

	log.Println("Station inserted successfully")
	return nil
}

func GetAllStation(db *sql.DB) ([]model.Station, error) {
	sqlStatement := `SELECT * FROM transport.station ORDER BY name ASC;`

	res, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var stations []model.Station
	for res.Next() {
		var station model.Station
		err := res.Scan(&station.Id, &station.Name, &station.Lat, &station.Long)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		stations = append(stations, station)
	}

	return stations, nil
}

func DeleteStation(db *sql.DB, id string) error {
	sqlStatement := `DELETE FROM transport.station WHERE id=$1`

	newId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = db.Exec(sqlStatement, newId)
	if err != nil {
		return err
	}

	log.Println("Station Deleted successfully")
	return nil
}

func SelectRouteFromSourceOrDestination(db *sql.DB, stationId string) ([]model.Route, error) {

	// sqlStatement := `SELECT * FROM transport.route WHERE id=(SELECT route_id FROM transport.routeStations WHERE station_id = $1);`
	// sqlStatement := `SELECT route_id FROM transport.routeStations WHERE station_id = $1;`

	// newId, err := strconv.Atoi(stationId)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return nil, err
	// }

	// res, err := db.Query(sqlStatement, newId)
	// if err != nil {
	// 	return nil, err
	// }
	// defer res.Close()

	// var routes []model.Route
	// for res.Next() {
	// 	var routeId int
	// 	err := res.Scan(&routeId)
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		return nil, err
	// 	}

	// 	res2, err := db.Query(`SELECT * FROM transport.route WHERE id=$1`, routeId)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	defer res2.Close()

	// 	var route model.Route
	// 	for res2.Next() {

	// 		err = res2.Scan(&route.Id, &route.Name, &route.Status, &route.Source, &route.Destination)
	// 		if err != nil {
	// 			log.Println(err.Error())
	// 			return nil, err
	// 		}
	// 		routes = append(routes, route)
	// 	}
	// }

	// log.Println(routes)
	// return routes, nil
	return nil, nil
}
