package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dnsRemoveCmd represents the remove command
var dnsRemoveCmd = &cobra.Command{
	Use:     "remove <domain name> <record name | id>",
	Short:   "Remove a domain record",
	Aliases: gRemoveAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return cmd.Usage()
		}
		id, err := removeRecord(args[0], args[1])
		if err != nil {
			return err
		}
		fmt.Println(id)
		return nil
	},
}

func removeRecord(domainName, record string) (int64, error) {
	id, err := getRecordIDByName(domainName, record)
	if err != nil {
		return 0, err
	}
	if err := csDNS.DeleteRecord(domainName, id); err != nil {
		return 0, err
	}

	return id, nil
}

func init() {
	dnsCmd.AddCommand(dnsRemoveCmd)
}
