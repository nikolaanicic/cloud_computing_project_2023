package responsemodels

import (
	"encoding/json"
	"fmt"
)

type InsertUserResponse struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *InsertUserResponse) String() string {
	return fmt.Sprintf("(%s, %s, %s)", r.Name, r.LastName, r.Username)
}

func (r InsertUserResponse) AsJson() []byte {
	data, _ := json.Marshal(r)

	return data
}

func NewInsertUserRequest(name, lastName, username, password string) InsertUserResponse {
	return InsertUserResponse{
		Name:     name,
		LastName: lastName,
		Username: username,
		Password: password,
	}
}
