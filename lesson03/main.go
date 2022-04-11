/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: go_web
   @Date: 2022/4/11 22:13
*/

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, Golang!",
	})
}

func main() {
	// 创建一个路由引擎
	r := gin.Default()
	// 指定用户使用 GET 请求访问 /hello 时，执行 sayHello 这个函数
	r.GET("/hello", sayHello)

	// GET 请求
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})
	// POST 请求
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})
	// PUT 请求
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})
	// DELETE 请求
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})

	// 启动服务
	err := r.Run(":9090")
	if err != nil {
		fmt.Printf("Http server failed, err: %v\n", err)
		return
	}
}
