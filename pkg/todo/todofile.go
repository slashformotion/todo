package todo

import (
	"fmt"
	"strings"
)

const DefaultExtension string = ".todo"

type TodoFile struct {
	Path  string
	tasks []*Task
}

func NewTodoFile(path string) (*TodoFile, error) {
	err := CheckTodoFileName(path)
	if err != nil {
		return nil, err
	}
	t := &TodoFile{
		Path:  path,
		tasks: make([]*Task, 0),
	}
	return t, nil
}

func CheckTodoFileName(path string) error {
	if len(path) == 0 {
		return fmt.Errorf("path can't be an empty string")
	}
	if !strings.HasSuffix(path, DefaultExtension) {
		return fmt.Errorf("file extension should be '%v'. got='%v'", DefaultExtension, path)
	}
	return nil
}

func (t *TodoFile) Append(task *Task) {
	t.tasks = append(t.tasks, task)
}

func (t *TodoFile) AppendMultiples(tasks []*Task) {
	t.tasks = append(t.tasks, tasks...)
}

func (t *TodoFile) RenderToFile() (res string) {
	for _, task := range t.tasks {
		res += fmt.Sprintf("- [%s] %v\n", task.getTaskStateChar(), task.Name)
	}
	return res
}

func (t *TodoFile) RenderToScreen() (res string) {
	for i, task := range t.tasks {
		res += fmt.Sprintf("%03d [%s] %v\n", i+1, task.getTaskStateChar(), task.Name)
	}
	return res
}

func (t *TodoFile) RemoveTask(index int) (*Task, error) {
	var rmTask *Task
	if index <= 0 || index > len(t.tasks) {
		return nil, fmt.Errorf("index out of bounds ([1,%d]). go=%d", len(t.tasks), index)
	}
	t.tasks, rmTask = removeTask(t.tasks, index-1)
	return rmTask, nil
}

func removeTask(stack []*Task, index int) ([]*Task, *Task) {
	// Remove the element at index i from stack.
	var rmTask *Task = stack[index]
	stack[index] = stack[len(stack)-1] // Copy last element to index i.
	stack[len(stack)-1] = nil          // Erase last element (write nil).
	stack = stack[:len(stack)-1]       // Truncate slice.
	return stack, rmTask
}

func (t *TodoFile) MarkAsCompleted(index int) (*Task, error) {
	task, err := t.executeActionOnSpecificTask(index, MarkTaskAsFinished)
	return task, err
}

// Mark the task at index as uncompleted
func (t *TodoFile) MarkAsUncompleted(index int) (*Task, error) {
	task, err := t.executeActionOnSpecificTask(index, MarkTaskAsUnfinished)
	return task, err
}

// Return the a *Task corresponding to the correct index
func (t *TodoFile) getTask(index int) *Task {
	return t.tasks[index]
}

// Execute the f func(*Task) err on the task located at the index i
// Returns a task pointer and a potential error
func (t *TodoFile) executeActionOnSpecificTask(i int, f func(*Task) error) (*Task, error) {
	if i <= 0 || i > len(t.tasks) {
		return nil, fmt.Errorf("index out of bounds ([1,%d]). go=%d", len(t.tasks), i)
	}
	task := t.getTask(i - 1)
	err := f(task)
	return task, err
}
