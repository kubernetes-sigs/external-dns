package cmd

import (
	"fmt"

	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

// affinitygroupCmd represents the affinitygroup command
var affinitygroupCmd = &cobra.Command{
	Use:   "affinitygroup",
	Short: "Affinity groups management",
}

func getAffinityGroupByName(name string) (*egoscale.AffinityGroup, error) {
	aff := &egoscale.AffinityGroup{}

	id, err := egoscale.ParseUUID(name)
	if err != nil {
		aff.ID = id
	} else {
		aff.Name = name
	}

	if err := cs.GetWithContext(gContext, aff); err != nil {
		if e, ok := err.(*egoscale.ErrorResponse); ok && e.ErrorCode == egoscale.ParamError {
			return nil, fmt.Errorf("missing Affinity Group %q", name)
		}

		return nil, err
	}

	return aff, nil
}

func init() {
	RootCmd.AddCommand(affinitygroupCmd)
}
