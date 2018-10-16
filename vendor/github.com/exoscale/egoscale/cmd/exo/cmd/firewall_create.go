package cmd

import (
	"os"

	"github.com/exoscale/egoscale"
	"github.com/exoscale/egoscale/cmd/exo/table"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var firewallCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create security group",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Usage()
		}
		desc, err := cmd.Flags().GetString("description")
		if err != nil {
			return err
		}
		return firewallCreate(args[0], desc)
	},
}

func firewallCreate(name, desc string) error {
	req := &egoscale.CreateSecurityGroup{Name: name}

	if desc != "" {
		req.Description = desc
	}

	resp, err := cs.RequestWithContext(gContext, req)
	if err != nil {
		return err
	}

	sgResp := resp.(*egoscale.SecurityGroup)

	table := table.NewTable(os.Stdout)
	if desc == "" {
		table.SetHeader([]string{"Name", "ID"})
		table.Append([]string{sgResp.Name, sgResp.ID.String()})
	} else {
		table.SetHeader([]string{"Name", "Description", "ID"})
		table.Append([]string{sgResp.Name, sgResp.Description, sgResp.ID.String()})
	}
	table.Render()
	return nil
}

func init() {
	firewallCreateCmd.Flags().StringP("description", "d", "", "Security group description")
	firewallCmd.AddCommand(firewallCreateCmd)
}
