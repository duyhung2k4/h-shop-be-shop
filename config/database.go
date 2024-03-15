package config

import (
	"app/model"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectPostgresql(migrate bool) error {
	var err error
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return err
	}

	if migrate {
		errMigrate := db.AutoMigrate(
			&model.Shop{},
			&model.TypeProduct{},
		)
		if errMigrate != nil {
			return errMigrate
		}
	}

	return nil
}

func connectRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     urlRedis,
		Password: "",
		DB:       0,
	})
}
