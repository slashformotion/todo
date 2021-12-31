package todo_test

import (
	"testing"

	"github.com/slashformotion/todo/pkg/todo"
)

func TestNewTask(t *testing.T) {
	_, err := todo.NewTask("", false)
	if err == nil {
		t.Errorf("creating a *Todo should fail on empty string")
	}

	_, err = todo.NewTask("valid string", false)
	if err != nil {
		t.Errorf("creating a *Todo should not fail on valid string")
	}

	task, _ := todo.NewTask("voila", false)
	if task.IsFinished() {
		t.Errorf("'Finished' field should be set to false, got=%v", task.IsFinished())
	}

	task, _ = todo.NewTask("voila", true)
	if !task.IsFinished() {
		t.Errorf("'Finished' field should be set to true, got=%v", task.IsFinished())
	}

}

func TestFinished(t *testing.T) {
	task, _ := todo.NewTask("voila", false)
	task.Finished()
	if !task.IsFinished() {
		t.Error("'.Complete' method should set the 'Finished' field to true")
	}
}

func TestIsFinished(t *testing.T) {
	task, _ := todo.NewTask("voila", false)
	task.Finished()
	if !task.IsFinished() {
		t.Error("'.Complete' method should set the 'Finished' field to true")
	}
}

func TestMarkTaskAsFinished(t *testing.T) {
	task, _ := todo.NewTask("voila", false)
	err := todo.MarkTaskAsFinished(task)
	if err != nil {
		t.Error("'MarkTaskAsFinished' function should not return an error")
	}
	if !task.IsFinished() {
		t.Error("'MarkTaskAsFinished' function should set the 'Finished' field to true")
	}
}

func TestMarkTaskAsUnfinished(t *testing.T) {
	task, _ := todo.NewTask("voila", true)
	err := todo.MarkTaskAsUnfinished(task)
	if err != nil {
		t.Error("'MarkTaskAsUnfinished' function should not return an error")
	}
	if task.IsFinished() {
		t.Error("'MarkTaskAsUnfinished' function should set the 'Finished' field to false")
	}
}
