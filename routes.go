package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func setupRoutes(r *gin.Engine) {
	//Oauth Authenticaton and Callbacks
	r.GET("/auth/github/callback", providerCallback)
	r.GET("/auth/github", providerAuth)

	//Index Route
	r.GET("/", indexHandler)

	//Api endpoints
	a := r.Group("api")
	a.GET("/activity", getActivity)
	a.GET("/users", getUsers)
}

func indexHandler(c *gin.Context) {
	t, err := template.ParseFiles(filepath.Join("web", "index.tmpl"))
	if err != nil {
		panic(err)
	}
	t.Execute(c.Writer, templates)
}

func getActivity(c *gin.Context) {
	c.JSON(200, gin.H{"activity": "here"})
}

func getUsers(c *gin.Context) {
	u := UserList{}
	err := u.List()
	if err != nil {
		c.JSON(400, gin.H{"error": err})
	}

	c.JSON(200, u)
}

func providerCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		fmt.Fprintln(c.Writer, err)
		return
	}

	//Add user to the user table
	u := User{
		Name:        user.Name,
		Username:    user.RawData["login"].(string),
		AvatarUrl:   user.AvatarURL,
		AccessToken: user.AccessToken,
		ProfileUrl:  user.RawData["url"].(string),
		Email:       user.Email,
		Joined:      user.RawData["created_at"].(string),
		Raw:         user.RawData,
	}

	err = u.Create()
	if err != nil {
		c.JSON(200, gin.H{"error": err})
		return
	}

	c.JSON(200, u)
}

func providerAuth(c *gin.Context) {
	gothic.GetProviderName = getProviderName
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func getProviderName(req *http.Request) (string, error) {
	return "github", nil
}
