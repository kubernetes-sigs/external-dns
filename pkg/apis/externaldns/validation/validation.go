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
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
)

// ValidateConfig performs validation on the Config object
func ValidateConfig(cfg *externaldns.Config) error {
	// TODO: Should probably return field.ErrorList
	validators := []func(*externaldns.Config) error{
		preValidateConfig,
		validateConfigForProvider,
		validateHostnameConfig,
		validateTXTRegistryConfig,
		validateLabelSelectors,
		validateAnnotationPrefix,
		validateKubeAPILimits,
		validatePTRConfig,
	}
	for _, validate := range validators {
		if err := validate(cfg); err != nil {
			return err
		}
	}
	return nil
}

// validateHostnameConfig ensures an FQDN template is set when hostname annotations are ignored.
func validateHostnameConfig(cfg *externaldns.Config) error {
	if cfg.IgnoreHostnameAnnotation && len(cfg.FQDNTemplate) == 0 {
		return errors.New("FQDN Template must be set if ignoring annotations")
	}
	return nil
}

// validateTXTRegistryConfig ensures the TXT registry prefix and suffix are not set together.
func validateTXTRegistryConfig(cfg *externaldns.Config) error {
	if len(cfg.TXTPrefix) > 0 && len(cfg.TXTSuffix) > 0 {
		return errors.New("txt-prefix and txt-suffix are mutual exclusive")
	}
	return nil
}

// validateLabelSelectors ensures the label and annotation filters are valid selectors.
func validateLabelSelectors(cfg *externaldns.Config) error {
	if _, err := labels.Parse(cfg.LabelFilter); err != nil {
		return errors.New("--label-filter does not specify a valid label selector")
	}
	if _, err := metav1.ParseToLabelSelector(cfg.AnnotationFilter); err != nil {
		return errors.New("--annotation-filter does not specify a valid label selector")
	}
	return nil
}

// validateAnnotationPrefix ensures the annotation prefix is set and ends with a slash.
func validateAnnotationPrefix(cfg *externaldns.Config) error {
	if cfg.AnnotationPrefix == "" {
		return errors.New("--annotation-prefix cannot be empty")
	}
	if !strings.HasSuffix(cfg.AnnotationPrefix, "/") {
		return errors.New("--annotation-prefix must end with '/'")
	}
	return nil
}

// validateKubeAPILimits ensures the Kubernetes API QPS and burst limits are positive.
func validateKubeAPILimits(cfg *externaldns.Config) error {
	if cfg.KubeAPIQPS <= 0 {
		return errors.New("--kube-api-qps must be greater than 0")
	}
	if cfg.KubeAPIBurst <= 0 {
		return errors.New("--kube-api-burst must be greater than 0")
	}
	return nil
}

// validatePTRConfig ensures PTR is a managed record type when PTR creation is enabled.
func validatePTRConfig(cfg *externaldns.Config) error {
	if cfg.CreatePTR && !cfg.IsPTRSupported() {
		return errors.New("--create-ptr requires PTR in --managed-record-types")
	}
	return nil
}

func preValidateConfig(cfg *externaldns.Config) error {
	if cfg.LogFormat != externaldns.LogFormatText && cfg.LogFormat != externaldns.LogFormatJSON {
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
	case externaldns.ProviderAzure:
		return validateConfigForAzure(cfg)
	case externaldns.ProviderRFC2136:
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
