package responsemodels

import (
	"encoding/json"
	"fmt"
)

type UserResponse struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
}

func (r *UserResponse) String() string {
	return fmt.Sprintf("(%s, %s, %s)", r.Name, r.LastName, r.Username)
}

func (r UserResponse) AsJson() []byte {
	data, _ := json.Marshal(r)

	return data
}

func NewUserResponse(name, lastName, username, password string) UserResponse {
	return UserResponse{
		Name:     name,
		LastName: lastName,
		Username: username,
	}
}
