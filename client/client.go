package client

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func InitDB() (err error) {
	dsn := "root:cs5224@tcp(127.0.0.1:3308)/food?charset=utf8mb4&parseTime=True&loc=Local"
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
