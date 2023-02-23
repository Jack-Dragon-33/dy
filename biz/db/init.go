package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

const (
	mysql_driver = "root:915+sxl..@tcp(127.0.0.1:3306)/dy?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	DB = Init()
)

func Init() *gorm.DB {
	DB, err := gorm.Open(mysql.Open(mysql_driver), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	fmt.Printf("DB: %v\n", DB)
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}
	return DB
}
