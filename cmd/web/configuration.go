package main

import "flag"

func setupConfig() config {
	var conf config
	// parse flags if exist
	flag.StringVar(&conf.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&conf.staticDir, "static-dir", "./ui/static", "Path to static assets")
	//If any errors are encountered during parsing the application will be terminated.
	flag.Parse()
	return conf
}
