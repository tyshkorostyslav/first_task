package usecase

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	repository "github.com/tyshkorostyslav/first_task/repository"
	models "github.com/tyshkorostyslav/first_task/repository/models"
)

func AvailableLearmingMaterials(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "couldn't get a database"})
		return
	}
	learningMaterials, err := repository.ReadLearningMaterial(db, true)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found!"})
	}
	learningMaterialsAll, err := repository.ReadLearningMaterial(db, false)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found!"})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   learningMaterials,
		"stats":  strconv.Itoa(len(learningMaterials)) + " of " + strconv.Itoa(len(learningMaterialsAll)) + " learning materials are avaliable",
	})
}

func AvailableBooks(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "couldn't get a database"})
		return
	}
	books, err := repository.ReadBook(db, true)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found!"})
	}
	booksAll, err := repository.ReadBook(db, false)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found!"})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   books,
		"stats":  strconv.Itoa(len(books)) + " of " + strconv.Itoa(len(booksAll)) + " books are avaliable",
	})
}

func AvailablePages(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "couldn't get a database"})
		return
	}
	pages, err := repository.ReadPage(db, true)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found!"})
	}
	pagesAll, err := repository.ReadPage(db, false)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found!"})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   pages,
		"stats":  strconv.Itoa(len(pages)) + " of " + strconv.Itoa(len(pagesAll)) + " pages are avaliable",
	})
}

func AddUser(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "couldn't get a database"})
		return
	}
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(400)
		return
	}
	tx := db.Begin()
	err := repository.CreateUser(tx, user)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "User was not created, because of server error!",
		})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User was created!"})
}

func AddLearningMaterial(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "couldn't get a database"})
		return
	}
	var learningMaterial models.LearningMaterial

	if err := c.BindJSON(&learningMaterial); err != nil {
		c.AbortWithStatus(400)
		return
	}
	tx := db.Begin()
	err := repository.CreateLearningMaterial(tx, learningMaterial)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Learning Material was not created, because of server error!",
		})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Learning Material was created!"})
}

func AddBook(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "couldn't get a database"})
		return
	}
	var book models.LearningMaterial

	if err := c.BindJSON(&book); err != nil {
		c.AbortWithStatus(400)
		return
	}
	tx := db.Begin()
	err := repository.CreateBook(tx, book)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Book was not created, because of server error!",
		})
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Book was created!"})
}

func AddPage(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "couldn't get a database"})
		return
	}
	var page models.LearningMaterial

	if err := c.BindJSON(&page); err != nil {
		c.AbortWithStatus(400)
		return
	}
	tx := db.Begin()
	err := repository.CreatePage(tx, page)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Page was not created, because of server error!",
		})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Page was created!"})
}

func Commitment(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "couldn't get a database"})
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
