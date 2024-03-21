package database

import (
	"busproject/model"
	"database/sql"
	"log"
	"strconv"
)

func InsertDriver(db *sql.DB, driver model.Driver) error {
	sqlStatement,err :=db.Prepare(`INSERT INTO transport.driver (name, phone, gender, dob) VALUES ($1, $2, $3, $4)`)
	if err != nil {
		return err
	}
	defer sqlStatement.Close()

	_, err = sqlStatement.Exec(driver.Name, driver.Phone, driver.Gender, driver.Dob)
	return err
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
			return nil, err
		}
		
		driver.Dob = driver.Dob[:10]
		drivers = append(drivers, driver)
	}

	return drivers, nil
}

func DeleteDriver(db *sql.DB, id string) error {
	sqlStatement,err := db.Prepare(`DELETE FROM transport.driver WHERE id=$1`)
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
