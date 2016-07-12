// Copyright 2016 Rob Thornton, All Rights Reserved.
// Governed by a Simplified BSD license.  See the LICENSE file or you may
// find a copy at: http://www.freebsd.org/copyright/freebsd-license.html

package main

import "net/http"

func errHandler(f func(w http.ResponseWriter, r *http.Response) error) http.HandlerFunc {
	return nil
}

func handleNewTask(w http.ResponseWriter, r *http.Response) error {
	return nil
}
func handleAllTasks(w http.ResponseWriter, r *http.Response) error {
	return nil
}
func handleQueryTask(w http.ResponseWriter, r *http.Response) error {
	return nil
}
func handleUpdateTask(w http.ResponseWriter, r *http.Response) error {
	return nil
}
