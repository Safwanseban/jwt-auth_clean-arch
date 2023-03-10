package http

import (
	"github.com/Safwanseban/jwt-auth/pkg/api/handlers"
	"github.com/gin-gonic/gin"
)

func NewServerHttp(router *gin.Engine) *gin.Engine {
	router = gin.Default()
	router.POST("/signup", handlers.Signup)
	router.GET("/login",handlers.Login)

	return router

}
