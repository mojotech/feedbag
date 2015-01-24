package main

import (
	"fmt"
	"os"

	"feedbag/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"feedbag/Godeps/_workspace/src/github.com/markbates/goth"
	"feedbag/Godeps/_workspace/src/github.com/markbates/goth/providers/github"
)

func main() {
	//Setup gin
	r := gin.Default()
	setupRoutes(r)

	// Setup Goth Authentication
	goth.UseProviders(
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:3000/auth/github/callback"),
	)

	port := os.Getenv("PORT")

	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%s", port))
}
