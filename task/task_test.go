// Copyright 2016 Rob Thornton, All Rights Reserved.
// Governed by a Simplified BSD license.  See the LICENSE file or you may
// find a copy at: http://www.freebsd.org/copyright/freebsd-license.html

package task_test

import (
	"testing"

	"github.com/rthornton128/todo/task"
)

func TestNewTask(t *testing.T) {
	desc := "test task"
	ts, err := task.NewTask(desc)
	if err != nil {
		t.Fatal("err not nil")
	}
	if ts == nil {
		t.Fatal("task was nil")
	}
	if ts.Desc != desc {
		t.Fatalf("expected %v, got %v", desc, ts.Desc)
	}
}

func TestNewTaskNilDesc(t *testing.T) {
	_, err := task.NewTask("")
	if err == nil {
		t.Fatal("err is nil")
	}
}

func TestNewManager(t *testing.T) {
	m, err := task.NewManager(":memory:")
	if m == nil {
		t.Fatal("manager is nil")
	}
	if err != nil {
		t.Fatal("err not nil")
	}
	all, err := m.All()
	if err != nil {
		t.Fatal("err not nil:", err)
	}
	if len(all) != 0 {
		t.Fatalf("got %v tasks", len(all))
	}
}

func TestManagerStore(t *testing.T) {
	m, err := task.NewManager(":memory:")
	if err != nil {
		t.Fatal("err not nil")
	}
	desc := "test task"
	nt, err := task.NewTask(desc)
	if err != nil {
		t.Fatal("failed to create task")
	}
	err = m.Store(nt)
	if err != nil {
		t.Fatal(err)
	}
	all, err := m.All()
	if err != nil {
		t.Fatal("err not nil:", err)
	}
	if len(all) != 1 {
		t.Fatalf("got %v tasks", len(all))
	}
	if all[0].Desc != desc {
		t.Fatalf("expected %v, got %v", desc, all[0].Desc)
	}
}

func TestManagerTwoStores(t *testing.T) {
	m, err := task.NewManager(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	t0, _ := task.NewTask("task0")
	t1, _ := task.NewTask("task1")
	err = m.Store(t0)
	if err != nil {
		t.Fatal(err)
	}
	err = m.Store(t1)
	if err != nil {
		t.Fatal(err)
	}
	all, err := m.All()
	if err != nil {
		t.Fatal("err not nil:", err)
	}
	if len(all) != 2 {
		t.Fatalf("got %v tasks", len(all))
	}
	if all[1].Desc != "task1" {
		t.Fatalf("got %v", all[1].Desc)
	}
}

func TestManagerTwoStoresAndQuery(t *testing.T) {
	m, err := task.NewManager(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	t0, _ := task.NewTask("task0")
	t1, _ := task.NewTask("task1")
	m.Store(t0)
	m.Store(t1)
	r, err := m.Query(t0.ID)
	if err != nil {
		t.Fatal(err)
	}
	if r.Desc != t0.Desc {
		t.Fatalf("expected %v, got %v", t0.Desc, r.Desc)
	}
}
