package controllers

import (
	"busproject/database"
	"busproject/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) CreateRouteHandler(w http.ResponseWriter, r *http.Request) {
	var route model.Route
	err := json.NewDecoder(r.Body).Decode(&route)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = database.InsertRoute(c.DB, route)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(route)
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(route.Name + "Inserted Successfully"))

}

func (c *Controller) GetAllRouteHandler(w http.ResponseWriter, r *http.Request) {

	routes, err := database.GetAllRoute(c.DB)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(routes)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) DeleteRouteHandler(w http.ResponseWriter, r *http.Request) {

	err := database.DeleteRoute(c.DB, mux.Vars(r)["id"])
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Route Deleted"))
}
