package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Output sample YAML manifest",
	Long: `Output sample YAML manifest
	
# TODO:`,
	Run: func(cmd *cobra.Command, args []string) {
		output()
	},
}

func init() {
	rootCmd.AddCommand(outputCmd)
}

func output() {
	fmt.Println("TODO:")
}
