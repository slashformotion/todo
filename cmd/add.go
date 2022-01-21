/*
Copyright © 2021 SLASHFORMOTION <slashformotion@protonmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/slashformotion/todo/internal"
	"github.com/slashformotion/todo/pkg/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your .todo file.",
	Long: `Example:
	$ todo add "fix the ship"
This command will add  a new task named "fix the ship" to your .todo file.`,
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
	t, err := internal.GetTodofile(path)
	if err != nil {
		return err
	}
	newTask, err := todo.NewTask(name, false)
	if err != nil {
		return err
	}
	t.Append(newTask)
	err = internal.SaveTodoFile(t)
	if err != nil {
		return err
	}
	fmt.Printf("Task %q added, now listing\n\n", newTask.Name)
	fmt.Println(t.RenderToScreen())
	return nil
}
