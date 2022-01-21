/*
Copyright © 2021 SLASHFORMOTION <slashformotion@protonmail.com>

*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/slashformotion/todo/internal"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove the task corresponding to the number provided",
	Long: `Example:
	$ todo rm 2
This command will remove the second task of your .todo file.`,
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

	t, err := internal.GetTodofile(path)
	if err != nil {
		return err
	}

	rmTask, err := t.RemoveTask(index)
	if err != nil {
		fmt.Printf("index=%v does not exist, please run 'todo ls'\n", index)
		os.Exit(0)
	}

	err = internal.SaveTodoFile(t)
	if err != nil {
		return err
	}
	fmt.Printf("Task %q removed, now listing\n\n", rmTask.Name)
	fmt.Println(t.RenderToScreen())
	return nil
}
