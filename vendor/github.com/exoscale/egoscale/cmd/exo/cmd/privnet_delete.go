package cmd

import (
	"fmt"

	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var privnetDeleteCmd = &cobra.Command{
	Use:     "delete <name | id>",
	Short:   "Delete private network",
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
			if !askQuestion(fmt.Sprintf("sure you want to delete %q private network", args[0])) {
				return nil
			}
		}

		return deletePrivnet(args[0], force)
	},
}

func deletePrivnet(name string, force bool) error {
	addrReq := &egoscale.DeleteNetwork{}
	var err error
	network, err := getNetworkByName(name)
	if err != nil {
		return err
	}
	addrReq.ID = network.ID
	addrReq.Forced = &force
	if err := cs.BooleanRequestWithContext(gContext, addrReq); err != nil {
		return err
	}
	println(addrReq.ID)
	return nil
}

func init() {
	privnetDeleteCmd.Flags().BoolP("force", "f", false, "Attempt to remove private network without prompting for confirmation")
	privnetCmd.AddCommand(privnetDeleteCmd)
}
