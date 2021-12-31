package todo

import "fmt"

// Those are constant slices, believe me... Explaination here => https://qvault.io/golang/golang-constant-maps-slices/

// CONSTANT: Return the caracter signaling that the task is done
//     Explaination here => https://qvault.io/golang/golang-constant-maps-slices/
func FinishedCharacters() []string { return []string{"X", "x"} }

// CONSTANT: Return the caracter signaling that the task is not done
//     Explaination here => https://qvault.io/golang/golang-constant-maps-slices/
func UnfinishedCharacters() []string { return []string{" ", ""} }

func getStateChar(b bool) string {
	if b {
		return FinishedCharacters()[0]
	} else {
		return UnfinishedCharacters()[0]
	}
}

// Task represents a Task in .todo files
// 'Name' is the name of the task
// 'Finished' is self-explainatory
type Task struct {
	Name     string
	finished bool
}

func (t *Task) getTaskStateChar() string {
	return getStateChar(t.finished)
}

// Create a new Task with a name and a state of completion (boolean)
func NewTask(name string, done bool) (*Task, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("name should be not be an empty string")
	}
	t := &Task{
		Name:     name,
		finished: done,
	}
	return t, nil
}

// Mark the task as Finished
func (t *Task) Finished() {
	t.finished = true
}

// Mark the task as Unfinished
func (t *Task) Unfinished() {
	t.finished = false
}

// return the true is the task is finished
func (t *Task) IsFinished() bool {
	return t.finished
}

// Mark the task as Finished
func MarkTaskAsFinished(task *Task) error {
	task.Finished()
	return nil
}

// Mark the task as Unfinished
func MarkTaskAsUnfinished(task *Task) error {
	task.Unfinished()
	return nil
}
