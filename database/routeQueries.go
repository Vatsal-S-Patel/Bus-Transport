package database

import (
	"busproject/model"
	"database/sql"
	"log"
)

func InsertRoute(db *sql.DB, route model.Route) error {
	sqlStatement := `INSERT INTO transport.route (id, name, status, source, destination) VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Exec(sqlStatement, route.Id, route.Name, route.Status, route.Source, route.Destination)
	if err != nil {
		return err
	}

	log.Println("Route inserted successfully")
	return nil
}
