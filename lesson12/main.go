/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: go_web
   @Date: 2022/4/11 22:39
*/

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取 form 表单提交的参数
func main() {
	// 创建一个默认路由引擎
	r := gin.Default()
	// 加载模板文件
	r.LoadHTMLFiles("./login.html", "./index.html")
	// GET 请求
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// /login POST 请求
	r.POST("/login", func(c *gin.Context) {
		// 获取 form 表单提交的数据
		// 方法一
		//username := c.PostForm("username")	// 取到就返回值，取不到就返回空字符串
		//password := c.PostForm("password")
		// 方法二
		//username := c.DefaultPostForm("username", "somebody")
		//password := c.DefaultPostForm("password", "***")
		// 方法三
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "somebody"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "***"
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"UserName": username,
			"Password": password,
		})

	})
	// 启动服务
	err := r.Run(":9090")
	if err != nil {
		fmt.Printf("Http server start failed, err: %v\n", err)
		return
	}
}
