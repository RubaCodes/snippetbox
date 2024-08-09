package main

import (
	"log"
	"net/http"
	"os"
)

type config struct {
	addr      string
	staticDir string
}
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	//start-up
	conf := setupConfig()
	// Crate custom loggers
	infoLog := log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR \t", log.Ldate|log.Ltime|log.Lshortfile)
	//create a struct to hold global dependency SINGLETONS
	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	// instanciate a new httpServer with the custom configurations
	srv := &http.Server{
		Addr:     conf.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(conf.staticDir),
	}
	infoLog.Printf("Starting server on %s", conf.addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
