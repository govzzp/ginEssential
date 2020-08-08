package model

import (
	"github.com/jinzhu/gorm"
)

type User_info struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Telephone string `gorm:"size:255;not null;unique"`
}
