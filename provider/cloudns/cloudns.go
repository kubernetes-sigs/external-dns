/*
Copyright 2022 The Kubernetes Authors.
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

package cloudns

import (
	"context"
	"fmt"
	"os"
	"strconv"

	cloudns "github.com/wmarchesi123/cloudns-go"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type ClouDNSProvider struct {
	provider.BaseProvider
	client       *cloudns.Client
	context      context.Context
	domainFilter endpoint.DomainFilter
	zoneIDFilter provider.ZoneIDFilter
	ownerID      string
	dryRun       bool
	testing      bool
}

type ClouDNSConfig struct {
	Context      context.Context
	DomainFilter endpoint.DomainFilter
	ZoneIDFilter provider.ZoneIDFilter
	OwnerID      string
	DryRun       bool
	Testing      bool
}

func NewClouDNSProvider(config ClouDNSConfig) (*ClouDNSProvider, error) {

	var client *cloudns.Client

	log.Info("Creating ClouDNS Provider")

	loginType, ok := os.LookupEnv("CLOUDNS_LOGIN_TYPE")
	if !ok {
		return nil, fmt.Errorf("CLOUDNS_LOGIN_TYPE is not set")
	}
	if loginType != "user-id" && loginType != "sub-user" && loginType != "sub-user-name" {
		return nil, fmt.Errorf("CLOUDNS_LOGIN_TYPE is not valid")
	}

	userPassword, ok := os.LookupEnv("CLOUDNS_USER_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("CLOUDNS_USER_PASSWORD is not set")
	}

	switch loginType {
	case "user-id":
		log.Info("Using user-id login type")

		userIDString, ok := os.LookupEnv("CLOUDNS_USER_ID")
		if !ok {
			return nil, fmt.Errorf("CLOUDNS_USER_ID is not set")
		}

		userIDInt, error := strconv.Atoi(userIDString)
		if error != nil {
			return nil, fmt.Errorf("CLOUDNS_USER_ID is not a valid integer")
		}

		c, error := cloudns.New(
			cloudns.AuthUserID(userIDInt, userPassword),
		)
		if error != nil {
			return nil, fmt.Errorf("error creating ClouDNS client: %s", error)
		}

		client = c
		log.Info("Authenticated with ClouDNS using user-id login type")
	case "sub-user":
		log.Info("Using sub-user login type")

		subUYserIDString, ok := os.LookupEnv("CLOUDNS_SUB_USER_ID")
		if !ok {
			return nil, fmt.Errorf("CLOUDNS_SUB_USER_ID is not set")
		}

		subUserIDInt, error := strconv.Atoi(subUYserIDString)
		if error != nil {
			return nil, fmt.Errorf("CLOUDNS_SUB_USER_ID is not a valid integer")
		}

		c, error := cloudns.New(
			cloudns.AuthSubUserID(subUserIDInt, userPassword),
		)
		if error != nil {
			return nil, fmt.Errorf("error creating ClouDNS client: %s", error)
		}

		client = c
		log.Info("Authenticated with ClouDNS using sub-user login type")
	case "sub-user-name":
		log.Info("Using sub-user-name login type")

		subUserName, ok := os.LookupEnv("CLOUDNS_SUB_USER_NAME")
		if !ok {
			return nil, fmt.Errorf("CLOUDNS_SUB_USER_NAME is not set")
		}

		c, error := cloudns.New(
			cloudns.AuthSubUserName(subUserName, userPassword),
		)
		if error != nil {
			return nil, fmt.Errorf("error creating ClouDNS client: %s", error)
		}

		client = c
		log.Info("Authenticated with ClouDNS using sub-user-name login type")
	}

	provider := &ClouDNSProvider{
		client:       client,
		context:      config.Context,
		domainFilter: config.DomainFilter,
		zoneIDFilter: config.ZoneIDFilter,
		ownerID:      config.OwnerID,
		dryRun:       config.DryRun,
		testing:      config.Testing,
	}

	return provider, nil
}

func (p *ClouDNSProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return nil, nil
}

func (p *ClouDNSProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	return nil
}
