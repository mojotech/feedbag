package main

import "feedbag/Godeps/_workspace/src/github.com/gin-gonic/gin"

func setupRoutes(r *gin.Engine) {
	a := r.Group("api")
	a.GET("/activity", getActivity)
}

func getActivity(c *gin.Context) {
	c.JSON(200, gin.H{"activity": "here"})
}
