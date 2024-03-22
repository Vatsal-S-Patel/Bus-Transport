package database

import (
	"busproject/model"
	"database/sql"
	"log"
	"strconv"
)

func InsertBus(db *sql.DB, bus model.Bus) error {
	sqlStatement,err := db.Prepare(`INSERT INTO transport.bus (id,registration_number, model, capacity) VALUES ($1, $2, $3,$4)`)
	if err != nil {
		return err
	}
	defer sqlStatement.Close()
	_, err = sqlStatement.Exec(bus.Id,bus.RegistrationNumber, bus.Model, bus.Capacity)
	if err != nil {
		return err
	}
	err = UpdateLiveBus(db,model.BusStatus{
		BusId: bus.Id,
		LastUpdated: "00:00",
	})
	return err
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
	sqlStatement,err := db.Prepare(`DELETE FROM transport.bus WHERE id=$1`)
	if err != nil{
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

// this functions will only be invoked by socket only...
func UpdateLiveBus(db *sql.DB, data model.BusStatus) error {

	sqlStatement:= `INSERT INTO transport.busstatus(bus_id,lat,long,last_updated,traffic,status,last_station_order) VALUES($1,$2,$3,$4,$5,$6,$7) ON CONFLICT (bus_id) DO UPDATE SET lat = $2,long = $3,last_updated = $4,traffic = $5,status = $6,last_station_order = $7`
	_, err := db.Exec(sqlStatement,data.BusId, data.Lat, data.Long, data.LastUpdated, data.Status, data.Status, data.LastStationOrder)
	return err
}

func ChangeBusStatus(db *sql.DB,busid int,status int)error{
	sqlStatement,err :=db.Prepare(`UPDATE transport.busstatus SET status = $1 WHERE bus_id = $2`)
	if err != nil {
		return err
	}
	defer sqlStatement.Close()

	_,err = sqlStatement.Exec(status,busid)
	return err
}
