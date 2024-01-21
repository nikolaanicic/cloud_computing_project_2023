package models

import "encoding/json"

type User struct {
	ID       int64  `json:"id"`
	Name     string `jsin:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Rentals  int64  `json:"rentals"`
}

func (u User) AsJson() []byte {
	data, _ := json.Marshal(u)

	return data
}

func (u User) FieldTypes() []string {
	return []string{"int64", "string", "string", "string", "string", "int64"}
}

func (u User) DataFields() []any {
	return []any{u.Name, u.LastName, u.Username, u.Password}
}

func (u *User) SetFields(fields []any) {
	id, _ := fields[0].(*int64)
	name, _ := fields[1].(*string)
	lastName, _ := fields[2].(*string)
	username, _ := fields[3].(*string)
	password, _ := fields[4].(*string)

	rentals, _ := fields[5].(*int64)

	u.ID = *id
	u.Name = *name
	u.LastName = *lastName
	u.Username = *username
	u.Rentals = *rentals
	u.Password = *password
}

func NewUser(id int64, name, lastName, username string, password string, rentals int64) *User {
	return &User{
		ID:       id,
		Name:     name,
		LastName: lastName,
		Username: username,
		Password: password,
		Rentals:  rentals,
	}
}
