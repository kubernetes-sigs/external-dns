package cmd

import (
	"os"
	"strconv"
	"strings"

	"github.com/exoscale/egoscale/cmd/exo/table"
	"github.com/spf13/cobra"
)

func init() {
	templateCmd.AddCommand(templateListCmd)
}

// templateListCmd represents the list command
var templateListCmd = &cobra.Command{
	Use:     "list [keyword]",
	Short:   "List all available templates",
	Aliases: gListAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		t := table.NewTable(os.Stdout)
		err := listTemplates(t, args)
		if err == nil {
			t.Render()
		}
		return err
	},
}

func listTemplates(t *table.Table, filters []string) error {
	zoneID, err := getZoneIDByName(gCurrentAccount.DefaultZone)
	if err != nil {
		return err
	}

	templates, err := findTemplates(zoneID, filters...)
	if err != nil {
		return err
	}

	t.SetHeader([]string{"Operating System", "Disk", "Release Date", "ID"})
	for _, template := range templates {
		sz := strconv.FormatInt(template.Size>>30, 10)
		if sz == "10" && strings.HasPrefix(template.Name, "Linux") {
			sz = ""
		}
		t.Append([]string{template.Name, sz, template.Created, template.ID.String()})
	}

	return nil
}
