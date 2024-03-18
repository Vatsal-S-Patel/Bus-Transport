package socket

import (
	"busproject/database"
	"busproject/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
)

// var m = map[string]net.Addr{}

func InitSocket(db *sql.DB) *socketio.Server {
	server := socketio.NewServer(&engineio.Options{
		PingInterval: 5 * time.Second,
		PingTimeout:  10 * time.Second,
	})

	server.OnConnect("", func(s socketio.Conn) error {
		// fmt.Println(s.RemoteAddr(), " is connected")
		fmt.Println("connected:", s.ID())
		// m[s.ID()] = s.RemoteAddr()
		// fmt.Println(m)
		// s.Close()
		fmt.Println("now active connections are", server.Count())
		return nil
	})

	server.OnEvent("/", "update", func(s socketio.Conn, msg string, routeid int) {
		// fmt.Println(s.RemoteAddr(), " got socket for update the bus")

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
			s.Emit("err", err.Error())
		}
		server.BroadcastToRoom("/", fmt.Sprintf("%d", routeid), "update", data, routeid)
		server.BroadcastToRoom("/", fmt.Sprintf("%dbus", data.BusId), "update", data, routeid)
	})

	server.OnEvent("/", "busSelected", func(s socketio.Conn, busId int) {
		// fmt.Println(s.RemoteAddr(), " got socket for bus selected")
		server.LeaveAllRooms("/", s)
		s.Join(fmt.Sprintf("%dbus", busId))
		if len(s.Rooms()) == 0 {
			s.Emit("roomJoined", fmt.Sprint("{code:200,message:'no rooms to join !1',", "ids:,", s.Rooms(), "}"))
			return
		}
		s.Emit("roomJoined", fmt.Sprint("{code:200,message:'you joined the rooms',", "ids:,", s.Rooms(), "}"))
	})

	server.OnEvent("/", "sourceSelected", func(s socketio.Conn, routeId []int) {
		// fmt.Println(s.RemoteAddr(), " got socket for source selected")
		server.LeaveAllRooms("/", s)
		for _, v := range routeId {
			s.Join(fmt.Sprintf("%d", v))
		}
		if len(s.Rooms()) == 0 {
			s.Emit("roomJoined", fmt.Sprint("{code:200,message:'no rooms to join !1',", "ids:,", s.Rooms(), "}"))
			return
		}
		s.Emit("roomJoined", fmt.Sprint("{code:200,message:'you joined the rooms',", "ids:,", s.Rooms(), "}"))
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) {
		fmt.Println("got bye from client")
		// fmt.Println(s.RemoteAddr(), " is disconnected with bye")
		s.Emit("bye", "bye bye")
		// delete(m, s.ID())
		err := s.Close()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("now active connections are", server.Count())
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
		// fmt.Println(s.RemoteAddr(), " is disconnectd due to error", e)

		// fmt.Println(m)
		err := s.Close()
		fmt.Println("now active connections are", server.Count())
		// print("hello")
		if err != nil {
			fmt.Println(err)
		}
	})

	server.OnDisconnect("", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
		server.LeaveAllRooms("/", s)
		// fmt.Println(s.RemoteAddr(), " is disconnected")
		// delete(m, s.ID())
		fmt.Println("now active connections are", server.Count())
		err := s.Close()
		if err != nil {
			fmt.Println(err)
		}
		s.Emit("disconnect", reason)
	})

	server.OnEvent("/", "disconnect", func(s socketio.Conn, reason string) {
		fmt.Println("a client is disconnected", reason)
		// fmt.Println(s.RemoteAddr(), " is disconnected")
		server.LeaveAllRooms("/", s)
		// delete(m, s.ID())
		err := s.Close()
		fmt.Println("now connections are ", server.Count())
		if err != nil {
			fmt.Println(err)
		}
	})

	go func() {
		fmt.Println("socket is listening")
		server.Serve()
		defer server.Close()

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
