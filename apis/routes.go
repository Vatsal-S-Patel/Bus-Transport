package apis

import (
	"busproject/configs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *App) InitializeRoutes() {

	r := mux.NewRouter()

	r.Use(corsMiddleware)

	apiRoute := r.PathPrefix("/api").Subrouter()

	busRouter := apiRoute.PathPrefix("/bus").Subrouter()
	driverRouter := apiRoute.PathPrefix("/driver").Subrouter()
	routeRouter := apiRoute.PathPrefix("/route").Subrouter()
	routeStationRouter := apiRoute.PathPrefix("/routeStation").Subrouter()
	scheduleRouter := apiRoute.PathPrefix("/schedule").Subrouter()
	stationRouter := apiRoute.PathPrefix("/station").Subrouter()

	busRouter.HandleFunc("/", app.controller.CreateBusHandler).Methods("POST")
	busRouter.HandleFunc("/", app.controller.GetAllBusHandler).Methods("GET")
	busRouter.HandleFunc("/{id}", app.controller.DeleteBusHandler).Methods("POST")

	driverRouter.HandleFunc("/", app.controller.CreateDriverHandler).Methods("POST")
	driverRouter.HandleFunc("/", app.controller.GetAllDriverHandler).Methods("GET")
	driverRouter.HandleFunc("/{id}", app.controller.DeleteDriverHandler).Methods("POST")

	routeRouter.HandleFunc("/", app.controller.CreateRouteHandler).Methods("POST")
	routeRouter.HandleFunc("/", app.controller.GetAllRouteHandler).Methods("GET")
	routeRouter.HandleFunc("/{id}", app.controller.DeleteRouteHandler).Methods("POST")

	routeStationRouter.HandleFunc("/", app.controller.CreateRouteStationHandler).Methods("POST")
	routeStationRouter.HandleFunc("/", app.controller.GetAllRouteStationHandler).Methods("GET")

	scheduleRouter.HandleFunc("/", app.controller.CreateScheduleHandler).Methods("POST")
	scheduleRouter.HandleFunc("/GetUpcomingBus/{id}", app.controller.GetUpcomingBus).Methods("GET")
	scheduleRouter.HandleFunc("/", app.controller.GetAllScheduleHandler).Methods("GET")
	scheduleRouter.HandleFunc("/{id}", app.controller.DeleteScheduleHandler).Methods("POST")

	stationRouter.HandleFunc("/", app.controller.CreateStationHandler).Methods("POST")
	stationRouter.HandleFunc("/", app.controller.GetAllStationHandler).Methods("GET")
	stationRouter.HandleFunc("/{id}", app.controller.DeleteStationHandler).Methods("POST")
	// stationRouter.HandleFunc("/routeFromStation/{id}", app.controller.SelectRouteFromSourceOrDestination).Methods("GET")

	// For All Entries from CSVs
	// busRouter.HandleFunc("/all", app.controller.CreateAllHandler).Methods("POST")

	server_port, err := configs.GetEnv("SERVER_PORT")
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("INFO: Server started on PORT:" + server_port)
	http.ListenAndServe(":"+server_port, r)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Allow preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
