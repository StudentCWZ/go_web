/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: go_web
   @Date: 2022/4/11 22:51
*/

package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// User 定义模型
type User struct {
	gorm.Model   // 内嵌 gorm.model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"` // 零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小 255
	MemberHumber *string `gorm:"unique;not null"` // 设置会员号 member number 唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给 address 字段创建名为 addr 的索引
	IgnoreMe     string  `gorm:"-"`               // 忽略本字段
}

// Animal 使用 `AnimalID` 作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

// TableName 自定义 table 名 (唯一指定表名)
func (Animal) TableName() string {
	return "animal_users"
}

// TableName 自定义 table 名 (唯一指定表名)
//func (u User) TableName() string {
//	// 条件判断
//	if u.Role == "admin" {
//		return "admin_users"
//	} else {
//		return "users"
//	}
//}

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
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "SMS_" + defaultTableName
	}
	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)
	// 创建表
	db.AutoMigrate(&User{})
	fmt.Println("Create table success!")
	// 创建表
	db.AutoMigrate(&Animal{})
	fmt.Println("Create table success!")
	// 使用 User 结构体创建名为 user_info 的表
	//db.Table("user_info").CreateTable(&User{})
}
