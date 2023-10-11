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

package adguard

import (
	"errors"

	"sigs.k8s.io/external-dns/endpoint"
)

var (
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidEndpoint = errors.New("invalid endpoint")
)

type Config struct {
	// Username is the username used to authenticate with Adguard Home.
	// Required.
	Username string
	// Password is the plain text password used to authenticate.
	// Required.
	Password string
	// Server where Adguard Home can be reached.
	// Required
	Server       string
	DomainFilter endpoint.DomainFilter
}

func (ac Config) validate() error {
	if ac.Username == "" {
		return ErrInvalidUsername
	}
	if ac.Password == "" {
		return ErrInvalidPassword
	}
	if ac.Server == "" {
		return ErrInvalidEndpoint
	}

	return nil
}
