package main

import (
	"flag"
	"log"
	"os"

	"github.com/mojotech/feedbag/feedbag"
)

var (
	configPort   = flag.String("port", "3000", "Port to run the server on")
	templatesDir = flag.String("templates", "./templates", "Path to your templates directory")
)

func main() {
	//Parse flags
	flag.Parse()

	// Configure port for server to run on
	port := *configPort
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	err := feedbag.Start(port, *templatesDir)
	if err != nil {
		log.Fatalln("Feedbag failed to start:", err)
	}
}
