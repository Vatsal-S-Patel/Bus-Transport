package controllers

import (
	"busproject/database"
	"busproject/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Controller) CreateRouteStationHandler(w http.ResponseWriter, r *http.Request) {
	var routeStation model.RouteStation
	err := json.NewDecoder(r.Body).Decode(&routeStation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	err = database.InsertRouteStation(c.DB, routeStation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(routeStation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("RouteStation Inserted Successfully"))
}

func (c *Controller) GetAllRouteStationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	routeStations, err := database.GetAllRouteStation(c.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.OutputStruct{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(routeStations)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (c *Controller) CreateMappingHandler(mapping model.RouteStationMerged, mappingId int) int {

	if len(mapping.RouteStationArray) == 0 {
		return http.StatusBadRequest
	}

	insertQuery := "INSERT INTO transport.routestations (route_id, station_id, station_order) VALUES "

	for _, v := range mapping.RouteStationArray {
		insertQuery += fmt.Sprintf("(%d, %d, %d),", mappingId, v.StationId, v.StationOrder)
	}

	err := database.InsertAllRouteStation(c.DB, insertQuery)
	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest
	}

	log.Println("Mapping Successful")
	return http.StatusOK
}
