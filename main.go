package main

import (
	"food/client"
	"food/router"
	"github.com/go-redis/redis"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	err := client.InitDB()
	if err != nil {
		log.Fatalf("error creating db: %s", err)
		return
	}

	err = client.InitRedis()
	if err != nil {
		log.Fatalf("error creating redis: %v", err)
		return
	}
	defer func(RedisCli *redis.Client) {
		err := RedisCli.Close()
		if err != nil {
			log.Fatalf("error closing redis: %v", err)
		}
	}(client.RedisCli)

	r := gin.Default()
	router.InitRouters(r)

	err = r.Run()
	if err != nil {
		log.Fatalf("init gin error, error=%v", err)
	}
}
