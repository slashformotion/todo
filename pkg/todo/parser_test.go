package todo_test

import (
	"testing"

	"github.com/slashformotion/todo/pkg/todo"
)

const validTestString string = `- [] a task list item
- [] list _syntax_ required
- [] normal **formatting**
-[x]more relaxed about syntax than GFM
- [x]so inconsistencies should be ok
- [] incomplete
- [x] completed`

func TestValidParse(t *testing.T) {
	tasks, err := todo.Parse(validTestString)
	if err != nil {
		t.Errorf("Parse should not return an error on valid string: got=%v", err)
	}
	if len(tasks) != 7 {
		t.Errorf("Parse didn't give the right number of task. Expected 7 got=%d", len(tasks))
	}
}
