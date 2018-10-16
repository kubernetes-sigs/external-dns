package cmd

import (
	"fmt"
	"os"

	"github.com/exoscale/egoscale"
	"github.com/exoscale/egoscale/cmd/exo/table"
	"github.com/spf13/cobra"
)

// sshCreateCmd represents the create command
var sshCreateCmd = &cobra.Command{
	Use:     "create <name>",
	Short:   "Create SSH key pair",
	Aliases: gCreateAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Usage()
		}
		keyPair, err := createSSHKey(args[0])
		if err != nil {
			return err
		}
		displayResult(keyPair)
		return nil
	},
}

func createSSHKey(name string) (*egoscale.SSHKeyPair, error) {
	resp, err := cs.RequestWithContext(gContext, &egoscale.CreateSSHKeyPair{
		Name: name,
	})
	if err != nil {
		return nil, err
	}

	sshKeyPair, ok := resp.(*egoscale.SSHKeyPair)
	if !ok {
		return nil, fmt.Errorf("wrong type expected %q, got %T", "egoscale.CreateSSHKeyPairResponse", resp)
	}

	return sshKeyPair, nil
}

func displayResult(sshKeyPair *egoscale.SSHKeyPair) {
	table := table.NewTable(os.Stdout)
	table.SetHeader([]string{"Name", "Fingerprint"})
	table.Append([]string{sshKeyPair.Name, sshKeyPair.Fingerprint})
	table.Render()

	fmt.Println(sshKeyPair.PrivateKey)
}

func init() {
	sshkeyCmd.AddCommand(sshCreateCmd)
}
