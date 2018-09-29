package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var configDeleteCmd = &cobra.Command{
	Use:     "delete <account name>",
	Short:   "Delete an account from config file",
	Aliases: gDeleteAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Usage()
		}
		if gAllAccount == nil {
			return fmt.Errorf("no accounts defined")
		}
		if !isAccountExist(args[0]) {
			return fmt.Errorf("account %q doesn't exist", args[0])
		}

		if args[0] == gAllAccount.DefaultAccount {
			return fmt.Errorf("cannot delete the default account")
		}

		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			return err
		}

		if !force {
			if !askQuestion(fmt.Sprintf("Do you really want to delete the config for %q?", args[0])) {
				return nil
			}
		}

		pos := 0
		for i, acc := range gAllAccount.Accounts {
			if acc.Name == args[0] {
				pos = i
				break
			}
		}

		gAllAccount.Accounts = append(gAllAccount.Accounts[:pos], gAllAccount.Accounts[pos+1:]...)

		if err := addAccount(viper.ConfigFileUsed(), nil); err != nil {
			return err
		}

		println(args[0])
		return nil
	},
}

func init() {
	configDeleteCmd.Flags().BoolP("force", "f", false, "Attempt to remove an account without prompting for confirmation")
	configCmd.AddCommand(configDeleteCmd)
}
