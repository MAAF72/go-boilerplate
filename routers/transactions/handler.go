package transactions

import (
	"net/http"

	"github.com/MAAF72/go-boilerplate/models"
	"github.com/MAAF72/go-boilerplate/services"
	"github.com/gin-gonic/gin"
)

// CreateTransaction create transaction
func CreateTransaction(c *gin.Context) {
	var request models.TransactionCreateRequest
	var err error

	if err = c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	transaction, err := services.Instance.CreateTransaction(c, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transaction,
	})
}

// GetTransactionByID get transaction by id
func GetTransactionByID(c *gin.Context) {
	var transaction models.Transaction
	var err error

	if err = c.ShouldBindUri(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	transaction, err = services.Instance.GetTransactionByID(c, transaction.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// GetTransactionDetailByID get transaction detail by id
func GetTransactionDetailByID(c *gin.Context) {
	var transactionDetail models.TransactionDetail
	var err error

	if err = c.ShouldBindUri(&transactionDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	transactionDetail, err = services.Instance.GetTransactionDetailByID(c, transactionDetail.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transactionDetail)
}

// UpdateTransactionByID update transaction by id
func UpdateTransactionByID(c *gin.Context) {
	var changeSet models.TransactionChangeSet
	var err error

	if err = c.BindJSON(&changeSet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	id := c.Param("id")
	transaction, err := services.Instance.UpdateTransactionByID(c, id, changeSet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// DeleteTransactionByID delete transaction by id
func DeleteTransactionByID(c *gin.Context) {
	var err error

	id := c.Param("id")
	err = services.Instance.DeleteTransactionByID(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, true)
}

// TransactionApplyPromo transaction apply promo
func TransactionApplyPromo(c *gin.Context) {
	var transactionDetail models.TransactionDetail
	var err error

	if err = c.ShouldBindUri(&transactionDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	transactionDetail, err = services.Instance.TransactionApplyPromo(c, transactionDetail.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transactionDetail)
}
