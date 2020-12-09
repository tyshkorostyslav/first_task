package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	repository "github.com/tyshkorostyslav/first_task/repository/models"
)

func CreateUser(tx *gorm.DB, user repository.User) error {
	if err := tx.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func CreateLearningMaterial(tx *gorm.DB, learningMaterial repository.LearningMaterial) error {
	if err := tx.Create(&learningMaterial).Error; err != nil {
		return err
	}
	return nil
}

func CreateBook(tx *gorm.DB, book repository.LearningMaterial) error {
	if err := tx.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func CreatePage(tx *gorm.DB, page repository.LearningMaterial) error {
	if err := tx.Create(&page).Error; err != nil {
		return err
	}
	return nil
}

func ReadUser(db *gorm.DB) ([]repository.User, error) {
	var users []repository.User

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func ReadLearningMaterial(db *gorm.DB, available bool) ([]repository.LearningMaterial, error) {
	var learningMaterials []repository.LearningMaterial
	if available {
		if err := db.Find(&learningMaterials, repository.LearningMaterial{OwnerID: 0}).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.Find(&learningMaterials).Error; err != nil {
			return nil, err
		}
	}

	return learningMaterials, nil
}

func ReadBook(db *gorm.DB, available bool) ([]repository.LearningMaterial, error) {
	var books []repository.LearningMaterial

	if available {
		if err := db.Find(&books, repository.LearningMaterial{OwnerID: 0}).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.Find(&books).Error; err != nil {
			return nil, err
		}
	}
	return books, nil
}

func ReadPage(db *gorm.DB, available bool) ([]repository.LearningMaterial, error) {
	var pages []repository.LearningMaterial

	if available {
		if err := db.Find(&pages, repository.LearningMaterial{OwnerID: 0}).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.Find(&pages).Error; err != nil {
			return nil, err
		}
	}
	return pages, nil
}

func UpdateUser(tx *gorm.DB, userID int, userName string, hashedPword string) error {
	var user repository.User

	if err := tx.First(&user, userID).Error; err != nil {
		return err
	}

	if userName != "" {
		if err := tx.Model(&user).Update("name", userName).Error; err != nil {
			return err
		}
	}

	if hashedPword != "" {
		if err := tx.Model(&user).Update("pword", hashedPword).Error; err != nil {
			return err
		}
	}

	return nil
}

func UpdateLearningMaterial(tx *gorm.DB, learningMaterialName string, ownerId int) error {
	var learningMaterial repository.LearningMaterial

	if err := tx.Where(&repository.LearningMaterial{
		Name: learningMaterialName,
	}).First(&learningMaterial).Error; err != nil {
		return err
	}
	if learningMaterial.OwnerID == 0 {
		if err := tx.Model(&learningMaterial).Update("OwnerId", ownerId).Error; err != nil {
			return err
		}
	} else {
		err := errors.New("This learning material is already taken.")
		return err
	}

	return nil
}

func UpdateBook(tx *gorm.DB, bookName string, ownerId int) error {
	var book repository.LearningMaterial

	if err := tx.Where(&repository.LearningMaterial{
		Name: bookName,
	}).First(&book).Error; err != nil {
		return err
	}
	if book.OwnerID == 0 {
		if err := tx.Model(&book).Update("OwnerId", ownerId).Error; err != nil {
			return err
		}
	} else {
		err := errors.New("This book is already taken.")
		return err
	}
	return nil
}

func UpdatePage(tx *gorm.DB, pageName string, ownerId int) error {
	var page repository.LearningMaterial

	if err := tx.Where(&repository.LearningMaterial{
		Name: pageName,
	}).First(&page).Error; err != nil {
		return err
	}
	if page.OwnerID == 0 {
		if err := tx.Model(&page).Update("OwnerId", ownerId).Error; err != nil {
			return err
		}
	} else {
		err := errors.New("This page is already taken.")
		return err
	}
	return nil
}

func DeleteUser(tx *gorm.DB, userID int) error {
	var user repository.User

	if err := tx.First(&user, userID).Error; err != nil {
		return err
	}

	if err := tx.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteLearningMaterial(tx *gorm.DB, learningMaterialID int) error {
	var learningMaterial repository.LearningMaterial

	if err := tx.First(&learningMaterial, learningMaterialID).Error; err != nil {
		return err
	}

	if err := tx.Delete(&learningMaterial).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBook(tx *gorm.DB, bookID int) error {
	var book repository.LearningMaterial
	if err := tx.First(&book, bookID).Error; err != nil {
		return err
	}

	if err := tx.Delete(&book).Error; err != nil {
		return err
	}
	return nil
}

func DeletePage(tx *gorm.DB, pageID int) error {
	var page repository.LearningMaterial
	if err := tx.First(&page, pageID).Error; err != nil {
		return err
	}

	if err := tx.Delete(&page).Error; err != nil {
		return err
	}
	return nil
}
