package config

import (
	"flag"
	"log"

	"github.com/go-chi/jwtauth/v5"
)

func init() {
	var migrate bool = false
	flag.BoolVar(&migrate, "db", true, "Migrate Database?")

	loadEnv()
	jwt = jwtauth.New("HS256", []byte("h-shop"), nil)

	if err := connectPostgresql(migrate); err != nil {
		log.Fatalf("Error connect Postgresql: %v", err)
	}
	connectRedis()
}
