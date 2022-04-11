/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: go_web
   @Date: 2022/4/11 22:36
*/

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建默认路由
	r := gin.Default()
	// GET 请求
	r.GET("/json", func(c *gin.Context) {
		// 方法一：使用 map
		//data := map[string]interface{}{
		//	"name":    "小王子",
		//	"message": "hello, world!",
		//	"age":     18,
		//}
		// 使用 gin.H
		data := gin.H{
			"name":    "小王子",
			"message": "hello, world!",
			"age":     18,
		}
		c.JSON(http.StatusOK, data)
	})
	// 方法二：结构体，灵活使用 tag 来对结构体做定制化操作
	type msg struct {
		Name    string `json:"name""`
		Message string `json:"message"`
		Age     int    `json:"age"`
	}
	// GET 请求
	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			"小王子",
			"Hello, Golang!",
			18,
		}
		c.JSON(http.StatusOK, data) // json 序列化
	})
	// 启动服务
	err := r.Run(":9090")
	// 错误判断
	if err != nil {
		fmt.Printf("Http server start failed, err: %v\n", err)
		return
	}
}
