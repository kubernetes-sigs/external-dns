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

	log.Info("Creating ClouDNS Provider")

	loginType, ok := os.LookupEnv("CLOUDNS_LOGIN_TYPE")
	if !ok {
		return nil, fmt.Errorf("CLOUDNS_LOGIN_TYPE is not set")
	}
	if loginType != "user-id" && loginType != "sub-user" && loginType != "sub-user-name" {
		return nil, fmt.Errorf("CLOUDNS_LOGIN_TYPE is not valid")
	}

	switch loginType {
	case "user-id":
		log.Info("Using user-id login type")
	case "sub-user":
		log.Info("Using sub-user login type")
	case "sub-user-name":
		log.Info("Using sub-user-name login type")
	}

	return nil, nil
}

func (p *ClouDNSProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return nil, nil
}

func (p *ClouDNSProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	return nil
}
