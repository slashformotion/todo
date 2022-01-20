package internal

import (
	"fmt"
	"os"

	"github.com/slashformotion/todo/pkg/todo"
	"github.com/spf13/afero"
)

func GetTodofile(path string) (*todo.TodoFile, error) {
	todofile, err := todo.NewTodoFile(path)
	if err != nil {
		return nil, err
	}
	file, err := Fs.OpenFile(path, os.O_RDWR, os.ModeAppend)
	if err != nil {
		return todofile, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}

	}()
	// load todofile
	readAndLoadFile(todofile, file)
	return todofile, nil
}

// Save and Close the file
func SaveTodoFile(t *todo.TodoFile) error {
	file, err := Fs.OpenFile(t.Path, os.O_RDWR, os.ModeAppend)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}

	}()
	content := t.RenderToFile()
	overwriteFile(file)
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func readAndLoadFile(t *todo.TodoFile, f afero.File) error {
	fileContent, err := afero.ReadAll(f)
	if err != nil {
		return fmt.Errorf("error while reading %q", f.Name())
	}
	tasks, err := todo.Parse(string(fileContent))
	if err != nil {
		return err
	}
	t.AppendMultiples(tasks)
	return nil
}

func overwriteFile(f afero.File) error {
	err := f.Truncate(0)
	if err != nil {
		return err
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		return err
	}
	return nil
}
