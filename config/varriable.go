package config

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const (
	APP_PORT    = "APP_PORT"
	DB_HOST     = "DB_HOST"
	DB_PORT     = "DB_PORT"
	DB_NAME     = "DB_NAME"
	DB_PASSWORD = "DB_PASSWORD"
	DB_USER     = "DB_USER"
	URL_REDIS   = "URL_REDIS"
)

var (
	appPort    string
	dbHost     string
	dbPort     string
	dbName     string
	dbPassword string
	dbUser     string
	urlRedis   string

	db  *gorm.DB
	rdb *redis.Client
	jwt *jwtauth.JWTAuth
)
