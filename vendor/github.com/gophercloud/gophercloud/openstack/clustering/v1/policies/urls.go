package policies

import "github.com/gophercloud/gophercloud"

const (
	apiVersion = "v1"
	apiName    = "policies"
)

func policyListURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}
