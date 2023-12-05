/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/715047274/jenkinsCmd/internal"

	"github.com/spf13/cobra"
)

// buildResultCmd represents the buildResult command
var buildResultCmd = &cobra.Command{
	Use:   "buildResult",
	Short: "Generate HTML email from Cypress result",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("buildResult called")
		inputFile, _ := cmd.Flags().GetString("input")
		outputDir, _ := cmd.Flags().GetString("output")

		// Call your internal function to generate HTML from Cypress test results
		internal.GenerateHTML(inputFile, outputDir)
	},
}

func init() {
	// Add flags to the generate command
	buildResultCmd.Flags().StringP("input", "i", "", "Input file path (Cypress test results in JSON format)")
	buildResultCmd.Flags().StringP("output", "o", "", "Output directory path for generated HTML email")

	rootCmd.AddCommand(buildResultCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildResultCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildResultCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
