/*
Copyright © 2021 SLASHFORMOTION <slashformotion@protonmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/slashformotion/todo/internal"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List the tasks in your .todo file.",
	Long: `Example:
	$ todo ls
This command will list the tasks your .todo file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := actionLs(FilePath)
		return err
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func actionLs(path string) error {
	tf, err := internal.GetTodofile(path)
	if err != nil {
		return err
	}
	content := tf.RenderToScreen()
	fmt.Print(content)
	return nil
}
