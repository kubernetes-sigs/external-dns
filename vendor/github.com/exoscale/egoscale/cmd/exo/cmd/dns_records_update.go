package cmd

import (
	"fmt"

	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

func init() {

	for i := egoscale.A; i <= egoscale.URL; i++ {

		var cmdUpdateRecord = &cobra.Command{
			Use:   fmt.Sprintf("%s <domain name> <record name | id>", egoscale.Record.String(i)),
			Short: fmt.Sprintf("Update %s record type to a domain", egoscale.Record.String(i)),
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) < 2 {
					return cmd.Usage()
				}

				recordID, err := getRecordIDByName(args[0], args[1])
				if err != nil {
					return err
				}

				name, err := cmd.Flags().GetString("name")
				if err != nil {
					return err
				}
				addr, err := cmd.Flags().GetString("content")
				if err != nil {
					return err
				}
				ttl, err := cmd.Flags().GetInt("ttl")
				if err != nil {
					return err
				}

				domain, err := csDNS.GetDomain(args[0])
				if err != nil {
					return err
				}

				_, err = csDNS.UpdateRecord(args[0], egoscale.UpdateDNSRecord{
					ID:         recordID,
					DomainID:   domain.ID,
					TTL:        ttl,
					RecordType: egoscale.Record.String(i),
					Name:       name,
					Content:    addr,
				})
				if err != nil {
					return err
				}
				fmt.Printf("Record %q was updated successfully to %q\n", cmd.Name(), args[0])
				return nil
			},
		}

		cmdUpdateRecord.Flags().StringP("name", "n", "", "Update name")
		cmdUpdateRecord.Flags().StringP("content", "c", "", "Update Content")
		cmdUpdateRecord.Flags().IntP("ttl", "t", 0, "Update ttl")
		cmdUpdateRecord.Flags().IntP("priority", "p", 0, "Update priority")

		dnsUpdateCmd.AddCommand(cmdUpdateRecord)
	}
}
