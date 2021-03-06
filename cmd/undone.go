/*
Copyright © 2021 SLASHFORMOTION <slashformotion@protonmail.com>

*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/slashformotion/todo/internal"
	"github.com/spf13/cobra"
)

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "Mark the task corresponding to the number provided as unfinished",
	Long: `Example:
	$ todo done 1
This mark the first task as unfinished`,
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

	t, err := internal.GetTodofile(path)
	if err != nil {
		return err
	}
	task, err := t.MarkAsUncompleted(index)
	if err != nil {
		return err
	}
	err = internal.SaveTodoFile(t)
	if err != nil {
		return err
	}
	fmt.Printf("Task %q Marked as not complete, now listing\n\n", task.Name)
	fmt.Println(t.RenderToScreen())
	return nil
}
