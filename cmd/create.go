package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a custom template",
	Long: `This command help you generate a custom template for future.
It will release in next version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This feature will release in next version.")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
