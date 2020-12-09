package repository

import "github.com/jinzhu/gorm"

type LearningMaterial struct {
	gorm.Model
	Type    string `db:"type" json:"type"`
	Name    string `db:"name" json:"name"`
	OwnerID int    `db:"ownerID" json:"OwnerID"`
	LMID    int    `db:"lmid" json:"lmid"`
}
