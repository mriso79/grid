package main

import (
	controllers "grid/controllers"
	d "grid/drivers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Initialize CouchDb
	db, err := d.GetConnection()

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", controllers.HomeInfo).Methods("GET")

	r.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		controllers.NewUser(w, req, db)
	}).Methods("POST")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
