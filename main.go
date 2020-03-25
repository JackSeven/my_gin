package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"  // 需要导入这个mysql库
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
			users, err := getAll()
			if err != nil {
				log.Fatal(err)
			}
			c.JSON(http.StatusOK, gin.H{"state":1,"data":users})
		}

	})

	router.Run(":6001")
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