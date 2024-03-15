package config

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	return db
}

func GetRDB() *redis.Client {
	return rdb
}

func GetAppPort() string {
	return appPort
}

func GetJWT() *jwtauth.JWTAuth {
	return jwt
}
