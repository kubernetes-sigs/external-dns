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
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
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
}

func CreateDefaultSession(cfg *externaldns.Config) *session.Session {
	result, err := newSession(
		AWSSessionConfig{
			AssumeRole:           cfg.AWSAssumeRole,
			AssumeRoleExternalID: cfg.AWSAssumeRoleExternalID,
			APIRetries:           cfg.AWSAPIRetries,
		},
	)
	if err != nil {
		logrus.Fatal(err)
	}
	return result
}

func CreateSessions(cfg *externaldns.Config) map[string]*session.Session {
	result := make(map[string]*session.Session)

	if len(cfg.AWSProfiles) == 0 || (len(cfg.AWSProfiles) == 1 && cfg.AWSProfiles[0] == "") {
		session, err := newSession(
			AWSSessionConfig{
				AssumeRole:           cfg.AWSAssumeRole,
				AssumeRoleExternalID: cfg.AWSAssumeRoleExternalID,
				APIRetries:           cfg.AWSAPIRetries,
			},
		)
		if err != nil {
			logrus.Fatal(err)
		}
		result[defaultAWSProfile] = session
	} else {
		for _, profile := range cfg.AWSProfiles {
			session, err := newSession(
				AWSSessionConfig{
					AssumeRole:           cfg.AWSAssumeRole,
					AssumeRoleExternalID: cfg.AWSAssumeRoleExternalID,
					APIRetries:           cfg.AWSAPIRetries,
					Profile:              profile,
				},
			)
			if err != nil {
				logrus.Fatal(err)
			}
			result[profile] = session
		}
	}
	return result
}

func newSession(awsConfig AWSSessionConfig) (*session.Session, error) {
	config := aws.NewConfig().WithMaxRetries(awsConfig.APIRetries)

	config.WithHTTPClient(
		instrumented_http.NewClient(config.HTTPClient, &instrumented_http.Callbacks{
			PathProcessor: func(path string) string {
				parts := strings.Split(path, "/")
				return parts[len(parts)-1]
			},
		}),
	)

	session, err := session.NewSessionWithOptions(session.Options{
		Config:            *config,
		SharedConfigState: session.SharedConfigEnable,
		Profile:           awsConfig.Profile,
	})
	if err != nil {
		return nil, fmt.Errorf("instantiating AWS session: %w", err)
	}

	if awsConfig.AssumeRole != "" {
		if awsConfig.AssumeRoleExternalID != "" {
			logrus.Infof("Assuming role: %s with external id %s", awsConfig.AssumeRole, awsConfig.AssumeRoleExternalID)
			session.Config.WithCredentials(stscreds.NewCredentials(session, awsConfig.AssumeRole, func(p *stscreds.AssumeRoleProvider) {
				p.ExternalID = &awsConfig.AssumeRoleExternalID
			}))
		} else {
			logrus.Infof("Assuming role: %s", awsConfig.AssumeRole)
			session.Config.WithCredentials(stscreds.NewCredentials(session, awsConfig.AssumeRole))
		}
	}

	session.Handlers.Build.PushBack(request.MakeAddToUserAgentHandler("ExternalDNS", externaldns.Version))

	return session, nil
}
