/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	b "dotenv-builder/internal/builder"
	"dotenv-builder/internal/config"

	"github.com/spf13/cobra"
)

// getenvCmd represents the getenv command
var getenvCmd = &cobra.Command{
	Use:   "getenv",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		cfg := config.InitConfig(&path)

		b.RunBuilder(cfg)
	},
}

func init() {
	rootCmd.AddCommand(getenvCmd)
	getenvCmd.Flags().StringP("path", "p", "", "path to search secrets")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getenvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getenvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
