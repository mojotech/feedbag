package main

import (
	"flag"
	"fmt"
	"os"

	"feedbag/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"feedbag/Godeps/_workspace/src/github.com/markbates/goth"
	"feedbag/Godeps/_workspace/src/github.com/markbates/goth/providers/github"
)

var (
	configPort   = flag.String("port", "3000", "Port to run the server on")
	templatesDir = flag.String("templates", "./templates", "Path to your templates directory")
)

func main() {
	//Setup gin
	r := gin.Default()
	setupRoutes(r)

	// Setup Goth Authentication
	goth.UseProviders(
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:3000/auth/github/callback"),
	)

	//Configure port for server to run on
	port := *configPort
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	// Listen and Serve on port from ENV or flag
	r.Run(fmt.Sprintf(":%s", port))
}
