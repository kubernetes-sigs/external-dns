package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var configListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List available accounts",
	Aliases: gListAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		if gAllAccount == nil {
			return fmt.Errorf("no accounts defined")
		}
		listAccounts()
		return nil
	},
}

func init() {
	configCmd.AddCommand(configListCmd)
}
