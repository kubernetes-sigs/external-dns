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
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
)

// Validator validates a Config, returning any problems as a field.ErrorList.
type Validator interface {
	Validate(cfg *externaldns.Config) field.ErrorList
}

// validatorFunc adapts a plain function to the Validator interface.
type validatorFunc func(cfg *externaldns.Config) field.ErrorList

func (f validatorFunc) Validate(cfg *externaldns.Config) field.ErrorList {
	return f(cfg)
}

// ValidateConfig performs validation on the Config object
func ValidateConfig(cfg *externaldns.Config) error {
	validators := []Validator{
		validatorFunc(preValidateConfig),
		validatorFunc(validateConfigForProvider),
		validatorFunc(validateHostnameConfig),
		validatorFunc(validateTXTRegistryConfig),
		validatorFunc(validateLabelSelectors),
		validatorFunc(validateAnnotationPrefix),
		validatorFunc(validateKubeAPILimits),
		validatorFunc(validatePTRConfig),
	}

	var errs field.ErrorList
	for _, v := range validators {
		errs = append(errs, v.Validate(cfg)...)
	}
	return errs.ToAggregate()
}

// validateHostnameConfig ensures an FQDN template is set when hostname annotations are ignored.
func validateHostnameConfig(cfg *externaldns.Config) field.ErrorList {
	if cfg.IgnoreHostnameAnnotation && len(cfg.FQDNTemplate) == 0 {
		return field.ErrorList{field.Required(field.NewPath("fqdn-template"), "FQDN Template must be set if ignoring annotations")}
	}
	return nil
}

// validateTXTRegistryConfig ensures the TXT registry prefix and suffix are not set together.
func validateTXTRegistryConfig(cfg *externaldns.Config) field.ErrorList {
	if len(cfg.TXTPrefix) > 0 && len(cfg.TXTSuffix) > 0 {
		return field.ErrorList{field.Invalid(field.NewPath("txt-prefix"), cfg.TXTPrefix, "txt-prefix and txt-suffix are mutual exclusive")}
	}
	return nil
}

// validateLabelSelectors ensures the label and annotation filters are valid selectors.
func validateLabelSelectors(cfg *externaldns.Config) field.ErrorList {
	var errs field.ErrorList
	if _, err := labels.Parse(cfg.LabelFilter); err != nil {
		errs = append(errs, field.Invalid(field.NewPath("label-filter"), cfg.LabelFilter, "does not specify a valid label selector"))
	}
	if _, err := metav1.ParseToLabelSelector(cfg.AnnotationFilter); err != nil {
		errs = append(errs, field.Invalid(field.NewPath("annotation-filter"), cfg.AnnotationFilter, "does not specify a valid label selector"))
	}
	return errs
}

// validateAnnotationPrefix ensures the annotation prefix is set and ends with a slash.
func validateAnnotationPrefix(cfg *externaldns.Config) field.ErrorList {
	path := field.NewPath("annotation-prefix")
	if cfg.AnnotationPrefix == "" {
		return field.ErrorList{field.Required(path, "--annotation-prefix cannot be empty")}
	}
	if !strings.HasSuffix(cfg.AnnotationPrefix, "/") {
		return field.ErrorList{field.Invalid(path, cfg.AnnotationPrefix, "--annotation-prefix must end with '/'")}
	}
	return nil
}

// validateKubeAPILimits ensures the Kubernetes API QPS and burst limits are positive.
func validateKubeAPILimits(cfg *externaldns.Config) field.ErrorList {
	var errs field.ErrorList
	if cfg.KubeAPIQPS <= 0 {
		errs = append(errs, field.Invalid(field.NewPath("kube-api-qps"), cfg.KubeAPIQPS, "must be greater than 0"))
	}
	if cfg.KubeAPIBurst <= 0 {
		errs = append(errs, field.Invalid(field.NewPath("kube-api-burst"), cfg.KubeAPIBurst, "must be greater than 0"))
	}
	return errs
}

// validatePTRConfig ensures PTR is a managed record type when PTR creation is enabled.
func validatePTRConfig(cfg *externaldns.Config) field.ErrorList {
	if cfg.CreatePTR && !cfg.IsPTRSupported() {
		return field.ErrorList{field.Invalid(field.NewPath("create-ptr"), cfg.CreatePTR, "--create-ptr requires PTR in --managed-record-types")}
	}
	return nil
}

func preValidateConfig(cfg *externaldns.Config) field.ErrorList {
	var errs field.ErrorList
	if cfg.LogFormat != externaldns.LogFormatText && cfg.LogFormat != externaldns.LogFormatJSON {
		errs = append(errs, field.Invalid(field.NewPath("log-format"), cfg.LogFormat, "unsupported log format"))
	}
	if len(cfg.Sources) == 0 {
		errs = append(errs, field.Required(field.NewPath("source"), "no sources specified"))
	}
	if cfg.Provider == "" {
		errs = append(errs, field.Required(field.NewPath("provider"), "no provider specified"))
	}
	return errs
}

func validateConfigForProvider(cfg *externaldns.Config) field.ErrorList {
	switch cfg.Provider {
	case externaldns.ProviderAzure:
		return validateConfigForAzure(cfg)
	case externaldns.ProviderRFC2136:
		return validateConfigForRfc2136(cfg)
	default:
		return nil
	}
}

func validateConfigForAzure(cfg *externaldns.Config) field.ErrorList {
	if cfg.AzureConfigFile == "" {
		return field.ErrorList{field.Required(field.NewPath("azure-config-file"), "no Azure config file specified")}
	}
	return nil
}

func validateConfigForRfc2136(cfg *externaldns.Config) field.ErrorList {
	var errs field.ErrorList
	if cfg.RFC2136MinTTL < 0 {
		errs = append(errs, field.Invalid(field.NewPath("rfc2136-min-ttl"), cfg.RFC2136MinTTL, "TTL specified for rfc2136 is negative"))
	}
	if cfg.RFC2136Insecure && cfg.RFC2136GSSTSIG {
		errs = append(errs, field.Invalid(field.NewPath("rfc2136-insecure"), cfg.RFC2136Insecure, "--rfc2136-insecure and --rfc2136-gss-tsig are mutually exclusive arguments"))
	}
	if cfg.RFC2136GSSTSIG {
		if cfg.RFC2136KerberosPassword == "" || cfg.RFC2136KerberosUsername == "" || cfg.RFC2136KerberosRealm == "" {
			errs = append(errs, field.Required(field.NewPath("rfc2136-kerberos-realm"), "--rfc2136-kerberos-realm, --rfc2136-kerberos-username, and --rfc2136-kerberos-password are required when specifying --rfc2136-gss-tsig option"))
		}
	}
	if cfg.RFC2136BatchChangeSize < 1 {
		errs = append(errs, field.Invalid(field.NewPath("rfc2136-batch-change-size"), cfg.RFC2136BatchChangeSize, "batch size specified for rfc2136 cannot be less than 1"))
	}
	return errs
}
