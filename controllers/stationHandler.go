package controllers

import (
	"busproject/database"
	"busproject/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) CreateStationHandler(w http.ResponseWriter, r *http.Request) {
	var station model.Station
	err := json.NewDecoder(r.Body).Decode(&station)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	err = database.InsertStation(c.DB, station)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusOK, Message: "station is created"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
}

func (c *Controller) GetAllStationHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	routes, err := database.GetAllStation(c.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusOK, Message: "station is fetched", Data: routes})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

}

func (c *Controller) DeleteStationHandler(w http.ResponseWriter, r *http.Request) {

	err := database.DeleteStation(c.DB, mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusOK, Message: "station is deleted"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
}

func (c *Controller) SelectRouteFromSourceOrDestination(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	routeStations, err := database.SelectRouteFromSourceOrDestination(c.DB, mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusOK, Message: "station is fetched", Data: routeStations})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

}
