package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:localhost123@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	DB = db
	return DB.DB().Ping()
}

func CloseDb() {
	err := DB.Close()
	if err != nil {
		fmt.Printf("Close mysql server failed, err: %v\n", err)
		return
	}
}
