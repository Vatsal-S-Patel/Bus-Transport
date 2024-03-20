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
		OutputToClient(w,http.StatusInternalServerError,err.Error(),nil)
		return
	}

	err = database.InsertRouteStation(c.DB, routeStation)
	if err != nil {
		OutputToClient(w,http.StatusInternalServerError,err.Error(),nil)
		return
	}

	OutputToClient(w,http.StatusOK,"route station is created",nil)
}

func (c *Controller) GetAllRouteStationHandler(w http.ResponseWriter, r *http.Request) {

	routeStations, err := database.GetAllRouteStation(c.DB)
	if err != nil {
		OutputToClient(w,http.StatusInternalServerError,err.Error(),nil)
		return
	}

	OutputToClient(w,http.StatusOK,"route station is fetched",routeStations)
}

func (c *Controller) CreateMappingHandler(mapping model.RouteStationMerged, mappingId int) int {

	if len(mapping.RouteStationArray) == 0 {
		return http.StatusInternalServerError
	}

	insertQuery := "INSERT INTO transport.routestations (route_id, station_id, station_order) VALUES "

	for _, v := range mapping.RouteStationArray {
		insertQuery += fmt.Sprintf("(%d, %d, %d),", mappingId, v.StationId, v.StationOrder)
	}

	err := database.InsertAllRouteStation(c.DB, insertQuery)
	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError
	}

	return http.StatusOK
}
