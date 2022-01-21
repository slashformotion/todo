/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/slashformotion/todo/internal"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark the task corresponding to the number provided as finished",
	Long: `Example:
	$ todo done 1
This mark the first task as finished`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := actionDone(args[0], FilePath)
		return err
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func actionDone(taskIndex, path string) error {
	index, err := strconv.Atoi(strings.TrimSpace(taskIndex))
	if err != nil {
		return fmt.Errorf("%q is not a number", taskIndex)
	}

	t, err := internal.GetTodofile(path)
	if err != nil {
		return err
	}
	task, err := t.MarkAsCompleted(index)
	if err != nil {
		return err
	}
	err = internal.SaveTodoFile(t)
	if err != nil {
		return err
	}
	fmt.Printf("Task %q Marked as Completed, now listing\n\n", task.Name)
	fmt.Println(t.RenderToScreen())
	return nil
}
