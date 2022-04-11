/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: go_web
   @Date: 2022/4/11 22:52
*/

package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 定义模型
type User struct {
	ID int64
	// sql.NullString 实现了 Scanner/Valuer 接口
	Name sql.NullString `gorm:"default:'admin'"` // 字段默认值
	Age  int64
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
	// 创建表 --> 自动迁移(关联表)
	db.AutoMigrate(&User{})
	// 创建记录
	//u := User{Name: "q1mi", Age: 18} // 在代码层面创建一个 User 对象
	//u := User{Age: 28}            // Name 默认为空字符串(如果结构体中有设置默认值就为默认值)
	//u := User{Name: "", Age: 38}  // 如果结构体中有设置默认值，当 Name 传入空字符串，还是会被默认值替换
	//u := User{Age: 48}            // Name 为 *string 指针类型的测试用例
	//u := User{Name: new(string), Age: 58} // 使用指针方式将零值或者空字符串存入设有默认值的数据库
	//u := User{Age: 68}            // Name 为 sql.NullString 类型的测试用例(实现 Scanner/Valuer 接口)
	u := User{Name: sql.NullString{Valid: true}, Age: 78} // 实现 Scanner/Valuer 接口方式将空字符串存入设有默认值的数据库
	fmt.Println(db.NewRecord(&u))                         // 判断主键是否为空
	//db.Create(&u)                 // 数据插入
	db.Debug().Create(&u)         // 打印具体信息
	fmt.Println(db.NewRecord(&u)) // 判断主键是否为空
	fmt.Println("Insert data success!")
}
