package main

import (
	"github.com/spf13/cobra"
)

// loadCmd represents the load command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	SilenceUsage: false,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}
