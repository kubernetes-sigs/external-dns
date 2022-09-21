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
	"io/ioutil"
	"strings"

	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// config represents common config items for Azure DNS and Azure Private DNS
type config struct {
	Cloud                       string            `json:"cloud" yaml:"cloud"`
	Environment                 azure.Environment `json:"-" yaml:"-"`
	TenantID                    string            `json:"tenantId" yaml:"tenantId"`
	SubscriptionID              string            `json:"subscriptionId" yaml:"subscriptionId"`
	ResourceGroup               string            `json:"resourceGroup" yaml:"resourceGroup"`
	Location                    string            `json:"location" yaml:"location"`
	ClientID                    string            `json:"aadClientId" yaml:"aadClientId"`
	ClientSecret                string            `json:"aadClientSecret" yaml:"aadClientSecret"`
	UseManagedIdentityExtension bool              `json:"useManagedIdentityExtension" yaml:"useManagedIdentityExtension"`
	UserAssignedIdentityID      string            `json:"userAssignedIdentityID" yaml:"userAssignedIdentityID"`
}

func getConfig(configFile, resourceGroup, userAssignedIdentityClientID string) (*config, error) {
	contents, err := ioutil.ReadFile(configFile)
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

	var environment azure.Environment
	if cfg.Cloud == "" {
		environment = azure.PublicCloud
	} else {
		environment, err = azure.EnvironmentFromName(cfg.Cloud)
		if err != nil {
			return nil, fmt.Errorf("invalid cloud value '%s': %v", cfg.Cloud, err)
		}
	}
	cfg.Environment = environment

	return cfg, nil
}

// getAccessToken retrieves Azure API access token.
func getAccessToken(cfg config, environment azure.Environment) (*adal.ServicePrincipalToken, error) {
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
		oauthConfig, err := adal.NewOAuthConfig(environment.ActiveDirectoryEndpoint, cfg.TenantID)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve OAuth config: %v", err)
		}

		token, err := adal.NewServicePrincipalToken(*oauthConfig, cfg.ClientID, cfg.ClientSecret, environment.ResourceManagerEndpoint)
		if err != nil {
			return nil, fmt.Errorf("failed to create service principal token: %v", err)
		}
		return token, nil
	}

	// Try to retrieve token with MSI.
	if cfg.UseManagedIdentityExtension {
		log.Info("Using managed identity extension to retrieve access token for Azure API.")

		if cfg.UserAssignedIdentityID != "" {
			log.Infof("Resolving to user assigned identity, client id is %s.", cfg.UserAssignedIdentityID)
			token, err := adal.NewServicePrincipalTokenFromManagedIdentity(environment.ServiceManagementEndpoint, &adal.ManagedIdentityOptions{
				ClientID: cfg.UserAssignedIdentityID,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to create the managed service identity token: %v", err)
			}
			return token, nil
		}

		log.Info("Resolving to system assigned identity.")
		token, err := adal.NewServicePrincipalTokenFromManagedIdentity(environment.ServiceManagementEndpoint, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create the managed service identity token: %v", err)
		}
		return token, nil
	}

	return nil, fmt.Errorf("no credentials provided for Azure API")
}
