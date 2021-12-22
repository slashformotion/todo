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

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	t, err := todo.NewTodoFile(path)
	if err != nil {
		return err
	}
	f, err := os.Open(t.Path)
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

	t.AppendMultiples(tasks)
	content := t.RenderToScreen()
	fmt.Print(content)
	return nil
}
