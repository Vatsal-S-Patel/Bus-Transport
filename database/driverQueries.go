package database

import (
	"busproject/model"
	"database/sql"
	"log"
)

func InsertDriver(db *sql.DB, driver model.Driver) error {
	sqlStatement := `INSERT INTO transport.driver (name, phone, gender, dob) VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(sqlStatement, driver.Name, driver.Phone, driver.Gender, driver.Dob)
	if err != nil {
		return err
	}

	log.Println("Driver inserted successfully")
	return nil
}
