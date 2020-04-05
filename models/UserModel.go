package models

import (
	"fmt"
)

// 用户搜索
type UserSearch struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	Keyword    string `json:"keyword"`
	EnginId    int    `json:"engin_id"`
	EnginName  string `json:"engin_name"`
	CreateTime int    `json:"create_time"`
}

// 用户信息
type UserInfo struct {
	Id         int       `json:"id"`
	UserName   string    `json:"user_name"`
	NickName   string    `json:"nick_name"`
	UserPwd    string    `json:"user_pwd"`
	Salt       string    `json:"salt"`
	UserImgUrl string    `json:"user_img_url"`
	Sex        string    `json:"sex"`
	Region     string    `json:"region"`
	Email      string    `json:"email"`
	Mobile     string    `json:"mobile"`
	RegisterIp string    `json:"register_ip"`
	CreateTime int    `json:"create_time"`
	UpdateTime int    `json:"update_time"`
}


// 添加用户搜索记录
func UsAdd(us UserSearch) (id int64, err error)  {

	id, err = Insert("insert into qk_user_search (user_id, keyword, engin_id, engin_name, create_time) values (?,?,?,?,?)",
		us.UserId, us.Keyword, us.EnginId, us.EnginName, us.CreateTime)

	if id ==0 {
		id = 0
		err = nil
	}

	return
}

// 用户搜索记录
func UsList(w map[string]int) (resRows []UserSearch) {

	InitDatabase()

	fmt.Println("uuuuuuurrrrrrrrr")
	UserRow(1)

	rows, err := Mydb.Query("select * from qk_user_search where 1=1  order by id desc limit ? offset ?", w["limit"], w["offset"])

	if err != nil {
		fmt.Println(err)
		return
	}

	resRows = make([]UserSearch, 0)
	for rows.Next()  {
		var us UserSearch
		rows.Scan(
			&us.Id,
			&us.UserId,
			&us.Keyword,
			&us.EnginId,
			&us.EnginName,
			&us.CreateTime,
		)

		resRows = append(resRows, us)
		defer rows.Close()
	}
	return

}

// 获取一条用户信息
func UserRow(id int) (userRes []UserInfo)  {

	InitDatabase()

	rows, err := Mydb.Query("select * from qk_user_info where id = ?", id)
	if err != nil {
		fmt.Println(err)
		return
	}

	userRes = make([]UserInfo, 0)

	for rows.Next()  {
		var ui UserInfo
		rows.Scan(
			&ui.Id,
			&ui.UserName,
			&ui.NickName,
			&ui.UserPwd,
			&ui.Salt,
			&ui.UserImgUrl,
			&ui.Sex,
			&ui.Region,
			&ui.Email,
			&ui.Mobile,
			&ui.RegisterIp,
			&ui.CreateTime,
			&ui.UpdateTime,
		)

		userRes = append(userRes, ui)
	}

	return
}