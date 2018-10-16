package cmd

import (
	"io/ioutil"
	"os"

	"github.com/exoscale/egoscale"
	"github.com/exoscale/egoscale/cmd/exo/table"
	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:     "upload <name> <public key file>",
	Short:   "Upload SSH key pair from given public key",
	Aliases: gUploadAlias,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return cmd.Usage()
		}
		return uploadSSHKey(args[0], args[1])
	},
}

func uploadSSHKey(name, publicKeyPath string) error {
	pbk, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return err
	}

	resp, err := cs.RequestWithContext(gContext, &egoscale.RegisterSSHKeyPair{
		Name:      name,
		PublicKey: string(pbk),
	})

	if err != nil {
		return err
	}

	keyPair := resp.(*egoscale.SSHKeyPair)
	table := table.NewTable(os.Stdout)
	table.SetHeader([]string{"Name", "Fingerprint"})
	table.Append([]string{keyPair.Name, keyPair.Fingerprint})
	table.Render()
	return nil
}

func init() {
	sshkeyCmd.AddCommand(uploadCmd)
}
