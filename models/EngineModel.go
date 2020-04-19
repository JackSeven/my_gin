package models

import "fmt"

type EngineInfo struct {
	Id         int    `json:"id"`
	EngineName string `json:"engine_name"`
	EnginLogo  string `json:"engin_logo"`
	EnginNick  string `json:"engin_nick"`
	Sort       int    `json:"sort"`
}


func Lists() (rows []EngineInfo) {

	result, err := Mydb.Query("select * from qk_engine_info")
	if err != nil {
		fmt.Println("query engine info faild")
		return
	}

	rows = make([]EngineInfo, 0)

	for result.Next()  {
		var row EngineInfo
		result.Scan(
			&row.Id,
			&row.EngineName,
			&row.EnginLogo,
			&row.EnginNick,
			&row.Sort,
		)
		rows = append(rows, row)
	}

	return
}