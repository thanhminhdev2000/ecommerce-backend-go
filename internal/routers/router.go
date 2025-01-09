package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/thanhminhdev2000/ecommerce-backend-go/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/user/1", c.NewUserController().GetUserByID)
	}

	return r
}

