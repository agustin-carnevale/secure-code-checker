package scan

import (
	"fmt"
	"os/exec"
)

// TODO: this is not working as it is now, polish this command to run ESlint with the security plugin
func runESLintSecurityCheck(filePath string) {
	cmd := exec.Command("npx", "eslint", "--no-ignore", "--config", "./eslint.config.js", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("[ESLint-Security] Error running ESLint on %s: %v\n", filePath, err)
	}
	if len(output) > 0 {
		logMessage := fmt.Sprintf("%s\n%s", filePath, string(output))
		printESLintAlert(logMessage)

	}
}
