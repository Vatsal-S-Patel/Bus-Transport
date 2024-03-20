package database

import (
	"busproject/model"
	"database/sql"
	"log"
	"strconv"
)

func InsertDriver(db *sql.DB, driver model.Driver) error {
	sqlStatement := `INSERT INTO transport.driver (name, phone, gender, dob) VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(sqlStatement, driver.Name, driver.Phone, driver.Gender, driver.Dob)
	if err != nil {
		return err
	}

	return nil
}

func GetAllDriver(db *sql.DB) ([]model.Driver, error) {
	sqlStatement := `SELECT * FROM transport.driver;`
	var drivers []model.Driver

	res, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var driver model.Driver

		err := res.Scan(&driver.Id, &driver.Name, &driver.Phone, &driver.Gender, &driver.Dob)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		
		driver.Dob = driver.Dob[:10]
		drivers = append(drivers, driver)
	}

	return drivers, nil
}

func DeleteDriver(db *sql.DB, id string) error {
	sqlStatement := `DELETE FROM transport.driver WHERE id=$1`

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
