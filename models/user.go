package models

import "encoding/json"

type User struct {
	ID       int64  `json:"id"`
	Name     string `jsin:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
}

func (u User) AsJson() []byte {
	data, _ := json.Marshal(u)

	return data
}

func (u User) FieldTypes() []string {
	return []string{"int64", "string", "string", "string"}
}

func (u User) DataFields() []any {
	return []any{u.ID, u.Name, u.LastName, u.Username}
}

func (u *User) SetFields(fields []any) {
	id, _ := fields[0].(*int64)
	name, _ := fields[1].(*string)
	lastName, _ := fields[2].(*string)
	username, _ := fields[3].(*string)

	u.ID = *id
	u.Name = *name
	u.LastName = *lastName
	u.Username = *username
}

func NewUser(id int64, name, lastName, username string) *User {
	return &User{
		ID:       id,
		Name:     name,
		LastName: lastName,
		Username: username,
	}
}
