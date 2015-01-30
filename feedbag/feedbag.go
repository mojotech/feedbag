package feedbag

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tommy351/gin-cors"
)

var (
	dbmap        = setupDb()
	activityChan = make(chan []Activity)
	TemplatesDir string
	Templates    []*Template
)

//Raw Json type for Type Converter
type RawJson map[string]interface{}

type TypeConverter struct{}

func Start(port, templatesDir string) error {
	defer dbmap.Db.Close()

	// Process our templates
	TemplatesDir = templatesDir
	Templates = getTemplates(TemplatesDir)

	// Setup Goth Authentication
	goth.UseProviders(
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:3000/auth/github/callback"),
	)

	// Setup Socket.io server and related activity fetching
	socketServer, err := SetupSocketIO()
	if err != nil {
		return err
	}

	err = StartSocketPusher(socketServer, activityChan)
	if err != nil {
		return err
	}

	err = StartExistingUsers(activityChan)
	if err != nil {
		return err
	}

	// Start up gin and its friends
	r := gin.Default()
	r.Use(cors.Middleware(cors.Options{AllowCredentials: true}))
	SetupRoutes(r, socketServer)
	r.Run(fmt.Sprintf(":%s", port))

	return nil
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func (me TypeConverter) ToDb(val interface{}) (interface{}, error) {
	switch t := val.(type) {
	case RawJson:
		b, err := json.Marshal(t)
		if err != nil {
			return "", err
		}
		return string(b), nil

	case ActivityPayload:
		b, err := json.Marshal(t)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	return val, nil
}

func (me TypeConverter) FromDb(target interface{}) (gorp.CustomScanner, bool) {
	switch target.(type) {
	case *RawJson:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("FromDb: Unable to convert Json to *string")
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}
		return gorp.CustomScanner{new(string), target, binder}, true

	case *ActivityPayload:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("FromDb: Unable to convert Json to *string")
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}
		return gorp.CustomScanner{new(string), target, binder}, true
	}

	return gorp.CustomScanner{}, false
}

func setupDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("sqlite3", "./feedbag.bin")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	dbmap.TypeConverter = TypeConverter{}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	dbmap.AddTableWithName(Activity{}, "activities").SetKeys(true, "Id")
	dbmap.AddTableWithName(ActivityPayload{}, "events").SetKeys(true, "Id").ColMap("GithubId").SetUnique(true)

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}