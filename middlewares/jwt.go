package middlewares

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/MAAF72/go-boilerplate/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// TokenExpireDuration jwt token expire duration
var TokenExpireDuration time.Duration

// JWTSecret jwt secret key
var JWTSecret []byte

// JWTClaims jwt claims struct
type JWTClaims struct {
	jwt.StandardClaims
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
}

// InitJWT init jwt
func InitJWT() {
	intTokenExpireDuration, err := strconv.Atoi(os.Getenv("JWT_TOKEN_EXPIRE_DURATION"))
	if err != nil {
		log.Fatal("Error loading env JWT_TOKEN_EXPIRE_DURATION")
	}

	TokenExpireDuration = time.Duration(intTokenExpireDuration) * time.Second

	JWTSecret = []byte(os.Getenv("JWT_SECRET"))
}

// GenerateToken generate jwt token
func GenerateToken(user models.User) (string, error) {
	c := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "auth-app",
		},
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(JWTSecret)
}

// ParseToken parse jwt token
func ParseToken(tokenString string) (*JWTClaims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware jwt auth middleware
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "empty authorization header",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) < 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "invalid header",
			})
			c.Abort()
			return
		}

		prefix := parts[0]
		tokenString := parts[1]

		if prefix != "Bearer" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "invalid header",
			})
			c.Abort()
			return
		}

		claims, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "invalid token",
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
