package main

import (
	"app/config"
	"app/router"
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

func main() {

	go func() {
		grpcServer := grpc.NewServer()
		listenerGRPC, err := net.Listen("tcp", ":20002")

		if err != nil {
			log.Fatalln(listenerGRPC)
		}

		log.Fatalln(grpcServer.Serve(listenerGRPC))
	}()

	server := http.Server{
		Addr:           ":" + config.GetAppPort(),
		Handler:        router.Router(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatalln(server.ListenAndServe())
}
