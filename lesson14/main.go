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

type UserInfo struct {
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	// 创建一个默认路由引擎
	r := gin.Default()
	// 加载模板文件
	r.LoadHTMLFiles("./index.html")
	// GET 请求
	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//u := UserInfo{
		//	UserName: username,
		//	Password: password,
		//}
		var u UserInfo          // 声明一个 UserInfo 类型变量 u
		err := c.ShouldBind(&u) // 注意：这里要传入指针变量
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	// GET 请求
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// POST 请求
	r.POST("/form", func(c *gin.Context) {
		var u UserInfo          // 声明一个 UserInfo 类型变量 u
		err := c.ShouldBind(&u) // 注意：这里要传入指针变量
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	// POST 请求
	r.POST("/json", func(c *gin.Context) {
		var u UserInfo          // 声明一个 UserInfo 类型变量 u
		err := c.ShouldBind(&u) // 注意：这里要传入指针变量
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	// 启动服务
	err := r.Run(":9090")
	if err != nil {
		fmt.Printf("Http server start failed, err: %v\n", err)
		return
	}
}
