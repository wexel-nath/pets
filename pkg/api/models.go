package api

import "github.com/wexel-nath/pets/pkg/user"

type UserApiModel struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Status    int64  `json:"userStatus"`
}

func NewUserApiModel(u user.User) UserApiModel {
	return UserApiModel{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		Phone:     u.Phone,
		Status:    u.Status,
	}
}

func (model UserApiModel) ToUser() user.User {
	return user.User{
		ID:        model.ID,
		Username:  model.Username,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Email:     model.Email,
		Password:  model.Password,
		Phone:     model.Phone,
		Status:    model.Status,
	}
}
