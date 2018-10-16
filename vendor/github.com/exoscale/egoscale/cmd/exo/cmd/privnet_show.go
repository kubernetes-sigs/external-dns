package cmd

import (
	"os"

	"github.com/exoscale/egoscale"
	"github.com/exoscale/egoscale/cmd/exo/table"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var privnetShowCmd = &cobra.Command{
	Use:   "show <privnet name | id>",
	Short: "Show a private network details",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Usage()
		}

		network, err := getNetworkByName(args[0])
		if err != nil {
			return err
		}

		vms, err := privnetDetails(network)
		if err != nil {
			return err
		}

		table := table.NewTable(os.Stdout)
		table.SetHeader([]string{"Zone", "Name", "Virtual Machine", "Virtual Machine ID"})
		zone := network.ZoneName
		name := network.Name
		if len(vms) > 0 {
			for _, vm := range vms {
				table.Append([]string{zone, name, vm.Name, vm.ID.String()})
				zone = ""
				name = ""
			}
		} else {
			table.Append([]string{zone, network.Name, "", ""})
		}
		table.Render()

		return nil
	},
}

func privnetDetails(network *egoscale.Network) ([]egoscale.VirtualMachine, error) {
	vms, err := cs.ListWithContext(gContext, &egoscale.VirtualMachine{
		ZoneID: network.ZoneID,
	})
	if err != nil {
		return nil, err
	}

	var vmsRes []egoscale.VirtualMachine
	for _, v := range vms {
		vm := v.(*egoscale.VirtualMachine)

		nic := vm.NicByNetworkID(*network.ID)
		if nic != nil {
			vmsRes = append(vmsRes, *vm)
		}
	}

	return vmsRes, nil
}

func init() {
	privnetCmd.AddCommand(privnetShowCmd)
}
