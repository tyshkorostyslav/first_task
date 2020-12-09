package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tyshkorostyslav/first_task/delivery"
	repository "github.com/tyshkorostyslav/first_task/repository/models"
)

var router *gin.Engine

func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&repository.User{}) {
		db.CreateTable(&repository.User{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&repository.User{})
	}

	if !db.HasTable(&repository.LearningMaterial{}) {
		db.CreateTable(&repository.LearningMaterial{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&repository.LearningMaterial{})
	}

	return db
}

// ApiMiddleware will add the db connection to the context
func ApiMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("databaseConn", db)
		c.Next()
	}
}

func main() {
	db := InitDb()
	defer db.Close()

	router = gin.Default()
	router.Use(ApiMiddleware(db))

	delivery.Endpoints(router)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	// Start serving the application
	router.Run(":" + port)

}
