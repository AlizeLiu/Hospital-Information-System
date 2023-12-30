package main

import (
	"backend/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

func main() {
	db := common.InitDB()
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database,err: " + err.Error())
	}
	defer sqlDB.Close()
	r := gin.Default()
	r = collectRoute(r)
	panic(r.Run())

}
