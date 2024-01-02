package model

import (
	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	DID           string `gorm:"primaryKey"`
	DGender       string // 例如：男、女等
	DPassword     string
	DName         string
	DPost         string
	DSection      string // 医生科室
	DPhone        int    // 电话号码
	DEmail        string
	DCard         int    // 身份证号
	DPrice        int    // 价格
	DIntroduction string // 医生简介，例如：很厉害的医生
}

type User struct {
	gorm.Model
	Account        string `gorm:"unique"`
	Username       string `gorm:"unique"`
	Password       string `gorm:"unique"`
	Role           string `gorm:"unique"`
	Gender         string `gorm:"unique"`
	Email          string `gorm:"unique"`
	Phone_number   string `gorm:"varchar(11);not null;unique"`
	Id_card_number string `gorm:"unique"`
	Date_of_birth  string `gorm:"unique"`
}

type Meta struct {
	Title       string `json:"title"`
	RequireAuth bool   `json:"requireAuth"`
}
