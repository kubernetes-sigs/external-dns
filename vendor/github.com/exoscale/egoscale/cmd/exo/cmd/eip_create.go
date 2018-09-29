package cmd

import (
	"os"

	"github.com/exoscale/egoscale"
	"github.com/exoscale/egoscale/cmd/exo/table"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var eipCreateCmd = &cobra.Command{
	Use:     "create [zone name | zone id]",
	Short:   "Create EIP",
	Aliases: gCreateAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		zone := gCurrentAccount.DefaultZone
		if len(args) >= 1 {
			zone = args[0]
		}
		return associateIPAddress(zone)
	},
}

func associateIPAddress(name string) error {
	ipReq := egoscale.AssociateIPAddress{}
	var err error
	ipReq.ZoneID, err = getZoneIDByName(name)
	if err != nil {
		return err
	}

	resp, err := cs.RequestWithContext(gContext, &ipReq)
	if err != nil {
		return err
	}

	ipResp := resp.(*egoscale.IPAddress)

	table := table.NewTable(os.Stdout)
	table.SetHeader([]string{"Zone", "IP", "ID"})

	table.Append([]string{ipResp.ZoneName, ipResp.IPAddress.String(), ipResp.ID.String()})

	table.Render()
	return nil
}

func init() {
	eipCmd.AddCommand(eipCreateCmd)
}
