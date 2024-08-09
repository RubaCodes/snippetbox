package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes(staticDir string) *http.ServeMux {
	mux := http.NewServeMux()
	/** File Server Configuration*/
	fileServer := http.FileServer(http.Dir(staticDir))
	//For matching paths, we strip the "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// handlers

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	return mux
}
