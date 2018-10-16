package cmd

import (
	"os"

	"github.com/exoscale/egoscale"
	"github.com/exoscale/egoscale/cmd/exo/table"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var affinitygroupListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List affinity group",
	Aliases: gListAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		return displayAffinitygroup()
	},
}

func displayAffinitygroup() error {
	resp, err := cs.RequestWithContext(gContext, &egoscale.ListAffinityGroups{})
	if err != nil {
		return nil
	}

	affinityGroups := resp.(*egoscale.ListAffinityGroupsResponse).AffinityGroup

	table := table.NewTable(os.Stdout)
	table.SetHeader([]string{"Name", "Description", "ID"})

	for _, affinitygroup := range affinityGroups {
		table.Append([]string{affinitygroup.Name, affinitygroup.Description, affinitygroup.ID.String()})
	}

	table.Render()

	return nil
}

func init() {
	affinitygroupCmd.AddCommand(affinitygroupListCmd)
}
