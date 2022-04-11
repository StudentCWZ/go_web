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
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo ---> 数据表
type UserInfo struct {
	Id     uint
	Name   string
	Gender string
	Hobby  string
}

// gorm 操作
func main() {
	// 连接 Mysql 数据库
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:13306)/sql_demo?"+
		"charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("Connect mysql database failed, err: %v\n", err)
		panic(err)
	}
	fmt.Println("Connect mysql database success!")
	defer func(db *gorm.DB) {
		err = db.Close()
		if err != nil {
			fmt.Printf("Close mysql database failed, err: %v\n", err)
			return
		}
	}(db)
	// 创建表 --> 自动迁移(把结构体和数据表进行对应)
	db.AutoMigrate(&UserInfo{})
	// 创建数据行
	//u1 := UserInfo{1, "七米", "男", "蛙泳"}
	//// 数据插入
	//db.Create(&u1)
	//fmt.Println("Insert data success!")
	// 数据查询
	var user UserInfo
	db.First(&user) // 查询表中的第一条数据保存到 user 中
	fmt.Printf("user: %#v\n", user)
	// 更新
	db.Model(&user).Update("hobby", "双色球")
	fmt.Printf("user: %#v\n", user)
	// 删除
	db.Delete(&user)
	fmt.Println("Delete data success!")
}
