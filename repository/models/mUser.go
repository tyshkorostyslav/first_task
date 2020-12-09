package repository

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name        string `db:"name" json:"name"`
	HashedPword string `db:"pword" json:"pword"`
}
