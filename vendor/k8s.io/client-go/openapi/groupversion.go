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

package openapi

import (
	"context"
<<<<<<< HEAD

	openapi_v3 "github.com/google/gnostic/openapiv3"
	"google.golang.org/protobuf/proto"
	"k8s.io/kube-openapi/pkg/handler3"
)

const openAPIV3mimePb = "application/com.github.proto-openapi.spec.v3@v1.0+protobuf"

type GroupVersion interface {
	Schema() (*openapi_v3.Document, error)
}

type groupversion struct {
	client *client
	item   handler3.OpenAPIV3DiscoveryGroupVersion
}

func newGroupVersion(client *client, item handler3.OpenAPIV3DiscoveryGroupVersion) *groupversion {
	return &groupversion{client: client, item: item}
}

func (g *groupversion) Schema() (*openapi_v3.Document, error) {
	data, err := g.client.restClient.Get().
		RequestURI(g.item.ServerRelativeURL).
		SetHeader("Accept", openAPIV3mimePb).
		Do(context.TODO()).
		Raw()

	if err != nil {
		return nil, err
	}

	document := &openapi_v3.Document{}
	if err := proto.Unmarshal(data, document); err != nil {
		return nil, err
	}

	return document, nil
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"net/url"

	"k8s.io/kube-openapi/pkg/handler3"
)

const ContentTypeOpenAPIV3PB = "application/com.github.proto-openapi.spec.v3@v1.0+protobuf"

type GroupVersion interface {
	Schema(contentType string) ([]byte, error)
}

type groupversion struct {
	client          *client
	item            handler3.OpenAPIV3DiscoveryGroupVersion
	useClientPrefix bool
}

func newGroupVersion(client *client, item handler3.OpenAPIV3DiscoveryGroupVersion, useClientPrefix bool) *groupversion {
	return &groupversion{client: client, item: item, useClientPrefix: useClientPrefix}
}

func (g *groupversion) Schema(contentType string) ([]byte, error) {
	if !g.useClientPrefix {
		return g.client.restClient.Get().
			RequestURI(g.item.ServerRelativeURL).
			SetHeader("Accept", contentType).
			Do(context.TODO()).
			Raw()
	}

	locator, err := url.Parse(g.item.ServerRelativeURL)
	if err != nil {
		return nil, err
	}

	path := g.client.restClient.Get().
		AbsPath(locator.Path).
		SetHeader("Accept", contentType)

	// Other than root endpoints(openapiv3/apis), resources have hash query parameter to support etags.
	// However, absPath does not support handling query parameters internally,
	// so that hash query parameter is added manually
	for k, value := range locator.Query() {
		for _, v := range value {
			path.Param(k, v)
		}
	}

	return path.Do(context.TODO()).Raw()
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}
