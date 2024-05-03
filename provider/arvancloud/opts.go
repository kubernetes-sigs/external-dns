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

package arvancloud

type providerOptions struct {
	apiVersion       string
	dnsPerPage       int
	domainPerPage    int
	enableCloudProxy bool
	dryRun           bool
}

type Option interface {
	apply(*providerOptions)
}

type proxyOption bool

func (p proxyOption) apply(opts *providerOptions) {
	opts.enableCloudProxy = bool(p)
}

// WithEnableCloudProxy Enable proxy cloud service
func WithEnableCloudProxy() Option {
	return proxyOption(true)
}

// WithDisableCloudProxy Disable proxy cloud service
func WithDisableCloudProxy() Option {
	return proxyOption(false)
}

type domainPerPageOption int

func (d domainPerPageOption) apply(opts *providerOptions) {
	opts.domainPerPage = int(d)
}

// WithDomainPerPage Add maximum records per page for domain (Default: 15)
func WithDomainPerPage(count int) Option {
	return domainPerPageOption(count)
}

type dnsPerPageOption int

func (d dnsPerPageOption) apply(opts *providerOptions) {
	opts.dnsPerPage = int(d)
}

// WithDnsPerPage Add maximum records per page for dns (Default: 15)
func WithDnsPerPage(count int) Option {
	return dnsPerPageOption(count)
}

type apiVersionOption string

func (a apiVersionOption) apply(opts *providerOptions) {
	opts.apiVersion = string(a)
}

// WithApiVersion Set api Arvan api version
func WithApiVersion(version string) Option {
	return apiVersionOption(version)
}
