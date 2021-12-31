package todo_test

import (
	"testing"

	"github.com/slashformotion/todo/pkg/todo"
)

func TestNewTodoFileValid(t *testing.T) {
	tf, err := todo.NewTodoFile("./todo.todo")
	if err != nil {
		t.Errorf("Valid path should not return an error. got=%v", err)
	}
	if tf == nil {
		t.Errorf("Valid path should not return a nil *TodoFile. got=%v", t)
	}

}

func TestNewTodoFileInvalid(t *testing.T) {
	tf, err := todo.NewTodoFile("./todo.to")
	if err == nil {
		t.Errorf("Invalid path should  return an error. got=%v", err)
	}
	if tf != nil {
		t.Errorf("Invalid path should return a nil *TodoFile. got=%v", tf)
	}
}

func TestAppend(t *testing.T) {
	tf, _ := todo.NewTodoFile("./todo.todo")
	task, _ := todo.NewTask("fix the ship", false)
	tf.Append(task)
}

func TestAppendMultiples(t *testing.T) {
	tf, _ := todo.NewTodoFile("./todo.todo")
	task1, _ := todo.NewTask("fix the ship", false)
	task2, _ := todo.NewTask("fix the ship", false)
	tasks := []*todo.Task{task1, task2}
	tf.AppendMultiples(tasks)
}
