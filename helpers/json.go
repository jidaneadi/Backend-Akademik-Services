package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ReadReqBody(rq *http.Request, data interface{}) {
	if rq.Body == nil {
		PanicErr(errors.New("request body nil"))
	}
	decoder := json.NewDecoder(rq.Body)
	err := decoder.Decode(data)
	PanicErr(err)
}

func WriteToResBody(w http.ResponseWriter, rs interface{}) {
	w.Header().Add("Content-Type", "application/json")
	e := json.NewEncoder(w)
	err := e.Encode(rs)
	PanicErr(err)
}
