package feedbag

import (
	"database/sql"
	"net/http"

	"github.com/fogcreek/logging"
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
	"github.com/markbates/goth/gothic"
	"github.com/mojotech/feedbag/feedbag/tmpl"
)

func SetupRoutes(r *gin.Engine, s *socketio.Server) {
	//Oauth Authenticaton and Callbacks
	r.GET("/auth/github/callback", providerCallback)
	r.GET("/auth/github", providerAuth)

	//Template Route
	r.GET("/templates", templateHandler)

	//Socket.io Route
	r.GET("/socket.io/", func(c *gin.Context) {
		s.ServeHTTP(c.Writer, c.Request)
	})

	//Api endpoints
	a := r.Group("api")
	a.GET("/activity", getActivity)
	a.GET("/users", getUsers)
}

func templateHandler(c *gin.Context) {
	var err error
	Templates, err = tmpl.ParseDir(TemplatesDir)
	if err != nil {
		logging.ErrorWithTags([]string{"templates"}, "Failed to load templates.", err.Error())
	}

	c.JSON(200, Templates)
}

func getActivity(c *gin.Context) {
	a := ActivityList{}
	err := a.List()
	if err != nil {
		logging.WarnWithTags([]string{"api"}, "Activity endpoint route failed.", err.Error())
		c.JSON(400, gin.H{"error": err})
	}

	c.JSON(200, a)
}

func getUsers(c *gin.Context) {
	u := UserList{}
	err := u.List()
	if err != nil {
		logging.WarnWithTags([]string{"api"}, "Users endpoint route failed.", err.Error())
		c.JSON(400, gin.H{"error": err})
	}

	c.JSON(200, u)
}

func providerCallback(c *gin.Context) {
	// Run user auth using the gothic library
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		logging.ErrorWithTags([]string{"github", "api"}, "Failed to create user from callback", err.Error())
	}

	u := User{}

	err = u.GetByUsername(user.RawData["login"].(string))
	if err != nil {
		if err != sql.ErrNoRows {
			logging.ErrorWithTags([]string{"api"}, "Failed to read from user table", err.Error())
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
		if err != nil {
			logging.ErrorWithTags([]string{"db"}, "Failed to update user row", err.Error())
		}
	} else {
		err = u.Create()
		if err != nil {
			logging.ErrorWithTags([]string{"db"}, "Failed to create new user row", err.Error())
		}

		//Add the user's go routine
		StartUserRoutine(u, activityChan)
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
