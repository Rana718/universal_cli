package cmd

import (
	"fmt"
	"gos/cmd/helper"
	"gos/cmd/utils"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(exCmd)
	exCmd.Flags().StringP("path", "p", "", "Specify a project name")
	rootCmd.AddCommand(vsCmd)
	vsCmd.Flags().StringP("path", "p", "", "Specify a project name")

}

var exCmd = &cobra.Command{
	Use:   "ex",
	Short: "Open the path in Explorer",
	Long:  `Open the path in Explorer`,
	Run: func(cmd *cobra.Command, args []string) {
		paths, _ := cmd.Flags().GetString("path")
		loadedPaths := utils.LoadPaths()

		finalPath, err := helper.ResolvePath(paths, args, loadedPaths)

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		err = exec.Command("explorer", finalPath).Run()
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

		finalPath, err := helper.ResolvePath(paths, args, loadedPaths)

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		err = exec.Command("code", finalPath).Run()

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
