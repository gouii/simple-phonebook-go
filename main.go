package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gouii/simple-phonebook-go/httphandler/phonebook"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gouii/simple-phonebook-go/app"
	"github.com/mmuflih/envgo/conf"
	"github.com/mmuflih/go-httplib/httplib"
)

var config conf.Config
var db *sql.DB

func init() {
	var err error
	fmt.Println(`        _                     _               _   `)
	fmt.Println(`  _ __ | |_   ___  _ _   ___ | |__  ___  ___ | |__`)
	fmt.Println(` | '_ \| ' \ / _ \| ' \ / -_)| '_ \/ _ \/ _ \| / /`)
	fmt.Println(` | .__/|_||_|\___/|_||_|\___||_.__/\___/\___/|_\_\`)
	fmt.Println(` |_|                                              `)
	fmt.Println(``)
	config = conf.NewConfig()
	err, db = app.NewMysqlConn(config)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	muxrr := httplib.NewMuxRequestReader()
	mainroute := mux.NewRouter()

	apiv1 := mainroute.PathPrefix("/api/v1").Subrouter()

	phonebook.NewRoute(apiv1, muxrr, db)

	/** cors */
	headersVal := config.GetStrings("cors.allowed_headers")
	methodsVal := config.GetStrings("cors.allowed_methods")
	originsVal := config.GetStrings("cors.allowed_origins")
	headers := handlers.AllowedHeaders(headersVal)
	methods := handlers.AllowedMethods(methodsVal)
	origins := handlers.AllowedOrigins(originsVal)
	cors := handlers.CORS(headers, methods, origins)
	apiCors := cors(httplib.Logger(mainroute))
	app.Logger("Application is running at ", time.Now().Format("2006-01-02 15:04:05.000"))
	app.Logger("Server listen on", config.GetString(`server.address`))
	log.Fatal(http.ListenAndServe(config.GetString(`server.address`), apiCors))
}
