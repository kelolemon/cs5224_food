package router

import (
	"food/method"

	"github.com/gin-gonic/gin"
)

func InitRouters(e *gin.Engine) {
	// r := e.Group("")
	api := e.Group("api")
	api.POST("/user", method.CreateNewUser)
	api.POST("/login", method.CreateLoginSession)
	api.GET("/user", method.GetUser)
}
