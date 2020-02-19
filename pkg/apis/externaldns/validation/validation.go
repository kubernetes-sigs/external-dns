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

package validation

import (
	"errors"
	"fmt"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
)

// ValidateConfig performs validation on the Config object
func ValidateConfig(cfg *externaldns.Config) error {
	// TODO: Should probably return field.ErrorList
	if cfg.LogFormat != "text" && cfg.LogFormat != "json" {
		return fmt.Errorf("unsupported log format: %s", cfg.LogFormat)
	}
	if len(cfg.Sources) == 0 {
		return errors.New("no sources specified")
	}
	if cfg.Provider == "" {
		return errors.New("no provider specified")
	}

	// Azure provider specific validations
	if cfg.Provider == "azure" {
		if cfg.AzureConfigFile == "" {
			return errors.New("no Azure config file specified")
		}
	}

	// Akamai provider specific validations
	if cfg.Provider == "akamai" {
		if cfg.AkamaiServiceConsumerDomain == "" {
			return errors.New("no Akamai ServiceConsumerDomain specified")
		}
		if cfg.AkamaiClientToken == "" {
			return errors.New("no Akamai client token specified")
		}
		if cfg.AkamaiClientSecret == "" {
			return errors.New("no Akamai client secret specified")
		}
		if cfg.AkamaiAccessToken == "" {
			return errors.New("no Akamai access token specified")
		}
	}

	// Infoblox provider specific validations
	if cfg.Provider == "infoblox" {
		if cfg.InfobloxGridHost == "" {
			return errors.New("no Infoblox Grid Manager host specified")
		}
		if cfg.InfobloxWapiPassword == "" {
			return errors.New("no Infoblox WAPI password specified")
		}
	}

	if cfg.Provider == "dyn" {
		if cfg.DynUsername == "" {
			return errors.New("no Dyn username specified")
		}
		if cfg.DynCustomerName == "" {
			return errors.New("no Dyn customer name specified")
		}

		if cfg.DynMinTTLSeconds < 0 {
			return errors.New("TTL specified for Dyn is negative")
		}
	}

	if cfg.Provider == "rfc2136" {
		if cfg.RFC2136MinTTL < 0 {
			return errors.New("TTL specified for rfc2136 is negative")
		}
	}

	if cfg.IgnoreHostnameAnnotation && cfg.FQDNTemplate == "" {
		return errors.New("FQDN Template must be set if ignoring annotations")
	}
	return nil
}
