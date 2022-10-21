package items

import (
	"net/http"

	"github.com/MAAF72/go-boilerplate/models"
	"github.com/MAAF72/go-boilerplate/services"
	"github.com/gin-gonic/gin"
)

// CreateItem create item
func CreateItem(c *gin.Context) {
	var request models.ItemCreateRequest
	var err error

	if err = c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	item, err := services.Instance.CreateItem(c, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}

// GetItemByID get item by id
func GetItemByID(c *gin.Context) {
	var item models.Item
	var err error

	if err = c.ShouldBindUri(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	item, err = services.Instance.GetItemByID(c, item.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, item)
}

// UpdateItemByID update item by id
func UpdateItemByID(c *gin.Context) {
	var changeSet models.ItemChangeSet
	var err error

	if err = c.BindJSON(&changeSet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	id := c.Param("id")
	item, err := services.Instance.UpdateItemByID(c, id, changeSet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DeleteItemByID delete item by id
func DeleteItemByID(c *gin.Context) {
	var err error

	id := c.Param("id")
	err = services.Instance.DeleteItemByID(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, true)
}
