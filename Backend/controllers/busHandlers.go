package controllers

import (
	"busproject/database"
	"busproject/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// inserts a bus into the database and a new bus will be created
// TOD: if a bus is added a entry to bus status table need to be inserted with status unassigned
func (c *Controller) CreateBusHandler(w http.ResponseWriter, r *http.Request) {
	var bus model.Bus

	err := json.NewDecoder(r.Body).Decode(&bus)
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = database.InsertBus(c.DB, bus)
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = database.UpdateLiveBus(c.DB, model.BusStatus{
		BusId:       bus.Id,
		LastUpdated: "00:00",
	})
	if err != nil {
		log.Println(err.Error())
		return
	}

	OutputToClient(w, http.StatusOK, "bus inserted successfull", nil)
}

// return each bus currently in table for update and delete purpose
func (c *Controller) GetAllBusHandler(w http.ResponseWriter, r *http.Request) {

	buses, err := database.GetAllBus(c.DB)
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	OutputToClient(w, http.StatusOK, "bus is fetched", buses)
}

func (c *Controller) DeleteBusHandler(w http.ResponseWriter, r *http.Request) {
	var id string

	if v, ok := mux.Vars(r)["id"]; !ok {
		OutputToClient(w, http.StatusBadRequest, "please specify id of bus to delete", nil)
		return
	} else {
		id = v
	}

	err := database.DeleteBus(c.DB, id)
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	OutputToClient(w, http.StatusOK, "bus deleted", nil)
}

func (c *Controller) UpdateLiveBus(w http.ResponseWriter, r *http.Request) {
	var data model.BusStatus

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = database.UpdateLiveBus(c.DB, data)
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	OutputToClient(w, http.StatusOK, "bus updated", nil)
}

