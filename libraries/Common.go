package libraries

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"strconv"
	"time"
)

// 返回成功json
func ReturnSuccess(c *gin.Context, data interface{}, msg string)  {
	c.JSON(http.StatusOK, gin.H{
		"state" : true,
		"data" : data,
		"msg" : msg,
	})
	return
}

// 返回失败json
func ReturnError(c *gin.Context, data interface{}, msg string)  {
	c.JSON(http.StatusOK, gin.H{
		"state" : false,
		"data" : data,
		"msg" : msg,
	})
	return
}


// 默认页面
const INDEX = 1

// 默认分页大小
const PAGESIZE = 10

// 分页计算
func GetPageInfo(i string, p string, getAll bool) (limit int, offset int)  {

	index, _ := strconv.Atoi(i)
	pageSize, _ := strconv.Atoi(p)

	limit = PAGESIZE
	offset = 0

	if getAll {
		limit = 0
		return
	}
	if index <=1 {
		index = 1
	}
	if pageSize <= 0 {
		pageSize = PAGESIZE
	}

	limit = pageSize
	offset = (index - 1) * pageSize

	return
}

// 获取mac地址
func GetMacAddr() (macAddr string)  {

	netInterfaces, err := net.Interfaces()

	if err!=nil {
		fmt.Println("获取mac地址失败")
		return
	}

	for _, v:= range netInterfaces {
		macAddr = v.HardwareAddr.String()

		if macAddr != "" {
			return
		}
	}

	return
}


// 计算执行时间差
func ExecTime(tags string, start int64) (res int64)  {

	if tags == "start" {
		res := time.Now().UnixNano()
		return res

	}else{
		end := time.Now().UnixNano()
		diff := (end - start) / 1000
		fmt.Println("相差微妙数：", diff)
		res = diff
		return
	}

}