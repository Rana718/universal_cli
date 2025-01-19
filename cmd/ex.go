/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(exCmd)
	rootCmd.AddCommand(vsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// exCmd represents the ex command
var exCmd = &cobra.Command{
	Use:   "ex",
	Short: "Open the path in Explorer",
	Long: `Open the path in Explorer`,
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		var finalPath string
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
	Long: `Open the path in Visual Studio Code`,
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		var finalPath string

		if path == "." {
			currentDir, err := os.Getwd()
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			finalPath = currentDir
		}else{
			finalPath = path
		}
		err := exec.Command("code", finalPath).Run()

		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
				fmt.Println("Path opened in Visual Studio Code: ", finalPath)
				return
			}
			fmt.Println("error: ", err)
		}else{
			fmt.Println("Path opened in Visual Studio Code: ", finalPath)
		}
	},
}


