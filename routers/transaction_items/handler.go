package transactionitems

import (
	"net/http"

	"github.com/MAAF72/go-boilerplate/models"
	"github.com/MAAF72/go-boilerplate/services"
	"github.com/gin-gonic/gin"
)

// CreateTransactionItem create transaction item
func CreateTransactionItem(c *gin.Context) {
	var request models.TransactionItemCreateRequest
	var err error

	if err = c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	transactionItem, err := services.Instance.CreateTransactionItem(c, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactionItem,
	})
}

// GetTransactionItemByID get transaction item by id
func GetTransactionItemByID(c *gin.Context) {
	var transactionItem models.TransactionItem
	var err error

	if err = c.ShouldBindUri(&transactionItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	transactionItem, err = services.Instance.GetTransactionItemByID(c, transactionItem.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transactionItem)
}

// UpdateTransactionItemByID update transaction item by id
func UpdateTransactionItemByID(c *gin.Context) {
	var changeSet models.TransactionItemChangeSet
	var err error

	if err = c.BindJSON(&changeSet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	id := c.Param("id")
	transactionItem, err := services.Instance.UpdateTransactionItemByID(c, id, changeSet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transactionItem)
}

// DeleteTransactionItemByID delete transaction item by id
func DeleteTransactionItemByID(c *gin.Context) {
	var err error

	id := c.Param("id")
	err = services.Instance.DeleteTransactionItemByID(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, true)
}
