/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	b "dotenv-builder/internal/builder"
	"dotenv-builder/internal/config"

	"github.com/spf13/cobra"
)

// envCmd represents the getenv command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Create env file for your project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		cfg := config.InitConfig(&path)

		b.RunBuilder(cfg)
	},
}

func init() {
	CreateCmd.AddCommand(envCmd)
	envCmd.Flags().StringP("path", "p", "", "path to search secrets")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// envCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// envCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
