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

	"k8s.io/apimachinery/pkg/labels"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
)

// ValidateConfig performs validation on the Config object
func ValidateConfig(cfg *externaldns.Config) error {
	// TODO: Should probably return field.ErrorList

	if err := preValidateConfig(cfg); err != nil {
		return err
	}

	if err := validateConfigForProvider(cfg); err != nil {
		return err
	}

	if cfg.IgnoreHostnameAnnotation && cfg.FQDNTemplate == "" {
		return errors.New("FQDN Template must be set if ignoring annotations")
	}

	if len(cfg.TXTPrefix) > 0 && len(cfg.TXTSuffix) > 0 {
		return errors.New("txt-prefix and txt-suffix are mutual exclusive")
	}

	_, err := labels.Parse(cfg.LabelFilter)
	if err != nil {
		return errors.New("--label-filter does not specify a valid label selector")
	}
	return nil
}

func preValidateConfig(cfg *externaldns.Config) error {
	if cfg.LogFormat != "text" && cfg.LogFormat != "json" {
		return fmt.Errorf("unsupported log format: %s", cfg.LogFormat)
	}
	if len(cfg.Sources) == 0 {
		return errors.New("no sources specified")
	}
	if cfg.Provider == "" {
		return errors.New("no provider specified")
	}
	return nil
}

func validateConfigForProvider(cfg *externaldns.Config) error {
	switch cfg.Provider {
	case "azure":
		return validateConfigForAzure(cfg)
	case "akamai":
		return validateConfigForAkamai(cfg)
	case "rfc2136":
		return validateConfigForRfc2136(cfg)
	default:
		return nil
	}
}

func validateConfigForAzure(cfg *externaldns.Config) error {
	if cfg.AzureConfigFile == "" {
		return errors.New("no Azure config file specified")
	}
	return nil
}

func validateConfigForAkamai(cfg *externaldns.Config) error {
	if cfg.AkamaiServiceConsumerDomain == "" && cfg.AkamaiEdgercPath != "" {
		return errors.New("no Akamai ServiceConsumerDomain specified")
	}
	if cfg.AkamaiClientToken == "" && cfg.AkamaiEdgercPath != "" {
		return errors.New("no Akamai client token specified")
	}
	if cfg.AkamaiClientSecret == "" && cfg.AkamaiEdgercPath != "" {
		return errors.New("no Akamai client secret specified")
	}
	if cfg.AkamaiAccessToken == "" && cfg.AkamaiEdgercPath != "" {
		return errors.New("no Akamai access token specified")
	}
	return nil
}

func validateConfigForRfc2136(cfg *externaldns.Config) error {
	if cfg.RFC2136MinTTL < 0 {
		return errors.New("TTL specified for rfc2136 is negative")
	}
	if cfg.RFC2136Insecure && cfg.RFC2136GSSTSIG {
		return errors.New("--rfc2136-insecure and --rfc2136-gss-tsig are mutually exclusive arguments")
	}
	if cfg.RFC2136GSSTSIG {
		if cfg.RFC2136KerberosPassword == "" || cfg.RFC2136KerberosUsername == "" || cfg.RFC2136KerberosRealm == "" {
			return errors.New("--rfc2136-kerberos-realm, --rfc2136-kerberos-username, and --rfc2136-kerberos-password are required when specifying --rfc2136-gss-tsig option")
		}
	}
	if cfg.RFC2136BatchChangeSize < 1 {
		return errors.New("batch size specified for rfc2136 cannot be less than 1")
	}
	return nil
}
