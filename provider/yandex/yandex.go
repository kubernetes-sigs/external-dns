package yandex

import (
	"context"
	"errors"
	"strings"

	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/iamkey"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	YandexAuthorizationTypeInstanceServiceAccount = "instance-service-account"
	YandexAuthorizationTypeOAuthToken             = "iam-token"
	YandexAuthorizationTypeKey                    = "iam-key-file"
)

type YandexConfig struct {
	DomainFilter            endpoint.DomainFilter
	ZoneNameFilter          endpoint.DomainFilter
	DryRun                  bool
	AuthorizationType       string
	AuthorizationOAuthToken string
	AuthorizationKeyFile    string
}

type YandexProvider struct {
	provider.BaseProvider
	sdk *ycsdk.SDK
}

func NewYandexProvider(ctx context.Context, cfg *YandexConfig) (*YandexProvider, error) {
	creds, err := cfg.credentials()
	if err != nil {
		return nil, err
	}

	sdk, err := ycsdk.Build(ctx, ycsdk.Config{
		Credentials: creds,
	})

	if err != nil {
		return nil, err
	}

	return &YandexProvider{sdk: sdk}, nil
}

func (y *YandexProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return nil, nil
}

func (y *YandexProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	return nil
}

func (cfg *YandexConfig) credentials() (ycsdk.Credentials, error) {
	auth := strings.TrimSpace(cfg.AuthorizationType)

	switch auth {
	case YandexAuthorizationTypeInstanceServiceAccount:
		return ycsdk.InstanceServiceAccount(), nil
	case YandexAuthorizationTypeOAuthToken:
		return ycsdk.OAuthToken(cfg.AuthorizationOAuthToken), nil
	case YandexAuthorizationTypeKey:
		key, err := iamkey.ReadFromJSONFile(cfg.AuthorizationKeyFile)
		if err != nil {
			return nil, err
		}
		return ycsdk.ServiceAccountKey(key)
	default:
		return nil, errors.New("unsupported authorization type")
	}
}
