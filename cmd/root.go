/*
Copyright © 2021 SLASHFORMOTION <slashformotion@protonmail.com>

*/
package cmd

import (
	"os"

	"github.com/slashformotion/todo/pkg/todo"
	"github.com/spf13/cobra"
)

const DefautFileName string = "todo" + todo.DefaultExtension
const DefaultFilePath = "./" + DefautFileName

// Represent a file path
var FilePath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A CLI to manage .todo files.",
	Long:  `A CLI too to manage .todo file, please head to https://github.com/slashformotion/todo to learn more about todo files.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&FilePath, "path", DefaultFilePath, "path to the .todo file")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
