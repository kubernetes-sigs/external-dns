package cmd

import (
	"fmt"
	"net"

	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

// dissociateCmd represents the dissociate command
var eipDissociateCmd = &cobra.Command{
	Use:     "dissociate <IP address> <instance name | instance id> [instance name | instance id] [...]",
	Short:   "Dissociate an IP from instance(s)",
	Aliases: gDissociateAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return cmd.Usage()
		}
		for _, arg := range args[1:] {
			if err := dissociateIP(args[0], arg); err != nil {
				return err
			}
		}
		return nil
	},
}

func dissociateIP(ipAddr, instance string) error {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return fmt.Errorf("invalid IP address %q", ipAddr)
	}

	vm, err := getVMWithNameOrID(instance)
	if err != nil {
		return err
	}

	defaultNic := vm.DefaultNic()

	if defaultNic == nil {
		return fmt.Errorf("the instance %q has no default NIC", vm.ID)
	}

	eipID, err := getSecondaryIP(defaultNic, ip)
	if err != nil {
		return err
	}

	req := &egoscale.RemoveIPFromNic{ID: eipID}

	if err := cs.BooleanRequestWithContext(gContext, req); err != nil {
		return err
	}
	println(req.ID)
	return nil
}

func getSecondaryIP(nic *egoscale.Nic, ip net.IP) (*egoscale.UUID, error) {
	for _, sIP := range nic.SecondaryIP {
		if sIP.IPAddress.Equal(ip) {
			return sIP.ID, nil
		}
	}
	return nil, fmt.Errorf("elastic IP %q not found", ip)
}

func init() {
	eipCmd.AddCommand(eipDissociateCmd)
}
