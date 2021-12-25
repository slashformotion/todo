package todo

import "fmt"

// Those are constant slices, believe me... Explaination here => https://qvault.io/golang/golang-constant-maps-slices/

// Return the caracter signaling that the task is done
func DoneCharacters() []string { return []string{"X", "x"} }

// Return the caracter signaling that the task is not done
func TodoCharacter() []string { return []string{" ", ""} }

func getStateChar(b bool) string {
	if b {
		return DoneCharacters()[0]
	} else {
		return TodoCharacter()[0]
	}
}

// Task represents a Task in .todo files
// 'Name' is the name of the task
// 'Completed' is self-explainatory
type Task struct {
	Name      string
	Completed bool
}

func (t *Task) getTaskStateChar() string {
	return getStateChar(t.Completed)
}

// Create a new Task with a name and a state of completion (boolean)
func New(name string, done bool) (*Task, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("name should be not be an empty string")
	}
	t := &Task{
		Name:      name,
		Completed: done,
	}
	return t, nil
}

// Mark the task as complete
func (t *Task) Complete() {
	t.Completed = true
}

// Mark the task as uncomplete
func (t *Task) Uncompleted() {
	t.Completed = false
}

func MarkTaskAsCompleted(task *Task) error {
	task.Complete()
	return nil
}

func MarkTaskAsUncompleted(task *Task) error {
	task.Uncompleted()
	return nil
}
