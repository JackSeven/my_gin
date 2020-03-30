package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type mysqlConfig struct{
	Mysql *mysqlNode `json:mysql`
}

type mysqlNode struct {
	DataSource      string `json:"data_source"`
	MaxIdleConns    int    `json:"max_idle_conns"`
	MaxOpenConns    int    `json:"max_open_conns"`
	ConnMaxLifeTime int64  `json:"conn_max_life_time"`
}


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

		fmt.Println("init mysqlconfig is ", MysqlConfig.Mysql.ConnMaxLifeTime)

		return true
	}


}