/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: go_web
   @Date: 2022/4/11 22:35
*/

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// 静态文件：
// html 页面上用到的样式文件：css 文件 、js 文件、图片
func main() {
	// 创建路由引擎
	r := gin.Default()
	// 加载静态文件
	r.Static("/statics", "./statics")
	// gin 框架中给模板添加自定义函数
	// safe 函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// 模板解析
	r.LoadHTMLGlob("templates/**/*")
	// GET 请求
	r.GET("/posts/index", func(c *gin.Context) {
		// HTTP 请求
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ // 模板渲染
			"title": "/posts/index.tmpl",
		})
	})
	// GET 请求
	r.GET("/users/index", func(c *gin.Context) {
		// HTTP 请求
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{ // 模板渲染
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})
	// 返回从网上下载的模板
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})
	// 启动服务
	err := r.Run(":9090")
	if err != nil {
		fmt.Printf("Http server start failed, err: %v \n", err)
		return
	}
}
