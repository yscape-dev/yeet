package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync files to the destination server",
	Long:  `Sync files from the current working directory to the destination server using rsync.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runSync(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}

func runSync() error {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	// Get the base name of the current directory
	dirName := filepath.Base(cwd)

	// Construct the destination path
	destPath := fmt.Sprintf("%s@%s:~/%s", currentConfig.Destination.Username, currentConfig.Destination.Host, dirName)

	// Construct rsync command
	args := []string{
		"-avz",
		"--delete",
		"-e", fmt.Sprintf("ssh -p %d", currentConfig.Destination.Port),
	}

	// Add include patterns
	for _, include := range currentConfig.Sync.Include {
		args = append(args, "--include", include)
	}

	// Add exclude pattern
	if currentConfig.Sync.Exclude != "" {
		args = append(args, "--exclude", currentConfig.Sync.Exclude)
	}

	// Add source and destination
	args = append(args, ".", destPath)

	// Run rsync command
	cmd := exec.Command("rsync", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Syncing files to destination server...")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("rsync failed: %w", err)
	}

	fmt.Println("Sync completed successfully.")
	return nil
}