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
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	err = database.InsertRoute(c.DB, routeWithStationOrder.Route)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	mappingStatusCode := c.CreateMappingHandler(routeWithStationOrder, routeWithStationOrder.Id)
	if mappingStatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: "internal server at 34 in routehandler.go"})
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
	err = json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusOK, Message:"routes is created"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
		// return
}

func (c *Controller) GetAllRouteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	routes, err := database.GetAllRoute(c.DB)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusOK, Message:err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusOK,Message: "route is fetched",Data: routes})
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusOK, Message:err.Error()})
		return
	}

}

func (c *Controller) DeleteRouteHandler(w http.ResponseWriter, r *http.Request) {

	err := database.DeleteRoute(c.DB, mux.Vars(r)["id"])
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusOK, Message:err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusOK, Message:"routes is deleted"})
}
