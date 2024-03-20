package main

import (
	"busproject/apis"
	"busproject/configs"
	"busproject/database"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	err := configs.ReadEnv()
	if err != nil {
		log.Println(err.Error())
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Println("ERROR: DB Connection Error", err)
		return
	}
	defer db.Close()

	app := apis.NewApp(db)
	r := app.InitializeRoutes()

	log.Println(startServer(r))
}

func startServer(r *mux.Router) error {
	server_port, err := configs.GetEnv("SERVER_PORT")
	
	if errors.Is(err, configs.ErrDataNotExist) {
		server_port = "8080"
	}else if err != nil {
		return err
	}

	log.Println("starting server at ",server_port)
	return http.ListenAndServe(":"+server_port, r)
}
