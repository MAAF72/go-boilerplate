package transactionitems

import (
	"github.com/MAAF72/go-boilerplate/middlewares"
	"github.com/gin-gonic/gin"
)

// TransactionItems router group const
const (
	PREFIX string = "/transaction-items"
)

// Routes routes transaction items API
func Routes(mainRouter *gin.RouterGroup) {
	r := mainRouter.Group(PREFIX, middlewares.JWTAuthMiddleware())
	{
		r.POST("/", CreateTransactionItem)
		r.GET("/:id", GetTransactionItemByID)
		r.PUT("/:id", UpdateTransactionItemByID)
		r.DELETE("/:id", DeleteTransactionItemByID)
	}
}
