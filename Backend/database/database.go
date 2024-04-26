package database

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	connStr,ok:= os.LookupEnv("DBCONN_URL")
	if !ok {
		connStr = `postgres://manav:clash@localhost:5432/bacancy?sslmode=disable`
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
