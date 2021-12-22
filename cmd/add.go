/*
Copyright © 2021 SLASHFORMOTION <slashformotion@protonmail.com>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/slashformotion/todo/pkg/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("add called")
		err := actionAdd(FilePath, args[0])
		return err
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func actionAdd(path, name string) error {
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
	newTask, err := todo.New(name, false)
	if err != nil {
		return err
	}

	t.AppendMultiples(tasks)
	t.Append(newTask)

	f, err = overwriteFile(f)
	if err != nil {
		return err
	}

	f.WriteString(t.RenderToFile())
	fmt.Printf("Task %q added, now listing\n\n", newTask.Name)
	fmt.Println(t.RenderToScreen())
	return nil
}
