package csvfileinsert

import (
	"busproject/model"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func InsertAll(fil string, routes *[]model.Route, station *[]model.Station, schedule *[]model.Schedule, routeStation *[]model.RouteStation, driver *[]model.Driver, bus *[]model.Bus) {

	file, err := os.Open(fil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println(err.Error())
		return
	}

	if routes != nil {
		for _, record := range records {

			id, _ := strconv.Atoi(record[0])
			source, _ := strconv.Atoi(record[2])
			destination, _ := strconv.Atoi(record[3])
			status, _ := strconv.Atoi(record[4])

			data := model.Route{
				Id:          id,
				Name:        record[1],
				Source:      source,
				Destination: destination,
				Status:      status,
			}

			*routes = append(*routes, data)

		}
	} else if routeStation != nil {
		for _, record := range records {
			
			route_id, _ := strconv.Atoi(record[0])
			station_id, _ := strconv.Atoi(record[1])
			order, _ := strconv.Atoi(record[2])

			data := model.RouteStation{
				RouteId:      route_id,
				StationId:    station_id,
				StationOrder: order,
			}

			*routeStation = append(*routeStation, data)

		}

	} else if schedule != nil {
		for _, record := range records {

			id, _ := strconv.Atoi(record[0])
			bus_id, _ := strconv.Atoi(record[1])
			route_id, _ := strconv.Atoi(record[2])

			data := model.Schedule{
				Id:            id,
				BusId:         bus_id,
				RouteId:       route_id,
				DepartureTime: record[3],
			}

			*schedule = append(*schedule, data)

		}
	} else if station != nil {
		for _, record := range records {
			
			id, _ := strconv.Atoi(record[0])
			lat, _ := strconv.ParseFloat(record[2], 64)
			long, _ := strconv.ParseFloat(record[3], 64)

			data := model.Station{
				Id:   id,
				Name: record[1],
				Lat:  lat,
				Long: long,
			}

			*station = append(*station, data)

		}
	} else if driver != nil {
		for _, record := range records {
			
			id, _ := strconv.Atoi(record[0])
			gender, _ := strconv.Atoi(record[3])

			data := model.Driver{
				Id:     id,
				Name:   record[1],
				Phone:  record[2],
				Gender: gender,
				Dob:    record[4],
			}

			*driver = append(*driver, data)

		}
	} else if bus != nil {
		for _, record := range records {
			
			id, _ := strconv.Atoi(record[0])
			capacity, _ := strconv.Atoi(record[1])

			data := model.Bus{
				Id:                 id,
				Capacity:           capacity,
				Model:              record[2],
				RegistrationNumber: record[3],
			}

			*bus = append(*bus, data)
			
		}
	}
}
