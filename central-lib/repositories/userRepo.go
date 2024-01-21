package repositories

import (
	"crypto/sha512"
	"encoding/hex"
	"rac_oblak_proj/data_context"
	"rac_oblak_proj/mapper"
	"rac_oblak_proj/models"
	requestmodels "rac_oblak_proj/request_models"
	responsemodels "rac_oblak_proj/response_models"
)

type UserRepo struct {
	ctx *data_context.DataContext
}

func NewUserRepo(ctx *data_context.DataContext) *UserRepo {
	return &UserRepo{ctx}
}

func (r *UserRepo) hashUser(username, password string) string {
	data := []byte(username)
	data = append(data, []byte(password)...)

	hashed := sha512.Sum512(data)

	return hex.EncodeToString(hashed[:])
}

func (r *UserRepo) ValidatePassword(hashed string, username string, password string) bool {
	return r.hashUser(username, password) == hashed
}

func (r *UserRepo) Insert(user requestmodels.InsertUserRequest) (responsemodels.UserResponse, error) {
	newUser, err := mapper.Map[requestmodels.InsertUserRequest, models.User](user)

	if err != nil {
		return responsemodels.UserResponse{}, err
	}

	newUser.Password = r.hashUser(user.Username, user.Password)
	query := "INSERT INTO users (name, last_name, username, password) VALUES (?,?,?,?)"

	affected, err := data_context.ExecuteInsert[models.User](r.ctx, query, newUser)

	if affected != 1 || err != nil {
		return responsemodels.UserResponse{}, err
	}

	return mapper.Map[models.User, responsemodels.UserResponse](newUser)
}

func (r *UserRepo) GetByUsername(username string) (models.User, error) {

	query := "SELECT * from users where username = ?;"

	result, err := data_context.ExecuteQuery[models.User](r.ctx, query, username)

	if err != nil {
		return models.User{}, err
	}

	return result[0], nil
}
