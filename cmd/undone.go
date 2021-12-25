/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := actionNotDone(args[0], FilePath)
		return err
	},
}

func init() {
	rootCmd.AddCommand(undoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// undoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// undoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func actionNotDone(taskIndex, path string) error {
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
	task, err := t.MarkAsUncompleted(index)
	if err != nil {
		fmt.Printf("index=%v does not exist, please run 'todo ls'\n", index)
		os.Exit(0)
	}

	f, err = overwriteFile(f)
	if err != nil {
		return err
	}

	f.WriteString(t.RenderToFile())
	fmt.Printf("Task %q Marked as not complete, now listing\n\n", task.Name)
	fmt.Println(t.RenderToScreen())
	return nil
}
