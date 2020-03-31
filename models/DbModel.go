package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // 需要导入这个mysql库
	"fmt"
	"my_gin/config"
)



// 数据库对象指针
var Mydb *sql.DB

func InitDatabase() bool {

	config.LoadConfig()
	fmt.Println("init Database....")

	if db, err:=sql.Open("mysql", config.DbConfig.Mysql.DataSource); err !=nil {
		fmt.Println("init Database failed")
		return false
	}else{
		Mydb = db
		defer Mydb.Close()
		fmt.Println("init Database success")
	}

	return true
}
