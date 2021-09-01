package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DTGlov/go-books/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

//indicates the version of the app
const version ="1.0.0" 

//info about the configuration
type config struct {
port int 
env string
db struct{
	dsn string
}
}

type AppStatus struct{
	Status string   `json:"status"`
	Environment string  `json:"environment"`
	Version string `json:"version"`
}

type application struct{
	config config
	errorLog *log.Logger
	infoLog  *log.Logger
	books    *mysql.BookModel
}

func main() {
var cfg config
flag.IntVar(&cfg.port,"port",4000,"Server port to listen on")
flag.StringVar(&cfg.env,"env","development","Application environment (development|production")
flag.StringVar(&cfg.db.dsn,"dsn","chip:jagersenn@/dbbooks?parseTime=true","MySQL Connection String")
flag.Parse()

infoLog := log.New(os.Stdout, "INFO\t",log.Ldate|log.Ltime|log.LUTC)
errorLog := log.New(os.Stderr,"ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)

	db,err := openDB(cfg)
	if err !=nil{
		errorLog.Fatal(err)
	}
	defer db.Close()

app := &application{
	config: cfg,
	errorLog: errorLog,
	infoLog: infoLog,
	books: &mysql.BookModel{DB:db},
}

srv := &http.Server{
	Addr: fmt.Sprintf(":%d",cfg.port),
	Handler: app.routes(),
	IdleTimeout: time.Minute,
	ReadTimeout: 10 * time.Second,
	WriteTimeout: 30 * time.Second,
}

infoLog.Println("Starting server on port", cfg.port)

err = srv.ListenAndServe()
if err!=nil{
	errorLog.Fatal(err)
}
}

func openDB(cfg config)(*sql.DB,error){
db,err :=sql.Open("mysql",cfg.db.dsn)
if err !=nil{
return nil,err
}
ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
defer cancel()
err = db.PingContext(ctx)
if err !=nil{
	return nil,err
}
return db,nil
}