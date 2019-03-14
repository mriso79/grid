package utils

import (
	uuid "github.com/satori/go.uuid"
)

//GetUuid returns an uuid
func GetUUID() uuid.UUID {

	//Generates unique UUID
	u1 := uuid.Must(uuid.NewV4())

	return u1
}

//UUIDFromString take a string and return an UUID
func UUIDFromString(s string) (uuid.UUID, error) {
	return uuid.FromString(s)
}
