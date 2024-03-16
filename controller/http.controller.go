package controller

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Error   error       `json:"error"`
}

type MetaResponse struct {
	Data     []interface{} `json:"data"`
	Message  string        `json:"message"`
	Status   int           `json:"status"`
	Error    error         `json:"error"`
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
	Total    int           `json:"total"`
}

func badRequest(w http.ResponseWriter, r *http.Request, err error) {
	res := Response{
		Data:    nil,
		Message: err.Error(),
		Status:  400,
		Error:   err,
	}
	w.WriteHeader(http.StatusBadRequest)
	render.JSON(w, r, res)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	res := Response{
		Data:    nil,
		Message: err.Error(),
		Status:  500,
		Error:   err,
	}
	w.WriteHeader(http.StatusInternalServerError)
	render.JSON(w, r, res)
}

func handleError(w http.ResponseWriter, r *http.Request, err error, status int) {
	res := Response{
		Data:    nil,
		Message: err.Error(),
		Status:  status,
		Error:   err,
	}
	w.WriteHeader(status)
	render.JSON(w, r, res)
}
