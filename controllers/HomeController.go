package controllers

import (
	"encoding/json"
	m "grid/models"
	"net/http"
)

//HomeInfo json
func HomeInfo(w http.ResponseWriter, r *http.Request) {
	var home m.Home

	home.Title = "System"
	home.Status = "Ok"
	js, err := json.Marshal(home)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
