package cmd

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

var defaultCIDR = egoscale.MustParseCIDR("0.0.0.0/0")
var defaultCIDR6 = egoscale.MustParseCIDR("::/0")

var firewallCmd = &cobra.Command{
	Use:   "firewall",
	Short: "Security groups management",
}

func init() {
	RootCmd.AddCommand(firewallCmd)
}

// Utils func for the firewall family

func formatRules(name string, rule *egoscale.IngressRule) []string {
	var source string
	if rule.CIDR != nil {
		source = fmt.Sprintf("CIDR %s", rule.CIDR)
	} else {
		ss := make([]string, len(rule.UserSecurityGroupList))
		for i, usg := range rule.UserSecurityGroupList {
			ss[i] = usg.String()
		}
		source = fmt.Sprintf("SG %s", strings.Join(ss, ","))
	}

	var ports string
	if rule.Protocol == "icmp" || rule.Protocol == "icmpv6" {
		c := icmpCode((uint16(rule.IcmpType) << 8) | uint16(rule.IcmpCode))
		t := c.icmpType()

		desc := c.StringFormatted()
		if desc == "" {
			desc = t.StringFormatted()
		}
		ports = fmt.Sprintf("%d, %d (%s)", rule.IcmpType, rule.IcmpCode, desc)
	} else if rule.StartPort == rule.EndPort {
		p := port(rule.StartPort)
		if p.StringFormatted() != "" {
			ports = fmt.Sprintf("%d (%s)", rule.StartPort, p.String())
		} else {
			ports = fmt.Sprintf("%d", rule.StartPort)
		}
	} else {
		ports = fmt.Sprintf("%d-%d", rule.StartPort, rule.EndPort)
	}

	return []string{name, source, rule.Protocol, ports, rule.Description, rule.RuleID.String()}
}

func getSecurityGroupByNameOrID(name string) (*egoscale.SecurityGroup, error) {
	sg := &egoscale.SecurityGroup{}

	id, err := egoscale.ParseUUID(name)

	if err != nil {
		sg.Name = name
	} else {
		sg.ID = id
	}

	if err := cs.GetWithContext(gContext, sg); err != nil {
		return nil, err
	}
	return sg, nil

}

func getMyCIDR(isIpv6 bool) (*egoscale.CIDR, error) {

	cidrMask := 32
	dnsServer := "resolver1.opendns.com"

	if isIpv6 {
		dnsServer = "resolver2.ipv6-sandbox.opendns.com"
		cidrMask = 128
	}

	resolver := net.Resolver{
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return net.Dial("udp", dnsServer+":53")
		},
		PreferGo: true,
	}

	ips, err := resolver.LookupIPAddr(context.Background(), "myip.opendns.com")
	if err != nil {
		return nil, err
	}

	if len(ips) < 1 {
		return nil, fmt.Errorf("no IP addresses were found using OpenDNS")
	}

	return egoscale.ParseCIDR(fmt.Sprintf("%s/%d", ips[0].IP, cidrMask))
}
