package users

import (
	"github.com/MAAF72/go-boilerplate/middlewares"
	"github.com/gin-gonic/gin"
)

// Users router group const
const (
	PREFIX string = "/users"
)

// Routes routes users API
func Routes(mainRouter *gin.RouterGroup) {
	r := mainRouter.Group(PREFIX, middlewares.JWTAuthMiddleware())
	{
		r.GET("/:id", GetUserByID)
		r.PUT("/:id", UpdateUserByID)
		r.DELETE("/:id", DeleteUserByID)
	}
}
