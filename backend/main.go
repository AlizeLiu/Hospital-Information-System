package main

import (
	"backend/db_mysql"
	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

func main() {
	/*
		// 1.创建路由
		r := gin.Default()
		// 2.绑定路由规则，执行的函数
		// gin.Context，封装了request和response
		r.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello World!")
			c.String(http.StatusOK, "hello Go!")
		})
		// 3.监听端口，默认在8080
		// Run("")里面不指定端口号默认为8080
		r.Run(":8000")
	*/
	db_mysql.SqlInsert_User("杰克Jack", "000000", "Patient")
}
