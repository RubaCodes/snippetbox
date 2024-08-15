package main

import "flag"

func setupConfig() config {
	var conf config
	// parse flags if exist
	flag.StringVar(&conf.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&conf.dbConn, "dbconn", "postgres://admin:password@localhost:5432/snippetbox?sslmode=disable", "Postgres connection string")
	//If any errors are encountered during parsing the application will be terminated.
	flag.Parse()
	return conf
}
