package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:     "show <account name>",
	Short:   "Show an account details",
	Aliases: gShowAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		if gAllAccount == nil {
			return fmt.Errorf("no accounts are defined")
		}

		account := gCurrentAccount.Name
		if len(args) > 0 {
			account = args[0]
		}

		if !isAccountExist(account) {
			return fmt.Errorf("account %q does not exist", account)
		}

		acc := getAccountByName(account)
		if acc == nil {
			return fmt.Errorf("account %q was not found", account)
		}

		secret := strings.Repeat("×", len(acc.Secret)/4)

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
		fmt.Fprintf(w, "Name:\t%s\n", acc.Name)                // nolint: errcheck
		fmt.Fprintf(w, "API Key:\t%s\n", acc.Key)              // nolint: errcheck
		fmt.Fprintf(w, "API Secret:\t%s\n", secret)            // nolint: errcheck
		fmt.Fprintf(w, "Account:\t%s\n", acc.Account)          // nolint: errcheck
		fmt.Fprintf(w, "Default zone:\t%s\n", acc.DefaultZone) // nolint: errcheck
		if acc.DefaultTemplate != "" {
			println("Default template:", acc.DefaultTemplate) // nolint: errcheck
		}
		if acc.Endpoint != defaultEndpoint {
			fmt.Fprintf(w, "Endpoint:\t%s\n", acc.Endpoint)        // nolint: errcheck
			fmt.Fprintf(w, "DNS Endpoint:\t%s\n", acc.DNSEndpoint) // nolint: errcheck
		}
		return w.Flush()
	},
}

func init() {
	configCmd.AddCommand(showCmd)
}
