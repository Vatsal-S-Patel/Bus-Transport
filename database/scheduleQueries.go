package database

import (
	"busproject/model"
	"database/sql"
	"errors"
	"log"
	"strconv"
)

func InsertSchedule(db *sql.DB, schedule model.Schedule) error {
	sqlStatement := `INSERT INTO transport.schedule (id, bus_id, route_id, departure_time) VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(sqlStatement, schedule.Id, schedule.BusId, schedule.RouteId, schedule.DepartureTime)
	if err != nil {
		return err
	}

	// log.Println("Schedule inserted successfully")
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

	// log.Println("Schedule Deleted successfully")
	return nil
}

func GetUpcomingBus(db *sql.DB, source, destination int) ([]model.UpcomingBus, error) {
	// fmt.Println(source, destination)
	// if any of source or destination is 0 means they are not provided by client
	if source == 0 {
		return nil, errors.New("source need to be specified")
	}
	var sqlStatement string
	var result *sql.Rows
	var err error
	if destination == 0 {
		sqlStatement = `SELECT f.bus_id,route_name,source,destination,departure_time,b.lat,b.long,b.last_station_order FROM (SELECT bus_id,route_name,source,destination,departure_time FROM bustransportsystem WHERE station_id = $1 and status = 1 AND departure_time >= CURRENT_TIME) AS f LEFT JOIN transport.busstatus as b ON f.bus_id = b.bus_id WHERE b.status = 1 OR b.status IS NULL ORDER BY departure_time ASC;`
		result, err = db.Query(sqlStatement, source)
	} else {
		sqlStatement = `SELECT o.bus_id,route_name,source,destination,departure_time,b.lat,b.long,b.last_station_order FROM (SELECT s.bus_id,s.route_name,s.source,s.destination,s.departure_time FROM (SELECT * FROM bustransportsystem WHERE station_id = $1 and status = 1) as f INNER JOIN (SELECT * FROM bustransportsystem WHERE station_id = $2 and status = 1) as s ON f.route_id = s.route_id WHERE f.bus_id = s.bus_id AND f.station_order < s.station_order) as o LEFT JOIN transport.busstatus as b ON o.bus_id = b.bus_id WHERE b.status = 1 OR b.status IS NULL AND departure_time >= CURRENT_TIME ORDER BY departure_time ASC;`
		result, err = db.Query(sqlStatement, source, destination)
	}
	// sqlStatement := `SELECT "name","source",destination FROM (SELECT route_id FROM transport.routestations INNER JOIN transport.station ON station_id = transport.station.id WHERE transport.station.id = $1 ) as r INNER JOIN transport.route ON transport.route.id = r.route_id WHERE status = 1;`
	// sqlStatement := `SELECT "name","source",destination FROM transport.routestations inner join transport.route on route_id = transport.route.id WHERE station_id = $1;`

	// show each route wich are active and having a bus departure time after the current time only
	// sqlStatement := `SELECT "name","source",destination FROM (SELECT "id","name","source",destination FROM transport.routestations inner join transport.route on route_id = transport.route.id WHERE station_id = $1 AND transport.route.status = 1) as b  INNER JOIN transport.schedule ON b."id" = transport.schedule.route_id where departure_time >= CURRENT_TIME;`
	// sqlStatement := `
	// SELECT * FROM (SELECT "name","source",destination,bus_id FROM (SELECT "id","name","source",destination FROM transport.routestations inner join transport.route on route_id = transport.route.id WHERE station_id = $1 AND STATUS = 1) as b  INNER JOIN transport.schedule ON b."id" = transport.schedule.route_id where departure_time >= CURRENT_TIME) AS d LEFT JOIN transport.busstatus ON d.bus_id = transport.busstatus.bus_id WHERE transport.busstatus.status = 1 OR transport.busstatus.status is null;`
	// sqlStatement := `SELECT name,
	// (select name from transport.station where transport.station.id = source LIMIT 1),
	// (select name from transport.station where transport.station.id = destination LIMIT 1) from (SELECT "name","source",destination FROM transport.routestations inner join transport.route on route_id = transport.route.id WHERE station_id = $1) as x;
	// `

	// result, err := db.Query(sqlStatement, source, destination)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("sorry no data available currently")
	}
	defer result.Close()

	var busOutput []model.UpcomingBus
	var dummy model.UpcomingBus
	for result.Next() {
		result.Scan(&dummy.Bus_id, &dummy.Name, &dummy.Source, &dummy.Destination, &dummy.DepartureTime, &dummy.Lat, &dummy.Long, &dummy.LastStationOrder)
		busOutput = append(busOutput, dummy)
		// log.Println(busOutput)
	}

	// log.Println("data is fetched")
	if len(busOutput) == 0 {
		return nil, errors.New("sorry no bus available")
	}
	return busOutput, nil
}
