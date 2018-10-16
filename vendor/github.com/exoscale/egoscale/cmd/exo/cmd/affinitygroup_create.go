package cmd

import (
	"os"

	"github.com/exoscale/egoscale"
	"github.com/exoscale/egoscale/cmd/exo/table"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var affinitygroupCreateCmd = &cobra.Command{
	Use:     "create <name>",
	Short:   "Create affinity group",
	Aliases: gCreateAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Usage()
		}

		desc, err := cmd.Flags().GetString("description")
		if err != nil {
			return err
		}

		return createAffinityGroup(args[0], desc)
	},
}

func createAffinityGroup(name, desc string) error {
	resp, err := cs.RequestWithContext(gContext, &egoscale.CreateAffinityGroup{Name: name, Description: desc, Type: "host anti-affinity"})
	if err != nil {
		return err
	}

	affinityGroup := resp.(*egoscale.AffinityGroup)

	table := table.NewTable(os.Stdout)
	table.SetHeader([]string{"Name", "Description", "ID"})
	table.Append([]string{affinityGroup.Name, affinityGroup.Description, affinityGroup.ID.String()})
	table.Render()
	return nil
}

func init() {
	affinitygroupCreateCmd.Flags().StringP("description", "d", "", "affinity group description")
	affinitygroupCmd.AddCommand(affinitygroupCreateCmd)
}
