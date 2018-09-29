package cmd

import (
	"fmt"
	"net"

	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

// eipCmd represents the eip command
var eipCmd = &cobra.Command{
	Use:   "eip",
	Short: "Elastic IPs management",
}

func getEIPIDByIP(ipAddr string) (*egoscale.UUID, error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return nil, fmt.Errorf("invalid IP address %q", ipAddr)
	}

	eips, err := cs.ListWithContext(gContext, &egoscale.IPAddress{IsElastic: true})
	if err != nil {
		return nil, err
	}

	for _, e := range eips {
		eip := e.(*egoscale.IPAddress)
		if eip.IPAddress.Equal(ip) {
			return eip.ID, nil
		}
	}

	return nil, fmt.Errorf("elastic IP %q not found", ipAddr)
}

func init() {
	RootCmd.AddCommand(eipCmd)
}
