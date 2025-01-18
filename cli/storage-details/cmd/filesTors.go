/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// filesTorsCmd represents the filesTors command
var filesTorsCmd = &cobra.Command{
	Use:   "filesTors",
	Short: "Lists all files and directories in the current directory",
	Long:  `Lists all files and directories in the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("filesTors called")
		var output = ExecuteShellCommand("ls -la")
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(filesTorsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// filesTorsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// filesTorsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
