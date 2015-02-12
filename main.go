package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/mojotech/feedbag/feedbag"
)

var (
	port         string
	templatesDir string
	publicDir    string
)

func flags() {
	flag.StringVar(&port, "port", "3000", "Port to run the server on")
	flag.StringVar(&templatesDir, "templates", "", "Path to your templates directory")
	flag.StringVar(&publicDir, "public", "", "Path to your public assets folder")

	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.Parse()
}

func folders() {
	_, file, _, _ := runtime.Caller(0)
	here := filepath.Dir(file)

	if publicDir == "" {
		publicDir = filepath.Join(here, "/public")
	}

	if templatesDir == "" {
		templatesDir = filepath.Join(here, "/templates")
	}
}

func init() {
	flags()
	folders()
}

func main() {
	// Configure port for server to run on
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	err := feedbag.Start(port, templatesDir, publicDir)
	if err != nil {
		log.Fatalln("Feedbag failed to start:", err)
	}
}
