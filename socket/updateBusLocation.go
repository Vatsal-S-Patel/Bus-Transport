package socket

import (
	"busproject/database"
	"busproject/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

func InitSocket(db *sql.DB) *socketio.Server {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "update", func(s socketio.Conn, msg string, routeid int) {
		// fmt.Println("notice:", msg)
		var data = model.BusStatus{}
		err := json.Unmarshal([]byte(msg), &data)
		if err != nil {
			log.Println(err)
			s.Emit("err", err.Error())
		}
		err = database.UpdateLiveBus(db, data)
		if err != nil {
			fmt.Println(err.Error())
			s.Emit("err",err.Error())
		}
		server.BroadcastToRoom("/", fmt.Sprintf("%d", routeid), "update", data)
		server.BroadcastToRoom("/", fmt.Sprint(data.BusId), "update", data)
	})

	server.OnEvent("/", "busSelected", func(s socketio.Conn, busId int) {
		s.LeaveAll()
		s.Join(fmt.Sprint(busId))
		s.Emit("roomJoined", fmt.Sprint("{code:200,message:'you joined the rooms',", "ids:,", s.Rooms(), "}"))
	})

	server.OnEvent("/", "sourceSelected", func(s socketio.Conn, routeId []int) {
		fmt.Println(routeId)
		// var routeId []int
		// json.Unmarshal([]byte(input),&routeId)
		s.LeaveAll()
		for _, v := range routeId {
			s.Join(fmt.Sprintf("%d", v))
		}
		s.Emit("roomJoined", fmt.Sprint("{code:200,message:'you joined the rooms',", "ids:,", s.Rooms(), "}"))
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
