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

var m = map[string]int{}

func InitSocket(db *sql.DB) *socketio.Server {
	server := socketio.NewServer(&engineio.Options{
		PingInterval: 1 * time.Second,
		PingTimeout:  10 * time.Second,
	})

	server.OnConnect("", func(s socketio.Conn) error {
		fmt.Println("connected:", s.ID())
		fmt.Println("now active connections are", server.Count())
		return nil
	})

	server.OnEvent("/", "update", func(s socketio.Conn, msg string, routeid int) {
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
		server.LeaveAllRooms("/", s)
		s.Join(fmt.Sprintf("%dbus", busId))
		if len(s.Rooms()) == 0 {
			s.Emit("roomJoined", fmt.Sprint("{code:200,message:'no rooms to join !1',", "ids:,", s.Rooms(), "}"))
			return
		}
		s.Emit("roomJoined", fmt.Sprint("{code:200,message:'you joined the rooms',", "ids:,", s.Rooms(), "}"))
	})

	server.OnEvent("/", "sourceSelected", func(s socketio.Conn, routeId []int) {
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
		if v,ok := m[s.ID()];ok{
			database.ChangeBusStatus(db,v,0)
			log.Println("a bus with id ",v ," disconnected")
			delete(m,s.ID())
		}
		s.Emit("bye", "bye bye")
		err := s.Close()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("now active connections are", server.Count())
	})

	server.OnEvent("/", "bus", func(s socketio.Conn, busid int) {
		err := database.ChangeBusStatus(db, busid, 1)
		if err != nil {
			log.Println(err)
			s.Close()
			s.Emit("error", err.Error)
			return
		}
		m[s.ID()] = busid
		log.Println("a bus is connected")
		s.Emit("hello", "you have been registered")
	})

	server.OnEvent("/","map",func(s socketio.Conn){
		s.Emit("map",m)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
		if v, ok := m[s.ID()]; ok {
			database.ChangeBusStatus(db, v, 0)
			delete(m, s.ID())
		}
		err := s.Close()
		fmt.Println("now active connections are", server.Count())
		if err != nil {
			fmt.Println(err)
		}
	})

	server.OnDisconnect("", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
		server.LeaveAllRooms("/", s)
		if v, ok := m[s.ID()]; ok {
			database.ChangeBusStatus(db, v, 0)
			delete(m, s.ID())
		}
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
		if v,ok := m[s.ID()];ok{
			database.ChangeBusStatus(db,v,0)
			log.Println("a bus with id ",v ," disconnected")
			delete(m,s.ID())
		}
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
