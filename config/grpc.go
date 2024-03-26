package config

import (
	"log"

	"google.golang.org/grpc"
)

func connectGPRC() {
	var errProfile error
	clientProfile, errProfile = grpc.Dial("host.docker.internal:20001", grpc.WithInsecure())
	if errProfile != nil {
		log.Fatalln(errProfile)
	}
}
