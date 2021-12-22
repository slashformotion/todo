/*
Copyright © 2021 SLASHFORMOTION <slashformotion@protonmail.com>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/slashformotion/todo/pkg/todo"
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
	t, err := todo.NewTodoFile(path)
	if err != nil {
		return err
	}
	content := t.RenderToFile()
	if _, err := os.Stat(t.Path); err == nil {
		fmt.Printf("The file %q already exists.\nPlease delete the file %q before using this command again\n", t.Path, t.Path)
		os.Exit(0)
	}
	f, err := os.Create(t.Path)
	if err != nil {
		fmt.Printf("Error while creating %q\n", t.Path)
		os.Exit(1)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}

	}()
	f.WriteString(content)

	return nil

}
