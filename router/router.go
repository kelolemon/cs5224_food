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
	api.GET("/diners", method.GetAllDiners)
	api.POST("/diner", method.CreateDiner)
	api.POST("/comment", method.CreateComment)
	api.GET("/comment", method.GetComments)
	api.POST("/dialog", method.CreateDialog)
	api.GET("/dialog", method.GetDialogs)
	api.GET("/ads", method.GetAds)
}
