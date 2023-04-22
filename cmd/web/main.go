package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/apex/gateway"
	"github.com/alexedwards/scs/v2"

	"github.com/bertoxic/bert/drivers"
	"github.com/bertoxic/bert/helpers"
	"github.com/bertoxic/bert/internal/config"
	"github.com/bertoxic/bert/internal/handlers"
	"github.com/bertoxic/bert/internal/render"
	"github.com/bertoxic/bert/models"
)

const portNumber = ":8080"

var sessions *scs.SessionManager
var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger


func main() {
	listener := gateway.ListenAndServe
	if portNumber != ":8081" {
        listener = http.ListenAndServe
    }
	db, err := run()
	if err != nil {
		log.Println(err)
	}

	defer db.SQL.Close()

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/example", handlers.Repo.Example)
	// from := "we@me.com"
	// auth := smtp.PlainAuth("",from,"","localhost")
	// err =smtp.SendMail("localhost:1025",auth,from,[]string{"h@hw.com"},[]byte("hello world"))

	defer close(app.MailChan)
	listenForMail()

	// msg := models.MailData{
	// 	To:      "me@gmail.com",
	// 	From:    "ok@fmail.com",
	// 	Subject: "Hellllo peoplre",
	// 	Content: "",
	// }
	// app.MailChan  <-msg
	// if err != nil {
	// 	log.Println(err)
	// }

	// srv := &http.Server{
	// 	Addr:    portNumber,
	// 	Handler: routes(&app),
	// }

	//_ = http.ListenAndServe(portNumber, nil)
	//err = srv.ListenAndServe()
	listener(portNumber,routes(&app))
	log.Fatal(err)
}

func run() (*drivers.DB, error) {
	gob.Register(models.Reservation{})
	gob.Register(models.Restriction{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(map[string]int{})
	//Read flags
	inProduction := *flag.Bool("production", true,"Application is in production")
	useCache := *flag.Bool("cache", false,"Use template cache")
	dbHost := flag.String("dbhost","localhost","Database Host")
	dbName := flag.String("dbname","","Database name")
	dbUser := flag.String("dbuser","","Database user")
	dbPass := flag.String("dbpass","","Database password")
	dbPort := flag.String("dbport","5432","Database port")
	dbSSL := flag.String("dbssl","disable","Database ssl settings (disable, prefer, require)")

	flag.Parse() 
	if *dbName=="" || *dbUser ==""{
		fmt.Println("required field absent in run.sh")
	}

	app.InProduction = inProduction
	app.UserCache= useCache     

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	sessions = scs.New()
	sessions.Cookie.Persist = true
	sessions.Lifetime = 24 * time.Hour
	sessions.Cookie.Secure = app.InProduction
	sessions.Cookie.SameSite = http.SameSiteLaxMode
	app.Sessions = sessions

	mailChan := make (chan models.MailData)
	app.MailChan = mailChan
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",*dbHost,*dbPort,*dbName,*dbUser,*dbPass, *dbSSL )
	log.Println("Connecting to database")
	db, err := drivers.ConnectSQL(connectionString)

	if err != nil {
		log.Fatal("cannot connnect to data base")
	}
	log.Println("just connected to the database")
	tc, err := render.CreateTemplateCache()

	render.NewRenderer(&app)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	app.TemplateCache = tc
 	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	helpers.NewHelpers(&app)
	return db, nil
}
