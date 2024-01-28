package requestmodels

import (
	"encoding/json"
	"fmt"
)

type UserSignUpRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *UserSignUpRequest) String() string {
	return fmt.Sprintf("(%s, %s, %s)", r.Name, r.LastName, r.Username)
}

func (r UserSignUpRequest) AsJson() []byte {
	data, _ := json.Marshal(r)

	return data
}

func NewInsertUserRequest(name, lastName, username, password string) UserSignUpRequest {
	return UserSignUpRequest{
		Name:     name,
		LastName: lastName,
		Username: username,
		Password: password,
	}
}
