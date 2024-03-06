package controllers

import (
	"busproject/database"
	"busproject/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) CreateBusHandler(w http.ResponseWriter, r *http.Request) {
	var bus model.Bus
	err := json.NewDecoder(r.Body).Decode(&bus)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	err = database.InsertBus(c.DB, bus)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	// log.Println("bus created ", bus)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusOK, Message: "bus Inserted Successfully"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	// json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusOK, Message: "bus Inserted Successfully"})

}

func (c *Controller) GetAllBusHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	buses, err := database.GetAllBus(c.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	// err = json.NewEncoder(w).Encode(buses)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusOK, Message: "bus is fetched", Data: buses})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
}

func (c *Controller) DeleteBusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	err := database.DeleteBus(c.DB, mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusOK, Message: "bus deleted"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
}

func (c *Controller) UpdateLiveBus(w http.ResponseWriter, r *http.Request) {
	var data model.BusStatus

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	err = database.UpdateLiveBus(c.DB, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusOK, Message: "bus updated"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

}

// func (c *Controller) CreateAllHandler(w http.ResponseWriter, r *http.Request) {

// 	var schedules []model.Schedule

// 	database.InsertAll("csvs/Bus_Route_Shedule - Sheet1.csv", nil, nil, &schedules, nil, nil, nil)
// 	log.Println(schedules)
// 	for _, schedule := range schedules {
// 		err := database.InsertSchedule(c.DB, schedule)
// 		if err != nil {
// 			log.Println(err.Error())
// 			return
// 		}
// 	}
// 	log.Println("For lopp chal gaya")

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Successfull All"))

// }
