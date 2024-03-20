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
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = database.InsertStation(c.DB, station)
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	OutputToClient(w, http.StatusOK, "station is created", nil)
}

func (c *Controller) GetAllStationHandler(w http.ResponseWriter, r *http.Request) {
	
	routes, err := database.GetAllStation(c.DB)
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	OutputToClient(w, http.StatusOK,"station is fetched", routes)
}

func (c *Controller) DeleteStationHandler(w http.ResponseWriter, r *http.Request) {

	err := database.DeleteStation(c.DB, mux.Vars(r)["id"])
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	OutputToClient(w, http.StatusOK, "station is deleted", nil)
}

