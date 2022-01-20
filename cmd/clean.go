/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/slashformotion/todo/internal"
	"github.com/slashformotion/todo/pkg/todo"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Erase the content of your .todo file",
	Long: `Exemple:
	$ todo clean
This will erase the content fo you .todo file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := actionClean(FilePath)
		return err
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func actionClean(path string) error {
	t, err := todo.NewTodoFile(path)
	if err != nil {
		return err
	}
	err = internal.SaveTodoFile(t)
	if err != nil {
		return err
	}

	fmt.Println()
	return nil
}
