package utils

import (
	"app/config"
	"context"
	"net/http"
	"strings"
)

type jwtUtils struct{}

type JwtUtils interface {
	JwtEncode(data map[string]interface{}) (string, error)
	JwtDecode(tokenString string) (map[string]interface{}, error)
	GetMapData(r *http.Request) (map[string]interface{}, error)
}

func (j *jwtUtils) JwtEncode(data map[string]interface{}) (string, error) {
	_, tokenString, err := config.GetJWT().Encode(data)
	return tokenString, err
}

func (j *jwtUtils) JwtDecode(tokenString string) (map[string]interface{}, error) {
	var dataMap map[string]interface{}
	jwt, err := config.GetJWT().Decode(tokenString)
	if err != nil {
		return dataMap, err
	}

	dataMap, errMap := jwt.AsMap(context.Background())
	return dataMap, errMap
}

func (j *jwtUtils) GetMapData(r *http.Request) (map[string]interface{}, error) {
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	mapData, err := j.JwtDecode(tokenString)

	return mapData, err
}

func NewJwtUtils() JwtUtils {
	return &jwtUtils{}
}
