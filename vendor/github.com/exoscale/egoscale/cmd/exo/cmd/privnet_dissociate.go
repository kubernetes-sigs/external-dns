package cmd

import (
	"fmt"

	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

// dissociateCmd represents the dissociate command
var dissociateCmd = &cobra.Command{
	Use:     "dissociate <privnet name | id> <vm name | vm id> [vm name | vm id] [...]",
	Short:   "Dissociate a private network from instance(s)",
	Aliases: gDissociateAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return cmd.Usage()
		}

		network, err := getNetworkByName(args[0])
		if err != nil {
			return err
		}

		for _, vm := range args[1:] {
			resp, err := dissociatePrivNet(network, vm)
			if err != nil {
				return err
			}
			fmt.Println(resp)
		}
		return nil
	},
}

func dissociatePrivNet(privnet *egoscale.Network, vmName string) (*egoscale.UUID, error) {
	vm, err := getVMWithNameOrID(vmName)
	if err != nil {
		return nil, err
	}

	nic := vm.NicByNetworkID(*privnet.ID)
	if nic == nil {
		return nil, fmt.Errorf("no nics found for network %q", privnet.ID)
	}

	_, err = cs.RequestWithContext(gContext, &egoscale.RemoveNicFromVirtualMachine{NicID: nic.ID, VirtualMachineID: vm.ID})
	if err != nil {
		return nil, err
	}

	return nic.ID, nil
}

func init() {
	privnetCmd.AddCommand(dissociateCmd)
}
