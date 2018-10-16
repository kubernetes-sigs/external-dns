package cmd

import (
	"fmt"
	"os"

	"github.com/exoscale/egoscale"
	"github.com/exoscale/egoscale/cmd/exo/table"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var privnetListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List private networks",
	Aliases: gListAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		zone, err := cmd.Flags().GetString("zone")
		if err != nil {
			return err
		}
		table := table.NewTable(os.Stdout)
		table.SetHeader([]string{"zone", "Name", "ID", "Associated Virtual machine"})
		if err := listPrivnets(zone, table); err != nil {
			return err
		}
		table.Render()
		return nil
	},
}

func listPrivnets(zone string, table *table.Table) error {
	pnReq := &egoscale.Network{}

	if zone != "" {
		var err error
		pnReq.Type = "Isolated"
		pnReq.ZoneID, err = getZoneIDByName(zone)
		if err != nil {
			return err
		}
		pnReq.CanUseForDeploy = true
		pns, err := cs.ListWithContext(gContext, pnReq)
		if err != nil {
			return err
		}

		var zone string
		for i, pNet := range pns {
			pn := pNet.(*egoscale.Network)
			if i == 0 {
				zone = pn.ZoneName
			}

			vms, err := privnetDetails(pn)
			if err != nil {
				return err
			}

			vmNum := fmt.Sprintf("%d", len(vms))

			table.Append([]string{zone, pn.Name, pn.ID.String(), vmNum})

			zone = ""
		}
		return nil
	}

	zones := &egoscale.Zone{}
	zs, err := cs.ListWithContext(gContext, zones)
	if err != nil {
		return err
	}

	for _, z := range zs {
		zID := z.(*egoscale.Zone).Name
		if err := listPrivnets(zID, table); err != nil {
			return err
		}
	}
	return nil
}

func init() {
	privnetListCmd.Flags().StringP("zone", "z", "", "Show Private Network from given zone")
	privnetCmd.AddCommand(privnetListCmd)
}
