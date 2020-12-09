package repository

import (
	"github.com/jinzhu/gorm"
	repository "github.com/tyshkorostyslav/first_task/repository/models"
)

func createUser(tx *gorm.DB, user repository.User) error {
	if err := tx.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func createLearningMaterial(tx *gorm.DB, learningMaterial repository.LearningMaterial) error {
	if err := tx.Create(&learningMaterial).Error; err != nil {
		return err
	}
	return nil
}

func createBook(tx *gorm.DB, book repository.LearningMaterial) error {
	if err := tx.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func createPage(tx *gorm.DB, page repository.LearningMaterial) error {
	if err := tx.Create(&page).Error; err != nil {
		return err
	}
	return nil
}

func readUser(tx *gorm.DB) error {
	var users []repository.User

	if err := tx.Find(&users).Error; err != nil {
		return err
	}
	return nil
}

func readLearningMaterial(tx *gorm.DB, available bool) error {
	var learningMaterials []repository.LearningMaterial
	if available {
		if err := tx.Find(&learningMaterials, repository.LearningMaterial{OwnerID: 0}).Error; err != nil {
			return err
		}
	} else {
		if err := tx.Find(&learningMaterials).Error; err != nil {
			return err
		}
	}

	return nil
}

func readBook(tx *gorm.DB, available bool) error {
	var books []repository.LearningMaterial

	if available {
		if err := tx.Find(&books, repository.LearningMaterial{OwnerID: 0}).Error; err != nil {
			return err
		}
	} else {
		if err := tx.Find(&books).Error; err != nil {
			return err
		}
	}
	return nil
}

func readPage(tx *gorm.DB, available bool) error {
	var pages []repository.LearningMaterial

	if available {
		if err := tx.Find(&pages, repository.LearningMaterial{OwnerID: 0}).Error; err != nil {
			return err
		}
	} else {
		if err := tx.Find(&pages).Error; err != nil {
			return err
		}
	}
	return nil
}

func updateUser(tx *gorm.DB, userID int, userName string, hashedPword string) error {
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

func updateLearningMaterial(tx *gorm.DB, learningMaterialID int, ownerId int) error {
	var learningMaterial repository.LearningMaterial

	if err := tx.First(&learningMaterial, learningMaterialID).Error; err != nil {
		return err
	}

	if err := tx.Model(&learningMaterial).Update("OwnerId", ownerId).Error; err != nil {
		return err
	}
	return nil
}

func updateBook(tx *gorm.DB, bookID int, ownerId int) error {
	var book repository.LearningMaterial

	if err := tx.First(&book, bookID).Error; err != nil {
		return err
	}

	if err := tx.Model(&book).Update("OwnerId", ownerId).Error; err != nil {
		return err
	}
	return nil
}

func updatePage(tx *gorm.DB, pageID int, ownerId int) error {
	var page repository.LearningMaterial

	if err := tx.First(&page, pageID).Error; err != nil {
		return err
	}

	if err := tx.Model(&page).Update("OwnerId", ownerId).Error; err != nil {
		return err
	}
	return nil
}

func deleteUser(tx *gorm.DB, userID int) error {
	var user repository.User

	if err := tx.First(&user, userID).Error; err != nil {
		return err
	}

	if err := tx.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func deleteLearningMaterial(tx *gorm.DB, learningMaterialID int) error {
	var learningMaterial repository.LearningMaterial

	if err := tx.First(&learningMaterial, learningMaterialID).Error; err != nil {
		return err
	}

	if err := tx.Delete(&learningMaterial).Error; err != nil {
		return err
	}
	return nil
}

func deleteBook(tx *gorm.DB, bookID int) error {
	var book repository.LearningMaterial
	if err := tx.First(&book, bookID).Error; err != nil {
		return err
	}

	if err := tx.Delete(&book).Error; err != nil {
		return err
	}
	return nil
}

func deletePage(tx *gorm.DB, pageID int) error {
	var page repository.LearningMaterial
	if err := tx.First(&page, pageID).Error; err != nil {
		return err
	}

	if err := tx.Delete(&page).Error; err != nil {
		return err
	}
	return nil
}
