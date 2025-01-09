package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhminhdev2000/ecommerce-backend-go/internal/service"
	response "github.com/thanhminhdev2000/ecommerce-backend-go/pkg/repsponse"
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
	response.SuccessResponse(c, 20001, []string{"Thanh Minh", "Developer"})
}