package controllers

import (
	"encoding/json"
	m "grid/models"
	u "grid/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	couchdb "github.com/rhinoman/couchdb-go"
)

//NewUser creates a new user - posting a json User struct
func NewUser(w http.ResponseWriter, req *http.Request, db *couchdb.Database) (string, error) {

	decoder := json.NewDecoder(req.Body)
	var newUser m.User
	err := decoder.Decode(&newUser)
	if err != nil {
		panic(err)
	}

	//Generates unique UUID
	u1 := u.GetUUID()
	currentTime := time.Now()

	newUser.CreatedAt = currentTime.Format("2006-01-02 15:04:05")
	newUser.UpdatedAt = currentTime.Format("2006-01-02 15:04:05")
	newUser.Doctype = "user"

	return db.Save(newUser, u1.String(), "")

}

//UpdateUser updates a user - posting a json User struct
func UpdateUser(w http.ResponseWriter, req *http.Request, db *couchdb.Database) (string, error) {

	decoder := json.NewDecoder(req.Body)
	var User m.User
	err := decoder.Decode(&User)

	if err != nil {
		panic(err)
	}

	response := m.Responses{}

	//Generates unique UUID
	currentTime := time.Now()

	User.UpdatedAt = currentTime.Format("2006-01-02 15:04:05")
	User.Doctype = "user"

	if User.Rev == "" {
		response.Status = "Error"
		response.Message = "Empty Rev Value"

		u.SendJSONResponse(w, response)

		return "", err
	}

	return db.Save(User, User.Id, User.Rev)

}

//GetAllUsers Returns All Users Data
func GetAllUsers(w http.ResponseWriter, req *http.Request, db *couchdb.Database) {
	result := m.ViewResponse{}
	//now try to query the view
	err := db.GetView("users", "users", &result, nil)

	if err != nil {
		panic(err)
	}

	u.SendJSONResponse(w, result)
}

//GetUser Returns User Data by ID
func GetUser(w http.ResponseWriter, req *http.Request, db *couchdb.Database) {
	vars := mux.Vars(req)
	user := m.User{}

	rev, err := db.Read(vars["key"], &user, nil)

	if err != nil {
		panic(err)
	}

	if rev == "" {
		panic(rev)
	}

	u.SendJSONResponse(w, user)
}
