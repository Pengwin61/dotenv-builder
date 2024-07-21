/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"dotenv-builder/internal/config"
	"dotenv-builder/internal/creater"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var structureCmd = &cobra.Command{
	Use:   "structure",
	Short: "Create secrets and hierarchy for your projects",
	Long:  `Create secrets and hierarchy for your projects`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		languages, _ := cmd.Flags().GetString("languages")

		cfg := config.InitConfig(&path)
		creater.RunCreater(cfg, languages)
	},
}

func init() {
	CreateCmd.AddCommand(structureCmd)
	structureCmd.Flags().StringP("path", "p", "", "path to search secrets")
	structureCmd.Flags().StringP("languages", "l", "", "repo languages ")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
