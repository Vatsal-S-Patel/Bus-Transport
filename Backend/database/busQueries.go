package database

import (
	"busproject/model"
	"database/sql"
	"log"
	"strconv"
)

func InsertBus(db *sql.DB, bus model.Bus) error {
	sqlStatement := `INSERT INTO transport.bus (registration_number, model, capacity) VALUES ($1, $2, $3)`

	_, err := db.Exec(sqlStatement, bus.RegistrationNumber, bus.Model, bus.Capacity)
	if err != nil {
		return err
	}

	// log.Println("Bus inserted successfully")
	return nil
}

func GetAllBus(db *sql.DB) ([]model.Bus, error) {
	sqlStatement := `SELECT * FROM transport.bus;`

	res, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var buses []model.Bus
	for res.Next() {
		var bus model.Bus
		err := res.Scan(&bus.Id, &bus.RegistrationNumber, &bus.Model, &bus.Capacity)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		buses = append(buses, bus)
	}

	return buses, nil
}

func DeleteBus(db *sql.DB, id string) error {
	sqlStatement := `DELETE FROM transport.bus WHERE id=$1`
	newId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = db.Exec(sqlStatement, newId)
	if err != nil {
		return err
	}

	// log.Println("Bus Deleted successfully")
	return nil
}

// this function will only be invoked by socket only...
func UpdateLiveBus(db *sql.DB, data model.BusStatus) error {

	sqlQuery := `INSERT INTO transport.busstatus(bus_id,lat,long,last_updated,traffic,status,last_station_order) VALUES($1,$2,$3,$4,$5,$6,$7) ON CONFLICT (bus_id) DO UPDATE SET lat = $2,long = $3,last_updated = $4,traffic = $5,status = $6,last_station_order = $7`
	_, err := db.Exec(sqlQuery, data.BusId, data.Lat, data.Long, data.LastUpdated, data.Status, data.Status, data.LastStationOrder)
	return err
}

func ChangeBusStatus(db *sql.DB,busid int,status int)error{
	sqlQuery := `UPDATE transport.busstatus SET status = $1 WHERE bus_id = $2`
	_,err := db.Exec(sqlQuery,status,busid)
	return err
}

// func InsertAll(fil string, routes *[]model.Route, station *[]model.Station, schedule *[]model.Schedule, routeStation *[]model.RouteStation, driver *[]model.Driver, bus *[]model.Bus) {

// 	file, err := os.Open(fil)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return
// 	}
// 	defer file.Close()
// 	records, err := csv.NewReader(file).ReadAll()
// 	if err != nil {
// 		log.Println(err.Error())
// 		return
// 	}
// 	if routes != nil {
// 		for _, record := range records {
// 			id, _ := strconv.Atoi(record[0])
// 			source, _ := strconv.Atoi(record[2])
// 			destination, _ := strconv.Atoi(record[3])
// 			status, _ := strconv.Atoi(record[4])

// 			data := model.Route{
// 				Id:          id,
// 				Name:        record[1],
// 				Source:      source,
// 				Destination: destination,
// 				Status:      status,
// 			}
// 			*routes = append(*routes, data)
// 		}
// 	} else if routeStation != nil {
// 		for _, record := range records {
// 			route_id, _ := strconv.Atoi(record[0])
// 			station_id, _ := strconv.Atoi(record[1])
// 			order, _ := strconv.Atoi(record[2])

// 			data := model.RouteStation{
// 				RouteId:      route_id,
// 				StationId:    station_id,
// 				StationOrder: order,
// 			}
// 			*routeStation = append(*routeStation, data)
// 		}

// 	} else if schedule != nil {
// 		for _, record := range records {
// 			id, _ := strconv.Atoi(record[0])
// 			bus_id, _ := strconv.Atoi(record[1])
// 			route_id, _ := strconv.Atoi(record[2])

// 			data := model.Schedule{
// 				Id:            id,
// 				BusId:         bus_id,
// 				RouteId:       route_id,
// 				DepartureTime: record[3],
// 			}
// 			*schedule = append(*schedule, data)
// 		}
// 	} else if station != nil {
// 		for _, record := range records {
// 			id, _ := strconv.Atoi(record[0])
// 			lat, _ := strconv.ParseFloat(record[2], 64)
// 			long, _ := strconv.ParseFloat(record[3], 64)

// 			data := model.Station{
// 				Id:   id,
// 				Name: record[1],
// 				Lat:  lat,
// 				Long: long,
// 			}
// 			*station = append(*station, data)
// 		}
// 	} else if driver != nil {
// 		for _, record := range records {
// 			id, _ := strconv.Atoi(record[0])
// 			gender, _ := strconv.Atoi(record[3])

// 			data := model.Driver{
// 				Id:     id,
// 				Name:   record[1],
// 				Phone:  record[2],
// 				Gender: gender,
// 				Dob:    record[4],
// 			}
// 			*driver = append(*driver, data)
// 		}
// 	} else if bus != nil {
// 		for _, record := range records {
// 			id, _ := strconv.Atoi(record[0])
// 			capacity, _ := strconv.Atoi(record[1])

// 			data := model.Bus{
// 				Id:                 id,
// 				Capacity:           capacity,
// 				Model:              record[2],
// 				RegistrationNumber: record[3],
// 			}
// 			*bus = append(*bus, data)
// 		}
// 	}
// }
