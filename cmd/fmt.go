/*
Copyright © 2022 SLASHFORMOTION <slashformotion@protonmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/slashformotion/todo/internal"
	"github.com/spf13/cobra"
)

// fmtCmd represents the fmt command
var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Format you .todo file.",
	Long: `Example:
	$ todo fmt
This commad format your .todo file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := actionFmt(FilePath)
		return err
	},
}

func init() {
	rootCmd.AddCommand(fmtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fmtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fmtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func actionFmt(path string) error {
	t, err := internal.GetTodofile(path)
	if err != nil {
		return err
	}
	err = internal.SaveTodoFile(t)
	if err != nil {
		return err
	}
	fmt.Printf("%q formated, now listing\n\n", t.Path)
	fmt.Println(t.RenderToScreen())
	return nil
}
