package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	c "github.com/thanhminhdev2000/ecommerce-backend-go/internal/controller"
	"github.com/thanhminhdev2000/ecommerce-backend-go/internal/middlewares"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}
}

func CC (c *gin.Context) {
		fmt.Println("Before --> CC")
		c.Next()
		fmt.Println("After --> CC")
	}

func NewRouter() *gin.Engine {
	r := gin.Default()

	// use middleware
	r.Use(middlewares.AuthenMiddleware(), AA(), BB(), CC)
	v1 := r.Group("/v1")
	{
		v1.GET("/user/1", c.NewUserController().GetUserByID)
	}

	return r
}

