package cmd

import (
	"bufio"
	"os"

	"github.com/exoscale/egoscale"
	"github.com/exoscale/egoscale/cmd/exo/table"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var privnetCreateCmd = &cobra.Command{
	Use:     "create <name>",
	Short:   "Create private network",
	Aliases: gCreateAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		desc, err := cmd.Flags().GetString("description")
		if err != nil {
			return err
		}
		zone, err := cmd.Flags().GetString("zone")
		if err != nil {
			return err
		}

		if name != "" && zone == "" {
			zone = gCurrentAccount.DefaultZone
		}

		if name == "" && zone == "" {
			reader := bufio.NewReader(os.Stdin)
			if name == "" {
				name, err = readInput(reader, "Name", "")
				if err != nil {
					return err
				}
			}
			if desc == "" {
				desc, err = readInput(reader, "Description", "")
				if err != nil {
					return err
				}
			}
			if zone == "" {
				zone, err = readInput(reader, "Zone", gCurrentAccount.DefaultZone)
				if err != nil {
					return err
				}
			}
		}

		if isEmptyArgs(name, zone) {
			return cmd.Usage()
		}

		return privnetCreate(name, desc, zone)
	},
}

func isEmptyArgs(args ...string) bool {
	for _, arg := range args {
		if arg == "" {
			return true
		}
	}
	return false
}

func privnetCreate(name, desc, zoneName string) error {
	zone, err := getZoneIDByName(zoneName)
	if err != nil {
		return err
	}

	resp, err := cs.RequestWithContext(gContext, &egoscale.ListNetworkOfferings{ZoneID: zone, Name: "PrivNet"})
	if err != nil {
		return err
	}

	s := resp.(*egoscale.ListNetworkOfferingsResponse)

	req := &egoscale.CreateNetwork{
		DisplayText: desc,
		Name:        name,
		ZoneID:      zone,
	}
	if len(s.NetworkOffering) > 0 {
		req.NetworkOfferingID = s.NetworkOffering[0].ID
	}

	resp, err = cs.RequestWithContext(gContext, req)
	if err != nil {
		return err
	}

	newNet := resp.(*egoscale.Network)

	table := table.NewTable(os.Stdout)
	table.SetHeader([]string{"Name", "Description", "ID"})
	table.Append([]string{newNet.Name, newNet.DisplayText, newNet.ID.String()})
	table.Render()
	return nil
}

func init() {
	privnetCreateCmd.Flags().StringP("name", "n", "", "Private network name")
	privnetCreateCmd.Flags().StringP("description", "d", "", "Private network description")
	privnetCreateCmd.Flags().StringP("zone", "z", "", "Assign private network to a zone")
	privnetCmd.AddCommand(privnetCreateCmd)
}
