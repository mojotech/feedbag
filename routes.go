package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func setupRoutes(r *gin.Engine) {
	//Oauth Authenticaton and Callbacks
	r.GET("/auth/github/callback", providerCallback)
	r.GET("/auth/github", providerAuth)

	//Api endpoints
	a := r.Group("api")
	a.GET("/activity", getActivity)
	a.GET("/users", getUsers)
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
	// Run user auth using the gothic library
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	checkErr(err, "Failed to authenicate user")

	u := User{}

	err = u.GetByUsername(user.RawData["login"].(string))
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatalln("Failed to read from user table", err)
			return
		}
	}

	//Add user to the user table
	u.Name = user.Name
	u.Username = user.RawData["login"].(string)
	u.AvatarUrl = user.AvatarURL
	u.AccessToken = user.AccessToken
	u.ProfileUrl = user.RawData["url"].(string)
	u.Email = user.Email
	u.Joined = user.RawData["created_at"].(string)
	u.Raw = user.RawData

	if u.Id != 0 {
		u.UpdateTime()
		_, err = dbmap.Update(&u)
		checkErr(err, "Failed to update user row")
	} else {
		err = u.Create()
		checkErr(err, "Failed to create new user row")
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
