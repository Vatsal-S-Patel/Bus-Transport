package database

import (
	"busproject/model"
	"database/sql"
	"log"
)

func InsertSchedule(db *sql.DB, schedule model.Schedule) error {
	sqlStatement := `INSERT INTO transport.schedule (id, bus_id, route_id, departure_time) VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(sqlStatement, schedule.Id, schedule.BusId, schedule.RouteId, schedule.DepartureTime)
	if err != nil {
		return err
	}

	log.Println("Driver inserted successfully")
	return nil
}
