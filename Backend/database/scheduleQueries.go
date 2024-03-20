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

	return nil
}

func GetUpcomingBus(db *sql.DB, source, destination int) ([]model.UpcomingBus, error) {
	// if any of source or destination is 0 means they are not provided by client
	if source == 0 {
		return nil, errors.New("source need to be specified")
	}

	var sqlStatement string
	var result *sql.Rows
	var err error

	if destination == 0 {
		sqlStatement = `SELECT f.bus_id,route_id,route_name,source,destination,departure_time,b.lat,b.long,b.last_station_order,b.status,b.traffic FROM (SELECT bus_id,route_id,route_name,source,destination,departure_time,station_order FROM bustransportsystem WHERE station_id = $1 and status = 1) AS f LEFT JOIN transport.busstatus as b ON f.bus_id = b.bus_id WHERE (b.status = 0 AND departure_time >= current_time) OR (b.status = 1 AND b.last_station_order <= f.station_order) ORDER BY departure_time ASC;`
		result, err = db.Query(sqlStatement, source)
	} else {
		sqlStatement = `SELECT o.bus_id,route_id,route_name,source,destination,departure_time,b.lat,b.long,b.last_station_order,b.status,b.traffic FROM (SELECT s.bus_id,s.route_id,s.route_name,s.source,s.destination,s.departure_time FROM (SELECT * FROM bustransportsystem WHERE station_id = $1 and status = 1) as f INNER JOIN (SELECT * FROM bustransportsystem WHERE station_id = $2 and status = 1) as s ON f.route_id = s.route_id WHERE f.bus_id = s.bus_id AND f.station_order < s.station_order) as o LEFT JOIN transport.busstatus as b ON o.bus_id = b.bus_id WHERE b.status = 1 OR (b.status = 0 AND departure_time >= CURRENT_TIME) ORDER BY departure_time ASC;`
		result, err = db.Query(sqlStatement, source, destination)
	}

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
		result.Scan(&dummy.Bus_id, &dummy.Route_id, &dummy.Name, &dummy.Source, &dummy.Destination, &dummy.DepartureTime, &dummy.Lat, &dummy.Long, &dummy.LastStationOrder, &dummy.Status, &dummy.Traffic)
		busOutput = append(busOutput, dummy)
	}

	if len(busOutput) == 0 {
		return nil, errors.New("sorry no bus available")
	}

	return busOutput, nil
}

func GetUpcomingSpecialBus(db *sql.DB, source, destination int) ([]model.UpcomingSpecialBus, error) {
	var sqlStatement string
	var result *sql.Rows
	var err error
	//	sqlStatement = `SELECT v1.route_id as sourceRoute,v1.station_id as middle, v2.route_id as destinationRoute FROM
	//
	// (SELECT o.route_id,station_id,station_order from transport.routestations as o INNER JOIN
	//
	//	(SELECT route_id,station_order as myOrder FROM transport.routestations as r where r.station_id = $1) as s
	//	ON o.route_id = s.route_id where station_order > myOrder) as v1
	//	INNER JOIN
	//
	// (SELECT o.route_id,station_id,station_order from transport.routestations as o INNER JOIN
	//
	//	(SELECT route_id,station_order as myOrder FROM transport.routestations as r where r.station_id = $2) as s
	//	ON o.route_id = s.route_id where station_order < myOrder) as v2 ON v1.station_id = v2.station_id;`
	sqlStatement = `SELECT * from (SELECT v3.route_id as sourceRoute,v3.route_name as sourceRouteName,v3.station_id as junctionStation,v3.station_name as junctionName,v3.station_order as junctionOrder,v3.myOrder,v4.route_id as destinationRoute,v4.route_name as destinationRouteName 
FROM (select v1.route_id,v1.route_name,v2.station_id,v2.station_name,v2.station_order,v1.myOrder from 
               	(SELECT route_id,station_id,route_name,station_name,station_order as myOrder FROM bustransportsystem where station_id = $1) as v1 
									inner join 
								(SELECT route_id,station_id,route_name,station_name,station_order FROM bustransportsystem) as v2 
              	on v1.route_id = v2.route_id where station_order > myOrder) as v3
					INNER JOIN
					(select v1.route_id,v1.route_name,v2.station_id,v2.station_name,v2.station_order from 
           	(SELECT route_id,station_id,route_name,station_name,station_order as myOrder FROM bustransportsystem where station_id = $2) as v1 
							inner join 
						(SELECT route_id,station_id,route_name,station_name,station_order FROM bustransportsystem) as v2 
          on v1.route_id = v2.route_id where station_order < myOrder) as v4 
      ON v3.station_id = v4.station_id group by sourceRoute, sourceRouteName, junctionStation, junctionName, destinationRoute, destinationRouteName, junctionOrder,myOrder) as q
     where sourceroute <> destinationroute order by (junctionOrder - myOrder); `
	result, err = db.Query(sqlStatement, source, destination)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("sorry no data available currently")
	}
	defer result.Close()

	var busOutput []model.UpcomingSpecialBus
	var dummy model.UpcomingSpecialBus

	for result.Next() {
		result.Scan(&dummy.SourceRoute, &dummy.SourceRouteName, &dummy.JunctionStation,&dummy.JunctionName,&dummy.JunctionOrder,&dummy.MyOrder,&dummy.DestinationRoute,&dummy.DestinationRouteName)
		busOutput = append(busOutput, dummy)
	}

	if len(busOutput) == 0 {
		return nil, errors.New("sorry no route available")
	}

	return busOutput, nil
}
