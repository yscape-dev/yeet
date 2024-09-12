/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"yeet/internal/config"
)

// Constants
var defaultConfigFile = "config/yeet.yml"

// Variables
var configFile string
var currentConfig *config.Config
var workingDirectory string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yeet",
	Short: "yeet - cli tool for evaluating performance",
	Long:  `yeet is a cli tool for better performance evaluation. It is a simple tool that allows you to setup a stable environment for performance evaluation, sync your code to that environment, and then run your benchmarks or profiling tools.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		currentConfig, err = config.ParseConfig(workingDirectory + "/" + configFile)
		if err != nil {
			return fmt.Errorf("failed to parse config file: %w", err)
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", defaultConfigFile, "config file (default is config/yeet.yaml)")
	rootCmd.PersistentFlags().StringVarP(&workingDirectory, "working-directory", "w", ".", "working directory (default is current directory)")

	cobra.OnInitialize(func() {
		// This function is now empty as we've moved config parsing to PersistentPreRunE
	})
}
