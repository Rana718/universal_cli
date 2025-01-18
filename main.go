package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "gos",
		Short: "gos is a simple CLI tool for managing multiple tasks",
		Long: `gos is a command-line tool that helps you manage tasks efficiently.
Use 'gos --help' to see all available commands.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Use 'gos --help' to see available commands.")
		},
	}

	// Explorer command
	var exPath string
	var exCmd = &cobra.Command{
		Use:   "ex",
		Short: "Open Windows File Explorer at the specified path",
		Long:  "Use 'gos -ex --path <path>' to open File Explorer at the given path. Defaults to the current directory if no path is provided.",
		Run: func(cmd *cobra.Command, args []string) {
			path := exPath
			if path == "." {
				currentDir, err := os.Getwd()
				if err != nil {
					fmt.Println("Error getting current directory:", err)
					return
				}
				path = currentDir
			}

			err := exec.Command("explorer", path).Start()
			if err != nil {
				fmt.Println("Error opening File Explorer:", err)
			} else {
				fmt.Println("File Explorer opened at:", path)
			}
		},
	}
	exCmd.Flags().StringVarP(&exPath, "path", "p", ".", "Path to open in File Explorer")

	// VSCode command
	var coPath string
	var coCmd = &cobra.Command{
		Use:   "co",
		Short: "Open Visual Studio Code at the specified path",
		Long:  "Use 'gos -co --path <path>' to open VS Code at the given path. Defaults to the current directory if no path is provided.",
		Run: func(cmd *cobra.Command, args []string) {
			path := coPath
			if path == "." {
				currentDir, err := os.Getwd()
				if err != nil {
					fmt.Println("Error getting current directory:", err)
					return
				}
				path = currentDir
			}

			err := exec.Command("code", path).Start()
			if err != nil {
				fmt.Println("Error opening VS Code:", err)
				fmt.Println("Make sure 'code' is available in your PATH.")
			} else {
				fmt.Println("VS Code opened at:", path)
			}
		},
	}
	coCmd.Flags().StringVarP(&coPath, "path", "p", ".", "Path to open in VS Code")

	// Add the commands to the root command
	rootCmd.AddCommand(exCmd)
	rootCmd.AddCommand(coCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
