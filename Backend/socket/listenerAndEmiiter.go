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

// all emitters
func emitError(s socketio.Conn, err error) {
	if err != nil {
		log.Println(err)
		s.Emit("err", err.Error())
	}
}

func emitRoomJoined(s socketio.Conn) {
	if len(s.Rooms()) == 0 {
		s.Emit("roomJoined", fmt.Sprint("{code:200,message:'no rooms to join !1',", "ids:,", s.Rooms(), "}"))
		return
	}
	s.Emit("roomJoined", fmt.Sprint("{code:200,message:'you joined the rooms',", "ids:,", s.Rooms(), "}"))
}

func emitBye(s socketio.Conn) {
	s.Emit("bye", "bye bye")
	err := s.Close()
	emitError(s,err)
}

func emitHello(s socketio.Conn) {
	s.Emit("hello", "you have been registered")
}

func emitMap(s socketio.Conn){
	s.Emit("map",m)
}

func emitDisconnect(s socketio.Conn,reason string){
	s.Emit("disconnect", reason)
}
// On connect event
func listenForConnect(server *socketio.Server) {
	server.OnConnect("", func(s socketio.Conn) error {
		fmt.Println("connected:", s.ID())
		fmt.Println("now active connections are", server.Count())
		return nil
	})
}

// On disconnect event
func listenOnBye(server *socketio.Server, db *sql.DB) {
	server.OnEvent("/", "bye", func(s socketio.Conn) {
		fmt.Println("got bye from client")
		if v, ok := m[s.ID()]; ok {
			err := database.ChangeBusStatus(db, v, 0)
			emitError(s, err)
			log.Println("a bus with id ", v, " disconnected")
			delete(m, s.ID())
		}
		emitBye(s)
	})
}

func listenForDisconnect(server *socketio.Server,db *sql.DB){
	server.OnDisconnect("", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
		server.LeaveAllRooms("/", s)
		if v, ok := m[s.ID()]; ok {
			err := database.ChangeBusStatus(db, v, 0)
			emitError(s,err)
			delete(m, s.ID())
		}
		fmt.Println("now active connections are", server.Count())
		err := s.Close()
		emitError(s,err)
		emitDisconnect(s,reason)
	})
}

func listenOnDisconnect(server *socketio.Server,db *sql.DB){
	server.OnEvent("/", "disconnect", func(s socketio.Conn, reason string) {
		fmt.Println("a client is disconnected", reason)
		server.LeaveAllRooms("/", s)
		if v, ok := m[s.ID()]; ok {
			err := database.ChangeBusStatus(db, v, 0)
			emitError(s,err)
			delete(m, s.ID())
		}
		err := s.Close()
		emitError(s,err)
	})
}
// On error event
func listenForError(server *socketio.Server,db *sql.DB){
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
		if v, ok := m[s.ID()]; ok {
			err := database.ChangeBusStatus(db, v, 0)
			emitError(s,err)
			delete(m, s.ID())
		}
		err := s.Close()
		emitError(s,err)
	})
}
// for client
func listenOnBusSelected(server *socketio.Server) {
	server.OnEvent("/", "busSelected", func(s socketio.Conn, busId int) {
		server.LeaveAllRooms("/", s)
		s.Join(fmt.Sprintf("%dbus", busId))
		emitRoomJoined(s)
	})
}

func listenOnsourceSelected(server *socketio.Server) {
	server.OnEvent("/", "sourceSelected", func(s socketio.Conn, routeId []int) {
		server.LeaveAllRooms("/", s)
		for _, v := range routeId {
			s.Join(fmt.Sprintf("%d", v))
		}
		emitRoomJoined(s)
	})
}

// for client and bus
func listenOnUpdate(server *socketio.Server, db *sql.DB) {
	server.OnEvent("/", "update", func(s socketio.Conn, msg string, routeid int) {
		var data = model.BusStatus{}
		err := json.Unmarshal([]byte(msg), &data)
		emitError(s, err)
		err = database.UpdateLiveBus(db, data)
		emitError(s, err)
		server.BroadcastToRoom("/", fmt.Sprintf("%d", routeid), "update", data, routeid)
		server.BroadcastToRoom("/", fmt.Sprintf("%dbus", data.BusId), "update", data, routeid)
	})
}

// for bus
func listenOnBus(server *socketio.Server, db *sql.DB) {
	server.OnEvent("/", "bus", func(s socketio.Conn, busid int) {
		err := database.ChangeBusStatus(db, busid, 1)
		emitError(s, err)
		m[s.ID()] = busid
		log.Println("a bus is connected")
		emitHello(s)
	})
}

// for debug
func listenOnMap(server *socketio.Server){
	server.OnEvent("/", "map", func(s socketio.Conn) {
		emitMap(s)
	})
}
