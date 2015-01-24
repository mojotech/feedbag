package main

import (
	"fmt"
	"os"

	"feedbag/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	port := os.Getenv("PORT")

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%s", port))
}
