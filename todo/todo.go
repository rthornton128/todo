// Copyright 2016 Rob Thornton, All Rights Reserved.
// Governed by a Simplified BSD license.  See the LICENSE file or you may
// find a copy at: http://www.freebsd.org/copyright/freebsd-license.html

// Todo
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var addr = "localhost:8080"

func main() {
	// init handlers
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("/static")))
	r.HandleFunc("/tasks", errHandler(handleAllTasks)).Methods("GET")
	r.HandleFunc("/tasks", errHandler(handleNewTask)).Methods("POST")
	r.HandleFunc("/tasks/{id}", errHandler(handleQueryTask)).Methods("GET")
	r.HandleFunc("/tasks/{id}", errHandler(handleUpdateTask)).Methods("PUT")
	//http.Handle(r)

	// start server
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalf("listen and serve: %v", err)
	}
}
