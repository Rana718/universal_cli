/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gos/cmd/utils"
	"sort"

	"github.com/spf13/cobra"
)


var addPathCmd = &cobra.Command{
	Use:   "addPath [name] [path]",
	Short: "Add a path to the JSON file with a custom name",
	Long:  `Add a user-defined name and associated path to a JSON file for easy retrieval later.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		path := args[1]

		paths := utils.LoadPaths()
		paths[name] = path

		if err := utils.SavedPaths(paths); err != nil {
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
		
		if pathName == "" {
			names := make([]string, 0, len(paths))
			for name := range paths {
				names = append(names, name)
			}
			sort.Strings(names)
			fmt.Println("Stored paths:")
			for _, name := range names {
				fmt.Printf("  %s: %s\n", name, paths[name])
			}
			return
		}
		
		path, exists := paths[pathName]
		if !exists {
			fmt.Printf("No path found for name '%s'.\n", pathName)
			return
		}
		fmt.Printf("Path for '%s': %s\n", pathName, path)
	},
}

var remCmd = &cobra.Command{
	Use:   "rem",
	Short: "Remove the path associated with the given name",
	Long:  `Remove the path associated with a given name stored in the JSON file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 0 {
			fmt.Println("Please provide a name to remove")
			return
		}
		name := args[0]
		paths := utils.LoadPaths()
		_, exists := paths[name]
		if !exists {
			fmt.Printf("No path found for name '%s'.\n", name)
			return
		}

		delete(paths, name)

		if err := utils.SavedPaths(paths); err != nil {
			fmt.Println("Error :", err)
			return
		}

		fmt.Printf("Path '%s' removed.\n", name)
	},
}

func init() {
	rootCmd.AddCommand(addPathCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(remCmd)
	getCmd.Flags().StringVarP(&pathName, "path", "p", "", "Name of the stored path")
}
