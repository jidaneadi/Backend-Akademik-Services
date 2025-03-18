package exceptions

import (
	"fmt"
	"net/http"
	"project-sia/helpers"
	"project-sia/models/response"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, rq *http.Request, err interface{}) {

	if errorNotFound(w, rq, err) {
		return
	}
	if validationError(w, rq, err) {
		return
	}
	if errorBadRequest(w, rq, err) {
		return
	}
	if errorFileUnsupported(w, rq, err) {
		return
	}
	internalServerError(w, rq, err)
}

func validationError(w http.ResponseWriter, rq *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		res := response.Error{
			Meta: response.Meta{
				Code:    400,
				Status:  "BAD REQUEST",
				Message: exception.Error(),
			},
		}
		helpers.WriteToResBody(w, res)
		return true
	} else {
		return false
	}
}

func errorNotFound(w http.ResponseWriter, rq *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		res := response.Error{
			Meta: response.Meta{
				Code:    404,
				Status:  "NOT FOUND",
				Message: exception.Error,
			},
		}

		helpers.WriteToResBody(w, res)
		return true
	} else {
		return false
	}
}

func errorBadRequest(w http.ResponseWriter, rq *http.Request, err interface{}) bool {
	exception, ok := err.(ErrorBadRequest)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		res := response.Error{
			Meta: response.Meta{
				Code:    400,
				Status:  "NOT FOUND",
				Message: exception.Error,
			},
		}

		helpers.WriteToResBody(w, res)
		return true
	} else {
		return false
	}
}

func errorFileUnsupported(w http.ResponseWriter, rq *http.Request, err interface{}) bool {
	exception, ok := err.(ErrorUnsupported)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		res := response.Error{
			Meta: response.Meta{
				Code:    415,
				Status:  "FILE UNSUPPORTED",
				Message: exception.Error,
			},
		}

		helpers.WriteToResBody(w, res)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, rq *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	res := response.Error{
		Meta: response.Meta{
			Code:    500,
			Status:  "INTERNAL SERVER ERROR",
			Message: fmt.Sprintf("%v", err),
		},
	}

	helpers.WriteToResBody(w, res)
}
