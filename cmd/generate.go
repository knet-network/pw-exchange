package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generatee a link to exchange a password",
	Run:   generate,
}

func generate(cmd *cobra.Command, args []string) {
}
