package cmd

import (
	"os"

	"github.com/exoscale/egoscale"
	"github.com/exoscale/egoscale/cmd/exo/table"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var eipShowCmd = &cobra.Command{
	Use:     "show <ip address | eip id>",
	Short:   "Show an eip details",
	Aliases: gShowAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Usage()
		}

		id, err := egoscale.ParseUUID(args[0])
		if err != nil {
			id, err = getEIPIDByIP(args[0])
			if err != nil {
				return err
			}
		}

		ip, vms, err := eipDetails(id)
		if err != nil {
			return err
		}

		table := table.NewTable(os.Stdout)
		table.SetHeader([]string{"Zone", "IP", "Virtual Machine", "Virtual Machine ID"})

		zone := ip.ZoneName
		ipaddr := ip.IPAddress.String()
		if len(vms) > 0 {
			for _, vm := range vms {
				table.Append([]string{zone, ipaddr, vm.Name, vm.ID.String()})
				zone = ""
				ipaddr = ""
			}
		} else {
			table.Append([]string{zone, ipaddr})
		}
		table.Render()
		return nil
	},
}

func eipDetails(eip *egoscale.UUID) (*egoscale.IPAddress, []egoscale.VirtualMachine, error) {

	var eipID = eip

	addr := &egoscale.IPAddress{ID: eipID, IsElastic: true}
	if err := cs.GetWithContext(gContext, addr); err != nil {
		return nil, nil, err
	}

	vms, err := cs.ListWithContext(gContext, &egoscale.VirtualMachine{ZoneID: addr.ZoneID})
	if err != nil {
		return nil, nil, err
	}

	vmAssociated := []egoscale.VirtualMachine{}

	for _, value := range vms {
		vm := value.(*egoscale.VirtualMachine)
		nic := vm.DefaultNic()
		if nic == nil {
			continue
		}
		for _, sIP := range nic.SecondaryIP {
			if sIP.IPAddress.Equal(addr.IPAddress) {
				vmAssociated = append(vmAssociated, *vm)
			}
		}
	}

	return addr, vmAssociated, nil
}

func init() {
	eipCmd.AddCommand(eipShowCmd)
}
