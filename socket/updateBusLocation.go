package socket

import (
	"busproject/model"
	"encoding/json"
	"fmt"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

// Easier to get running with CORS. Thanks for help @Vindexus and @erkie
// var allowOriginFunc = func(r *http.Request) bool {
// 	return true
// }

func InitSocket() *socketio.Server {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "update", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		var data = model.BusStatus{}
		err := json.Unmarshal([]byte(msg), &data)
		if err != nil {
			log.Println(err)
			s.Emit("err", err.Error())
		}
		server.BroadcastToNamespace("/", "update", data)
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		fmt.Println("got bye from client")
		last := s.Context().(string)
		s.Emit("bye", "bye bye")
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go func() {
		fmt.Println("socket is listening")
		server.Serve()
	}()

	// http.Handle("/socket.io/", server)
	// http.Handle("/", http.FileServer(http.Dir("./asset")))
	// log.Println("Serving at localhost:8000...")
	// log.Fatal(http.ListenAndServe(":8000", nil))
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
