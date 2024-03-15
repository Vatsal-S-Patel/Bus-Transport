package database

import (
	"busproject/model"
	"database/sql"
	"log"
	"strconv"
)

func InsertRoute(db *sql.DB, route model.Route) error {
	sqlStatement := `INSERT INTO transport.route (id, name, status, source, destination) VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Exec(sqlStatement, route.Id, route.Name, route.Status, route.Source, route.Destination)
	if err != nil {
		return err
	}

	// log.Println("Route inserted successfully")
	return nil
}

func GetAllRoute(db *sql.DB) ([]model.Route, error) {
	sqlStatement := `SELECT * FROM transport.route;`

	res, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var routes []model.Route
	for res.Next() {
		var route model.Route
		err := res.Scan(&route.Id, &route.Name, &route.Status, &route.Source, &route.Destination)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		routes = append(routes, route)
	}

	return routes, nil
}

func DeleteRoute(db *sql.DB, id string) error {
	sqlStatement := `DELETE FROM transport.route WHERE id=$1`

	sqlStatement2 := `DELETE FROM transport.routestations WHERE route_id=$1`

	// sqlStatement3 := `DELETE FROM transport.schedule WHERE route_id=$1`

	newId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// _, err = db.Exec(sqlStatement3, newId)
	// if err != nil {
	// 	return err
	// }

	_, err = db.Exec(sqlStatement2, newId)
	if err != nil {
		return err
	}

	_, err = db.Exec(sqlStatement, newId)
	if err != nil {
		return err
	}

	// log.Println("Route Deleted successfully")
	return nil
}
