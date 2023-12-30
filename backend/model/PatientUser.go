package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account        string `gorm:"unique"`
	Username       string `gorm:"unique"`
	Password       string `gorm:"unique"`
	Role           string `gorm:"unique"`
	Email          string `gorm:"unique"`
	Phone_number   string `gorm:"varchar(11);not null;unique"`
	Id_card_number string `gorm:"unique"`
	Date_of_birth  string `gorm:"unique"`
}
