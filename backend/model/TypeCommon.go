package model

import (
	"gorm.io/gorm"
	"time"
)

var OrderNum int

type Check struct {
	gorm.Model
	ChId    string `gorm:"primaryKey"` //药物ID
	ChName  string
	ChPrice int //价格 反正测试环境随便编
}

type Drug struct {
	gorm.Model
	DrId     string `gorm:"primaryKey"` //药物ID
	DrName   string //药物名称
	DrNumber int    //库存
	DrUnit   string //单位（盒之类的）
	DrPrice  int    //价格 反正测试环境随便编
}

type Admin struct {
	gorm.Model
	RName     string `gorm:"primaryKey"`
	RPassword string `gorm:"default:123456"`
	Role      string
	Rstate    string
}
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
	DState        int
	DNum          int    `gorm:"default:30"`
	DNonWorkDays  string `gorm:"type:varchar(255);"`
}

type User struct {
	gorm.Model
	Account        string `gorm:"unique"` //脑抽了命名错了。。但是写数据进去了 迁移有点麻烦 PId对应这个字段
	Username       string `gorm:"unique"`
	Password       string
	Role           string
	Gender         string
	Email          string `gorm:"unique"`
	Phone_number   string `gorm:"varchar(11);not null;unique"`
	Id_card_number string `gorm:"unique"`
	Date_of_birth  string `gorm:"unique"`
}

type Registration struct {
	gorm.Model
	Account         string
	Patient         User `gorm:"foreignKey:account"`
	DID             string
	Doctor          Doctor `gorm:"foreignKey:DID"`
	ORecord         string
	OTime           time.Time
	MedicineRecord  string
	DiagnosisRecord string
	ODrugs          string
	OChecks         string
	OTotalPrice     string
	OCheckBuyData   []string `gorm:"type:json"` //检查明细
	ODrugBuyData    []string `gorm:"type:json"` //药物明细
}

type Meta struct {
	Title       string `json:"title"`
	RequireAuth bool   `json:"requireAuth"`
}
