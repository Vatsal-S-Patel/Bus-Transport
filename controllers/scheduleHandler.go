package controllers

import (
	"busproject/database"
	"busproject/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (c *Controller) CreateScheduleHandler(w http.ResponseWriter, r *http.Request) {
	var schedule model.Schedule
	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = database.InsertSchedule(c.DB, schedule)
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(schedule)
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.Write([]byte(string(rune(schedule.Id)) + "Inserted Successfully"))
}

func (c *Controller) GetAllScheduleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	routes, err := database.GetAllSchedule(c.DB)
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

func (c *Controller) DeleteScheduleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := database.DeleteSchedule(c.DB, mux.Vars(r)["id"])
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Schedule Deleted"))
}

func (c *Controller) GetUpcomingBus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var body map[string]string = map[string]string{}

	// err := json.NewDecoder(r.Body).Decode(&body)
	variable := mux.Vars(r)
	id, err := strconv.Atoi(variable["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ouput, err := database.GetUpcomingBus(c.DB, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ouput)
}
