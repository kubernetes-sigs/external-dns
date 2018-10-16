package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var dnsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add record to domain",
}

func init() {
	dnsCmd.AddCommand(dnsAddCmd)
}
