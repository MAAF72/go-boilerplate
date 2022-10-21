package routers

import (
	"github.com/MAAF72/go-boilerplate/routers/auth"
	"github.com/MAAF72/go-boilerplate/routers/items"
	transactionitems "github.com/MAAF72/go-boilerplate/routers/transaction_items"
	"github.com/MAAF72/go-boilerplate/routers/transactions"
	"github.com/MAAF72/go-boilerplate/routers/users"
	"github.com/gin-gonic/gin"
)

// Router const
const (
	VERSION = "/v1"
)

// RegisterRouters routes all API
func RegisterRouters(app *gin.Engine) {
	r := app.Group(VERSION)
	{
		auth.Routes(r)
		users.Routes(r)
		items.Routes(r)
		transactions.Routes(r)
		transactionitems.Routes(r)
	}
}
