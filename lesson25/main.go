/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: go_web
   @Date: 2022/4/11 22:57
*/

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() error {
	dsn := "root:localhost123@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	DB = db
	return DB.DB().Ping()
}

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := initMySQL()
	if err != nil {
		fmt.Printf("Conect mysql server failed, err: %v\n", err)
		return
	}
	defer func(DB *gorm.DB) { // 程序退出关闭数据库
		err := DB.Close()
		if err != nil {
			fmt.Printf("Close mysql server failed, err: %v\n", err)
			return
		}
	}(DB)
	// 模型绑定
	DB.AutoMigrate(&Todo{})
	// 创建一个默认的路由实例
	r := gin.Default()
	// 告诉 gin 框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// GET 请求
	// 告诉 gin 框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			// 前端页面填写待办事项，点击提交，会发请求到这里
			// 1. 从请求中把数据拿出来
			var todo Todo
			err := c.BindJSON(&todo)
			if err != nil {
				fmt.Printf("Get data failed, err: %v\n", err)
				return
			}
			// 2. 存入数据库
			err = DB.Create(&todo).Error
			if err != nil {
				fmt.Printf("Insert data failed, err: %v\n", err)
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				// 3. 返回响应
				c.JSON(http.StatusOK, todo)
			}
		})
		// 查看所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			// 查询 todo 这个表里的所有数据
			var todoList []Todo
			err := DB.Find(&todoList).Error
			if err != nil {
				fmt.Printf("Search data failed, err: %v\n", err)
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		// 查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "无效的 id",
				})
				return
			}
			var todo Todo
			err := DB.Where("id=?", id).First(&todo).Error
			if err != nil {
				fmt.Printf("Update data failed, err: %v\n", err)
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			} else {
				err = c.BindJSON(&todo)
				if err != nil {
					c.JSON(http.StatusOK, gin.H{
						"error": "无效的 variable",
					})
					return
				}
				err = DB.Save(&todo).Error
				if err != nil {
					c.JSON(http.StatusOK, gin.H{
						"error": err.Error(),
					})
					return
				} else {
					c.JSON(http.StatusOK, todo)
				}
			}
		})
		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "无效的 id",
				})
				return
			}
			err := DB.Where("id=?", id).Delete(Todo{}).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					id: "deleted",
				})
			}
		})
	}
	// 启动服务
	err = r.Run(":9090")
	if err != nil {
		fmt.Printf("Http server start failed, err: %v\n", err)
		return
	}
}
