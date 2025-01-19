/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gos/cmd/utils"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(exCmd)
	exCmd.Flags().StringP("path", "p", "", "Specify a project name")
	rootCmd.AddCommand(vsCmd)
	vsCmd.Flags().StringP("path", "p", "", "Specify a project name")

}

// exCmd represents the ex command
var exCmd = &cobra.Command{
	Use:   "ex",
	Short: "Open the path in Explorer",
	Long:  `Open the path in Explorer`,
	Run: func(cmd *cobra.Command, args []string) {
		paths, _ := cmd.Flags().GetString("path")
		loadedPaths := utils.LoadPaths()

		var finalPath string

		if paths != "" {
			path, exists := loadedPaths[paths]

			if !exists {
				fmt.Println("Error: No path found for name ", paths)
				return
			} else {
				finalPath = path
			}

		} else if len(args) > 0 {
			path := args[0]

			if path == "." {
				currentDir, err := os.Getwd()
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}
				finalPath = currentDir
			} else {
				finalPath = path
			}
		} else {
			fmt.Println("Error: Either provide a path as an argument or use the -p flag.")
			return
		}

		err := exec.Command("explorer", finalPath).Run()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
				fmt.Println("Path opened in Explorer: ", finalPath)
				return
			}
			fmt.Println("Unexpected error: ", err)

		} else {
			fmt.Println("Path opened in explorer: ", finalPath)
		}

	},
}

var vsCmd = &cobra.Command{
	Use:   "vs",
	Short: "Open the path in Visual Studio Code",
	Long:  `Open the path in Visual Studio Code`,
	Run: func(cmd *cobra.Command, args []string) {
		paths, _ := cmd.Flags().GetString("path")
		loadedPaths := utils.LoadPaths()

		var finalPath string

		if paths != "" {
			path, exists := loadedPaths[paths]

			if !exists {
				fmt.Println("Error: No path found for name ", paths)
				return
			} else {
				finalPath = path
			}

		} else if len(args) > 0 {
			path := args[0]

			if path == "." {
				currentDir, err := os.Getwd()
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}
				finalPath = currentDir
			} else {
				finalPath = path
			}
		} else {
			fmt.Println("Error: Either provide a path as an argument or use the -p flag.")
			return
		}

		err := exec.Command("code", finalPath).Run()

		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
				fmt.Println("Path opened in Visual Studio Code: ", finalPath)
				return
			}
			fmt.Println("error: ", err)
		} else {
			fmt.Println("Path opened in Visual Studio Code: ", finalPath)
		}
	},
}
