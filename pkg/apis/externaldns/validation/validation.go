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
	"slices"
	"strings"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"sigs.k8s.io/external-dns/endpoint"
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

	if cfg.IgnoreHostnameAnnotation && len(cfg.FQDNTemplate) == 0 {
		return errors.New("FQDN Template must be set if ignoring annotations")
	}

	if len(cfg.TXTPrefix) > 0 && len(cfg.TXTSuffix) > 0 {
		return errors.New("txt-prefix and txt-suffix are mutual exclusive")
	}

	_, err := labels.Parse(cfg.LabelFilter)
	if err != nil {
		return errors.New("--label-filter does not specify a valid label selector")
	}

	if _, err := metav1.ParseToLabelSelector(cfg.AnnotationFilter); err != nil {
		return errors.New("--annotation-filter does not specify a valid label selector")
	}

	if cfg.AnnotationPrefix == "" {
		return errors.New("--annotation-prefix cannot be empty")
	}
	if !strings.HasSuffix(cfg.AnnotationPrefix, "/") {
		return errors.New("--annotation-prefix must end with '/'")
	}

	if cfg.KubeAPIQPS <= 0 {
		return errors.New("--kube-api-qps must be greater than 0")
	}
	if cfg.KubeAPIBurst <= 0 {
		return errors.New("--kube-api-burst must be greater than 0")
	}

	if cfg.CreatePTR && !cfg.IsPTRSupported() {
		return errors.New("--create-ptr requires PTR in --managed-record-types")
	}

	if err := validateACMEDelegation(cfg); err != nil {
		return err
	}

	return nil
}

// validateACMEDelegation checks the consistency of the --acme-cname-delegation-* flags.
func validateACMEDelegation(cfg *externaldns.Config) error {
	if cfg.ACMEDelegationTargetTemplate == "" {
		if len(cfg.ACMEDelegationDomainFilter) > 0 || cfg.ACMEDelegationTTL != 0 {
			return errors.New("--acme-cname-delegation-domain-filter and --acme-cname-delegation-ttl require --acme-cname-delegation-target-template")
		}
		return nil
	}
	if cfg.ACMEDelegationTTL < 0 {
		return errors.New("--acme-cname-delegation-ttl cannot be negative")
	}
	if !slices.Contains(cfg.ManagedDNSRecordTypes, endpoint.RecordTypeCNAME) {
		return errors.New("--acme-cname-delegation-target-template requires CNAME in --managed-record-types")
	}
	if cfg.RegexDomainFilter != nil && cfg.RegexDomainFilter.String() != "" {
		log.Warn("--regex-domain-filter is set together with --acme-cname-delegation-target-template; the regex must also match the generated '_acme-challenge.*' names or they will be filtered out")
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
