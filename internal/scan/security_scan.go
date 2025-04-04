package scan

func RunAllSecurityChecks(tmpDir string) {
	// Add here all the necessary checks
	basicFilesScanCheck(tmpDir)
	runESLintSecurityCheck(tmpDir)
}
