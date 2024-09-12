/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// checkConnectionCmd represents the checkConnection command
var checkConnectionCmd = &cobra.Command{
	Use:   "check-connection",
	Short: "Check the connection to the evaluation host",
	Long:  `Check the connection to the evaluation host`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking connection to host")

		host := currentConfig.Destination.Host
		port := currentConfig.Destination.Port
		username := currentConfig.Destination.Username

		fmt.Printf("Attempting to connect to %s@%s:%d\n", username, host, port)

		// Construct the SSH command
		sshCommand := fmt.Sprintf("ssh -o StrictHostKeyChecking=yes -p %d %s@%s echo 'Connection test successful'", port, username, host)

		// Execute the SSH command
		output, err := exec.Command("sh", "-c", sshCommand).CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to connect: %v\n", err)
			fmt.Printf("Output: %s\n", output)
			return
		}

		fmt.Println("Successfully connected to the host")
		fmt.Printf("Remote output: %s\n", output)
	},
}

func init() {
	rootCmd.AddCommand(checkConnectionCmd)
}
