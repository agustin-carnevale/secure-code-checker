package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:   "firstpass",
	Short: "A CLI tool for scanning JavaScript/TypeScript repos for security threats.",
}

var versionCmd = &cobra.Command{
	Use:   "print-version",
	Short: "Print the version number of firstpass",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("firstpass version %s\n", version)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Initialize commands
func init() {
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(versionCmd)
}
