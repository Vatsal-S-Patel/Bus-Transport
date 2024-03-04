package controllers

import (
	"busproject/database"
	"busproject/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) CreateStationHandler(w http.ResponseWriter, r *http.Request) {
	var station model.Station
	err := json.NewDecoder(r.Body).Decode(&station)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = database.InsertStation(c.DB, station)
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(station)
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.Write([]byte(station.Name + "Inserted Successfully"))
}

func (c *Controller) GetAllStationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	routes, err := database.GetAllStation(c.DB)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(routes)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func (c *Controller) DeleteStationHandler(w http.ResponseWriter, r *http.Request) {

	err := database.DeleteStation(c.DB, mux.Vars(r)["id"])
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Station Deleted"))
}

func (c *Controller) SelectRouteFromSourceOrDestination(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	routeStations, err := database.SelectRouteFromSourceOrDestination(c.DB, mux.Vars(r)["id"])
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(routeStations)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
