/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package azure

import (
	"fmt"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// config represents common config items for Azure DNS and Azure Private DNS
type config struct {
	Cloud                        string `json:"cloud" yaml:"cloud"`
	TenantID                     string `json:"tenantId" yaml:"tenantId"`
	SubscriptionID               string `json:"subscriptionId" yaml:"subscriptionId"`
	ResourceGroup                string `json:"resourceGroup" yaml:"resourceGroup"`
	Location                     string `json:"location" yaml:"location"`
	ClientID                     string `json:"aadClientId" yaml:"aadClientId"`
	ClientSecret                 string `json:"aadClientSecret" yaml:"aadClientSecret"`
	UseManagedIdentityExtension  bool   `json:"useManagedIdentityExtension" yaml:"useManagedIdentityExtension"`
	UseWorkloadIdentityExtension bool   `json:"useWorkloadIdentityExtension" yaml:"useWorkloadIdentityExtension"`
	UserAssignedIdentityID       string `json:"userAssignedIdentityID" yaml:"userAssignedIdentityID"`
}

func getConfig(configFile, resourceGroup, userAssignedIdentityClientID string) (*config, error) {
	contents, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read Azure config file '%s': %v", configFile, err)
	}
	cfg := &config{}
	err = yaml.Unmarshal(contents, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read Azure config file '%s': %v", configFile, err)
	}

	// If a resource group was given, override what was present in the config file
	if resourceGroup != "" {
		cfg.ResourceGroup = resourceGroup
	}
	// If userAssignedIdentityClientID is provided explicitly, override existing one in config file
	if userAssignedIdentityClientID != "" {
		cfg.UserAssignedIdentityID = userAssignedIdentityClientID
	}
	return cfg, nil
}

// getAccessToken retrieves Azure API access token.
func getCredentials(cfg config) (azcore.TokenCredential, error) {
	cloudCfg, err := getCloudConfiguration(cfg.Cloud)
	if err != nil {
		return nil, fmt.Errorf("failed to get cloud configuration: %w", err)
	}

	// Try to retrieve token with service principal credentials.
	// Try to use service principal first, some AKS clusters are in an intermediate state that `UseManagedIdentityExtension` is `true`
	// and service principal exists. In this case, we still want to use service principal to authenticate.
	if len(cfg.ClientID) > 0 &&
		len(cfg.ClientSecret) > 0 &&
		// due to some historical reason, for pure MSI cluster,
		// they will use "msi" as placeholder in azure.json.
		// In this case, we shouldn't try to use SPN to authenticate.
		!strings.EqualFold(cfg.ClientID, "msi") &&
		!strings.EqualFold(cfg.ClientSecret, "msi") {
		log.Info("Using client_id+client_secret to retrieve access token for Azure API.")
		opts := &azidentity.ClientSecretCredentialOptions{
			ClientOptions: azcore.ClientOptions{
				Cloud: cloudCfg,
			},
		}
		cred, err := azidentity.NewClientSecretCredential(cfg.TenantID, cfg.ClientID, cfg.ClientSecret, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to create service principal token: %w", err)
		}
		return cred, nil
	}

	// Try to retrieve token with Workload Identity.
	if cfg.UseWorkloadIdentityExtension {
		log.Info("Using workload identity extension to retrieve access token for Azure API.")

		wiOpt := azidentity.WorkloadIdentityCredentialOptions{
			ClientOptions: azcore.ClientOptions{
				Cloud: cloudCfg,
			},
			// In a standard scenario, Client ID and Tenant ID are expected to be read from environment variables.
			// Though, in certain cases, it might be important to have an option to override those (e.g. when AZURE_TENANT_ID is not set
			// through a webhook or azure.workload.identity/client-id service account annotation is absent). When any of those values are
			// empty in our config, they will automatically be read from environment variables by azidentity
			TenantID: cfg.TenantID,
			ClientID: cfg.ClientID,
		}

		cred, err := azidentity.NewWorkloadIdentityCredential(&wiOpt)
		if err != nil {
			return nil, fmt.Errorf("failed to create a workload identity token: %w", err)
		}

		return cred, nil
	}

	// Try to retrieve token with MSI.
	if cfg.UseManagedIdentityExtension {
		log.Info("Using managed identity extension to retrieve access token for Azure API.")
		msiOpt := azidentity.ManagedIdentityCredentialOptions{
			ClientOptions: azcore.ClientOptions{
				Cloud: cloudCfg,
			},
		}
		if cfg.UserAssignedIdentityID != "" {
			msiOpt.ID = azidentity.ClientID(cfg.UserAssignedIdentityID)
		}
		cred, err := azidentity.NewManagedIdentityCredential(&msiOpt)
		if err != nil {
			return nil, fmt.Errorf("failed to create the managed service identity token: %w", err)
		}
		return cred, nil
	}

	return nil, fmt.Errorf("no credentials provided for Azure API")
}

func getCloudConfiguration(name string) (cloud.Configuration, error) {
	name = strings.ToUpper(name)
	switch name {
	case "AZURECLOUD", "AZUREPUBLICCLOUD", "":
		return cloud.AzurePublic, nil
	case "AZUREUSGOVERNMENT", "AZUREUSGOVERNMENTCLOUD":
		return cloud.AzureGovernment, nil
	case "AZURECHINACLOUD":
		return cloud.AzureChina, nil
	}
	return cloud.Configuration{}, fmt.Errorf("unknown cloud name: %s", name)
}
