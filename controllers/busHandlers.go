package controllers

import (
	"busproject/database"
	"busproject/model"
	"encoding/json"
	"log"
	"net/http"
)

func (c *Controller) CreateBusHandler(w http.ResponseWriter, r *http.Request) {
	var bus model.Bus
	err := json.NewDecoder(r.Body).Decode(&bus)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = database.InsertBus(c.DB, bus)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(bus)
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(bus.RegistrationNumber + "Inserted Successfully"))

}

func (c *Controller) CreateAllHandler(w http.ResponseWriter, r *http.Request) {

	var schedules []model.Schedule

	database.InsertAll("csvs/Bus_Route_Shedule - Sheet1.csv", nil, nil, &schedules, nil, nil, nil)
	log.Println(schedules)
	for _, schedule := range schedules {
		err := database.InsertSchedule(c.DB, schedule)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
	log.Println("For lopp chal gaya")

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Successfull All"))

}
