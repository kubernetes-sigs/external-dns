package cmd

import (
	"fmt"
	"os"

	"github.com/exoscale/egoscale"
	"github.com/exoscale/egoscale/cmd/exo/table"

	"github.com/spf13/cobra"
)

// zoneCmd represents the zone command
var zoneCmd = &cobra.Command{
	Use:   "zone",
	Short: "List all available zones",
	RunE: func(cmd *cobra.Command, args []string) error {
		return listZones()
	},
}

func listZones() error {
	zones, err := cs.ListWithContext(gContext, &egoscale.Zone{})
	if err != nil {
		return err
	}

	table := table.NewTable(os.Stdout)
	table.SetHeader([]string{"Name", "ID"})

	for _, zone := range zones {
		z := zone.(*egoscale.Zone)
		table.Append([]string{z.Name, z.ID.String()})
	}
	table.Render()
	return nil
}

func getZoneIDByName(name string) (*egoscale.UUID, error) {
	zone := egoscale.Zone{}

	id, err := egoscale.ParseUUID(name)
	if err != nil {
		zone.Name = name
	} else {
		zone.ID = id
	}

	if err := cs.GetWithContext(gContext, &zone); err != nil {
		if e, ok := err.(*egoscale.ErrorResponse); ok && e.ErrorCode == egoscale.ParamError {
			return nil, fmt.Errorf("missing Zone %q", name)
		}

		return nil, err
	}

	return zone.ID, nil
}

func init() {
	RootCmd.AddCommand(zoneCmd)
}
