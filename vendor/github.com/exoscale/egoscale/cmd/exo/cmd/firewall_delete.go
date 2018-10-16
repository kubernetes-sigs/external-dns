package cmd

import (
	"fmt"

	"github.com/exoscale/egoscale"

	"github.com/spf13/cobra"
)

func init() {
	firewallDeleteCmd.Flags().BoolP("force", "f", false, "Attempt to remove security group without prompting for confirmation")
	firewallCmd.AddCommand(firewallDeleteCmd)
}

// deleteCmd represents the delete command
var firewallDeleteCmd = &cobra.Command{
	Use:     "delete <security group name | id>",
	Short:   "Delete security group",
	Aliases: gDeleteAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Usage()
		}

		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			return err
		}

		for _, arg := range args {
			q := fmt.Sprintf("Are you sure you want to delete the security group: %q", arg)
			if !force && !askQuestion(q) {
				continue
			}

			output, err := firewallDelete(arg)
			if err != nil {
				return err
			}
			fmt.Println(output)
		}

		return nil
	},
}

func firewallDelete(name string) (*egoscale.UUID, error) {
	sg, err := getSecurityGroupByNameOrID(name)
	if err != nil {
		return nil, err
	}

	req := &egoscale.SecurityGroup{ID: sg.ID}
	if err := cs.Delete(req); err != nil {
		return nil, err
	}

	return sg.ID, nil
}
