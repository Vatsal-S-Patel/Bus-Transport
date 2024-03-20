package database

import (
	"busproject/model"
	"database/sql"
	"log"
	"strconv"
)

func InsertStation(db *sql.DB, station model.Station) error {
	sqlStatement := `INSERT INTO transport.station (id, name, lat, long) VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(sqlStatement, station.Id, station.Name, station.Lat, station.Long)
	if err != nil {
		return err
	}

	return nil
}

func GetAllStation(db *sql.DB) ([]model.Station, error) {
	sqlStatement := `SELECT * FROM transport.station ORDER BY name ASC;`

	res, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var stations []model.Station
	for res.Next() {
		var station model.Station

		err := res.Scan(&station.Id, &station.Name, &station.Lat, &station.Long)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		
		stations = append(stations, station)
	}

	return stations, nil
}

func DeleteStation(db *sql.DB, id string) error {
	sqlStatement := `DELETE FROM transport.station WHERE id=$1`

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
