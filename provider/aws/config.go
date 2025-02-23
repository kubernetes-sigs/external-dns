/*
Copyright 2023 The Kubernetes Authors.

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

package aws

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	awsv2 "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	stscredsv2 "github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/linki/instrumented_http"
	"github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
)

// AWSSessionConfig contains configuration to create a new AWS provider.
type AWSSessionConfig struct {
	AssumeRole           string
	AssumeRoleExternalID string
	APIRetries           int
	Profile              string
	DomainRolesMap       map[string]string
}

type AWSZoneConfig struct {
	Config         awsv2.Config
	HostedZoneName string
	Route53Config  Route53API
}

func CreateDefaultV2Config(cfg *externaldns.Config) *AWSZoneConfig {
	result, err := newV2Config(
		AWSSessionConfig{
			AssumeRole:           cfg.AWSAssumeRole,
			AssumeRoleExternalID: cfg.AWSAssumeRoleExternalID,
			APIRetries:           cfg.AWSAPIRetries,
			DomainRolesMap:       cfg.AWSDomainRoles,
		},
	)
	if err != nil {
		logrus.Fatal(err)
	}

	if len(result) == 0 {
		logrus.Fatal("No AWS credentials found")
	}

	return result[0]
}

func CreateV2Configs(cfg *externaldns.Config) map[string][]*AWSZoneConfig {
	result := make(map[string][]*AWSZoneConfig)
	if len(cfg.AWSProfiles) == 0 || (len(cfg.AWSProfiles) == 1 && cfg.AWSProfiles[0] == "") {
		cfg := CreateDefaultV2Config(cfg)
		result[defaultAWSProfile] = make([]*AWSZoneConfig, 0)
		result[defaultAWSProfile] = append(result[defaultAWSProfile], cfg)
	} else {
		for _, profile := range cfg.AWSProfiles {
			configs, err := newV2Config(
				AWSSessionConfig{
					AssumeRole:           cfg.AWSAssumeRole,
					AssumeRoleExternalID: cfg.AWSAssumeRoleExternalID,
					APIRetries:           cfg.AWSAPIRetries,
					Profile:              profile,
					DomainRolesMap:       cfg.AWSDomainRoles,
				},
			)
			if err != nil {
				logrus.Fatal(err)
			}
			result[profile] = configs
		}
	}
	return result
}

func newV2Config(awsConfig AWSSessionConfig) ([]*AWSZoneConfig, error) {
	hostedZonesConfigs := make([]*AWSZoneConfig, 0)
	defaultOpts := []func(*config.LoadOptions) error{
		config.WithRetryer(func() awsv2.Retryer {
			return retry.AddWithMaxAttempts(retry.NewStandard(), awsConfig.APIRetries)
		}),
		config.WithHTTPClient(instrumented_http.NewClient(&http.Client{}, &instrumented_http.Callbacks{
			PathProcessor: func(path string) string {
				parts := strings.Split(path, "/")
				return parts[len(parts)-1]
			},
		})),
		config.WithSharedConfigProfile(awsConfig.Profile),
	}

	cfg, err := config.LoadDefaultConfig(context.Background(), defaultOpts...)
	if err != nil {
		return nil, fmt.Errorf("instantiating AWS config: %w", err)
	}

	if awsConfig.DomainRolesMap == nil {
		hostedZonesConfigs = append(hostedZonesConfigs, &AWSZoneConfig{Config: cfg})
		return hostedZonesConfigs, nil
	}

	for dom, role := range awsConfig.DomainRolesMap {
		if role != "" {
			stsSvc := sts.NewFromConfig(cfg)
			var assumeRoleOpts []func(*stscredsv2.AssumeRoleOptions)
			if awsConfig.AssumeRoleExternalID != "" {
				logrus.Infof("Assuming role %s with external id", awsConfig.AssumeRole)
				logrus.Debugf("External id: %s", awsConfig.AssumeRoleExternalID)
				assumeRoleOpts = []func(*stscredsv2.AssumeRoleOptions){
					func(opts *stscredsv2.AssumeRoleOptions) {
						opts.ExternalID = &awsConfig.AssumeRoleExternalID
					},
				}
			} else {
				logrus.Infof("Assuming role: %s", role)
			}
			creds := stscredsv2.NewAssumeRoleProvider(stsSvc, role, assumeRoleOpts...)
			cfg.Credentials = awsv2.NewCredentialsCache(creds)

			hostedZonesConfigs = append(hostedZonesConfigs, &AWSZoneConfig{Config: cfg, HostedZoneName: dom})
		}
	}

	//if awsConfig.AssumeRole != "" {
	//	stsSvc := sts.NewFromConfig(cfg)
	//	var assumeRoleOpts []func(*stscredsv2.AssumeRoleOptions)
	//	if awsConfig.AssumeRoleExternalID != "" {
	//		logrus.Infof("Assuming role %s with external id", awsConfig.AssumeRole)
	//		logrus.Debugf("External id: %s", awsConfig.AssumeRoleExternalID)
	//		assumeRoleOpts = []func(*stscredsv2.AssumeRoleOptions){
	//			func(opts *stscredsv2.AssumeRoleOptions) {
	//				opts.ExternalID = &awsConfig.AssumeRoleExternalID
	//			},
	//		}
	//	} else {
	//		logrus.Infof("Assuming role: %s", awsConfig.AssumeRole)
	//	}
	//	creds := stscredsv2.NewAssumeRoleProvider(stsSvc, awsConfig.AssumeRole, assumeRoleOpts...)
	//	cfg.Credentials = awsv2.NewCredentialsCache(creds)
	//}

	return hostedZonesConfigs, nil
}
