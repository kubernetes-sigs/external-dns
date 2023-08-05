package configmap

import (
	"bufio"
	"context"
	"strings"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/source"
)

type ConfigMapProvider struct {
	provider.BaseProvider

	namespace    string
	name         string
	key          string
	kubeClient   kubernetes.Interface
	domainFilter endpoint.DomainFilter
	dryRun       bool
}

type host struct {
	ip       string
	hostname string
}

type hosts []*host

// NewConfigMapProvider is a factory function for configmap providers
func NewConfigMapProvider(clientGenerator source.ClientGenerator, domainFilter endpoint.DomainFilter, namespace, name, key string, dryRun bool) (provider.Provider, error) {
	kubeClient, err := clientGenerator.KubeClient()
	if err != nil {
		return nil, err
	}
	return &ConfigMapProvider{namespace: namespace, name: name, key: key, kubeClient: kubeClient, domainFilter: domainFilter, dryRun: dryRun}, nil
}

func (c *ConfigMapProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	cfg, err := c.kubeClient.CoreV1().ConfigMaps(c.namespace).Get(ctx, c.name, v1.GetOptions{})
	if err != nil {
		return err
	}
	hosts, err := parseHosts(cfg.Data[c.key])
	if err != nil {
		return err
	}
	for _, delete := range append(changes.Delete, changes.UpdateOld...) {
		for _, target := range delete.Targets {
			hosts = hosts.Delete(delete.DNSName, target)
		}
	}
	for _, create := range append(changes.Create, changes.UpdateNew...) {
		for _, target := range create.Targets {
			hosts = hosts.Create(create.DNSName, target)
		}
	}
	updateOptions := v1.UpdateOptions{}
	if c.dryRun {
		updateOptions.DryRun = []string{v1.DryRunAll}
	}
	_, err = c.kubeClient.CoreV1().ConfigMaps(c.namespace).Update(ctx, cfg, updateOptions)
	return err
}

func (c *ConfigMapProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	cfg, err := c.kubeClient.CoreV1().ConfigMaps(c.namespace).Get(ctx, c.name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	hosts, err := parseHosts(cfg.Data[c.key])
	if err != nil {
		return nil, err
	}
	endpoints := make([]*endpoint.Endpoint, len(hosts))
	for idx, host := range hosts {
		endpoints[idx] = endpoint.NewEndpoint(host.hostname, "A", host.ip)
	}
	return endpoints, nil
}

func parseHosts(raw string) (hosts, error) {
	hosts := hosts{}
	scanner := bufio.NewScanner(strings.NewReader(raw))
	for scanner.Scan() {
		pair := strings.Fields(scanner.Text())
		if len(pair) != 2 {
			log.Warnf("ConfigMap: '%s' is an invalid host line", scanner.Text())
			continue
		}
		hosts = append(hosts, &host{ip: pair[0], hostname: pair[1]})
	}
	return hosts, scanner.Err()
}

func (h hosts) Create(hostname, ip string) hosts {
	return append(h, &host{hostname: hostname, ip: ip})
}

func (h hosts) Delete(hostname, ip string) hosts {
	newHosts := hosts{}
	for _, host := range h {
		if host.hostname != hostname && host.ip != ip {
			newHosts = append(newHosts, host)
		}
	}
	return newHosts
}
