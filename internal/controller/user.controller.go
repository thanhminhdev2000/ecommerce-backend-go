package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhminhdev2000/ecommerce-backend-go/internal/service"
)

type UserController struct{
	userSerive *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userSerive: service.NewUserService(),
	}
}

// controller -> service -> repo -> model -> db
func (uc *UserController) GetUserByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userSerive.GetInfoUser(),
	})
}