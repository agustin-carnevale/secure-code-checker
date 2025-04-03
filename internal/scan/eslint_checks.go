package scan

import (
	"fmt"
	"os/exec"
)

func runESLintSecurity(filePath string) {
	cmd := exec.Command("npx", "eslint", "--no-eslintrc", "--rule", "security/detect-eval-with-expression:error", "--rule", "security/detect-object-injection:error", "--rule", "security/detect-unsafe-regex:error", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("[ESLint-Security] Error running ESLint on %s: %v\n", filePath, err)
	}
	if len(output) > 0 {
		logMessage := fmt.Sprintf("%s\n%s", filePath, string(output))
		printESLintAlert(logMessage)

	}
}
