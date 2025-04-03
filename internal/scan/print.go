package scan

import (
	"fmt"

	"github.com/fatih/color"
)

func printSecurityAlert(message string) {
	d := color.New(color.BgRed, color.Bold)
	d.Printf("[SECURITY ALERT]:")
	fmt.Printf(" %s\n", message)
}

func printESLintAlert(message string) {
	d := color.New(color.BgYellow, color.Bold)
	d.Printf("[ESLint-Security]:")
	fmt.Printf(" %s\n", message)
}
