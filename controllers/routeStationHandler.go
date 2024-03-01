package controllers

import (
	"busproject/database"
	"busproject/model"
	"encoding/json"
	"log"
	"net/http"
)

func (c *Controller) CreateRouteStationHandler(w http.ResponseWriter, r *http.Request) {
	var routeStation model.RouteStation
	err := json.NewDecoder(r.Body).Decode(&routeStation)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = database.InsertRouteStation(c.DB, routeStation)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(routeStation)
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("RouteStation Inserted Successfully"))
}

func (c *Controller) GetAllRouteStationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	routeStations, err := database.GetAllRouteStation(c.DB)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(routeStations)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
