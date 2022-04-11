/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: go_web
   @Date: 2022/4/11 22:53
*/

package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 定义模型
type User struct {
	gorm.Model // ID CreateAt UpdateAt DeleteAt
	Name       string
	Age        int64
	Active     bool
}

func main() {
	// 连接 MySQL 数据库
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
	// 创建表 --> 自动迁移
	db.AutoMigrate(&User{})
	// 数据插入
	//u1 := User{Name: "Mark", Age: 18, Active: true}
	//db.Create(&u1)
	//u2 := User{Name: "Jack", Age: 20, Active: true}
	//db.Create(&u2)
	// 查询
	var user User
	db.First(&user)
	fmt.Printf("user: %#v\n", user)
	// 更新
	user.Name = "Stone"
	user.Age = 99
	db.Debug().Save(&user) // 默认会修改所有字段
	fmt.Printf("user: %#v\n", user)
	db.Debug().Model(&user).Update("name", "stephen") // 更新单个字段
	fmt.Printf("user: %#v\n", user)
	m1 := map[string]interface{}{"name": "Bruce", "age": 28, "active": true}
	// 更新多个字段(map 列出的所有字段都会更新)
	db.Debug().Model(&user).Updates(m1)
	fmt.Printf("user: %#v\n", user)
	// 只更新 age 字段
	db.Debug().Model(&user).Select("age").Updates(m1)
	fmt.Printf("user: %#v\n", user)
	// 排除 active 字段，更新其他字段
	db.Debug().Model(&user).Omit("active").Updates(m1)
	fmt.Printf("user: %#v\n", user)
	// 只更新某个字段
	db.Debug().Model(&user).UpdateColumn("age", 30)
	fmt.Printf("user: %#v\n", user)
	// 查看数据库总共影响了多少行
	rowsNum := db.Debug().Model(&user).Updates(User{Name: "Tom", Age: 18}).RowsAffected
	fmt.Printf("rowsNum: %d\n", rowsNum)
	fmt.Printf("user: %#v\n", user)
	// 让 users 表中所有用户的年龄在原来基础上 +2
	db.Debug().Model(&User{}).Update("age", gorm.Expr("age  + ?", 2))
}
