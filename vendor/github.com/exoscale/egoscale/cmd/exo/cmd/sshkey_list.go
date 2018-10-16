package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/exoscale/egoscale"
	"github.com/exoscale/egoscale/cmd/exo/table"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List SSH key pairs",
	Aliases: gListAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		t := table.NewTable(os.Stdout)
		err := listSSHKey(t, args)
		if err == nil {
			t.Render()
		}

		return err
	},
}

func listSSHKey(t *table.Table, filters []string) error {
	sshKey := &egoscale.SSHKeyPair{}
	sshKeys, err := cs.ListWithContext(gContext, sshKey)
	if err != nil {
		return err
	}

	data := make([][]string, 0)

	for _, key := range sshKeys {
		k := key.(*egoscale.SSHKeyPair)

		keep := true
		if len(filters) > 0 {
			keep = false
			s := strings.ToLower(fmt.Sprintf("%s§%s", k.Name, k.Fingerprint))

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

		data = append(data, []string{k.Name, k.Fingerprint})

	}

	headers := []string{"Name", "Fingerprint"}
	if len(data) > 0 {
		t.SetHeader(headers)
	}
	if len(data) > 10 {
		t.SetFooter(headers)
	}

	t.AppendBulk(data)

	return nil
}

func init() {
	sshkeyCmd.AddCommand(listCmd)
}
