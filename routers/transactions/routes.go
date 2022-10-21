package transactions

import (
	"github.com/MAAF72/go-boilerplate/middlewares"
	"github.com/gin-gonic/gin"
)

// Transactions router group const
const (
	PREFIX string = "/transactions"
)

// Routes routes transactions API
func Routes(mainRouter *gin.RouterGroup) {
	r := mainRouter.Group(PREFIX, middlewares.JWTAuthMiddleware())
	{
		r.POST("/", CreateTransaction)
		r.POST("/:id/apply-promo", TransactionApplyPromo)
		r.GET("/:id", GetTransactionByID)
		r.GET("/:id/detail", GetTransactionDetailByID)
		r.PUT("/:id", UpdateTransactionByID)
		r.DELETE("/:id", DeleteTransactionByID)
	}
}
