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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db_user := os.Getenv("DB_USER")
	pword := os.Getenv("DB_PWORD")
	db_addr := os.Getenv("DB_ADDR")
	db_name := os.Getenv("DB_NAME")
	db := repository.InitDb(db_user, pword, db_addr, db_name)
	defer db.Close()

	router = gin.Default()
	router.Use(repository.ApiMiddleware(db))

	delivery.Endpoints(router)

	port := os.Getenv("PORT")

	// Start serving the application
	router.Run(":" + port)

}
