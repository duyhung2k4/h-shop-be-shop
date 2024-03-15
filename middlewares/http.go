package middlewares

import (
	"app/controller"
	"net/http"

	"github.com/go-chi/render"
)

func authServerError(w http.ResponseWriter, r *http.Request, err error) {
	res := controller.Response{
		Data:    nil,
		Message: err.Error(),
		Status:  401,
		Error:   err,
	}
	w.WriteHeader(http.StatusUnauthorized)
	render.JSON(w, r, res)
}
