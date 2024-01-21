package requestmodels

import "encoding/json"

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u UserLoginRequest) AsJson() []byte {
	data, _ := json.Marshal(u)

	return data
}

func NewUserLoginRequest(username, password string) UserLoginRequest {
	return UserLoginRequest{
		Username: username,
		Password: password,
	}
}
