package controllers

import (
	"busproject/database"
	"busproject/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) CreateScheduleHandler(w http.ResponseWriter, r *http.Request) {
	var schedule model.Schedule
	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	err = database.InsertSchedule(c.DB, schedule)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = 	json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusOK,Message: "shedule is created"})
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	// w.Write([]byte(string(rune(schedule.Id)) + "Inserted Successfully"))
}

func (c *Controller) GetAllScheduleHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	routes, err := database.GetAllSchedule(c.DB)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	
	w.WriteHeader(http.StatusOK)
	err = 	json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusOK,Message: "schedule is fetched",Data: routes})
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

}

func (c *Controller) DeleteScheduleHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	err := database.DeleteSchedule(c.DB, mux.Vars(r)["id"])
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusOK,Message: "shedule is deleted"})
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

}

func (c *Controller) GetUpcomingBus(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	var variable map[string]int = map[string]int{}

	err := json.NewDecoder(r.Body).Decode(&variable)
	// variable := r.URL.Query()
	log.Println(variable)
	// variable := mux.Vars(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	// source, err := strconv.Atoi(variable["source"])
	// var destinaiton = 0

	ouput, err := database.GetUpcomingBus(c.DB, variable["source"], variable["destination"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusOK,Message: "shedule is fetched",Data: ouput})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
}
