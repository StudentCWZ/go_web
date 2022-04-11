/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: go_web
   @Date: 2022/4/11 22:43
*/

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取请求的 path(URI) 参数
// 注意：URL 匹配不要冲突
func main() {
	// 创建一个默认路由引擎
	r := gin.Default()
	// GET 请求
	r.GET("/user/:name/:age", func(c *gin.Context) {
		// 获取路径参数
		name := c.Param("name")
		age := c.Param("age") // string 类型
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	// GET 请求
	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})
	// 启动服务
	err := r.Run(":9090")
	if err != nil {
		fmt.Printf("Http server start failed, err: %v \n", err)
		return
	}
}
