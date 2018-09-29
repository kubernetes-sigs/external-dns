package cmd

import (
	"fmt"

	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var affinitygroupDeleteCmd = &cobra.Command{
	Use:     "delete <name | id>",
	Short:   "Delete affinity group",
	Aliases: gDeleteAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Usage()
		}

		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			return err
		}

		if !force {
			if !askQuestion(fmt.Sprintf("sure you want to delete %q affinity group", args[0])) {
				return nil
			}
		}

		return deleteAffinityGroup(args[0])
	},
}

func deleteAffinityGroup(name string) error {
	aff, err := getAffinityGroupByName(name)
	if err != nil {
		return err
	}

	_, err = cs.Request(&egoscale.DeleteAffinityGroup{ID: aff.ID})
	if err != nil {
		return err
	}

	fmt.Println(aff.ID)

	return nil
}

func init() {
	affinitygroupDeleteCmd.Flags().BoolP("force", "f", false, "Attempt to remove affinity group without prompting for confirmation")
	affinitygroupCmd.AddCommand(affinitygroupDeleteCmd)
}
