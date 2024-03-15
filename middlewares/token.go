package middlewares

import (
	"app/config"
	"app/utils"
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type middlewares struct {
	utils utils.JwtUtils
	rdb   *redis.Client
}

type Middlewares interface {
	ValidateExpAccessToken() func(http.Handler) http.Handler
}

func (m *middlewares) ValidateExpAccessToken() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		funcHttp := func(w http.ResponseWriter, r *http.Request) {
			tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
			mapData, errMapData := m.utils.JwtDecode(tokenString)

			if errMapData != nil {
				authServerError(w, r, errMapData)
				return
			}

			log.Println(mapData)

			exp := mapData["exp"].(time.Time)

			if time.Now().Unix() > exp.Unix() {
				authServerError(w, r, errors.New("token expired"))
				return
			}

			profileId := strconv.Itoa(int(mapData["profile_id"].(float64)))

			accessToken, errKeyAccessToken := m.rdb.Get(context.Background(), "access_token:"+profileId).Result()
			if errKeyAccessToken != nil {
				authServerError(w, r, errKeyAccessToken)
				return
			}
			refreshToken, errKeyRefreshToken := m.rdb.Get(context.Background(), "refresh_token:"+profileId).Result()
			if errKeyRefreshToken != nil {
				authServerError(w, r, errKeyRefreshToken)
				return
			}

			if accessToken != tokenString && refreshToken != tokenString {
				authServerError(w, r, errors.New("token not exist"))
				return
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(funcHttp)
	}
}

func NewMiddlewares() Middlewares {
	return &middlewares{
		utils: utils.NewJwtUtils(),
		rdb:   config.GetRDB(),
	}
}
