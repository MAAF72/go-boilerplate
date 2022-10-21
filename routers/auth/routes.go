package auth

import (
	"github.com/MAAF72/go-boilerplate/middlewares"
	"github.com/gin-gonic/gin"
)

// Auth router group const
const (
	PREFIX string = "/auth"
)

// Routes routes auth API
func Routes(mainRouter *gin.RouterGroup) {
	r := mainRouter.Group(PREFIX)
	{
		r.POST("/register", Register)
		r.POST("/login", Login)
		r.GET("/validate", middlewares.JWTAuthMiddleware(), Validate)
	}
}
