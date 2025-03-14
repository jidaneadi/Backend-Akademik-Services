package helpers

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, data interface{}) {
	d := json.NewDecoder(r.Body)
	err := d.Decode(data)
	ErrorBadRequest(err)
}

func WriteToResponseBody(w http.ResponseWriter, m interface{}) {
	w.Header().Add("Content-Type", "application/json")
	e := json.NewEncoder(w)
	err := e.Encode(m)
	PanicErr(err)
}
