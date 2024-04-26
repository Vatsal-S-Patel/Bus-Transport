package main

import (
	"busproject/apis"
	"busproject/database"
	myLog "busproject/logging"
	"busproject/socket"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
)

func main() {
	var debug bool
	var wg sync.WaitGroup
	var r *mux.Router
	var httpServer *http.Server
	var logger *myLog.Logger

	flag.BoolVar(&debug, "debug", false, "start server in debug mode")
	flag.Parse()

	// err := godotenv.Load(".env")
	// if err != nil {
	// log.Println(err.Error())
	// return
	// }

	if debug {
		err := myLog.InitLogger()
		if err != nil {
			log.Println(err)
			return
		}
		err = os.Setenv("debug", "true")
		if err != nil {
			log.Println(err)
			return
		}
		logger, err = myLog.GetLogger()
		if err != nil {
			log.Println(err)
		}
		err = logger.LogThis("logging event is started yaay!!!")
		if err != nil {
			log.Println(err)
		}
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Println("ERROR: DB Connection Error", err)
		return
	}
	defer db.Close()

	server := socket.InitSocket(db)
	app := apis.NewApp(db)
	r = app.InitializeRoutes(server)
	httpServer = &http.Server{
		Handler: r,
	}

	wg.Add(1)
	go startServer(httpServer, server, &wg)

	context, cancel := context.WithCancel(context.Background())
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	wg.Add(1)
	go func() {
		cancel()
		gradullayShutDownServers(context, httpServer, server, interruptChan, &wg)
	}()

	wg.Wait()
	log.Println("everything is closed")
}

func startServer(httpServer *http.Server, socketserver *socketio.Server, wg *sync.WaitGroup) error {
	defer wg.Done()
	server_port, ok := os.LookupEnv("SERVER_PORT")

	// if errors.Is(err, configs.ErrDataNotExist) {
	if !ok {
		server_port = "8080"
	}

	httpServer.Addr = ":" + server_port

	go func() {
		log.Println("starting server at ", server_port)
		err := httpServer.ListenAndServe()
		log.Println(err)
	}()

	go func() {
		fmt.Println("socket is listening")
		err := socketserver.Serve()
		log.Println(err)
	}()

	return nil
}

func gradullayShutDownServers(ctx context.Context, httpServer *http.Server, socketServer *socketio.Server, c chan os.Signal, wg *sync.WaitGroup) error {
	defer wg.Done()

	<-c
	err := socketServer.Close()
	if err != nil {
		return err
	}
	log.Println("socket server is closed")

	err = httpServer.Shutdown(ctx)
	if err != nil {
		return err
	}

	log.Println("http server is closed")
	return nil
}
