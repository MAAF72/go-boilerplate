package users

import (
	"net/http"

	"github.com/MAAF72/go-boilerplate/models"
	"github.com/MAAF72/go-boilerplate/services"
	"github.com/gin-gonic/gin"
)

// GetUserByID get user by id
func GetUserByID(c *gin.Context) {
	var user models.User
	var err error

	if err = c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user, err = services.Instance.GetUserByID(c, user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserByID update user by id
func UpdateUserByID(c *gin.Context) {
	var changeSet models.UserChangeSet
	var err error

	if err = c.BindJSON(&changeSet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	id := c.Param("id")
	user, err := services.Instance.UpdateUserByID(c, id, changeSet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUserByID delete user by id
func DeleteUserByID(c *gin.Context) {
	var err error

	id := c.Param("id")
	err = services.Instance.DeleteUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, true)
}
