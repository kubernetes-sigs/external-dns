package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func init() {
	templateCmd.AddCommand(templateShowCmd)
}

// templateShowCmd represents the show command
var templateShowCmd = &cobra.Command{
	Use:   "show <template name | id>",
	Short: "Show a template",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("show expects one template by name or id")
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
		err := showTemplate(w, strings.Join(args, " "))
		if err == nil {
			return w.Flush()
		}
		return err
	},
}

func showTemplate(w io.Writer, name string) error {
	zoneID, err := getZoneIDByName(gCurrentAccount.DefaultZone)
	if err != nil {
		return err
	}

	template, err := getTemplateByName(zoneID, name)
	if err != nil {
		return err
	}

	username, usernameOk := template.Details["username"]

	fmt.Fprintf(w, "ID:\t%s\n", template.ID)              // nolint: errcheck
	fmt.Fprintf(w, "Name:\t%s\n", template.Name)          // nolint: errcheck
	fmt.Fprintf(w, "OS Type:\t%s\n", template.OsTypeName) // nolint: errcheck
	if usernameOk {
		fmt.Fprintf(w, "Username:\t%s\n", username) // nolint: errcheck
	}
	fmt.Fprintf(w, "Size:\t%d GiB\n", template.Size>>30)         // nolint: errcheck
	fmt.Fprintf(w, "Created:\t%s\n", template.Created)           // nolint: errcheck
	fmt.Fprintf(w, "Password?:\t%v\n", template.PasswordEnabled) // nolint: errcheck

	return nil
}
