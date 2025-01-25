/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gos/cmd/helper"
	"gos/cmd/shell"
	"gos/cmd/utils"

	"github.com/spf13/cobra"
)


var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Change the current directory",
	Long:  `Change the current directory to the specified directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		loadedPaths := utils.LoadPaths()

		if len(args) == 0 {
			fmt.Println("Error: Please provide a path name")
			return
		}

		finalPath, err := helper.ResolvePath(args[0], nil, loadedPaths)

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		err = shell.ChangeDirectory(finalPath)

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(cdCmd)
}
