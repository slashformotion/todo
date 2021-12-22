package todo_test

import (
	"testing"

	"github.com/slashformotion/todo/pkg/todo"
)

func TestNew(t *testing.T) {
	_, err := todo.New("", false)
	if err == nil {
		t.Errorf("creating a *Todo should fail on empty string")
	}

	_, err = todo.New("valid string", false)
	if err != nil {
		t.Errorf("creating a *Todo should not fail on valid string")
	}

	task, _ := todo.New("voila", false)
	if task.Completed {
		t.Errorf("'Completed' field should be set to false, got=%v", task.Completed)
	}

	task, _ = todo.New("voila", true)
	if !task.Completed {
		t.Errorf("'Completed' field should be set to true, got=%v", task.Completed)
	}

}

func TestCompleted(t *testing.T) {
	task, _ := todo.New("voila", false)
	task.Complete()
	if !task.Completed {
		t.Error("'.Complete' method should set the 'Completed' field to true")
	}
}
