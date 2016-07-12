// Copyright 2016 Rob Thornton, All Rights Reserved.
// Governed by a Simplified BSD license.  See the LICENSE file or you may
// find a copy at: http://www.freebsd.org/copyright/freebsd-license.html

// Package serve contains the HTTP handlers for Todo. It also provides
// an errorHandler middleware wrapper.
package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rthornton128/todo/task"
)

func errHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := f(w, r); err != nil {
			switch err.(type) {
			// database error = StatusInternalServerError
			// badRequest
			// ??
			}
		}
		// send StatusOk
	}
}

func handleNewTask(w http.ResponseWriter, r *http.Request) error {
	var t task.Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return err
	}
	return m.Store(&t)
}

func handleAllTasks(w http.ResponseWriter, r *http.Request) error {
	all, err := m.All()
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(all)
}

func handleQueryTask(w http.ResponseWriter, r *http.Request) error {
	tmp := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(tmp, 10, 64)
	if err != nil {
		return err
	}
	t, err := m.Query(id)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(t)
}

func handleUpdateTask(w http.ResponseWriter, r *http.Request) error {
	tmp := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(tmp, 10, 64)
	if err != nil {
		return err
	}
	var t task.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		return err
	}
	if t.ID != id {
		return errors.New("body and URI ID do not match")
	}
	return m.Store(&t)
}
