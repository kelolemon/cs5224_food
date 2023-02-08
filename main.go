package main

import (
	"food/client"
	"food/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	err := client.InitDB()
	if err != nil {
		log.Fatalf("error creating db: %s", err)
		return
	}

	r := gin.Default()
	router.InitRouters(r)

	err = r.Run()
	if err != nil {
		log.Fatalf("init gin error, error=%v", err)
	}
}
