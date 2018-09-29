package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "DNS cmd lets you host your zones and manage records",
}

// getRecordIDByName get record ID by name
func getRecordIDByName(domainName, recordName string) (int64, error) {
	records, err := csDNS.GetRecords(domainName)
	if err != nil {
		return 0, err
	}

	resRecID := []int64{}

	for _, r := range records {
		id := fmt.Sprintf("%d", r.ID)
		if id == recordName {
			return r.ID, nil
		}
		if recordName == r.Name {
			resRecID = append(resRecID, r.ID)
		}
	}
	if len(resRecID) > 1 {
		return 0, fmt.Errorf("more than one records were found")
	}
	if len(resRecID) == 1 {
		return resRecID[0], nil
	}

	return 0, fmt.Errorf("no records were found")
}

func init() {
	RootCmd.AddCommand(dnsCmd)
}
