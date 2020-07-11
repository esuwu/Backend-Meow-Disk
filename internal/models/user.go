package models

//easyjson:json
type User struct {
	ID 		 	int    `json:"id,omitempty"`
	Email 		string `json:"email"`
	Phone 	 	string `json:"phone"`
	Password    string `json:"password,omitempty"`
}