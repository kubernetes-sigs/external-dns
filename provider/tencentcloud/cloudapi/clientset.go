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

package cloudapi

import (
	"fmt"
	"sync"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	privatedns "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/privatedns/v20201028"
	"go.uber.org/ratelimit"
)

type TencentClientSetService interface {
	PrivateDnsCli(action string) *privatedns.Client
	DnsPodCli(action string) *dnspod.Client
}

func NewTencentClientSetService(region string, rate int, secretId string, secretKey string, internetEndpoint bool) *defaultTencentClientSetService {
	p := &defaultTencentClientSetService{
		Region:    region,
		RateLimit: rate,
	}
	cred := common.NewCredential(secretId, secretKey)

	privatednsProf := profile.NewClientProfile()
	if !internetEndpoint {
		privatednsProf.HttpProfile.Endpoint = "privatedns.internal.tencentcloudapi.com"
	}
	p.privateDnsClient, _ = privatedns.NewClient(cred, region, privatednsProf)

	dnsPodProf := profile.NewClientProfile()
	if !internetEndpoint {
		dnsPodProf.HttpProfile.Endpoint = "dnspod.internal.tencentcloudapi.com"
	}
	p.dnsPodClient, _ = dnspod.NewClient(cred, region, dnsPodProf)

	return p
}

type defaultTencentClientSetService struct {
	Region           string
	RateLimit        int
	RateLimitSyncMap sync.Map

	privateDnsClient *privatedns.Client
	dnsPodClient     *dnspod.Client
}

func (p *defaultTencentClientSetService) checkRateLimit(request, method string) {
	action := fmt.Sprintf("%s_%s", request, method)
	if rl, ok := p.RateLimitSyncMap.LoadOrStore(action, ratelimit.New(p.RateLimit, ratelimit.WithoutSlack)); ok {
		rl.(ratelimit.Limiter).Take()
	}
}

func (p *defaultTencentClientSetService) PrivateDnsCli(action string) *privatedns.Client {
	p.checkRateLimit("privateDns", action)
	return p.privateDnsClient
}

func (p *defaultTencentClientSetService) DnsPodCli(action string) *dnspod.Client {
	p.checkRateLimit("dnsPod", action)
	return p.dnsPodClient
}
