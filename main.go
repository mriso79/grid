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

	r.HandleFunc("/user", func(w http.ResponseWriter, req *http.Request) {
		controllers.NewUser(w, req, db)
	}).Methods("POST")

	r.HandleFunc("/user/update", func(w http.ResponseWriter, req *http.Request) {
		controllers.UpdateUser(w, req, db)
	}).Methods("POST")

	r.HandleFunc("/user/{key}", func(w http.ResponseWriter, req *http.Request) {
		controllers.GetUser(w, req, db)
	}).Methods("POST")

	r.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		controllers.GetAllUsers(w, req, db)
	}).Methods("GET")

	r.HandleFunc("/find", func(w http.ResponseWriter, req *http.Request) {
		controllers.SearchUser(w, req, db)
	}).Methods("POST")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
