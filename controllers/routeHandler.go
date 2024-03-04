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
	var route map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&route)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	log.Println(route)
	var routeData = model.Route{
		Id : int(route["id"].(float64)),
		Name : route["name"].(string),
		Status : int(route["status"].(float64)),
		Source : int(route["source"].(float64)),
		Destination: int(route["destination"].(float64)),
	}
	var array = route["Mapping"].([]interface{})
	log.Println(route , routeData, array)

	err = database.InsertRoute(c.DB, routeData)
	if err!= nil {
		log.Println(err.Error())
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	
	for _,v := range array{
		log.Println("inserting staiton with ",routeData.Id," name ", v.(map[string]interface{})["station_id"]," with order " , v.(map[string]interface{})["station_order"])
		err := database.InsertRouteStation(c.DB,model.RouteStation{RouteId: routeData.Id , StationId: int(v.(map[string]interface{})["station_id"].(float64)) , StationOrder: int(v.(map[string]interface{})["station_order"].(float64))})

		if err != nil {
			log.Println(err.Error())
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}

		// for k,y := range v.(map[interface{}]interface{}) {
			// log.Println("inserting staiton with ",routeData.Id," name ",k," with order " , int(y.(flaot64)))
		// }
	}
	// err = json.NewEncoder(w).Encode(route)
	// if err != nil {
		// log.Println(err.Error())
		// return
	// }

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(routeData.Name + "Inserted Successfully"))

}

func (c *Controller) GetAllRouteHandler(w http.ResponseWriter, r *http.Request) {

	routes, err := database.GetAllRoute(c.DB)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(routes)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
