package main

import (
	"fmt"
	"net/http"

	"feedbag/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"feedbag/Godeps/_workspace/src/github.com/markbates/goth/gothic"
)

func setupRoutes(r *gin.Engine) {
	//Oauth Authenticaton and Callbacks
	r.GET("/auth/github/callback", providerCallback)
	r.GET("/auth/github", providerAuth)

	//Api endpoints
	a := r.Group("api")
	a.GET("/activity", getActivity)
}

func getActivity(c *gin.Context) {
	c.JSON(200, gin.H{"activity": "here"})
}

func providerCallback(c *gin.Context) {
	// print our state string to the console
	fmt.Println(c.Request.URL.Query().Get("state"))

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		fmt.Fprintln(c.Writer, err)
		return

	}

	c.JSON(200, gin.H{"user": user})

}

func providerAuth(c *gin.Context) {
	gothic.GetProviderName = getProviderName
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func getProviderName(req *http.Request) (string, error) {
	return "github", nil
}
