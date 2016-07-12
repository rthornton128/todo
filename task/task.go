// Copyright 2016 Rob Thornton, All Rights Reserved.
// Governed by a Simplified BSD license.  See the LICENSE file or you may
// find a copy at: http://www.freebsd.org/copyright/freebsd-license.html

// Package task provides a task manager using an SQLite3 datastore.
// The manager can be either in-memory or on physical media.
package task

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3" // to register sqlite3 in sql.Open()
)

// Task has some kind of bs description
type Task struct {
	ID   int64
	Desc string
	Done bool
}

// NewTask returns a new Task. If desc is an empty string, error will be non-nil
func NewTask(desc string) (*Task, error) {
	if desc == "" {
		return nil, errors.New("empty description")
	}
	return &Task{Desc: desc}, nil
}

// Manager is an SQLite3 store.
type Manager struct {
	db *sql.DB
}

// NewManager takes the path to a new database store or ':memory:' may be used for
// an in-memory database
func NewManager(path string) (*Manager, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Tasks (ID INTEGER PRIMARY KEY,Desc TEXT,Done INT);")
	if err != nil {
		return nil, err
	}
	return &Manager{db: db}, nil
}

// All returns all currently stored tasks
func (m *Manager) All() ([]*Task, error) {
	s := "SELECT ID,Desc,Done FROM Tasks;"
	rows, err := m.db.Query(s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []*Task
	for rows.Next() {
		var t Task
		err = rows.Scan(&t.ID, &t.Desc, &t.Done)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	return tasks, rows.Err()
}

// Query searches the database for a task and returns the result
func (m *Manager) Query(id int64) (*Task, error) {
	s := "SELECT ID,Desc,Done FROM Tasks WHERE ID=?;"
	var t Task
	return &t, m.db.QueryRow(s, id).Scan(&t.ID, &t.Desc, &t.Done)
}

// Store inserts a new task into or updates an existing one in the database.
// Error will be a result of an SQL error
func (m *Manager) Store(t *Task) (err error) {
	s := "INSERT OR REPLACE INTO Tasks VALUES (?,?,?);"
	if t.ID == 0 {
		var r sql.Result
		r, err = m.db.Exec(s, nil, t.Desc, t.Done)
		if err != nil {
			return err
		}
		t.ID, err = r.LastInsertId()
	} else {
		_, err = m.db.Exec(s, t.ID, t.Desc, t.Done)
	}
	return err
}