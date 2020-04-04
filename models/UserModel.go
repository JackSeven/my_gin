package models

// 用户搜索
type UserSearch struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	Keyword    string `json:"keyword"`
	EnginId    int    `json:"engin_id"`
	EnginName  string `json:"engin_name"`
	CreateTime int    `json:"create_time"`
}


//添加用户搜索记录
func UsAdd(us UserSearch) (id int64, err error)  {

	id, err = Insert("insert into qk_user_search (user_id, keyword, engin_id, engin_name, create_time) values (?,?,?,?,?)",
		us.UserId, us.Keyword, us.EnginId, us.EnginName, us.CreateTime)

	if id ==0 {
		id = 0
		err = nil
	}

	return
}