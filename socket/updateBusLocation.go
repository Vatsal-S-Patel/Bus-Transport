package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func BusLocation(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Allow requests from any origin
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Allow specified methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")     // Allow specified headers

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		latitude := randomFloat(23, 24)
		longitude := randomFloat(72, 73)

		// Send bus location data to the client
		if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%f,%f", latitude, longitude))); err != nil {
			log.Println("write error:", err)
			return
		}

		requestData := map[string]interface{}{
			"bus_id":             1,
			"lat":                latitude,
			"long":               longitude,
			"last_updated":       "00:00",
			"last_station_order": 1,
			"status":             1,
			"traffic":            1,
		}

		// Marshal the data into JSON format
		jsonData, err := json.Marshal(requestData)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		res, err := http.Post("http://192.168.6.222:8080/api/bus/live/update", "application/json", bytes.NewBuffer(jsonData))

		log.Println(res, err)

		time.Sleep(time.Second)
	}

}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
