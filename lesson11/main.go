/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: go_web
   @Date: 2022/4/11 22:37
*/

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// querystring
func main() {
	// 创建一个默认路由
	r := gin.Default()
	// GET 请求
	// URL 的 ? 后面的是 querystring 参数
	// key=value 格式，多个 key-value 用 & 连接
	// eq: /web?query=小王子&age=18
	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器发送请求携带的 querystring 参数
		name := c.Query("query") // 通过 Query 获取请求中携带的 querystring 参数
		age := c.Query("age")
		//name := c.DefaultQuery("query", "somebody")	// 取不到就用指定的默认值
		//name, ok := c.GetQuery("query") // 取到返回(值, true)，取不到返回("", false)
		//if !ok {
		//	// 取不到
		//	name = "somebody"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	// 启动服务
	err := r.Run(":9090")
	if err != nil {
		fmt.Printf("Http server start failed, err: %v\n", err)
		return
	}
}
