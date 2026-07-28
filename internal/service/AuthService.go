package service

import (
	"time"

	"github.com/gaurav4288/go_tutorial/blogging_app/internal/dto"
	"github.com/gaurav4288/go_tutorial/blogging_app/internal/models"
	"github.com/gaurav4288/go_tutorial/blogging_app/internal/repository"
	"github.com/gaurav4288/go_tutorial/blogging_app/internal/utils"
	"github.com/gin-gonic/gin"
)

type (
	UserService interface {
		CreateUser(user *models.User) error
		GetUserByID(id string) (*models.User, error)
		UpdateUser(user *models.User) error
		DeleteUser(id string) error
	}

	userService struct {
		repository repository.UserRepository
	}
)

func (s *userService) Create(ctx *gin.Context, request dto.UserCreateRequest) error {
	user := request.ToModel()
	timestamp := time.Now()
	_, err := s.repository.IsUserExistsById(user.UserId)
	user.UpdatedAt = timestamp

	if err != nil { ////if we get error retrieving the details, that means the data corresponding to userId dont exist
		user.CreatedAt = timestamp
		//create uuid for userId
		user.UserId = utils.GenerateUUID()
		return s.repository.CreateUser(user)
	}

	return s.repository.UpdateUser(user)
}
