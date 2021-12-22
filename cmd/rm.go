/*
Copyright © 2021 SLASHFORMOTION <slashformotion@protonmail.com>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/slashformotion/todo/pkg/todo"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := actionRm(args[0], FilePath)
		return err
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func actionRm(taskIndex, path string) error {
	index, err := strconv.Atoi(strings.TrimSpace(taskIndex))
	if err != nil {
		return fmt.Errorf("%q is not a number", taskIndex)
	}

	t, err := todo.NewTodoFile(path)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(t.Path, os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("Error while opening %q\n", t.Path)
		os.Exit(1)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}

	}()
	fileContent, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("Error while reading %q\n", t.Path)
		os.Exit(1)
	}
	tasks, err := todo.Parse(string(fileContent))
	if err != nil {
		return err
	}
	t.AppendMultiples(tasks)
	rmTask, err := t.RemoveTask(index)
	if err != nil {
		fmt.Printf("index=%v does not exist, please run 'todo ls'\n", index)
		os.Exit(0)
	}

	f, err = overwriteFile(f)
	if err != nil {
		return err
	}

	f.WriteString(t.RenderToFile())
	fmt.Printf("Task %q removed, now listing\n\n", rmTask.Name)
	fmt.Println(t.RenderToScreen())
	return nil
}

func overwriteFile(f *os.File) (*os.File, error) {
	err := f.Truncate(0)
	if err != nil {
		return nil, err
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	return f, nil
}
