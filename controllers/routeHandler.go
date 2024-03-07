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

	// log.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&routeWithStationOrder)
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = database.InsertRoute(c.DB, routeWithStationOrder.Route)
	if err != nil {
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	mappingStatusCode := c.CreateMappingHandler(routeWithStationOrder, routeWithStationOrder.Id)
	if mappingStatusCode != http.StatusOK {
		OutputToClient(w, http.StatusInternalServerError, "internal server at 34 in routehandler.go", nil)
		return
	}

	OutputToClient(w, http.StatusOK, "routes is created", nil)
	// for _, v := range routeWithStationOrder.RouteStationArray {
	// 	log.Println("inserting staiton with ", routeWithStationOrder.Id, " name ", v.StationId, " with order ", v.StationOrder)
	// 	err := database.InsertRouteStation(c.DB, v)

	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// }
}

func (c *Controller) GetAllRouteHandler(w http.ResponseWriter, r *http.Request) {
	routes, err := database.GetAllRoute(c.DB)
	if err != nil {
		log.Println(err.Error())
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	OutputToClient(w, http.StatusOK,"route is fetched",routes)
}

func (c *Controller) DeleteRouteHandler(w http.ResponseWriter, r *http.Request) {

	err := database.DeleteRoute(c.DB, mux.Vars(r)["id"])
	if err != nil {
		log.Println(err.Error())
		OutputToClient(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	OutputToClient(w, http.StatusOK, "routes is deleted", nil)
}
