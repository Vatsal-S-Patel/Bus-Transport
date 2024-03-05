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
	var routeWithStationOrder model.RouteStationMerged

	err := json.NewDecoder(r.Body).Decode(&routeWithStationOrder)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	err = database.InsertRoute(c.DB, routeWithStationOrder.Route)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	for _, v := range routeWithStationOrder.RouteStation {
		log.Println("inserting staiton with ", routeWithStationOrder.Id, " name ", v.StationId, " with order ", v.StationOrder)
		err := database.InsertRouteStation(c.DB, v)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(routeWithStationOrder.Name + "Inserted Successfully"))

}

func (c *Controller) GetAllRouteHandler(w http.ResponseWriter, r *http.Request) {

	routes, err := database.GetAllRoute(c.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(routes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

}

func (c *Controller) DeleteRouteHandler(w http.ResponseWriter, r *http.Request) {

	err := database.DeleteRoute(c.DB, mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Errorstruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Route Deleted"))
}
