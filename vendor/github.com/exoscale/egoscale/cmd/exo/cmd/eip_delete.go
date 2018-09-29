package cmd

import (
	"fmt"
	"net"

	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var eipDeleteCmd = &cobra.Command{
	Use:     "delete <ip | eip id>",
	Short:   "Delete EIP",
	Aliases: gDeleteAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Usage()
		}

		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			return err
		}

		if !force {
			if !askQuestion(fmt.Sprintf("sure you want to delete %q EIP", args[0])) {
				return nil
			}
		}

		return deleteEip(args[0])
	},
}

func deleteEip(ip string) error {
	addrReq := &egoscale.DisassociateIPAddress{}

	ipAddr := net.ParseIP(ip)

	if ipAddr == nil {
		id, err := egoscale.ParseUUID(ip)
		if err != nil {
			return fmt.Errorf("delete the eip by ID or IP address, gotb %q", ip)
		}
		addrReq.ID = id
	} else {
		req := &egoscale.IPAddress{IPAddress: ipAddr, IsElastic: true}
		if err := cs.GetWithContext(gContext, req); err != nil {
			return err
		}
		addrReq.ID = req.ID
	}

	if err := cs.BooleanRequestWithContext(gContext, addrReq); err != nil {
		return err
	}
	println(addrReq.ID)
	return nil
}

func init() {
	eipDeleteCmd.Flags().BoolP("force", "f", false, "Attempt to remove EIP without prompting for confirmation")
	eipCmd.AddCommand(eipDeleteCmd)
}
