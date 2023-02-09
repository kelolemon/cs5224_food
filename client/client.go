package client

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB
var RedisCli *redis.Client

func InitDB() (err error) {
	dsn := "root:cs5224@tcp(127.0.0.1:3308)/food?charset=utf8mb4&parseTime=True&loc=Local"
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

func InitRedis() (err error) {
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6400",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err = RedisCli.Ping().Err()
	return err
}
