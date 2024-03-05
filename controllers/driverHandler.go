package controllers

import (
	"busproject/database"
	"busproject/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) CreateDriverHandler(w http.ResponseWriter, r *http.Request) {
	var driver model.Driver
	err := json.NewDecoder(r.Body).Decode(&driver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	err = database.InsertDriver(c.DB, driver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(driver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.Write([]byte(driver.Name + "Inserted Successfully"))

}

func (c *Controller) GetAllDriverHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	drivers, err := database.GetAllDriver(c.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(drivers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

}

func (c *Controller) DeleteDriverHandler(w http.ResponseWriter, r *http.Request) {

	err := database.DeleteDriver(c.DB, mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Driver Deleted"))
}
