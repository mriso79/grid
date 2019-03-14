package utils

import (
	"encoding/json"
	"net/http"
)

//SendJSONResponse sends a struct to JSON
func SendJSONResponse(w http.ResponseWriter, m interface{}) {
	js, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
