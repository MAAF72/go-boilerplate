package items

import (
	"github.com/MAAF72/go-boilerplate/middlewares"
	"github.com/gin-gonic/gin"
)

// Items router group const
const (
	PREFIX string = "/items"
)

// Routes routes items API
func Routes(mainRouter *gin.RouterGroup) {
	r := mainRouter.Group(PREFIX, middlewares.JWTAuthMiddleware())
	{
		r.POST("/", CreateItem)
		r.GET("/:id", GetItemByID)
		r.PUT("/:id", UpdateItemByID)
		r.DELETE("/:id", DeleteItemByID)
	}
}
