package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MAAF72/go-boilerplate/adapters"
	"github.com/MAAF72/go-boilerplate/middlewares"
	"github.com/MAAF72/go-boilerplate/routers"
	"github.com/MAAF72/go-boilerplate/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(os.Getenv("GIN_MODE"))
	app := gin.Default()

	adapters := adapters.Init()
	services.Init(adapters)

	middlewares.Init()

	routers.RegisterRouters(app)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	app.Run(fmt.Sprintf("%s:%s", host, port))
}
