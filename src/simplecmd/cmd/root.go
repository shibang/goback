package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "simplecmd",
	Short: "simplecmd is a command line tool",
	Long:  "simplecmd is a command line tool",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("cobra output")
	},
}

// Execute excute cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
