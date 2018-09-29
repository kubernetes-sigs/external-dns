package cmd

import (
	"errors"
	"os"

	"github.com/exoscale/egoscale"
	"github.com/exoscale/egoscale/cmd/exo/table"
	"github.com/spf13/cobra"
)

func init() {
	firewallCmd.AddCommand(firewallShow)
}

var firewallShow = &cobra.Command{
	Use:     "show <security group name | id>",
	Short:   "Show a security group rules details",
	Aliases: gListAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("show expects one security group by name or id")
		}

		t := table.NewTable(os.Stdout)
		err := firewallListRules(t, args[0])
		if err == nil {
			t.Render()
		}
		return err
	},
}

func firewallListRules(t *table.Table, name string) error {
	sg, err := getSecurityGroupByNameOrID(name)
	if err != nil {
		return err
	}

	t.SetHeader([]string{"Type", "Source", "Protocol", "Port", "Description", "ID"})

	heading := "INGRESS"
	for _, in := range sg.IngressRule {
		t.Append(formatRules(heading, &in))
		heading = ""
	}

	heading = "EGRESS"
	for _, out := range sg.EgressRule {
		t.Append(formatRules(heading, (*egoscale.IngressRule)(&out)))
		heading = ""
	}

	return nil
}
