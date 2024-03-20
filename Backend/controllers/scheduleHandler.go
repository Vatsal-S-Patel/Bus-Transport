package controllers

import (
	"busproject/database"
	"busproject/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) CreateScheduleHandler(w http.ResponseWriter, r *http.Request) {
	var schedule model.Schedule

	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = database.InsertSchedule(c.DB, schedule)
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	OutputToClient(w, http.StatusOK, "shedule is created", nil)
}

func (c *Controller) GetAllScheduleHandler(w http.ResponseWriter, r *http.Request) {

	routes, err := database.GetAllSchedule(c.DB)
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	OutputToClient(w, http.StatusOK, "schedule is fetched", routes)
}

func (c *Controller) DeleteScheduleHandler(w http.ResponseWriter, r *http.Request) {

	err := database.DeleteSchedule(c.DB, mux.Vars(r)["id"])
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	OutputToClient(w, http.StatusOK, "schedule is deleted", nil)
}

func (c *Controller) GetUpcomingBus(w http.ResponseWriter, r *http.Request) {
	var variable map[string]int = map[string]int{}

	err := json.NewDecoder(r.Body).Decode(&variable)
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	ouput, err := database.GetUpcomingBus(c.DB, variable["source"], variable["destination"])
	if err != nil && err.Error() == "sorry no bus available" {
		OutputToClient(w, http.StatusNotFound, err.Error(), nil)
		return
	} else if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	OutputToClient(w, http.StatusOK, "schedule is fetched", ouput)
}

func (c *Controller) GetUpcomingSpecialBus(w http.ResponseWriter, r *http.Request) {
	var variable map[string]int = map[string]int{}

	err := json.NewDecoder(r.Body).Decode(&variable)
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	output, err := database.GetUpcomingSpecialBus(c.DB, variable["source"], variable["destination"])
	if err != nil && err.Error() == "sorry no bus available" {
		OutputToClient(w, http.StatusNotFound, err.Error(), nil)
		return
	} else if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	OutputToClient(w, http.StatusOK, "schedule is fetched", output)
}
