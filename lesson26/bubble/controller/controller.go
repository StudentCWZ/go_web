package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web/lesson26/bubble/models"
	"net/http"
)

/*
	url --> controller --> logic --> models
	请求来了 --> 控制器 --> 业务逻辑 --> 模型层的增删改查
*/
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	// 前端页面填写待办事项，点击提交，会发请求到这里
	// 1. 从请求中把数据拿出来
	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		fmt.Printf("Get data failed, err: %v\n", err)
		return
	}
	// 2. 存入数据库
	err = models.CreateATodo(&todo)
	if err != nil {
		fmt.Printf("Insert data failed, err: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		// 3. 返回响应
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context) {
	// 查询 todo 这个表里的所有数据
	todoList, err := models.GetAllTodo()
	if err != nil {
		fmt.Printf("Search data failed, err: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的 id",
		})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		fmt.Printf("Update data failed, err: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的 variable",
		})
		return
	}
	err = models.UpdateATodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的 id",
		})
		return
	}
	err := models.DeleteATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			id: "deleted",
		})
	}
}
