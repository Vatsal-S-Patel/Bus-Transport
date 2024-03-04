package database

import (
	"busproject/model"
	"database/sql"
	"log"
	"strconv"
)

func InsertSchedule(db *sql.DB, schedule model.Schedule) error {
	sqlStatement := `INSERT INTO transport.schedule (id, bus_id, route_id, departure_time) VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(sqlStatement, schedule.Id, schedule.BusId, schedule.RouteId, schedule.DepartureTime)
	if err != nil {
		return err
	}

	log.Println("Schedule inserted successfully")
	return nil
}

func GetAllSchedule(db *sql.DB) ([]model.Schedule, error) {
	sqlStatement := `SELECT * FROM transport.schedule;`

	res, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var schedules []model.Schedule
	for res.Next() {
		var shcedule model.Schedule
		err := res.Scan(&shcedule.Id, &shcedule.BusId, &shcedule.RouteId, &shcedule.DepartureTime)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		schedules = append(schedules, shcedule)
	}

	return schedules, nil
}

func DeleteSchedule(db *sql.DB, id string) error {
	sqlStatement := `DELETE FROM transport.schedule WHERE id=$1`

	newId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = db.Exec(sqlStatement, newId)
	if err != nil {
		return err
	}

	log.Println("Schedule Deleted successfully")
	return nil
}

func GetUpcomingBus(db *sql.DB, id int) ([]model.UpcomingBus, error) {
	// sqlStatement := `SELECT "name","source",destination FROM (SELECT route_id FROM transport.routestations INNER JOIN transport.station ON station_id = transport.station.id WHERE transport.station.id = $1 ) as r INNER JOIN transport.route ON transport.route.id = r.route_id WHERE status = 1;`
	// sqlStatement := `SELECT "name","source",destination FROM transport.routestations inner join transport.route on route_id = transport.route.id WHERE station_id = $1;`
	sqlStatement := `SELECT "name","source",destination FROM (SELECT "id","name","source",destination FROM transport.routestations inner join transport.route on route_id = transport.route.id WHERE station_id = $1) as b  INNER JOIN transport.schedule ON b."id" = transport.schedule.route_id where departure_time >= CURRENT_TIME;`
	// sqlStatement := `SELECT name,
// (select name from transport.station where transport.station.id = source LIMIT 1),
// (select name from transport.station where transport.station.id = destination LIMIT 1) from (SELECT "name","source",destination FROM transport.routestations inner join transport.route on route_id = transport.route.id WHERE station_id = $1) as x;
// `

	result, err := db.Query(sqlStatement, id)
	if err != nil {
		return []model.UpcomingBus{}, err
	}

	var busOutput []model.UpcomingBus
	var dummy model.UpcomingBus
	for result.Next() {
		result.Scan(&dummy.Name, &dummy.Source, &dummy.Destination)
		busOutput = append(busOutput, dummy)
		// log.Println(busOutput)
	}

	// log.Println("data is fetched")
	return busOutput, nil
}
