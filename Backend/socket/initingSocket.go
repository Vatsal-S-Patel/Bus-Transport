package socket

import (
	"database/sql"
	"fmt"
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

	go func() {
		fmt.Println("socket is listening")
		server.Serve()
		defer server.Close()
	}()

	return server
}
