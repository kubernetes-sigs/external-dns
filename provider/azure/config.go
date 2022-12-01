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
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// config represents common config items for Azure DNS and Azure Private DNS
type config struct {
	Cloud                        string            `json:"cloud" yaml:"cloud"`
	Environment                  azure.Environment `json:"-" yaml:"-"`
	TenantID                     string            `json:"tenantId" yaml:"tenantId"`
	SubscriptionID               string            `json:"subscriptionId" yaml:"subscriptionId"`
	ResourceGroup                string            `json:"resourceGroup" yaml:"resourceGroup"`
	Location                     string            `json:"location" yaml:"location"`
	ClientID                     string            `json:"aadClientId" yaml:"aadClientId"`
	ClientSecret                 string            `json:"aadClientSecret" yaml:"aadClientSecret"`
	UseManagedIdentityExtension  bool              `json:"useManagedIdentityExtension" yaml:"useManagedIdentityExtension"`
	UseWorkloadIdentityExtension bool              `json:"useWorkloadIdentityExtension" yaml:"useWorkloadIdentityExtension"`
	UserAssignedIdentityID       string            `json:"userAssignedIdentityID" yaml:"userAssignedIdentityID"`
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

	// Try to retrieve token with Workload Identity.
	if cfg.UseWorkloadIdentityExtension {
		log.Info("Using workload identity extension to retrieve access token for Azure API.")

		token, err := getWIToken(environment, cfg)
		if err != nil {
			return nil, err
		}

		// adal does not offer methods to dynamically replace a federated token, thus we need to have a wrapper to make sure
		// we're using up-to-date secret while requesting an access token.
		// NOTE: There's no RefreshToken in the whole process (in fact, it's absent in AAD responses). An AccessToken can be
		// received only in exchange for a federated token.
		var refreshFunc adal.TokenRefresh = func(context context.Context, resource string) (*adal.Token, error) {
			newWIToken, err := getWIToken(environment, cfg)
			if err != nil {
				return nil, err
			}

			// An AccessToken gets populated into an spt only when .Refresh() is called. Normally, it's something that happens implicitly when
			// a first request to manipulate Azure resources is made. Since our goal here is only to receive a fresh AccessToken, we need to make
			// an explicit call.
			// .Refresh() itself results in a call to Oauth endpoint. During the process, a federated token is exchanged for an AccessToken.
			// RefreshToken is absent from responses.
			err = newWIToken.Refresh()
			if err != nil {
				return nil, err
			}

			accessToken := newWIToken.Token()

			return &accessToken, nil
		}

		token.SetCustomRefreshFunc(refreshFunc)

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

// getWIToken prepares a token for a Workload Identity-enabled setup
func getWIToken(environment azure.Environment, cfg config) (*adal.ServicePrincipalToken, error) {
	// NOTE: all related environment variables are described here: https://azure.github.io/azure-workload-identity/docs/installation/mutating-admission-webhook.html
	oauthConfig, err := adal.NewOAuthConfig(environment.ActiveDirectoryEndpoint, os.Getenv("AZURE_TENANT_ID"))
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve OAuth config: %v", err)
	}

	jwt, err := os.ReadFile(os.Getenv("AZURE_FEDERATED_TOKEN_FILE"))
	if err != nil {
		return nil, fmt.Errorf("failed to read a file with a federated token: %v", err)
	}

	// AZURE_CLIENT_ID will be empty in case azure.workload.identity/client-id annotation is not set
	// Thus, it's important to offer optional ClientID overrides
	clientID := os.Getenv("AZURE_CLIENT_ID")
	if cfg.UserAssignedIdentityID != "" {
		clientID = cfg.UserAssignedIdentityID
	}

	token, err := adal.NewServicePrincipalTokenFromFederatedToken(*oauthConfig, clientID, string(jwt), environment.ResourceManagerEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create a workload identity token: %v", err)
	}

	return token, nil
}
