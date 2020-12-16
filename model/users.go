package model

import "gorm.io/gorm"

// User is struct of user
type User struct {
	gorm.Model
	Name      string `gorm:"varchar(20);not null"`
	Password  string `gorm:"size:255;not null"`
	Email     string `gorm:"varchar(100);not null;unique"`
	Telephone string `gorm:"varchar(11);not null;unique"`
}
