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
		OutputToClient(w,http.StatusInternalServerError,err.Error(),nil)
		return
	}

	err = database.InsertDriver(c.DB, driver)
	if err != nil {
		OutputToClient(w,http.StatusInternalServerError,err.Error(),nil)
		return
	}

	OutputToClient(w,http.StatusOK,driver.Name+"driver is created",nil)
}

func (c *Controller) GetAllDriverHandler(w http.ResponseWriter, r *http.Request) {

	drivers, err := database.GetAllDriver(c.DB)
	if err != nil {
		OutputToClient(w,http.StatusInternalServerError,err.Error(),nil)
		return
	}

	OutputToClient(w,http.StatusOK,"driver is fetched",drivers)
}

func (c *Controller) DeleteDriverHandler(w http.ResponseWriter, r *http.Request) {

	err := database.DeleteDriver(c.DB, mux.Vars(r)["id"])
	if err != nil {
		OutputToClient(w,http.StatusInternalServerError,err.Error(),nil)
		return
	}

	OutputToClient(w,http.StatusOK,"driver is deleted",nil)
}
