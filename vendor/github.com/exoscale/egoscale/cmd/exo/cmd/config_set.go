package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setCmd represents the set command
var configSetCmd = &cobra.Command{
	Use:   "set <account name>",
	Short: "Set an account as default",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Usage()
		}
		if gAllAccount == nil {
			return fmt.Errorf("no accounts are defined")
		}

		if !isAccountExist(args[0]) {
			return fmt.Errorf("account %q does not exist", args[0])
		}

		viper.Set("defaultAccount", args[0])

		if err := addAccount(viper.ConfigFileUsed(), nil); err != nil {
			return err
		}

		println("Default profile set to", args[0])

		return nil
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
}
