package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/exoscale/egoscale"
	"github.com/exoscale/egoscale/cmd/exo/table"
	"github.com/spf13/cobra"
)

func init() {
	firewallCmd.AddCommand(&cobra.Command{
		Use:     "list [filter by name | id]*",
		Short:   "List security groups",
		Aliases: gListAlias,
		RunE: func(cmd *cobra.Command, args []string) error {
			t := table.NewTable(os.Stdout)
			err := firewallListSecurityGroups(t, args)
			if err == nil {
				t.Render()
			}
			return err
		},
	})
}

func firewallListSecurityGroups(t *table.Table, filters []string) error {
	sgs, err := cs.ListWithContext(gContext, &egoscale.SecurityGroup{})
	if err != nil {
		return err
	}

	data := make([][]string, 0)

	for _, s := range sgs {
		sg := s.(*egoscale.SecurityGroup)

		keep := true
		if len(filters) > 0 {
			keep = false
			s := strings.ToLower(
				fmt.Sprintf("%s#%s#%s", sg.ID, sg.Name, sg.Description))

			for _, filter := range filters {
				substr := strings.ToLower(filter)
				if strings.Contains(s, substr) {
					keep = true
					break
				}
			}
		}

		if !keep {
			continue
		}

		data = append(data, []string{
			sg.Name,
			sg.Description,
			sg.ID.String(),
		})
	}

	headers := []string{"Name", "Description", "ID"}
	if len(data) > 0 {
		t.SetHeader(headers)
	}
	t.AppendBulk(data)
	if len(data) > 10 {
		t.SetFooter(headers)
	}

	return nil
}
