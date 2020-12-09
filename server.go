package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tyshkorostyslav/first_task/delivery"
	repository "github.com/tyshkorostyslav/first_task/repository"
)

var router *gin.Engine

func main() {
	db := repository.InitDb()
	defer db.Close()

	router = gin.Default()
	router.Use(repository.ApiMiddleware(db))

	delivery.Endpoints(router)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	// Start serving the application
	router.Run(":" + port)

}
