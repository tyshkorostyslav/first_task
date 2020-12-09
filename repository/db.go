package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	repository "github.com/tyshkorostyslav/first_task/repository/models"
)

func InitDb(db_user string, pword string, db_addr string, db_name string) *gorm.DB {
	dsn := db_user + ":" + pword + "@" + db_addr + "/" + db_name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
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
