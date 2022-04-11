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
	gorm.Model // ID creatAt UpdateAt DeleteAt
	Name       string
	Age        int64
}

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
	// 创建表 --> 自动迁移
	db.AutoMigrate(&User{})
	// 数据插入
	//u1 := User{Name: "Mark", Age: 18}
	//u2 := User{Name: "Jack", Age: 20}
	//db.Debug().Create(&u1)
	//db.Debug().Create(&u2)
	//fmt.Println("Insert data success!")
	// 数据查询
	// 一般查询
	var u1, u2, u3, u4 User
	// new 和 make 的区别：new 基本数据类型，则返回相应的指针；make 给 channel、map、slice 申请内存，返回申请的类型
	// u := new(User)
	var users []User
	// 根据主键查询第一条记录
	db.First(&u1)
	fmt.Printf("u1: %#v\n", u1)
	// 随机获取一条记录
	db.Take(&u2)
	fmt.Printf("u2: %#v\n", u2)
	// 根据主键查询最后一条数据
	db.Last(&u3)
	fmt.Printf("u3: %#v\n", u3)
	// 查询指定的某条记录(仅当主键为整型时可用)
	db.First(&u4, 2)
	fmt.Printf("u4: %#v\n", u4)
	// 查询所有记录
	db.Debug().Find(&users)
	fmt.Printf("users: %#v\n", users)
	// where 条件
	var u5 User
	// Get first matched record
	db.Where("name = ?", "Mark").First(&u5)
	fmt.Printf("u5: %#v\n", u5)
	// SELECT * FROM users WHERE name = 'Mark' limit 1;
	// Get all matched records
	//db.Where("name = ?", "Mark").Find(&users)
	//fmt.Printf("users: %#v\n", users)
	//// SELECT * FROM users WHERE name = 'Mark';
	//// <>
	//db.Where("name <> ?", "Jack").Find(&users)
	//// SELECT * FROM users WHERE name <> 'Jack';
	//// IN
	//db.Where("name IN (?)", []string{"Mark", "Jack"}).Find(&users)
	//// SELECT * FROM users WHERE name in ('Mark','Jack');
	//// LIKE
	//db.Where("name LIKE ?", "%ark%").Find(&users)
	//// SELECT * FROM users WHERE name LIKE '%jin%';
	//// AND
	//db.Where("name = ? AND age >= ?", "Mark", "22").Find(&users)
	//// SELECT * FROM users WHERE name = 'Mark' AND age >= 22;

	// FirstOrInit
	var u6 User
	//db.Attrs(User{Age: 99}).FirstOrInit(&u6, User{Name: "Stone"})	// 记录找到，才会将参数赋值给 struct
	db.Assign(User{Age: 99}).FirstOrInit(&u6, User{Name: "Mark"}) // 不管记录是否找到，才会将参数赋值给 struct
	fmt.Printf("u6: %#v\n", u6)
}
