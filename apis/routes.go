package apis

import (
	"busproject/configs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *App) InitializeRoutes() {

	r := mux.NewRouter()
	apiRoute := r.PathPrefix("/api")
	busRouter := apiRoute.PathPrefix("/bus").Subrouter()
	// driverRouter := r.PathPrefix("/driver").Subrouter()
	// routeRouter := r.PathPrefix("/route").Subrouter()

	busRouter.HandleFunc("/", app.controller.CreateBusHandler).Methods("POST")
	busRouter.HandleFunc("/all", app.controller.CreateAllHandler).Methods("POST")
	// driverRouter.HandleFunc("/driver", controllers.CreateDriverHandler).Methods("POST")
	// routeRouter.HandleFunc("/route", controllers.CreateRouteHandler).Methods("POST")

	server_port, err := configs.GetEnv("SERVER_PORT")
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("INFO: Server started on PORT:" + server_port)
	http.ListenAndServe(":"+server_port, r)
}
