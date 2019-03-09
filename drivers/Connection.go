package drivers

import (
	"time"

	couchdb "github.com/rhinoman/couchdb-go"
)

/* This function get the Connection */
func GetConnection() (*couchdb.Database, error) {
	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("127.0.0.1", 5984, timeout)
	auth := couchdb.BasicAuth{Username: "admin", Password: "admin"}
	db := conn.SelectDB("infoball", &auth)

	if err != nil {
		return nil, err
	}

	return db, nil
}
