package controllers

import (
	"fmt"
	"net/http"
)

func HomeInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}
