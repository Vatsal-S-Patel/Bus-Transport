package database

import (
	"busproject/model"
	"database/sql"
	"log"
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
