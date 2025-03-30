package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        // m'ajoute les fiels ID, createdAt ...
	Username   string `gorm:"uniqueIndex;not null"` //uniqueIndex pour ne pas avoir de doublons
	Email      string `gorm:"uniqueIndex;not null"`
	Password   string `gorm:"not null"`
}
