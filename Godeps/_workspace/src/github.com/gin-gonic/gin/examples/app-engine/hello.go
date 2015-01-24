package hello

import (
	"feedbag/Godeps/_workspace/src/github.com/gin-gonic/gin" // This function's name is a must. App Engine uses it to drive the requests properly.
	"net/http"
)

func init() {
	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	// Define your handlers
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Handle all requests using net/http
	http.Handle("/", r)
}
