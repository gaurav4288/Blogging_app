package dto

import "github.com/gaurav4288/go_tutorial/blogging_app/internal/models"

type (
	UserResponse struct {
		UserId    string `json:"user_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}

	UserResponseWithUserId struct {
		UserResponse
		UserId string `bson:"user_id" json:"user_id"`
	}

	UserCreateRequest struct {
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
		Username  string `json:"username" validate:"required"`
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required,min=6"`
	}

	UserUpdateRequest struct {
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
		Email     string `json:"email" validate:"required,email"`
	}
)

func (u *UserCreateRequest) ToModel() *models.User {
	return &models.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
	}
}

func (u *UserUpdateRequest) ToModel() *models.User {
	return &models.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}

func ToUserResponse(user *models.User) *UserResponseWithUserId {
	resp := &UserResponseWithUserId{
		UserId: user.UserId,
	}
	resp.FirstName = user.FirstName
	resp.LastName = user.LastName
	resp.Email = user.Email
	return resp
}
