package main

import (
	"busproject/apis"
	"busproject/configs"
	"busproject/database"
	"log"
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
	app.InitializeRoutes()
}
