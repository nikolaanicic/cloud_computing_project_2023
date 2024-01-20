package requestmodels

import (
	"encoding/json"
	"fmt"
)

type InsertUserRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *InsertUserRequest) String() string {
	return fmt.Sprintf("(%s, %s, %s)", r.Name, r.LastName, r.Username)
}

func (r InsertUserRequest) AsJson() []byte {
	data, _ := json.Marshal(r)

	return data
}

func NewInsertUserRequest(name, lastName, username, password string) InsertUserRequest {
	return InsertUserRequest{
		Name:     name,
		LastName: lastName,
		Username: username,
		Password: password,
	}
}
