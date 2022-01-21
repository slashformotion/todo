package todo

import (
	"fmt"
	"regexp"
	"strings"
)

const parsingString string = `(?m)^-(?:\s|)\[(X|x|| )\](?:\s|)(.*)$`

func Parse(s string) (tasks []*Task, err error) {
	var rgx *regexp.Regexp = regexp.MustCompile(parsingString)
	matchs := rgx.FindAllStringSubmatch(s, -1)
	for _, match := range matchs {
		task, err := ProcessMatch(match)
		if err != nil {
			return nil, fmt.Errorf("can't process  '%s'", match[0])
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func ProcessMatch(match []string) (*Task, error) {
	if len(match) != 3 {
		return nil, fmt.Errorf("match args should be of size . got=%v", len(match))
	}
	name := strings.TrimSpace(match[2])
	if len(name) == 0 {
		return nil, fmt.Errorf("can't parse the .todo file correctly")
	}
	return NewTask(name, checkIfTodoIsCompleted(match[1]))
}

func checkIfTodoIsCompleted(s string) bool {
	for _, sym := range UnfinishedCharacters() {
		if s == sym {
			return false
		}
	}
	return true
}
