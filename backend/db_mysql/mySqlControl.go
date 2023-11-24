package db_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

var dsn = "root:ll3152067813@tcp(127.0.0.1:3306)/HiS_db"

func InitDB() {
	//数据库
	//用户名:密码@tcp(ip:3306)/数据库的名字
	//dsn := "root:ll3152067813@tcp(127.0.0.1:3306)/HiS_db"
	//连接数据集
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("open %s faild,err:%v\n", dsn, err)
		return
	}
	fmt.Println("connect success")
	db.SetMaxIdleConns(10)
	return
}

/*


	User表操作


*/

func SqlInsert_User(usr, passwd, role string) error {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return err
	}
	defer db.Close()
	InsertUser := "insert into user(Username,Password,Role) VALUES(?,?,?)" //写user表
	_, err = db.Exec(InsertUser, usr, passwd, role)
	if err != nil {
		return err
	}
	var userID int
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&userID)
	InsertUserRole := "INSERT INTO UserRole (UserID, Role) VALUES (?, ?)"
	_, err = db.Exec(InsertUserRole, userID, role) //关联userrole表
	if err != nil {
		return err
	}
	fmt.Printf("%v create successful\n", usr)
	return nil
}

func sqldelete_User() {

}

func SqlUpdate_User() {

}

func SqlQuery_User() {

}

/*
仅测试 不使用
_func main() {
	err := Insert("全志愿", "000000", "Doctor")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}
*/
