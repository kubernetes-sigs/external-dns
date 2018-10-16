package cmd

import (
	"fmt"

	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

// privnetAssociateCmd represents the associate command
var privnetAssociateCmd = &cobra.Command{
	Use:     "associate <privnet name | id> <vm name | vm id> [vm name | vm id] [...]",
	Short:   "Associate a private network to instance(s)",
	Aliases: gAssociateAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return cmd.Usage()
		}

		network, err := getNetworkByName(args[0])
		if err != nil {
			return err
		}

		for _, vm := range args[1:] {
			resp, err := associatePrivNet(network, vm)
			if err != nil {
				return err
			}
			fmt.Println(resp)
		}

		return nil
	},
}

func associatePrivNet(privnet *egoscale.Network, vmName string) (*egoscale.UUID, error) {
	vm, err := getVMWithNameOrID(vmName)
	if err != nil {
		return nil, err
	}

	req := &egoscale.AddNicToVirtualMachine{NetworkID: privnet.ID, VirtualMachineID: vm.ID}
	resp, err := cs.RequestWithContext(gContext, req)
	if err != nil {
		return nil, err
	}

	nic := resp.(*egoscale.VirtualMachine).NicByNetworkID(*privnet.ID)
	if nic == nil {
		return nil, fmt.Errorf("no nics found for network %q", privnet.ID)
	}

	return nic.ID, nil

}

func init() {
	privnetCmd.AddCommand(privnetAssociateCmd)
}
