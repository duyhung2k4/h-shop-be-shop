package utils

import (
	"app/config"
	"context"
)

type jwtUtils struct{}

type JwtUtils interface {
	JwtEncode(data map[string]interface{}) (string, error)
	JwtDecode(tokenString string) (map[string]interface{}, error)
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

func NewJwtUtils() JwtUtils {
	return &jwtUtils{}
}
