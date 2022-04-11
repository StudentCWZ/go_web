/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: go_web
   @Date: 2022/4/11 22:54
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
	// 删除
	//var u = User{}
	//u.ID = 1
	////u.Name = "Mark" // 主键为空，删除表里所有记录
	//db.Debug().Delete(&u)
	//fmt.Println("Delete data success!")
	//db.Debug().Where("name = ?", "Jack").Delete(User{})
	//db.Debug().Delete(User{}, "age = ?", 18)
	var u1 []User
	// 可以查询到软删除的记录
	db.Debug().Unscoped().Where("name = ?", "Jack").Find(&u1)
	fmt.Println(u1)
	// 物理删除(永久删除)
	db.Debug().Unscoped().Where("name = ?", "Jack").Delete(User{})
}
