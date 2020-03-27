package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 需要导入这个mysql库
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)



/*********************************************************************
// table struct
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` char(8) NOT NULL COMMENT '用户ID',
  `nickname` varchar(255) NOT NULL COMMENT '昵称',
  PRIMARY KEY (`id`),
  UNIQUE KEY `u_idx_1` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=133636 DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
**********************************************************************/

type UserInfo struct {
	Id int `json:"id"`
	Userid int `json:"user_id"`
	Nickname string `json:"nickname"`
}

func main() {
	router := gin.Default()

	router.GET("/curd", func(c *gin.Context) {
		action := c.Query("action")

		if action == "r" {
			fmt.Println(action)

			users, err := getAll()
			if err != nil {
				log.Fatal(err)
			}
			c.JSON(http.StatusOK, gin.H{"state":1,"data":users})

		}else if action == "i" {
			fmt.Println(action)

			id, err := userAdd(UserInfo{
				Userid:   2,
				Nickname: "用户2",
			})
			if err != nil {
				log.Fatal(err.Error())
			}

			c.JSON(http.StatusOK, gin.H{"state":1, "data":id})

		}

	})

	router.Run(":6001")
}




// 添加用户
func userAdd(user UserInfo) (Id int, err error)  {

	db, err := sql.Open("mysql", "root:anling123@tcp(127.0.0.01)/test?charset=utf8")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	// 插入方法 1
	//sql := "insert into user(user_id, nickname) values (?,?)"
	//rs, err := db.Exec(sql, user.Userid, user.Nickname)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//if id,_ :=rs.LastInsertId(); id >0 {
	//	fmt.Println("ok")
	//}
	//return


	// 插入方法 2
	stmt, err:= db.Prepare("insert into user(user_id, nickname) values (?, ?)")

	fmt.Println(user.Userid)
	fmt.Println(user.Nickname)
	// 执行sql
	rs, err := stmt.Exec(user.Userid, user.Nickname)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// 返回id
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}

	Id = int(id)

	defer stmt.Close()
	return

}


// 查询数据
func getAll() (users []UserInfo, err error)  {

	db, err := sql.Open("mysql", "root:anling123@tcp(127.0.0.1:3306)/test?charset=utf8")

	if err!=nil {
		log.Fatal(err.Error())
	}

	defer db.Close()


	rows, err := db.Query("select id, user_id, nickname from user where id <=128953")
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		var user UserInfo
		rows.Scan(&user.Id, &user.Userid, &user.Nickname)
		users = append(users, user)
	}

	defer rows.Close()
	return
}

