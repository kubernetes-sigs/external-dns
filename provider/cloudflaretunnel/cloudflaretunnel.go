package cloudflaretunnel

import (
	"context"
	"fmt"
	"os"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type CloudFlareTunnelProvider struct {
	provider.BaseProvider
	Client *cloudflare.API
	// domainFilter endpoint.DomainFilter
	accountId string
	DryRun    bool
	tunnelId  string
}

func NewCloudFlareTunnelProvider(domainFilter endpoint.DomainFilter, dryRun bool) (*CloudFlareTunnelProvider, error) {
	var (
		client *cloudflare.API
		err    error
	)

	token, ok := os.LookupEnv("CF_API_TOKEN")
	if ok {
		if strings.HasPrefix(token, "file:") {
			tokenBytes, err := os.ReadFile(strings.TrimPrefix(token, "file:"))
			if err != nil {
				return nil, fmt.Errorf("failed to read CF_API_TOKEN from file: %w", err)
			}
			token = string(tokenBytes)
		}
		client, err = cloudflare.NewWithAPIToken(token)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize cloudflare provider: %v", err)
		}
	} else {
		client, err = cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
		if err != nil {
			return nil, fmt.Errorf("failed to initialize cloudflare provider: %v", err)
		}
	}

	accountId, ok := os.LookupEnv("CF_ACCOUNT_ID")
	if !ok {
		return nil, fmt.Errorf("failed to get cloudflare account id: please set env, CF_ACCOUNT_ID")
	}

	tunnelId, ok := os.LookupEnv("CF_TUNNEL_ID")
	if !ok {
		return nil, fmt.Errorf("failed to get cloudflare tunnel id: please set env, CF_TUNNEL_ID")
	}

	provider := &CloudFlareTunnelProvider{
		Client:    client,
		accountId: accountId,
		DryRun:    dryRun,
		tunnelId:  tunnelId,
	}
	return provider, nil
}

func (p *CloudFlareTunnelProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	configResult, err := p.Client.GetTunnelConfiguration(ctx, cloudflare.AccountIdentifier(p.accountId), p.tunnelId)
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}
	for _, config := range configResult.Config.Ingress {
		endpoint := endpoint.NewEndpoint(config.Hostname, "CNAME", config.Service)
		endpoints = append(endpoints, endpoint)
	}
	return endpoints, nil
}

func (p *CloudFlareTunnelProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	oldConfigResult, err := p.Client.GetTunnelConfiguration(ctx, cloudflare.AccountIdentifier(p.accountId), p.tunnelId)
	if err != nil {
		return err
	}

	param := cloudflare.TunnelConfigurationParams{TunnelID: p.tunnelId, Config: oldConfigResult.Config}
	for _, endpoint := range changes.Create {
		for _, target := range endpoint.Targets {
			param.Config.Ingress = append(param.Config.Ingress, cloudflare.UnvalidatedIngressRule{
				Hostname: endpoint.DNSName,
				Service:  target,
			})
		}
	}

	for i, desired := range changes.UpdateNew {
		current := changes.UpdateOld[i]
		add, remove, leave := provider.Difference(current.Targets, desired.Targets)

		for _, a := range remove {
			for i, v := range param.Config.Ingress {
				if v.Service == a {
					param.Config.Ingress = append(param.Config.Ingress[:i], param.Config.Ingress[i+1])
					break
				}
			}
		}

		for _, a := range add {
			param.Config.Ingress = append(param.Config.Ingress, cloudflare.UnvalidatedIngressRule{
				Hostname: desired.DNSName,
				Service:  a,
			})
		}

		for _, a := range leave {
			for i, v := range param.Config.Ingress {
				if v.Service == a {
					param.Config.Ingress[i] = cloudflare.UnvalidatedIngressRule{
						Hostname: desired.DNSName,
						Service:  a,
					}
					break
				}
			}
		}
	}

	for _, endpoint := range changes.Delete {
		for _, target := range endpoint.Targets {
			for i, v := range param.Config.Ingress {
				if v.Service == target {
					param.Config.Ingress = append(param.Config.Ingress[:i], param.Config.Ingress[i+1])
					break
				}
			}
		}
	}

	if p.DryRun {
		return nil
	}

	_, err = p.Client.UpdateTunnelConfiguration(ctx, cloudflare.AccountIdentifier(p.accountId), param)
	if err != nil {
		return err
	}
	return nil
}
