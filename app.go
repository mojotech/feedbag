package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	_ "github.com/mattn/go-sqlite3"
)

var (
	configPort   = flag.String("port", "3000", "Port to run the server on")
	templatesDir = flag.String("templates", "./templates", "Path to your templates directory")
	dbmap        = setupDb()
)

func main() {
	//Setup gin
	r := gin.Default()
	setupRoutes(r)

	// Close the database connection if we fail
	defer dbmap.Db.Close()

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

func setupDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("sqlite3", "./tmp/feedbag.bin")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
