package controllers

import (
	"encoding/json"
	m "grid/models"
	"net/http"
	"time"

	couchdb "github.com/rhinoman/couchdb-go"
	uuid "github.com/satori/go.uuid"
)

//NewUser creates a new user
func NewUser(w http.ResponseWriter, req *http.Request, db *couchdb.Database) (string, error) {

	decoder := json.NewDecoder(req.Body)
	var newUser m.User
	err := decoder.Decode(&newUser)
	if err != nil {
		panic(err)
	}

	//Generates unique UUID
	u1 := uuid.Must(uuid.NewV4())
	currentTime := time.Now()

	newUser.CreatedAt = currentTime.Format("2006-01-02 15:04:05")
	newUser.UpdatedAt = currentTime.Format("2006-01-02 15:04:05")
	newUser.Doctype = "user"

	return db.Save(newUser, u1.String(), "")

}
