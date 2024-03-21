package apis

import (
	"busproject/socket"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *App) InitializeRoutes() *mux.Router {

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
	busRouter.HandleFunc("/live/update", app.controller.UpdateLiveBus).Methods("POST")

	driverRouter.HandleFunc("/", app.controller.CreateDriverHandler).Methods("POST")
	driverRouter.HandleFunc("/", app.controller.GetAllDriverHandler).Methods("GET")
	driverRouter.HandleFunc("/{id}", app.controller.DeleteDriverHandler).Methods("POST")

	routeRouter.HandleFunc("/", app.controller.CreateRouteHandler).Methods("POST")
	routeRouter.HandleFunc("/", app.controller.GetAllRouteHandler).Methods("GET")
	routeRouter.HandleFunc("/{id}", app.controller.DeleteRouteHandler).Methods("POST")

	routeStationRouter.HandleFunc("/", app.controller.CreateRouteStationHandler).Methods("POST")
	routeStationRouter.HandleFunc("/", app.controller.GetAllRouteStationHandler).Methods("GET")

	scheduleRouter.HandleFunc("/", app.controller.CreateScheduleHandler).Methods("POST")
	scheduleRouter.HandleFunc("/GetUpcomingBus", app.controller.GetUpcomingBus).Methods("POST")
	scheduleRouter.HandleFunc("/GetUpcomingSpecialBus", app.controller.GetUpcomingSpecialBus).Methods("POST")
	scheduleRouter.HandleFunc("/", app.controller.GetAllScheduleHandler).Methods("GET")
	scheduleRouter.HandleFunc("/{id}", app.controller.DeleteScheduleHandler).Methods("POST")

	stationRouter.HandleFunc("/", app.controller.CreateStationHandler).Methods("POST")
	stationRouter.HandleFunc("/", app.controller.GetAllStationHandler).Methods("GET")
	stationRouter.HandleFunc("/{id}", app.controller.DeleteStationHandler).Methods("POST")

	server := socket.InitSocket(app.controller.DB)

	r.Handle("/socket.io/", server)

	return r

}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")

		// Allow preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// w.Header().Del("Origin")
		// Call the next handler
		next.ServeHTTP(w, r)

	})
}
