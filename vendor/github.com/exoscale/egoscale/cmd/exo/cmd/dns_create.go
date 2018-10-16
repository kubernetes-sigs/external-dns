package cmd

import (
	"fmt"

	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

// dnsCreateCmd represents the create command
var dnsCreateCmd = &cobra.Command{
	Use:     "create <domain name>",
	Short:   "Create a domain",
	Aliases: gCreateAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Usage()
		}
		resp, err := createDomain(args[0])
		if err != nil {
			return err
		}

		fmt.Printf("Domain %q was created successfully\n", resp.Name)
		return nil
	},
}

func createDomain(domainName string) (*egoscale.DNSDomain, error) {
	return csDNS.CreateDomain(domainName)
}

func init() {
	dnsCmd.AddCommand(dnsCreateCmd)
}
