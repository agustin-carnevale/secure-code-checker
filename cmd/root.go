package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "firstpass",
	Short: "A CLI tool for scanning JavaScript/TypeScript repos for security threats.",
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
}
