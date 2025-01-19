/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gos/cmd/utils"
	
)

// addPathCmd represents the addPath command
var addPathCmd = &cobra.Command{
	Use:   "addPath [name] [path]",
	Short: "Add a path to the JSON file with a custom name",
	Long:  `Add a user-defined name and associated path to a JSON file for easy retrieval later.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		path := args[1]

		paths := utils.LoadPaths()
		paths[name] = path

		if err := utils.SavedPaths(paths); err != nil{
			fmt.Println("Error :", err)
			return
		}
		fmt.Printf("Path '%s' added with name '%s'.\n", path, name)
	},
}

var pathName string
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the path associated with the given name",
	Long:  `Retrieve the path associated with a given name stored in the JSON file.`,
	Run: func(cmd *cobra.Command, args []string) {
		paths := utils.LoadPaths()
		path, exists := paths[pathName]
		if !exists {
			fmt.Printf("No path found for name '%s'.\n", pathName)
			return
		}
		fmt.Printf("Path for '%s': %s\n", pathName, path)
	},
}

func init() {
	rootCmd.AddCommand(addPathCmd)
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&pathName, "path", "p", "", "Name of the stored path")
	getCmd.MarkFlagRequired("path")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addPathCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addPathCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
