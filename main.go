package main

import (
	"app/config"
	"app/router"
	"log"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr:           ":" + config.GetAppPort(),
		Handler:        router.Router(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatalln(server.ListenAndServe())
}
