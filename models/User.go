package models

//User type
type User struct {
	Name      string `json:"Name,omitempty"`
	Email     string `json:"Email,omitempty"`
	Password  string `json:"Password,omitempty"`
	Doctype   string `json:"Doctype,omitempty"`
	Birthday  string `json:"Birthday,omitempty"`
	CreatedAt string `json:"CreatedAt,omitempty"`
	UpdatedAt string `json:"UpdatedAt,omitempty"`
}
