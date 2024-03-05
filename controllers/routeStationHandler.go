package controllers

import (
	"busproject/database"
	"busproject/model"
	"encoding/json"
	"net/http"
)

func (c *Controller) CreateRouteStationHandler(w http.ResponseWriter, r *http.Request) {
	var routeStation model.RouteStation
	err := json.NewDecoder(r.Body).Decode(&routeStation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	err = database.InsertRouteStation(c.DB, routeStation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)	
	err = json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusOK,Message: "route and station is created"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
}

func (c *Controller) GetAllRouteStationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	routeStations, err := database.GetAllRouteStation(c.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = 	json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusOK,Message: "route station is fetched",Data: routeStations})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

}
