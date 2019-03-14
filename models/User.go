package models

//User type
type User struct {
	Id        string `json:"_id,omitempty"`
	Rev       string `json:"_rev,omitempty"`
	Name      string `json:"Name,omitempty"`
	Email     string `json:"Email,omitempty"`
	Password  string `json:"Password,omitempty"`
	Doctype   string `json:"Doctype,omitempty"`
	Birthday  string `json:"Birthday,omitempty"`
	CreatedAt string `json:"CreatedAt,omitempty"`
	UpdatedAt string `json:"UpdatedAt,omitempty"`
}

//ViewResult Type
type ViewResult struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value User   `json:"value"`
}

//ViewResponse type
type ViewResponse struct {
	TotalRows int          `json:"total_rows"`
	Offset    int          `json:"offset"`
	Rows      []ViewResult `json:"rows,omitempty"`
}
