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

	log.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&routeWithStationOrder)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = database.InsertRoute(c.DB, routeWithStationOrder.Route)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mappingStatusCode := c.CreateMappingHandler(routeWithStationOrder, routeWithStationOrder.Id)
	if mappingStatusCode != http.StatusOK {
		w.WriteHeader(mappingStatusCode)
		w.Write([]byte(routeWithStationOrder.Name + "Something Bad Happened"))
		return
	}

	// for _, v := range routeWithStationOrder.RouteStationArray {
	// 	log.Println("inserting staiton with ", routeWithStationOrder.Id, " name ", v.StationId, " with order ", v.StationOrder)
	// 	err := database.InsertRouteStation(c.DB, v)

	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// }

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(routeWithStationOrder.Name + "Inserted Successfully"))
}

func (c *Controller) GetAllRouteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
