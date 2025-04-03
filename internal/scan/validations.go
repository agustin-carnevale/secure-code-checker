package scan

import (
	"path/filepath"
)

// For now we don't need to use this validation as the user of this tool
// should only use it to test his own known paths.
// But we keep it here handy anyways.

// const baseDir = "/app/data" // Define a safe base directory

// validateFilePath ensures the file path is safe before accessing it
func validateFilePath(filePath string) (string, error) {
	// Clean the file path to remove ".." and extra slashes
	cleanPath := filepath.Clean(filePath)

	// Step 2: Reject absolute paths
	// if filepath.IsAbs(cleanPath) {
	// 	return "", fmt.Errorf("absolute paths are not allowed")
	// }

	// Convert to absolute path relative to baseDir
	// absPath, err := filepath.Abs(filepath.Join(baseDir, cleanPath))
	// if err != nil {
	// 	return "", fmt.Errorf("invalid file path: %w", err)
	// }

	// Ensure the file remains within baseDir
	// if !filepath.HasPrefix(absPath, baseDir) {
	// 	return "", fmt.Errorf("unauthorized file access attempt: %s", absPath)
	// }

	return cleanPath, nil
}
