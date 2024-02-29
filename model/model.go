package model

type Bus struct {
	Id                 int    `json:"id" sql:"id"`
	RegistrationNumber string `json:"registration_number" sql:"registration_number"`
	Model              string `json:"model" sql:"model"`
	Capacity           int    `json:"capacity" sql:"capacity"`
}

type BusStatus struct {
	BusId            int     `json:"bus_id" sql:"bus_id"`
	Lat              float64 `json:"lat" sql:"lat"`
	Long             float64 `json:"long" sql:"long"`
	LastUpdated      float64 `json:"last_updated" sql:"last_updated"`
	Traffic          int     `json:"traffic" sql:"traffic"`
	Status           int     `json:"status" sql:"status"`
	LastStationOrder int     `json:"last_station_order" sql:"last_station_order"`
}

type Driver struct {
	Id     int    `json:"id" sql:"id"`
	Name   string `json:"name" sql:"name"`
	Phone  string `json:"phone" sql:"phone"`
	Gender int    `json:"gender" sql:"gender"`
	Dob    string `json:"dob" sql:"dob"`
}

type Route struct {
	Id          int    `json:"id" sql:"id"`
	Name        string `json:"name" sql:"name"`
	Status      int    `json:"status" sql:"status"`
	Source      int    `json:"source" sql:"source"`
	Destination int    `json:"destination" sql:"destination"`
}

type RouteStation struct {
	RouteId      int `json:"id" sql:"id"`
	StationId    int `json:"station_id" sql:"station_id"`
	StationOrder int `json:"station_order" sql:"station_order"`
}

type Schedule struct {
	Id            int    `json:"id" sql:"id"`
	BusId         int    `json:"bus_id" sql:"bus_id"`
	RouteId       int    `json:"route_id" sql:"route_id"`
	DepartureTime string `json:"dep" sql:"dep"`
}

type Station struct {
	Id   int     `json:"id" sql:"id"`
	Name string  `json:"name" sql:"name"`
	Lat  float64 `json:"lat" sql:"lat"`
	Long float64 `json:"long" sql:"long"`
}
