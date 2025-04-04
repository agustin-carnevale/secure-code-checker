package scan

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	es "github.com/evanw/esbuild/pkg/api"
)

func basicFilesScanCheck(rootDir string) {
	fmt.Println("Starting files scan...")
	// Traverse all the the files and folders from the rootDir
	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// For now only care about javascript/typescript files (js/ts)
		if !d.IsDir() && (strings.HasSuffix(d.Name(), ".js") || strings.HasSuffix(d.Name(), ".ts")) {
			// fmt.Printf("Scanning file: %s\n", path)
			scanFileAST(path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error scanning files: %v\n", err)
	}
}

func scanFileAST(filePath string) {
	content, err := os.ReadFile(filePath) // #nosec G304
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		return
	}

	result := es.Transform(string(content), es.TransformOptions{
		Loader:       es.LoaderTS,
		Format:       es.FormatESModule,
		MinifySyntax: true,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("[WARNING] Syntax errors found in %s\n", filePath)
		for _, e := range result.Errors {
			fmt.Printf("  - %s\n", e.Text)
		}
	}

	if strings.Contains(string(result.Code), "eval(") || strings.Contains(string(result.Code), "Function(") {
		// fmt.Printf("[ALERT] Potentially unsafe function found in %s\n", filePath)
		logMessage := fmt.Sprintf("Potentially unsafe function found in %s", filePath)
		printSecurityAlert(logMessage)
	}
}
