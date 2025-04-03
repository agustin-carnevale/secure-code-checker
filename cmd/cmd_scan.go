package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/agustin-carnevale/secure-code-checker/internal/scan"
	git "github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan [-d|-r] <path_or_url>",
	Short: "Scan files from a repository (-r) or local directory (-d)",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		fromDir, _ := cmd.Flags().GetBool("dir")
		fromRepo, _ := cmd.Flags().GetBool("repo")

		if !fromDir && !fromRepo {
			return errors.New("you must specify either -d for directory or -r for repository")
		}
		if fromDir && fromRepo {
			return errors.New("you can only specify -d or -r")
		}
		if len(args) < 1 {
			return errors.New("you must provide a path (for -d) or a URL (for -r)")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		pathOrURL := args[0]
		fromDir, _ := cmd.Flags().GetBool("dir")
		fromRepo, _ := cmd.Flags().GetBool("repo")

		if fromDir {
			scanFromDirectory(pathOrURL)
		} else if fromRepo {
			scanFromRepository(pathOrURL)
		}
	},
}

func scanFromDirectory(pathToDir string) {
	fmt.Printf("Scanning local directory: %s\n", pathToDir)
	scan.RunAllSecurityChecks(pathToDir)
}

func scanFromRepository(repoUrl string) {
	fmt.Printf("Scanning repository: %s\n", repoUrl)

	tempDir := filepath.Join(os.TempDir(), "codescanner_repo")
	os.RemoveAll(tempDir) // Clean up previous runs

	_, err := git.PlainClone(tempDir, false, &git.CloneOptions{
		URL:      repoUrl,
		Depth:    1,   // Shallow clone for speed
		Progress: nil, //os.Stdout,
	})
	if err != nil {
		fmt.Printf("Error cloning repository: %v\n", err)
		return
	}
	fmt.Printf("Repository cloned to: %s\n", tempDir)

	scan.RunAllSecurityChecks(tempDir)
}

// Initialize scan command flags
func init() {
	scanCmd.Flags().BoolP("dir", "d", false, "Scan a local directory")
	scanCmd.Flags().BoolP("repo", "r", false, "Scan a repository")

	// Ensure that only one flag is used
	scanCmd.MarkFlagsMutuallyExclusive("dir", "repo")
}
