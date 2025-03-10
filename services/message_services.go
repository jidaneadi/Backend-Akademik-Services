package services

import (
	"encoding/json"
	"net/http"
)

type Meta struct {
	Status  bool   `json:"status"`
	Code    uint16 `json:"code"`
	Message string `json:"message"`
}

func MessageError(c http.ResponseWriter, code uint16, message string, status bool) error {
	response := Meta{
		Status:  status,
		Code:    code,
		Message: message,
	}
	c.Header().Add("Content-Type", "application/json")
	c.WriteHeader(int(code))
	return json.NewEncoder(c).Encode(response)
}

func MessageSucces(c http.ResponseWriter, message string, status bool) error {
	res := Meta{
		Code:    200,
		Status:  true,
		Message: message,
	}
	c.Header().Add("Content-Type", "application/json")
	c.WriteHeader(200)
	return json.NewEncoder(c).Encode(res)
}
