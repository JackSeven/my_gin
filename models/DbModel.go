package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 需要导入这个mysql库
	"my_gin/config"
)


// 数据库对象指针
var Mydb *sql.DB

func init()  {
	// 加载默认数据库配置
	//config.LoadFile("db_mysql.pro.json","json")
	//config.LoadConfig()

	config.SetMysqlConfig()

	a:=config.DbConfig
	fmt.Println(a)

	// 连接数据库
	InitDatabase()
}

// 初始化数据库
func InitDatabase() bool {

	if Mydb != nil {
		return true
	}

	if db, err:=sql.Open("mysql", config.DbConfig.Mysql.DataSource); err !=nil {
		fmt.Println("init Database failed")
		return false
	}else{
		Mydb = db
		//defer Mydb.Close()
		fmt.Println("init Database success")
	}

	//dbClose()

	return true
}

// 公共插入方法
func Insert(sql string, args ...interface{}) (int64, error)  {

	result, err:= Mydb.Exec(sql, args...)
	if err!=nil {
		fmt.Println("insert exec err: ", sql)
		fmt.Println("insert exec err: ", err)
		return 0, err
	}

	//count, err:= result.RowRowsAffected()

	count, err := result.LastInsertId();

	if err != nil {
		fmt.Println("insert affects return false: ", err)
	}

	return count, err

}

// 延时关闭数据库链接
func dbClose()  {
	defer Mydb.Close()
}
