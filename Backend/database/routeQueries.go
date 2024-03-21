package database

import (
	"busproject/model"
	"database/sql"
	"log"
	"strconv"
)

func InsertRoute(db *sql.DB, route model.Route) error {
	sqlStatement,err:=db.Prepare(`INSERT INTO transport.route (id, name, status, source, destination) VALUES ($1, $2, $3, $4, $5)`)
	if err != nil {
		return err
	}
	defer sqlStatement.Close()

	_, err = sqlStatement.Exec(route.Id, route.Name, route.Status, route.Source, route.Destination)
	return err
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
	sqlStatement,err := db.Prepare(`DELETE FROM transport.route WHERE id=$1`)
	if err != nil {
		return err
	}
	defer sqlStatement.Close()

	sqlStatement2,err := db.Prepare(`DELETE FROM transport.routestations WHERE route_id=$1`)
	if err != nil {
		return err
	}
	defer sqlStatement2.Close()

	newId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = sqlStatement2.Exec(newId)
	if err != nil {
		return err
	}

	_, err = sqlStatement.Exec(newId)
	return err
}
