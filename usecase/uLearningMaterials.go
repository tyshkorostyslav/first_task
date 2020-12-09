package usecase

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	repository "github.com/tyshkorostyslav/first_task/repository"
)

func AvailableLearmingMaterials(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't get a database"})
		return
	}
	learningMaterials, err := repository.ReadLearningMaterial(db, true)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found!"})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": learningMaterials})
}

func AvailableBooks(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't get a database"})
		return
	}
	books, err := repository.ReadBook(db, true)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found!"})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": books})
}

func AvailablePages(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't get a database"})
		return
	}
	pages, err := repository.ReadPage(db, true)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found!"})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": pages})
}

func Commitment(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't get a database"})
		return
	}
	tx := db.Begin()
	userID, err := strconv.Atoi(c.PostForm("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	learningMaterial := c.PostForm("learningMaterial")
	book := c.PostForm("book")
	page := c.PostForm("page")
	err = repository.UpdateLearningMaterial(tx, learningMaterial, userID)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = repository.UpdateBook(tx, book, userID)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = repository.UpdatePage(tx, page, userID)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Comittment was successful!"})
}
