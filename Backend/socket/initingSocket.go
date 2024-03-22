package socket

import (
	"database/sql"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
)

var m = map[string]int{}

func InitSocket(db *sql.DB) *socketio.Server {
	server := socketio.NewServer(&engineio.Options{
		PingInterval: 1 * time.Second,
		PingTimeout:  10 * time.Second,
	},
	)

	listenForConnect(server)

	listenOnUpdate(server, db)

	listenOnBusSelected(server)

	listenOnsourceSelected(server)

	listenOnBye(server, db)

	listenOnBus(server, db)

	listenOnMap(server)

	listenForError(server, db)

	listenForDisconnect(server, db)

	listenOnDisconnect(server, db)

	return server
}

// 	for {
// 		latitude := randomFloat(23, 24)
// 		longitude := randomFloat(72, 73)

// 		// Send bus location data to the client
// 		if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%f,%f", latitude, longitude))); err != nil {
// 			log.Println("write error:", err)
// 			return
// 		}

// 		requestData := map[string]interface{}{
// 			"bus_id":             1,
// 			"lat":                latitude,
// 			"long":               longitude,
// 			"last_updated":       "00:00",
// 			"last_station_order": 1,
// 			"status":             1,
// 			"traffic":            1,
// 		}

// 		// Marshal the data into JSON format
// 		jsonData, err := json.Marshal(requestData)
// 		if err != nil {
// 			fmt.Println("Error marshalling JSON:", err)
// 			return
// 		}

// 		res, err := http.Post("http://192.168.6.222:8080/api/bus/live/update", "application/json", bytes.NewBuffer(jsonData))

// 		log.Println(res, err)

// 		time.Sleep(time.Second)
// 	}

// }
