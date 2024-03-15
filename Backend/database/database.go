package database

import (
	"busproject/configs"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	connStr, err := configs.GetEnv("DBCONN_URL")
	if err != nil {
		return nil, err
	}

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database Connection Established")

	return conn, nil
}
