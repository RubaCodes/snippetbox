package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form"
	_ "github.com/lib/pq"
	"snippetbox.rubacodes.com/internal/models"
)

type config struct {
	addr      string
	staticDir string
	dbConn    string
}
type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	snippets       *models.SnippetModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {
	//start-up
	conf := setupConfig()
	// Crate custom loggers
	infoLog := log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR \t", log.Ldate|log.Ltime|log.Lshortfile)
	db, err := openDB(conf.dbConn)
	if err != nil {
		errorLog.Fatal(err)
	}
	// We also defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exits.
	defer db.Close()
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}
	// Use the scs.New() function to initialize a new session manager. Then we
	// configure it to use our MySQL database as the session store, and set a
	// lifetime of 12 hours (so that sessions automatically expire 12 hours
	// after first being created).
	sessionManager := scs.New()
	sessionManager.Store = postgresstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour
	//create a struct to hold global dependency SINGLETONS
	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{
			DB: db,
		},
		templateCache:  templateCache,
		formDecoder:    form.NewDecoder(),
		sessionManager: sessionManager,
	}
	// instanciate a new httpServer with the custom configurations
	srv := &http.Server{
		Addr:     conf.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(conf.staticDir),
	}
	infoLog.Printf("Starting server on http://localhost%s", conf.addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
