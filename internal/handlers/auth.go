package handlers

import (
	"net/http"

	"github.com/gaurav4288/go_tutorial/blogging_app/internal/models"
	"github.com/gaurav4288/go_tutorial/blogging_app/internal/service"
	"github.com/gaurav4288/go_tutorial/blogging_app/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.UserService
}

func RefreshToken(c *gin.Context) {

}

func (h *AuthHandler) CreateUser(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.authService.CreateUser(&user)
	if err != nil {
		utils.ErrorResponse("registration failed", err)
		return
	}

	utils.SuccessResponse("created", gin.H{
		"status": http.StatusCreated,
		"user":   user,
	})
}

func GetUser(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func UpdateUser(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func DeleteUser(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func GetAllUsers(c *gin.Context) {

}

func LoginUser(c *gin.Context) {

}

func LogoutUser(c *gin.Context) {

}

func ChangePassword(c *gin.Context) {

}
