/*
Copyright © 2021 SLASHFORMOTION <slashformotion@protonmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/slashformotion/todo/internal"
	"github.com/slashformotion/todo/pkg/todo"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a brand new .todo file",
	Long: `Example:
	$ todo init
This will create a brand new .todo file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := actionInit(FilePath)
		return err
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func actionInit(path string) error {
	todofile, err := todo.NewTodoFile(path)
	if err != nil {
		return err
	}
	exists, err := internal.FileExists(path)
	if err != nil {
		return err
	}
	if empty, _ := afero.IsEmpty(internal.Fs, path); exists && !empty {
		return fmt.Errorf("can't init a file that exists and is not empty")
	}
	internal.SaveTodoFile(todofile)
	return nil
}
