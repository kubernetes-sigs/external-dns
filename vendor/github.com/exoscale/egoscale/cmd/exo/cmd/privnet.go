package cmd

import (
	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

// privnetCmd represents the pn command
var privnetCmd = &cobra.Command{
	Use:   "privnet",
	Short: "Private networks management",
}

func getNetworkByName(name string) (*egoscale.Network, error) {
	net := &egoscale.Network{
		Type:            "Isolated",
		CanUseForDeploy: true,
	}

	id, errUUID := egoscale.ParseUUID(name)
	if errUUID != nil {
		net.Name = name
	} else {
		net.ID = id
	}

	if err := cs.GetWithContext(gContext, net); err != nil {
		return nil, err
	}

	return net, nil
}

func init() {
	RootCmd.AddCommand(privnetCmd)
}
