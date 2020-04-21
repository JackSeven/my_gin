package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
)

// sql配置
var DbConfig *mysqlConfig

type mysqlConfig struct{
	Mysql *mysqlNode `json:mysql`
}

type mysqlNode struct {
	DataSource      string `json:"data_source"`
	MaxIdleConns    int    `json:"max_idle_conns"`
	MaxOpenConns    int    `json:"max_open_conns"`
	ConnMaxLifeTime int64  `json:"conn_max_life_time"`
}


// 默认配置
var ConfigPath = "./config/"

var sqlConfigFile = "db_mysql.json"

var sqlConfigType = "json"


// 读文件方式，加载sql配置
func LoadConfig() bool {

	// 读配置
	if buf, err := ioutil.ReadFile("config/db_mysql.json"); err != nil {
		fmt.Println("load config faild", err)
		return false

	}else{
		//fmt.Println("buf", buf)
		MysqlConfig := &mysqlConfig{}

		// 解析配置
		if err := json.Unmarshal(buf, MysqlConfig); err != nil {
			fmt.Println("read config faild", err)
			return false
		}

		DbConfig = MysqlConfig

		// **************** viper 配置文件 test ************************/
		//viper.SetConfigName("viper_test.json")
		//viper.AddConfigPath("./config/")
		//viper.SetConfigType("json")
		//if err := viper.ReadInConfig();err != nil {
		//	fmt.Printf("err:%s\n",err)
		//}
		//
		//fmt.Println(viper.GetStringMap("mysql"))
		//fmt.Println(viper.GetString("mysql.data_source"))
		//
		//viper.SetConfigName("viper_test.yaml")
		//viper.AddConfigPath("./config/")
		//viper.SetConfigType("yaml")
		//
		//viper.Set("test","test_valueok")
		//viper.WriteConfig()



		// **************** viper 配置文件 test ************************/


		return true
	}

}

// viper方法 加载sql配置
func SetMysqlConfig()(res bool)  {

	err := LoadFile(sqlConfigFile, sqlConfigType)
	if err !=nil {
		fmt.Println(err)
		return false
	}

	item := GetConfig("mysql")

	DbConfig = &mysqlConfig{
		Mysql: &mysqlNode{
			DataSource : item["data_source"].(string),
		},
	}

	//
	//viper.Reset()
	//viper.SetConfigName("db.mysql.ini")
	//viper.AddConfigPath("./config/")
	//viper.SetConfigType("ini")
	//
	//fmt.Println(item["data_source"])
	//fmt.Println(item["max_idle_conns"])
	//fmt.Println(item["max_open_conns"])
	//fmt.Println(item["conn_max_life_time"])
	//
	//
	//viper.Set("mysql.data_source", fmt.Sprintf("%s", item["data_source"]) )
	//viper.Set("mysql.max_idle_conns", fmt.Sprintf("%s", item["max_idle_conns"]) )
	//viper.Set("mysql.max_open_conns", fmt.Sprintf("%s", item["max_open_conns"]) )
	//viper.Set("mysql.conn_max_life_time", fmt.Sprintf("%s", item["conn_max_life_time"]) )
	//viper.WriteConfig()
	//
	//fmt.Println("写入完毕")

	return true
}


// 加载配置文件
func LoadFile(file_path string, file_type string) (err error)  {

	viper.SetConfigName(file_path)
	viper.AddConfigPath(ConfigPath)

	if len(file_type) > 0 {
		viper.SetConfigType(file_type)
	}else{
		viper.SetConfigType("ini")
	}

	if err := viper.ReadInConfig();err != nil {
		fmt.Printf("err:%s\n",err)
	}

	return
}

// 获取配置
func GetConfig(key string) (data map[string]interface{}) {
	return viper.GetStringMap(key)
}








