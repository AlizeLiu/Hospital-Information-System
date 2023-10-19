package db_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

var dsn = "root:ll3152067813@tcp(127.0.0.1:3306)/HiS_db"

func InitDB() {
	//数据库
	//用户名:密码@tcp(ip:端口)/数据库的名字
	//dsn := "root:ll3152067813@tcp(127.0.0.1:3306)/HiS_db"
	//连接数据集
	db, err := sql.Open("mysql", dsn) //open不会检验用户名和密码
	if err != nil {
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("open %s faild,err:%v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功")
	db.SetMaxIdleConns(10)
	return
}

func Insert(usr, passwd, role string) error {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return err
	}
	defer db.Close()
	InsertsqlStr := "insert into user(Username,Password,Role) VALUES(?,?,?)"
	_, err = db.Exec(InsertsqlStr, usr, passwd, role)

	if err != nil {
		return err
	}
	fmt.Printf("%v create successful", usr)
	return nil
}

/*
func main() {
	err := Insert("全志愿", "000000", "Doctor")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}
*/
