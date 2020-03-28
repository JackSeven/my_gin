package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // 需要导入这个mysql库
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
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

			// go 默认的这个 rand.Intn 是个假的随机，每次产生的随机数，都是可以复现
			//user_id := rand.Intn(1000)

			// 需要加入随机因素
			rand.Seed(time.Now().UnixNano())
			user_id := rand.Intn(1000)

			fmt.Println("rand user id :", user_id)

			// 添加 1
			id, err := userAdd1(UserInfo{
				Userid:  user_id ,
				Nickname: "用户" + fmt.Sprintf("%d",user_id),
			})
			if err != nil {
				log.Fatal(err.Error())
			}

			////	添加2
			//id, err := userAdd2(UserInfo{
			//	Userid:   rand.Intn(1000),
			//	Nickname: "用户方法2",
			//})
			//if err != nil {
			//	log.Fatal(err.Error())
			//}

			c.JSON(http.StatusOK, gin.H{"state":1, "data":id})

		}else if action == "u" {
			res, err := updateUser(UserInfo{
				Id:       133687,
				Userid:   66,
				Nickname: "修改的用户名666",
			})

			if err != nil {
				log.Fatal(err.Error())
			}

			c.JSON(http.StatusOK, gin.H{"state":1, "data": res})
		}else if action == "d" {
			id := c.Query("id")
			if id == "" {
				c.JSON(http.StatusOK, gin.H{"state":0, "data":id})
			}
			del_id, _ := strconv.Atoi(id)

			res,_ := userDel( int64(del_id) )

			c.JSON(http.StatusOK, gin.H{"state":1, "data":res})
		}

	})

	router.Run(":6001")
}

// 删除
func userDel(id int64) (res int64, err error)  {

	db, _ := sql.Open("mysql", "root:anling123@tcp(127.0.0.1)/test?charset=utf8")

	defer db.Close()

	rs, err := db.Exec("delete from user where id = ?", id)
	if err != nil {
		log.Fatal(err.Error())
	}

	res, err = rs.RowsAffected()
	if err != nil {
		log.Fatal(err.Error())
	}

	return
}


// 更新
func updateUser(user UserInfo) (res int64, err error)  {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/test?charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	stmt,err := db.Prepare("update user set nickname = ?, user_id=? where id = ?")
	if err != nil {
		log.Fatal(err.Error())
	}

	rs, err := stmt.Exec(user.Nickname, user.Userid, user.Id)
	if err != nil {
		log.Fatal(err.Error())
	}

	res, err = rs.RowsAffected()
	if err != nil {
		log.Fatal(err.Error())
	}

	defer stmt.Close()
	return
}

// 添加用户方法1
func userAdd1(user UserInfo) (Id int64, err error)  {

	db, err := sql.Open("mysql", "root:anling123@tcp(127.0.0.01)/test?charset=utf8")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	stmt, err:= db.Prepare("insert into user(user_id, nickname) values (?, ?)")

	fmt.Println("user.Userid is:")
	fmt.Println(user.Userid)
	fmt.Println(user.Nickname)

	// 执行sql
	rs, err := stmt.Exec(user.Userid, user.Nickname)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// 返回id
	Id, err = rs.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}

	defer stmt.Close()
	return
}



// 添加用户方法2
func userAdd2(user UserInfo) (id int64, err error)  {

	db, err := sql.Open("mysql", "root:anling123@tcp(127.0.0.01)/test?charset=utf8")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	fmt.Println(user.Userid)
	fmt.Println(user.Nickname)

	//插入方法
	sql := "insert into user(user_id, nickname) values (?, ?)"
	rs, err := db.Exec(sql, user.Userid, user.Nickname)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(fmt.Sprintf("type %T , value %d", id, id))

	// 返回id

	id, err = rs.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}

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

