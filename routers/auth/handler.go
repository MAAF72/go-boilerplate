package auth

import (
	"net/http"

	"github.com/MAAF72/go-boilerplate/middlewares"
	"github.com/MAAF72/go-boilerplate/models"
	"github.com/MAAF72/go-boilerplate/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register register
func Register(c *gin.Context) {
	var request models.UserRegisterRequest
	var err error

	if err = c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user, err := services.Instance.RegisterUser(c, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// Login login
func Login(c *gin.Context) {
	var request models.UserLoginRequest
	var err error

	if err = c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user, _ := services.Instance.GetUserByPhoneNumber(c, request.PhoneNumber)
	if user.ID == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "invalid credentials",
		})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(request.Password)) != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "wrong password",
		})
		return
	}

	tokenString, err := middlewares.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

// Validate validate
func Validate(c *gin.Context) {
	claims := c.MustGet("claims").(*middlewares.JWTClaims)

	c.JSON(http.StatusOK, gin.H{
		"data": claims,
	})
}
