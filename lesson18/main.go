/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: go_web
   @Date: 2022/4/11 22:50
*/

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
	Gin 中间件：Gin 框架允许开发者在处理请求过程中，加入用户自己的钩子(Hook)函数，这个钩子函数就叫中间件。中间件适合处理一些公共的业务逻辑，
比如登录认证、权限校验、数据分页、记录日志、耗时统计等。
*/

// handlerFunc
func indexHandler(c *gin.Context) {
	fmt.Println("index ...")
	name, ok := c.Get("name") // 从上下文中取值(跨中间件存取值)
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// 定义一个中间件 m1: 统计处理函数的耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	// 记时
	start := time.Now()
	//go funcXX // 在 funcXX 中只能使用 c 的拷贝：只能使用只读的对象
	c.Next() // 调用后续的处理函数
	//c.Abort() // 阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost time: %v\n", cost)
	fmt.Println("m1 out ...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Set("name", "q1mi") // 在上下文 c 中设置值
	c.Next()              // 调用后续的处理函数
	//c.Abort() // 阻止调用后续的处理函数
	fmt.Println("m2 out ...")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	// 连接数据库
	// 或者其他准备工作
	return func(c *gin.Context) {
		if doCheck {
			// 存放具体的逻辑
			// 是否登录的判断
			// if 是登录用户
			// c.Next()
			// else
			// c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	// 创建一个默认路由引擎
	r := gin.Default() // 默认使用了 Logger 和 Recovery 中间件
	//r := gin.New()
	// 全局注册中间件函数 m1, m2, authMiddleware
	r.Use(m1, m2, authMiddleware(true))
	// GET 请求
	r.GET("/index", indexHandler)
	// GET 请求
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	// GET 请求
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})
	// 路由组注册中间件方法一：
	xxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
		})
	}
	// 路由组注册中间件方法二：
	xx2Group := r.Group("/xx2")
	xx2Group.Use(authMiddleware(true))
	{
		xx2Group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xx2Group"})
		})
	}
	// 启动服务
	err := r.Run(":9090")
	if err != nil {
		fmt.Printf("Http server start failed, err: %v\n", err)
		return
	}
}
