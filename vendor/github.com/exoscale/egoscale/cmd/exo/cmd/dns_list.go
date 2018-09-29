package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/exoscale/egoscale/cmd/exo/table"
	"github.com/spf13/cobra"
)

// dnsListCmd represents the list command
var dnsListCmd = &cobra.Command{
	Use:     "list [domain name | id]*",
	Short:   "List domains",
	Aliases: gListAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		t := table.NewTable(os.Stdout)
		err := listDomains(t, args)
		if err == nil {
			t.Render()
		}
		return err
	},
}

func listDomains(t *table.Table, filters []string) error {
	domains, err := csDNS.GetDomains()
	if err != nil {
		return err
	}

	data := make([][]string, 0)
	for _, d := range domains {
		keep := true
		if len(filters) > 0 {
			keep = false
			s := strings.ToLower(
				fmt.Sprintf("%d#%s#%s", d.ID, d.Name, d.UnicodeName))

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

		unicodeName := ""
		if d.Name != d.UnicodeName {
			unicodeName = d.UnicodeName
		}
		data = append(data, []string{
			d.Name, unicodeName, strconv.FormatInt(d.ID, 10),
		})
	}

	headers := []string{"Name", "Unicode", "ID"}
	if len(data) > 0 {
		t.SetHeader(headers)
	}

	t.AppendBulk(data)
	if len(data) > 10 {
		t.SetFooter(headers)
	}

	return nil
}

func init() {
	dnsCmd.AddCommand(dnsListCmd)
}
