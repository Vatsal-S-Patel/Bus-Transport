package database

import (
	"busproject/model"
	"database/sql"
	"errors"
	"log"
	"strconv"
)

var (
	ErrSourceNotSpecified = errors.New("source need to be specified")
	ErrNoDataAvailable    = errors.New("sorry no data available currently")
	ErrNoBusAvailable     = errors.New("sorry no bus available")
)

func InsertSchedule(db *sql.DB, schedule model.Schedule) error {
	sqlStatement, err := db.Prepare(`INSERT INTO transport.schedule (id, bus_id, route_id, departure_time) VALUES ($1, $2, $3, $4)`)
	if err != nil {
		return err
	}
	defer sqlStatement.Close()

	_, err = sqlStatement.Exec(schedule.Id, schedule.BusId, schedule.RouteId, schedule.DepartureTime)
	return err
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
	sqlStatement, err := db.Prepare(`DELETE FROM transport.schedule WHERE id=$1`)
	if err != nil {
		return err
	}
	defer sqlStatement.Close()

	newId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = sqlStatement.Exec(newId)
	return err
}

func GetUpcomingBus(db *sql.DB, source, destination int) ([]model.UpcomingBus, error) {
	// if any of source or destination is 0 means they are not provided by client
	if source == 0 {
		return nil, ErrSourceNotSpecified
	}

	var sqlStatement *sql.Stmt
	var result *sql.Rows
	var err error

	if destination == 0 {
		sqlStatement, err = db.Prepare(`SELECT f.bus_id,route_id,route_name,source,destination,departure_time,b.lat,b.long,b.last_station_order,b.status,b.traffic 
							FROM (SELECT bus_id,route_id,route_name,source,destination,departure_time,station_order 
								FROM bustransportsystem WHERE station_id = $1 AND status = 1) AS f 
									LEFT JOIN 
								transport.busstatus AS b ON f.bus_id = b.bus_id 
									WHERE (b.status = 0) OR (b.status = 1 AND b.last_station_order <= f.station_order) 
										ORDER BY departure_time ASC;`)
		if err != nil {
			return []model.UpcomingBus{}, err
		}
		defer sqlStatement.Close()
		result, err = sqlStatement.Query(source)
	} else {
		sqlStatement, err = db.Prepare(`SELECT o.bus_id,route_id,route_name,source,destination,departure_time,b.lat,b.long,b.last_station_order,b.status,b.traffic 
							FROM (SELECT s.bus_id,s.route_id,s.route_name,s.source,s.destination,s.departure_time 
								FROM (SELECT * FROM bustransportsystem 
									WHERE station_id = $1 AND status = 1) AS f 
										INNER JOIN 
									(SELECT * FROM bustransportsystem 
										WHERE station_id = $2 AND status = 1) AS s 
									ON f.route_id = s.route_id 
										WHERE f.bus_id = s.bus_id AND f.station_order < s.station_order) AS o 
								LEFT JOIN transport.busstatus AS b 
									ON o.bus_id = b.bus_id 
										WHERE b.status = 1 OR (b.status = 0) 
										ORDER BY departure_time ASC;`)
		if err != nil {
			return []model.UpcomingBus{}, err
		}
		defer sqlStatement.Close()
		result, err = sqlStatement.Query(source, destination)
	}

	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, ErrNoDataAvailable
	}
	defer result.Close()

	var busOutput []model.UpcomingBus
	var dummy model.UpcomingBus

	for result.Next() {
		result.Scan(&dummy.Bus_id, &dummy.Route_id, &dummy.Name, &dummy.Source, &dummy.Destination, &dummy.DepartureTime, &dummy.Lat, &dummy.Long, &dummy.LastStationOrder, &dummy.Status, &dummy.Traffic)
		busOutput = append(busOutput, dummy)
	}

	if len(busOutput) == 0 {
		return nil, ErrNoBusAvailable
	}

	return busOutput, nil
}

func GetUpcomingSpecialBus(db *sql.DB, source, destination int) ([]model.UpcomingSpecialBus, error) {
	var sqlStatement *sql.Stmt
	var result *sql.Rows
	var err error
	var SrcDestMap = map[int]int{}

	sqlStatement, err = db.Prepare(`SELECT * FROM (SELECT v3.route_id AS sourceRoute,v3.route_name AS sourceRouteName,v3.station_id AS junctionStation,v3.station_name AS junctionName,v3.station_order AS junctionOrder,v3.myOrder,v4.route_id AS destinationRoute,v4.route_name AS destinationRouteName 
					FROM (SELECT v1.route_id,v1.route_name,v2.station_id,v2.station_name,v2.station_order,v1.myOrder 
					FROM (SELECT route_id,station_id,route_name,station_name,station_order AS myOrder FROM bustransportsystem WHERE station_id = $1) AS v1 
						 INNER JOIN 
						 (SELECT route_id,station_id,route_name,station_name,station_order FROM bustransportsystem) AS v2 
              			 on v1.route_id = v2.route_id where station_order > myOrder) as v3
					INNER JOIN
					(SELECT v1.route_id,v1.route_name,v2.station_id,v2.station_name,v2.station_order 
					FROM (SELECT route_id,station_id,route_name,station_name,station_order AS myOrder 
						 FROM bustransportsystem WHERE station_id = $2) AS v1 
						 INNER JOIN 
						(SELECT route_id,station_id,route_name,station_name,station_order FROM bustransportsystem) AS v2 
          				ON v1.route_id = v2.route_id WHERE station_order < myOrder) AS v4 
      				ON v3.station_id = v4.station_id 
					GROUP BY sourceRoute, sourceRouteName, junctionStation, junctionName, destinationRoute, destinationRouteName, junctionOrder,myOrder) AS q
     				WHERE sourceroute <> destinationroute ORDER BY (junctionOrder - myOrder); `)
	if err != nil {
		return []model.UpcomingSpecialBus{}, err
	}
	defer sqlStatement.Close()

	result, err = sqlStatement.Query(source, destination)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, ErrNoDataAvailable
	}
	defer result.Close()

	var busOutput []model.UpcomingSpecialBus
	var dummy model.UpcomingSpecialBus

	for result.Next() {
		result.Scan(&dummy.SourceRoute, &dummy.SourceRouteName, &dummy.JunctionStation, &dummy.JunctionName, &dummy.JunctionOrder, &dummy.MyOrder, &dummy.DestinationRoute, &dummy.DestinationRouteName)
		if v, ok := SrcDestMap[dummy.SourceRoute]; !ok || (ok && v != dummy.DestinationRoute) {
			SrcDestMap[dummy.SourceRoute] = dummy.DestinationRoute
			busOutput = append(busOutput, dummy)
		}
	}

	if len(busOutput) == 0 {
		return nil, ErrNoBusAvailable
	}

	return busOutput, nil
}
