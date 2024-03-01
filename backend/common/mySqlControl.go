package common

import (
	"backend/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 开启连接池
func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "His_db"
	username := "root"
	password := "ll3152067813"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err: " + err.Error())
	}
	//自动创建数据表
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Doctor{})
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.Drug{})
	db.AutoMigrate(&model.Check{})
	db.AutoMigrate(&model.Registration{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

/*func UpdateDoctorNum(DB *gorm.DB) {
	var doctors []model.Doctor
	if err := DB.Find(&doctors).Error; err != nil {
		fmt.Println("Failed to fetch doctors:", err)
		return
	}

	// 更新每个医生的dNum
	for _, doctor := range doctors {
		doctor.DNum = 30
		if err := DB.Save(&doctor).Error; err != nil {
			fmt.Println("Failed to update doctor:", doctor.DID, err)
		}
	}
	//fmt.Println("Updated all doctors' dNum")
}*/
