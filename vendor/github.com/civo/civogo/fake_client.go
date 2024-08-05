package civogo

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// FakeClient is a temporary storage structure for use when you don't want to communicate with a real Civo API server
type FakeClient struct {
	LastID                  int64
	Charges                 []Charge
	Domains                 []DNSDomain
	DomainRecords           []DNSRecord
	Firewalls               []Firewall
	FirewallRules           []FirewallRule
	InstanceSizes           []InstanceSize
	Instances               []Instance
	Clusters                []KubernetesCluster
	IP                      []IP
	Networks                []Network
	Volumes                 []Volume
	SSHKeys                 []SSHKey
	Webhooks                []Webhook
	DiskImage               []DiskImage
	Quota                   Quota
	Organisation            Organisation
	OrganisationAccounts    []Account
	OrganisationRoles       []Role
	OrganisationTeams       []Team
	OrganisationTeamMembers map[string][]TeamMember
	LoadBalancers           []LoadBalancer
	Pools                   []KubernetesPool
	PingErr                 error
	// Snapshots            []Snapshot
	// Templates            []Template
}

// Clienter is the interface the real civogo.Client and civogo.FakeClient implement
type Clienter interface {
	// Charges
	ListCharges(from, to time.Time) ([]Charge, error)

	// DNS
	ListDNSDomains() ([]DNSDomain, error)
	FindDNSDomain(search string) (*DNSDomain, error)
	CreateDNSDomain(name string) (*DNSDomain, error)
	GetDNSDomain(name string) (*DNSDomain, error)
	UpdateDNSDomain(d *DNSDomain, name string) (*DNSDomain, error)
	DeleteDNSDomain(d *DNSDomain) (*SimpleResponse, error)
	CreateDNSRecord(domainID string, r *DNSRecordConfig) (*DNSRecord, error)
	ListDNSRecords(dnsDomainID string) ([]DNSRecord, error)
	GetDNSRecord(domainID, domainRecordID string) (*DNSRecord, error)
	UpdateDNSRecord(r *DNSRecord, rc *DNSRecordConfig) (*DNSRecord, error)
	DeleteDNSRecord(r *DNSRecord) (*SimpleResponse, error)

	// Firewalls
	ListFirewalls() ([]Firewall, error)
	FindFirewall(search string) (*Firewall, error)
	NewFirewall(*FirewallConfig) (*FirewallResult, error)
	RenameFirewall(id string, f *FirewallConfig) (*SimpleResponse, error)
	DeleteFirewall(id string) (*SimpleResponse, error)
	NewFirewallRule(r *FirewallRuleConfig) (*FirewallRule, error)
	ListFirewallRules(id string) ([]FirewallRule, error)
	FindFirewallRule(firewallID string, search string) (*FirewallRule, error)
	DeleteFirewallRule(id string, ruleID string) (*SimpleResponse, error)

	// Instances
	ListInstances(page int, perPage int) (*PaginatedInstanceList, error)
	ListAllInstances() ([]Instance, error)
	FindInstance(search string) (*Instance, error)
	GetInstance(id string) (*Instance, error)
	NewInstanceConfig() (*InstanceConfig, error)
	CreateInstance(config *InstanceConfig) (*Instance, error)
	SetInstanceTags(i *Instance, tags string) (*SimpleResponse, error)
	UpdateInstance(i *Instance) (*SimpleResponse, error)
	DeleteInstance(id string) (*SimpleResponse, error)
	RebootInstance(id string) (*SimpleResponse, error)
	HardRebootInstance(id string) (*SimpleResponse, error)
	SoftRebootInstance(id string) (*SimpleResponse, error)
	StopInstance(id string) (*SimpleResponse, error)
	StartInstance(id string) (*SimpleResponse, error)
	GetInstanceConsoleURL(id string) (string, error)
	UpgradeInstance(id, newSize string) (*SimpleResponse, error)
	MovePublicIPToInstance(id, ipAddress string) (*SimpleResponse, error)
	SetInstanceFirewall(id, firewallID string) (*SimpleResponse, error)

	// Instance sizes
	ListInstanceSizes() ([]InstanceSize, error)
	FindInstanceSizes(search string) (*InstanceSize, error)

	// Clusters
	ListKubernetesClusters() (*PaginatedKubernetesClusters, error)
	FindKubernetesCluster(search string) (*KubernetesCluster, error)
	NewKubernetesClusters(kc *KubernetesClusterConfig) (*KubernetesCluster, error)
	GetKubernetesCluster(id string) (*KubernetesCluster, error)
	UpdateKubernetesCluster(id string, i *KubernetesClusterConfig) (*KubernetesCluster, error)
	ListKubernetesMarketplaceApplications() ([]KubernetesMarketplaceApplication, error)
	DeleteKubernetesCluster(id string) (*SimpleResponse, error)
	RecycleKubernetesCluster(id string, hostname string) (*SimpleResponse, error)
	ListAvailableKubernetesVersions() ([]KubernetesVersion, error)
	ListKubernetesClusterInstances(id string) ([]Instance, error)
	FindKubernetesClusterInstance(clusterID, search string) (*Instance, error)

	//Pools
	ListKubernetesClusterPools(cid string) ([]KubernetesPool, error)
	GetKubernetesClusterPool(cid, pid string) (*KubernetesPool, error)
	FindKubernetesClusterPool(cid, search string) (*KubernetesPool, error)
	DeleteKubernetesClusterPoolInstance(cid, pid, id string) (*SimpleResponse, error)
	UpdateKubernetesClusterPool(cid, pid string, config *KubernetesClusterPoolUpdateConfig) (*KubernetesPool, error)

	// Networks
	GetDefaultNetwork() (*Network, error)
	NewNetwork(label string) (*NetworkResult, error)
	CreateNetwork(configs NetworkConfig) (*NetworkResult, error)
	ListNetworks() ([]Network, error)
	FindNetwork(search string) (*Network, error)
	RenameNetwork(label, id string) (*NetworkResult, error)
	DeleteNetwork(id string) (*SimpleResponse, error)

	// Quota
	GetQuota() (*Quota, error)

	// Regions
	ListRegions() ([]Region, error)

	// Snapshots
	// CreateSnapshot(name string, r *SnapshotConfig) (*Snapshot, error)
	// ListSnapshots() ([]Snapshot, error)
	// FindSnapshot(search string) (*Snapshot, error)
	// DeleteSnapshot(name string) (*SimpleResponse, error)

	// SSHKeys
	ListSSHKeys() ([]SSHKey, error)
	NewSSHKey(name string, publicKey string) (*SimpleResponse, error)
	UpdateSSHKey(name string, sshKeyID string) (*SSHKey, error)
	FindSSHKey(search string) (*SSHKey, error)
	DeleteSSHKey(id string) (*SimpleResponse, error)

	// Templates
	// ListTemplates() ([]Template, error)
	// NewTemplate(conf *Template) (*SimpleResponse, error)
	// UpdateTemplate(id string, conf *Template) (*Template, error)
	// GetTemplateByCode(code string) (*Template, error)
	// FindTemplate(search string) (*Template, error)
	// DeleteTemplate(id string) (*SimpleResponse, error)

	// DiskImages
	ListDiskImages() ([]DiskImage, error)
	GetDiskImage(id string) (*DiskImage, error)
	FindDiskImage(search string) (*DiskImage, error)

	// Volumes
	ListVolumes() ([]Volume, error)
	GetVolume(id string) (*Volume, error)
	FindVolume(search string) (*Volume, error)
	NewVolume(v *VolumeConfig) (*VolumeResult, error)
	ResizeVolume(id string, size int) (*SimpleResponse, error)
	AttachVolume(id string, instance string) (*SimpleResponse, error)
	DetachVolume(id string) (*SimpleResponse, error)
	DeleteVolume(id string) (*SimpleResponse, error)

	// Webhooks
	CreateWebhook(r *WebhookConfig) (*Webhook, error)
	ListWebhooks() ([]Webhook, error)
	FindWebhook(search string) (*Webhook, error)
	UpdateWebhook(id string, r *WebhookConfig) (*Webhook, error)
	DeleteWebhook(id string) (*SimpleResponse, error)

	// Reserved IPs
	ListIPs() (*PaginatedIPs, error)
	FindIP(search string) (*IP, error)
	GetIP(id string) (*IP, error)
	NewIP(v *CreateIPRequest) (*IP, error)
	UpdateIP(id string, v *UpdateIPRequest) (*IP, error)
	DeleteIP(id string) (*SimpleResponse, error)
	AssignIP(id, resourceID, resourceType, region string) (*SimpleResponse, error)
	UnassignIP(id, region string) (*SimpleResponse, error)

	// LoadBalancer
	ListLoadBalancers() ([]LoadBalancer, error)
	GetLoadBalancer(id string) (*LoadBalancer, error)
	FindLoadBalancer(search string) (*LoadBalancer, error)
	CreateLoadBalancer(r *LoadBalancerConfig) (*LoadBalancer, error)
	UpdateLoadBalancer(id string, r *LoadBalancerUpdateConfig) (*LoadBalancer, error)
	DeleteLoadBalancer(id string) (*SimpleResponse, error)

	// Ping
	Ping() error
}

// NewFakeClient initializes a Client that doesn't attach to a
func NewFakeClient() (*FakeClient, error) {
	return &FakeClient{
		Quota: Quota{
			CPUCoreLimit:           10,
			InstanceCountLimit:     10,
			RAMMegabytesLimit:      100,
			DiskGigabytesLimit:     100,
			DiskVolumeCountLimit:   10,
			DiskSnapshotCountLimit: 10,
			PublicIPAddressLimit:   10,
			NetworkCountLimit:      10,
			SecurityGroupLimit:     10,
			SecurityGroupRuleLimit: 10,
		},
		InstanceSizes: []InstanceSize{
			{
				Name:          "g3.xsmall",
				CPUCores:      1,
				RAMMegabytes:  1024,
				DiskGigabytes: 10,
			},
			{
				Name:          "g3.small",
				CPUCores:      2,
				RAMMegabytes:  2048,
				DiskGigabytes: 20,
			},
			{
				Name:          "g3.medium",
				CPUCores:      4,
				RAMMegabytes:  4096,
				DiskGigabytes: 40,
			},
		},
		DiskImage: []DiskImage{
			{
				ID:           "b82168fe-66f6-4b38-a3b8-5283542d5475",
				Name:         "centos-7",
				Version:      "7",
				State:        "available",
				Distribution: "centos",
				Description:  "",
				Label:        "",
			},
			{
				ID:           "b82168fe-66f6-4b38-a3b8-52835425895",
				Name:         "debian-9",
				Version:      "9",
				State:        "available",
				Distribution: "debian",
				Description:  "",
				Label:        "",
			},
			{
				ID:           "b82168fe-66f6-4b38-a3b8-52835428965",
				Name:         "debian-10",
				Version:      "10",
				State:        "available",
				Distribution: "debian",
				Description:  "",
				Label:        "",
			},
			{
				ID:           "b82168fe-66f6-4b38-a3b8-528354282548",
				Name:         "ubuntu-20-4",
				Version:      "20.4",
				State:        "available",
				Distribution: "ubuntu",
				Description:  "",
				Label:        "",
			},
		},
	}, nil
}

// Ping implemented in a fake way for automated tests
func (c *FakeClient) Ping() error {
	if c.PingErr != nil {
		return c.PingErr
	}
	return nil
}

func (c *FakeClient) generateID() string {
	c.LastID++
	return strconv.FormatInt(c.LastID, 10)
}

func (c *FakeClient) generatePublicIP() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return fmt.Sprintf("%v.%v.%v.%v", r.Intn(256), r.Intn(256), r.Intn(256), r.Intn(256))
}

// ListCharges implemented in a fake way for automated tests
func (c *FakeClient) ListCharges(from, to time.Time) ([]Charge, error) {
	return []Charge{}, nil
}

// ListDNSDomains implemented in a fake way for automated tests
func (c *FakeClient) ListDNSDomains() ([]DNSDomain, error) {
	return c.Domains, nil
}

// FindDNSDomain implemented in a fake way for automated tests
func (c *FakeClient) FindDNSDomain(search string) (*DNSDomain, error) {
	for _, domain := range c.Domains {
		if strings.Contains(domain.Name, search) {
			return &domain, nil
		}
	}

	err := fmt.Errorf("unable to find %s, zero matches", search)
	return nil, ZeroMatchesError.wrap(err)
}

// CreateDNSDomain implemented in a fake way for automated tests
func (c *FakeClient) CreateDNSDomain(name string) (*DNSDomain, error) {
	domain := DNSDomain{
		ID:   c.generateID(),
		Name: name,
	}
	c.Domains = append(c.Domains, domain)
	return &domain, nil
}

// GetDNSDomain implemented in a fake way for automated tests
func (c *FakeClient) GetDNSDomain(name string) (*DNSDomain, error) {
	for _, domain := range c.Domains {
		if domain.Name == name {
			return &domain, nil
		}
	}

	return nil, ErrDNSDomainNotFound
}

// UpdateDNSDomain implemented in a fake way for automated tests
func (c *FakeClient) UpdateDNSDomain(d *DNSDomain, name string) (*DNSDomain, error) {
	for i, domain := range c.Domains {
		if domain.Name == d.Name {
			c.Domains[i] = *d
			return d, nil
		}
	}

	return nil, ErrDNSDomainNotFound
}

// DeleteDNSDomain implemented in a fake way for automated tests
func (c *FakeClient) DeleteDNSDomain(d *DNSDomain) (*SimpleResponse, error) {
	for i, domain := range c.Domains {
		if domain.Name == d.Name {
			c.Domains[len(c.Domains)-1], c.Domains[i] = c.Domains[i], c.Domains[len(c.Domains)-1]
			c.Domains = c.Domains[:len(c.Domains)-1]
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return nil, ErrDNSDomainNotFound
}

// CreateDNSRecord implemented in a fake way for automated tests
func (c *FakeClient) CreateDNSRecord(domainID string, r *DNSRecordConfig) (*DNSRecord, error) {
	record := DNSRecord{
		ID:          c.generateID(),
		DNSDomainID: domainID,
		Name:        r.Name,
		Value:       r.Value,
		Type:        r.Type,
	}

	c.DomainRecords = append(c.DomainRecords, record)
	return &record, nil
}

// ListDNSRecords implemented in a fake way for automated tests
func (c *FakeClient) ListDNSRecords(dnsDomainID string) ([]DNSRecord, error) {
	return c.DomainRecords, nil
}

// GetDNSRecord implemented in a fake way for automated tests
func (c *FakeClient) GetDNSRecord(domainID, domainRecordID string) (*DNSRecord, error) {
	for _, record := range c.DomainRecords {
		if record.ID == domainRecordID && record.DNSDomainID == domainID {
			return &record, nil
		}
	}

	return nil, ErrDNSRecordNotFound
}

// UpdateDNSRecord implemented in a fake way for automated tests
func (c *FakeClient) UpdateDNSRecord(r *DNSRecord, rc *DNSRecordConfig) (*DNSRecord, error) {
	for i, record := range c.DomainRecords {
		if record.ID == r.ID {
			record := DNSRecord{
				ID:          c.generateID(),
				DNSDomainID: record.DNSDomainID,
				Name:        rc.Name,
				Value:       rc.Value,
				Type:        rc.Type,
			}

			c.DomainRecords[i] = record
			return &record, nil
		}
	}

	return nil, ErrDNSRecordNotFound
}

// DeleteDNSRecord implemented in a fake way for automated tests
func (c *FakeClient) DeleteDNSRecord(r *DNSRecord) (*SimpleResponse, error) {
	for i, record := range c.DomainRecords {
		if record.ID == r.ID {
			c.DomainRecords[len(c.DomainRecords)-1], c.DomainRecords[i] = c.DomainRecords[i], c.DomainRecords[len(c.DomainRecords)-1]
			c.DomainRecords = c.DomainRecords[:len(c.DomainRecords)-1]
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return nil, ErrDNSRecordNotFound
}

// ListFirewalls implemented in a fake way for automated tests
func (c *FakeClient) ListFirewalls() ([]Firewall, error) {
	return c.Firewalls, nil
}

// FindFirewall implemented in a fake way for automated tests
func (c *FakeClient) FindFirewall(search string) (*Firewall, error) {
	for _, firewall := range c.Firewalls {
		if strings.Contains(firewall.Name, search) {
			return &firewall, nil
		}
	}

	err := fmt.Errorf("unable to find %s, zero matches", search)
	return nil, ZeroMatchesError.wrap(err)
}

// NewFirewall implemented in a fake way for automated tests
func (c *FakeClient) NewFirewall(*FirewallConfig) (*FirewallResult, error) {
	firewall := Firewall{
		ID:   c.generateID(),
		Name: "fw-name",
	}
	c.Firewalls = append(c.Firewalls, firewall)

	return &FirewallResult{
		ID:     firewall.ID,
		Name:   firewall.Name,
		Result: "success",
	}, nil
}

// RenameFirewall implemented in a fake way for automated tests
func (c *FakeClient) RenameFirewall(id string, f *FirewallConfig) (*SimpleResponse, error) {
	for i, firewall := range c.Firewalls {
		if firewall.ID == id {
			c.Firewalls[i].Name = f.Name
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	err := fmt.Errorf("unable to find %s, zero matches", id)
	return nil, ZeroMatchesError.wrap(err)
}

// DeleteFirewall implemented in a fake way for automated tests
func (c *FakeClient) DeleteFirewall(id string) (*SimpleResponse, error) {
	for i, firewall := range c.Firewalls {
		if firewall.ID == id {
			c.Firewalls[len(c.Firewalls)-1], c.Firewalls[i] = c.Firewalls[i], c.Firewalls[len(c.Firewalls)-1]
			c.Firewalls = c.Firewalls[:len(c.Firewalls)-1]
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return &SimpleResponse{Result: "failed"}, nil
}

// NewFirewallRule implemented in a fake way for automated tests
func (c *FakeClient) NewFirewallRule(r *FirewallRuleConfig) (*FirewallRule, error) {
	rule := FirewallRule{
		ID:        c.generateID(),
		Protocol:  r.Protocol,
		StartPort: r.StartPort,
		EndPort:   r.EndPort,
		Cidr:      r.Cidr,
		Label:     r.Label,
	}
	c.FirewallRules = append(c.FirewallRules, rule)
	return &rule, nil
}

// ListFirewallRules implemented in a fake way for automated tests
func (c *FakeClient) ListFirewallRules(id string) ([]FirewallRule, error) {
	return c.FirewallRules, nil
}

// FindFirewallRule implemented in a fake way for automated tests
func (c *FakeClient) FindFirewallRule(firewallID string, search string) (*FirewallRule, error) {
	for _, rule := range c.FirewallRules {
		if rule.FirewallID == firewallID && strings.Contains(rule.Label, search) {
			return &rule, nil
		}
	}

	err := fmt.Errorf("unable to find %s, zero matches", search)
	return nil, ZeroMatchesError.wrap(err)
}

// DeleteFirewallRule implemented in a fake way for automated tests
func (c *FakeClient) DeleteFirewallRule(id string, ruleID string) (*SimpleResponse, error) {
	for i, rule := range c.FirewallRules {
		if rule.ID == ruleID {
			c.FirewallRules[len(c.FirewallRules)-1], c.FirewallRules[i] = c.FirewallRules[i], c.FirewallRules[len(c.FirewallRules)-1]
			c.FirewallRules = c.FirewallRules[:len(c.FirewallRules)-1]
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return &SimpleResponse{Result: "failed"}, nil
}

// ListInstances implemented in a fake way for automated tests
func (c *FakeClient) ListInstances(page int, perPage int) (*PaginatedInstanceList, error) {
	return &PaginatedInstanceList{
		Items:   c.Instances,
		Page:    page,
		PerPage: perPage,
		Pages:   page,
	}, nil
}

// ListAllInstances implemented in a fake way for automated tests
func (c *FakeClient) ListAllInstances() ([]Instance, error) {
	return c.Instances, nil
}

// FindInstance implemented in a fake way for automated tests
func (c *FakeClient) FindInstance(search string) (*Instance, error) {
	for _, instance := range c.Instances {
		if strings.Contains(instance.Hostname, search) {
			return &instance, nil
		}
	}

	err := fmt.Errorf("unable to find %s, zero matches", search)
	return nil, ZeroMatchesError.wrap(err)
}

// GetInstance implemented in a fake way for automated tests
func (c *FakeClient) GetInstance(id string) (*Instance, error) {
	for _, instance := range c.Instances {
		if instance.ID == id {
			return &instance, nil
		}
	}

	err := fmt.Errorf("unable to find %s, zero matches", id)
	return nil, ZeroMatchesError.wrap(err)
}

// NewInstanceConfig implemented in a fake way for automated tests
func (c *FakeClient) NewInstanceConfig() (*InstanceConfig, error) {
	return &InstanceConfig{}, nil
}

// CreateInstance implemented in a fake way for automated tests
func (c *FakeClient) CreateInstance(config *InstanceConfig) (*Instance, error) {
	instance := Instance{
		ID:          c.generateID(),
		Hostname:    config.Hostname,
		Size:        config.Size,
		Region:      config.Region,
		TemplateID:  config.TemplateID,
		InitialUser: config.InitialUser,
		SSHKey:      config.SSHKeyID,
		Tags:        config.Tags,
		PublicIP:    c.generatePublicIP(),
	}
	c.Instances = append(c.Instances, instance)
	return &instance, nil
}

// SetInstanceTags implemented in a fake way for automated tests
func (c *FakeClient) SetInstanceTags(i *Instance, tags string) (*SimpleResponse, error) {
	for idx, instance := range c.Instances {
		if instance.ID == i.ID {
			c.Instances[idx].Tags = strings.Split(tags, " ")
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return &SimpleResponse{Result: "failed"}, nil
}

// UpdateInstance implemented in a fake way for automated tests
func (c *FakeClient) UpdateInstance(i *Instance) (*SimpleResponse, error) {
	for idx, instance := range c.Instances {
		if instance.ID == i.ID {
			c.Instances[idx] = *i
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return &SimpleResponse{Result: "failed"}, nil
}

// DeleteInstance implemented in a fake way for automated tests
func (c *FakeClient) DeleteInstance(id string) (*SimpleResponse, error) {
	for i, instance := range c.Instances {
		if instance.ID == id {
			c.Instances[len(c.Instances)-1], c.Instances[i] = c.Instances[i], c.Instances[len(c.Instances)-1]
			c.Instances = c.Instances[:len(c.Instances)-1]
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return &SimpleResponse{Result: "failed"}, nil
}

// RebootInstance implemented in a fake way for automated tests
func (c *FakeClient) RebootInstance(id string) (*SimpleResponse, error) {
	return &SimpleResponse{Result: "success"}, nil
}

// HardRebootInstance implemented in a fake way for automated tests
func (c *FakeClient) HardRebootInstance(id string) (*SimpleResponse, error) {
	return &SimpleResponse{Result: "success"}, nil
}

// SoftRebootInstance implemented in a fake way for automated tests
func (c *FakeClient) SoftRebootInstance(id string) (*SimpleResponse, error) {
	return &SimpleResponse{Result: "success"}, nil
}

// StopInstance implemented in a fake way for automated tests
func (c *FakeClient) StopInstance(id string) (*SimpleResponse, error) {
	return &SimpleResponse{Result: "success"}, nil
}

// StartInstance implemented in a fake way for automated tests
func (c *FakeClient) StartInstance(id string) (*SimpleResponse, error) {
	return &SimpleResponse{Result: "success"}, nil
}

// GetInstanceConsoleURL implemented in a fake way for automated tests
func (c *FakeClient) GetInstanceConsoleURL(id string) (string, error) {
	return fmt.Sprintf("https://console.example.com/%s", id), nil
}

// UpgradeInstance implemented in a fake way for automated tests
func (c *FakeClient) UpgradeInstance(id, newSize string) (*SimpleResponse, error) {
	for idx, instance := range c.Instances {
		if instance.ID == id {
			c.Instances[idx].Size = newSize
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return &SimpleResponse{Result: "failed"}, nil
}

// MovePublicIPToInstance implemented in a fake way for automated tests
func (c *FakeClient) MovePublicIPToInstance(id, ipAddress string) (*SimpleResponse, error) {
	oldIndex := -1
	for idx, instance := range c.Instances {
		if instance.PublicIP == ipAddress {
			oldIndex = idx
		}
	}

	newIndex := -1
	for idx, instance := range c.Instances {
		if instance.ID == id {
			newIndex = idx
		}
	}

	if oldIndex == -1 || newIndex == -1 {
		return &SimpleResponse{Result: "failed"}, nil
	}

	c.Instances[newIndex].PublicIP = c.Instances[oldIndex].PublicIP
	c.Instances[oldIndex].PublicIP = ""

	return &SimpleResponse{Result: "success"}, nil
}

// SetInstanceFirewall implemented in a fake way for automated tests
func (c *FakeClient) SetInstanceFirewall(id, firewallID string) (*SimpleResponse, error) {
	for idx, instance := range c.Instances {
		if instance.ID == id {
			c.Instances[idx].FirewallID = firewallID
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return &SimpleResponse{Result: "failed"}, nil
}

// ListInstanceSizes implemented in a fake way for automated tests
func (c *FakeClient) ListInstanceSizes() ([]InstanceSize, error) {
	return c.InstanceSizes, nil
}

// FindInstanceSizes implemented in a fake way for automated tests
func (c *FakeClient) FindInstanceSizes(search string) (*InstanceSize, error) {
	for _, size := range c.InstanceSizes {
		if strings.Contains(size.Name, search) {
			return &size, nil
		}
	}

	err := fmt.Errorf("unable to find %s, zero matches", search)
	return nil, ZeroMatchesError.wrap(err)
}

// ListKubernetesClusters implemented in a fake way for automated tests
func (c *FakeClient) ListKubernetesClusters() (*PaginatedKubernetesClusters, error) {
	return &PaginatedKubernetesClusters{
		Items:   c.Clusters,
		Page:    1,
		PerPage: 10,
		Pages:   1,
	}, nil
}

// FindKubernetesCluster implemented in a fake way for automated tests
func (c *FakeClient) FindKubernetesCluster(search string) (*KubernetesCluster, error) {
	for _, cluster := range c.Clusters {
		if strings.Contains(cluster.Name, search) || cluster.ID == search {
			return &cluster, nil
		}
	}

	err := fmt.Errorf("unable to find %s, zero matches", search)
	return nil, ZeroMatchesError.wrap(err)
}

// ListKubernetesClusterInstances implemented in a fake way for automated tests
func (c *FakeClient) ListKubernetesClusterInstances(id string) ([]Instance, error) {
	for _, cluster := range c.Clusters {
		if cluster.ID == id {
			instaces := make([]Instance, 0)
			for _, kins := range cluster.Instances {
				for _, instance := range c.Instances {
					if instance.ID == kins.ID {
						instaces = append(instaces, instance)
					}
				}
			}
			return instaces, nil
		}
	}

	err := fmt.Errorf("unable to find %s, zero matches", id)
	return nil, DatabaseKubernetesClusterNotFoundError.wrap(err)
}

// FindKubernetesClusterInstance implemented in a fake way for automated tests
func (c *FakeClient) FindKubernetesClusterInstance(clusterID, search string) (*Instance, error) {
	instances, err := c.ListKubernetesClusterInstances(clusterID)
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := Instance{}

	for _, instance := range instances {
		if instance.Hostname == search || instance.ID == search {
			exactMatch = true
			result = instance
		} else if strings.Contains(instance.Hostname, search) || strings.Contains(instance.ID, search) {
			if !exactMatch {
				result = instance
				partialMatchesCount++
			}
		}
	}

	if exactMatch || partialMatchesCount == 1 {
		return &result, nil
	} else if partialMatchesCount > 1 {
		err := fmt.Errorf("unable to find %s because there were multiple matches", search)
		return nil, MultipleMatchesError.wrap(err)
	} else {
		err := fmt.Errorf("unable to find %s, zero matches", search)
		return nil, ZeroMatchesError.wrap(err)
	}
}

// NewKubernetesClusters implemented in a fake way for automated tests
func (c *FakeClient) NewKubernetesClusters(kc *KubernetesClusterConfig) (*KubernetesCluster, error) {
	cluster := KubernetesCluster{
		ID:             c.generateID(),
		Name:           kc.Name,
		MasterIP:       c.generatePublicIP(),
		NumTargetNode:  kc.NumTargetNodes,
		TargetNodeSize: kc.TargetNodesSize,
		Ready:          true,
		Status:         "ACTIVE",
		Instances:      make([]KubernetesInstance, 0),
		Pools:          make([]KubernetesPool, 0),
	}
	pool := KubernetesPool{
		Instances: make([]KubernetesInstance, 0),
	}
	for i := 0; i < kc.NumTargetNodes; i++ {
		instance := KubernetesInstance{
			ID:       c.generateID(),
			Hostname: fmt.Sprintf("%s_pool_%d", kc.Name, i),
		}
		pool.Instances = append(pool.Instances, instance)
		cluster.Instances = append(pool.Instances, instance)
	}

	cluster.Pools = append(cluster.Pools, pool)
	c.Clusters = append(c.Clusters, cluster)
	return &cluster, nil
}

// GetKubernetesCluster implemented in a fake way for automated tests
func (c *FakeClient) GetKubernetesCluster(id string) (*KubernetesCluster, error) {
	for _, cluster := range c.Clusters {
		if cluster.ID == id {
			return &cluster, nil
		}
	}

	err := fmt.Errorf("unable to find %s, zero matches", id)
	return nil, ZeroMatchesError.wrap(err)
}

// UpdateKubernetesCluster implemented in a fake way for automated tests
func (c *FakeClient) UpdateKubernetesCluster(id string, kc *KubernetesClusterConfig) (*KubernetesCluster, error) {
	for i, cluster := range c.Clusters {
		if cluster.ID == id {
			c.Clusters[i].Name = kc.Name
			c.Clusters[i].NumTargetNode = kc.NumTargetNodes
			c.Clusters[i].TargetNodeSize = kc.TargetNodesSize
			return &cluster, nil
		}
	}

	err := fmt.Errorf("unable to find %s, zero matches", id)
	return nil, ZeroMatchesError.wrap(err)
}

// ListKubernetesMarketplaceApplications implemented in a fake way for automated tests
func (c *FakeClient) ListKubernetesMarketplaceApplications() ([]KubernetesMarketplaceApplication, error) {
	return []KubernetesMarketplaceApplication{}, nil
}

// DeleteKubernetesCluster implemented in a fake way for automated tests
func (c *FakeClient) DeleteKubernetesCluster(id string) (*SimpleResponse, error) {
	for i, cluster := range c.Clusters {
		if cluster.ID == id {
			c.Clusters[len(c.Clusters)-1], c.Clusters[i] = c.Clusters[i], c.Clusters[len(c.Clusters)-1]
			c.Clusters = c.Clusters[:len(c.Clusters)-1]
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return &SimpleResponse{Result: "failed"}, nil
}

// RecycleKubernetesCluster implemented in a fake way for automated tests
func (c *FakeClient) RecycleKubernetesCluster(id string, hostname string) (*SimpleResponse, error) {
	return &SimpleResponse{Result: "success"}, nil
}

// ListAvailableKubernetesVersions implemented in a fake way for automated tests
func (c *FakeClient) ListAvailableKubernetesVersions() ([]KubernetesVersion, error) {
	return []KubernetesVersion{
		{
			Version: "1.20+k3s1",
			Type:    "stable",
		},
	}, nil
}

// GetDefaultNetwork implemented in a fake way for automated tests
func (c *FakeClient) GetDefaultNetwork() (*Network, error) {
	for _, network := range c.Networks {
		if network.Default {
			return &network, nil
		}
	}

	err := fmt.Errorf("unable to find default network, zero matches")
	return nil, ZeroMatchesError.wrap(err)
}

// NewNetwork implemented in a fake way for automated tests
func (c *FakeClient) NewNetwork(label string) (*NetworkResult, error) {
	network := Network{
		ID:   c.generateID(),
		Name: label,
	}
	c.Networks = append(c.Networks, network)

	return &NetworkResult{
		ID:     network.ID,
		Label:  network.Name,
		Result: "success",
	}, nil

}

// CreateNetwork creates a new network within the FakeClient, including VLAN configurations
func (c *FakeClient) CreateNetwork(config NetworkConfig) (*NetworkResult, error) {
	networkID := c.generateID()

	// Prepare the new Network object
	newNetwork := Network{
		ID:          networkID,
		Name:        config.Label,
		Default:     config.Default == "true",
		CIDR:        config.CIDRv4,
		Label:       config.Label,
		IPv4Enabled: config.IPv4Enabled != nil && *config.IPv4Enabled,
	}

	// Handle VLAN configuration if present
	if config.VLanConfig != nil {
		newNetwork.VlanID = config.VLanConfig.VlanID
		newNetwork.HardwareAddr = config.VLanConfig.HardwareAddr
		newNetwork.GatewayIPv4 = config.VLanConfig.GatewayIPv4
		newNetwork.AllocationPoolV4Start = config.VLanConfig.AllocationPoolV4Start
		newNetwork.AllocationPoolV4Start = config.VLanConfig.AllocationPoolV4End
	}

	// Append the newly created network to the networks slice
	c.Networks = append(c.Networks, newNetwork)

	// Return a success result
	return &NetworkResult{
		ID:     networkID,
		Label:  config.Label,
		Result: "success",
	}, nil
}

// ListNetworks implemented in a fake way for automated tests
func (c *FakeClient) ListNetworks() ([]Network, error) {
	return c.Networks, nil
}

// FindNetwork implemented in a fake way for automated tests
func (c *FakeClient) FindNetwork(search string) (*Network, error) {
	for _, network := range c.Networks {
		if strings.Contains(network.Name, search) {
			return &network, nil
		}
	}

	err := fmt.Errorf("unable to find default network, zero matches")
	return nil, ZeroMatchesError.wrap(err)
}

// RenameNetwork implemented in a fake way for automated tests
func (c *FakeClient) RenameNetwork(label, id string) (*NetworkResult, error) {
	for i, network := range c.Networks {
		if network.ID == id {
			c.Networks[i].Label = label
			return &NetworkResult{
				ID:     network.ID,
				Label:  network.Label,
				Result: "success",
			}, nil
		}
	}

	err := fmt.Errorf("unable to find default network, zero matches")
	return nil, ZeroMatchesError.wrap(err)
}

// DeleteNetwork implemented in a fake way for automated tests
func (c *FakeClient) DeleteNetwork(id string) (*SimpleResponse, error) {
	for i, network := range c.Networks {
		if network.ID == id {
			c.Networks[len(c.Networks)-1], c.Networks[i] = c.Networks[i], c.Networks[len(c.Networks)-1]
			c.Networks = c.Networks[:len(c.Networks)-1]
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return &SimpleResponse{Result: "failed"}, nil
}

// GetQuota implemented in a fake way for automated tests
func (c *FakeClient) GetQuota() (*Quota, error) {
	return &c.Quota, nil
}

// ListRegions implemented in a fake way for automated tests
func (c *FakeClient) ListRegions() ([]Region, error) {
	return []Region{
		{
			Code:    "FAKE1",
			Name:    "Fake testing region",
			Default: true,
		},
	}, nil
}

// CreateSnapshot implemented in a fake way for automated tests
// func (c *FakeClient) CreateSnapshot(name string, r *SnapshotConfig) (*Snapshot, error) {
// 	snapshot := Snapshot{
// 		ID:         c.generateID(),
// 		Name:       name,
// 		InstanceID: r.InstanceID,
// 		Cron:       r.Cron,
// 	}
// 	c.Snapshots = append(c.Snapshots, snapshot)

// 	return &snapshot, nil
// }

// ListSnapshots implemented in a fake way for automated tests
// func (c *FakeClient) ListSnapshots() ([]Snapshot, error) {
// 	return c.Snapshots, nil
// }

// FindSnapshot implemented in a fake way for automated tests
// func (c *FakeClient) FindSnapshot(search string) (*Snapshot, error) {
// 	for _, snapshot := range c.Snapshots {
// 		if strings.Contains(snapshot.Name, search) {
// 			return &snapshot, nil
// 		}
// 	}

// 	err := fmt.Errorf("unable to find %s, zero matches", search)
// 	return nil, ZeroMatchesError.wrap(err)
// }

// DeleteSnapshot implemented in a fake way for automated tests
// func (c *FakeClient) DeleteSnapshot(name string) (*SimpleResponse, error) {
// 	for i, snapshot := range c.Snapshots {
// 		if snapshot.Name == name {
// 			c.Snapshots[len(c.Snapshots)-1], c.Snapshots[i] = c.Snapshots[i], c.Snapshots[len(c.Snapshots)-1]
// 			c.Snapshots = c.Snapshots[:len(c.Snapshots)-1]
// 			return &SimpleResponse{Result: "success"}, nil
// 		}
// 	}

// 	return &SimpleResponse{Result: "failed"}, nil
// }

// ListSSHKeys implemented in a fake way for automated tests
func (c *FakeClient) ListSSHKeys() ([]SSHKey, error) {
	return c.SSHKeys, nil
}

// NewSSHKey implemented in a fake way for automated tests
func (c *FakeClient) NewSSHKey(name string, publicKey string) (*SimpleResponse, error) {
	sshKey := SSHKey{
		Name:        name,
		Fingerprint: publicKey, // This is weird, but we're just storing a value
	}
	c.SSHKeys = append(c.SSHKeys, sshKey)
	return &SimpleResponse{Result: "success"}, nil
}

// UpdateSSHKey implemented in a fake way for automated tests
func (c *FakeClient) UpdateSSHKey(name string, sshKeyID string) (*SSHKey, error) {
	for i, sshKey := range c.SSHKeys {
		if sshKey.ID == sshKeyID {
			c.SSHKeys[i].Name = name
			return &sshKey, nil
		}
	}

	err := fmt.Errorf("unable to find SSH key %s, zero matches", sshKeyID)
	return nil, ZeroMatchesError.wrap(err)
}

// FindSSHKey implemented in a fake way for automated tests
func (c *FakeClient) FindSSHKey(search string) (*SSHKey, error) {
	for _, sshKey := range c.SSHKeys {
		if strings.Contains(sshKey.Name, search) {
			return &sshKey, nil
		}
	}

	err := fmt.Errorf("unable to find SSH key %s, zero matches", search)
	return nil, ZeroMatchesError.wrap(err)
}

// DeleteSSHKey implemented in a fake way for automated tests
func (c *FakeClient) DeleteSSHKey(id string) (*SimpleResponse, error) {
	for i, sshKey := range c.SSHKeys {
		if sshKey.ID == id {
			c.SSHKeys[len(c.SSHKeys)-1], c.SSHKeys[i] = c.SSHKeys[i], c.SSHKeys[len(c.SSHKeys)-1]
			c.SSHKeys = c.SSHKeys[:len(c.SSHKeys)-1]
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return &SimpleResponse{Result: "failed"}, nil
}

// ListTemplates implemented in a fake way for automated tests
// func (c *FakeClient) ListTemplates() ([]Template, error) {
// 	return c.Templates, nil
// }

// NewTemplate implemented in a fake way for automated tests
// func (c *FakeClient) NewTemplate(conf *Template) (*SimpleResponse, error) {
// 	template := Template{
// 		ID:          conf.ID,
// 		Name:        conf.Name,
// 		CloudConfig: conf.CloudConfig,
// 		ImageID:     conf.ImageID,
// 	}
// 	c.Templates = append(c.Templates, template)
// 	return &SimpleResponse{Result: "success"}, nil
// }

// UpdateTemplate implemented in a fake way for automated tests
// func (c *FakeClient) UpdateTemplate(id string, conf *Template) (*Template, error) {
// 	for i, template := range c.Templates {
// 		if template.ID == id {
// 			c.Templates[i].Name = conf.Name
// 			c.Templates[i].CloudConfig = conf.CloudConfig
// 			c.Templates[i].ImageID = conf.ImageID
// 			return &template, nil
// 		}
// 	}

// 	err := fmt.Errorf("unable to find SSH key %s, zero matches", id)
// 	return nil, ZeroMatchesError.wrap(err)
// }

// GetTemplateByCode implemented in a fake way for automated tests
// func (c *FakeClient) GetTemplateByCode(code string) (*Template, error) {
// 	for _, template := range c.Templates {
// 		if template.Code == code {
// 			return &template, nil
// 		}
// 	}

// 	err := fmt.Errorf("unable to find SSH key %s, zero matches", code)
// 	return nil, ZeroMatchesError.wrap(err)
// }

// FindTemplate implemented in a fake way for automated tests
// func (c *FakeClient) FindTemplate(search string) (*Template, error) {
// 	for _, template := range c.Templates {
// 		if strings.Contains(template.Name, search) {
// 			return &template, nil
// 		}
// 	}

// 	err := fmt.Errorf("unable to find template %s, zero matches", search)
// 	return nil, ZeroMatchesError.wrap(err)
// }

// DeleteTemplate implemented in a fake way for automated tests
// func (c *FakeClient) DeleteTemplate(id string) (*SimpleResponse, error) {
// 	for i, template := range c.Templates {
// 		if template.ID == id {
// 			c.Templates[len(c.Templates)-1], c.Templates[i] = c.Templates[i], c.Templates[len(c.Templates)-1]
// 			c.Templates = c.Templates[:len(c.Templates)-1]
// 			return &SimpleResponse{Result: "success"}, nil
// 		}
// 	}

// 	return &SimpleResponse{Result: "failed"}, nil
// }

// ListDiskImages implemented in a fake way for automated tests
func (c *FakeClient) ListDiskImages() ([]DiskImage, error) {
	return c.DiskImage, nil
}

// GetDiskImage implemented in a fake way for automated tests
func (c *FakeClient) GetDiskImage(id string) (*DiskImage, error) {
	for k, v := range c.DiskImage {
		if v.ID == id {
			return &c.DiskImage[k], nil
		}
	}

	err := fmt.Errorf("unable to find disk image %s, zero matches", id)
	return nil, ZeroMatchesError.wrap(err)
}

// FindDiskImage implemented in a fake way for automated tests
func (c *FakeClient) FindDiskImage(search string) (*DiskImage, error) {
	for _, diskimage := range c.DiskImage {
		if strings.Contains(diskimage.Name, search) || strings.Contains(diskimage.ID, search) {
			return &diskimage, nil
		}
	}

	err := fmt.Errorf("unable to find volume %s, zero matches", search)
	return nil, ZeroMatchesError.wrap(err)
}

// ListVolumes implemented in a fake way for automated tests
func (c *FakeClient) ListVolumes() ([]Volume, error) {
	return c.Volumes, nil
}

// GetVolume implemented in a fake way for automated tests
func (c *FakeClient) GetVolume(id string) (*Volume, error) {
	for _, volume := range c.Volumes {
		if volume.ID == id {
			return &volume, nil
		}
	}

	err := fmt.Errorf("unable to get volume %s", id)
	return nil, ZeroMatchesError.wrap(err)
}

// FindVolume implemented in a fake way for automated tests
func (c *FakeClient) FindVolume(search string) (*Volume, error) {
	for _, volume := range c.Volumes {
		if strings.Contains(volume.Name, search) || strings.Contains(volume.ID, search) {
			return &volume, nil
		}
	}

	err := fmt.Errorf("unable to find volume %s, zero matches", search)
	return nil, ZeroMatchesError.wrap(err)
}

// NewVolume implemented in a fake way for automated tests
func (c *FakeClient) NewVolume(v *VolumeConfig) (*VolumeResult, error) {
	volume := Volume{
		ID:            c.generateID(),
		Name:          v.Name,
		SizeGigabytes: v.SizeGigabytes,
		Status:        "available",
	}
	c.Volumes = append(c.Volumes, volume)

	return &VolumeResult{
		ID:     volume.ID,
		Name:   volume.Name,
		Result: "success",
	}, nil
}

// ResizeVolume implemented in a fake way for automated tests
func (c *FakeClient) ResizeVolume(id string, size int) (*SimpleResponse, error) {
	for i, volume := range c.Volumes {
		if volume.ID == id {
			c.Volumes[i].SizeGigabytes = size
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	err := fmt.Errorf("unable to find volume %s, zero matches", id)
	return nil, ZeroMatchesError.wrap(err)
}

// AttachVolume implemented in a fake way for automated tests
func (c *FakeClient) AttachVolume(id string, instance string) (*SimpleResponse, error) {
	for i, volume := range c.Volumes {
		if volume.ID == id {
			c.Volumes[i].InstanceID = instance
			c.Volumes[i].Status = "attached"
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	err := fmt.Errorf("unable to find volume %s, zero matches", id)
	return nil, ZeroMatchesError.wrap(err)
}

// DetachVolume implemented in a fake way for automated tests
func (c *FakeClient) DetachVolume(id string) (*SimpleResponse, error) {
	for i, volume := range c.Volumes {
		if volume.ID == id {
			c.Volumes[i].InstanceID = ""
			c.Volumes[i].Status = "available"
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	err := fmt.Errorf("unable to find volume %s, zero matches", id)
	return nil, ZeroMatchesError.wrap(err)
}

// DeleteVolume implemented in a fake way for automated tests
func (c *FakeClient) DeleteVolume(id string) (*SimpleResponse, error) {
	for i, volume := range c.Volumes {
		if volume.ID == id {
			c.Volumes[len(c.Volumes)-1], c.Volumes[i] = c.Volumes[i], c.Volumes[len(c.Volumes)-1]
			c.Volumes = c.Volumes[:len(c.Volumes)-1]
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return &SimpleResponse{Result: "failed"}, nil
}

// CreateWebhook implemented in a fake way for automated tests
func (c *FakeClient) CreateWebhook(r *WebhookConfig) (*Webhook, error) {
	webhook := Webhook{
		ID:     c.generateID(),
		Events: r.Events,
		Secret: r.Secret,
		URL:    r.URL,
	}
	c.Webhooks = append(c.Webhooks, webhook)

	return &webhook, nil
}

// ListWebhooks implemented in a fake way for automated tests
func (c *FakeClient) ListWebhooks() ([]Webhook, error) {
	return c.Webhooks, nil
}

// FindWebhook implemented in a fake way for automated tests
func (c *FakeClient) FindWebhook(search string) (*Webhook, error) {
	for _, webhook := range c.Webhooks {
		if strings.Contains(webhook.Secret, search) || strings.Contains(webhook.URL, search) {
			return &webhook, nil
		}
	}

	err := fmt.Errorf("unable to find %s, zero matches", search)
	return nil, ZeroMatchesError.wrap(err)
}

// UpdateWebhook implemented in a fake way for automated tests
func (c *FakeClient) UpdateWebhook(id string, r *WebhookConfig) (*Webhook, error) {
	for i, webhook := range c.Webhooks {
		if webhook.ID == id {
			c.Webhooks[i].Events = r.Events
			c.Webhooks[i].Secret = r.Secret
			c.Webhooks[i].URL = r.URL

			return &webhook, nil
		}
	}

	err := fmt.Errorf("unable to find %s, zero matches", id)
	return nil, ZeroMatchesError.wrap(err)
}

// DeleteWebhook implemented in a fake way for automated tests
func (c *FakeClient) DeleteWebhook(id string) (*SimpleResponse, error) {
	for i, webhook := range c.Webhooks {
		if webhook.ID == id {
			c.Webhooks[len(c.Webhooks)-1], c.Webhooks[i] = c.Webhooks[i], c.Webhooks[len(c.Webhooks)-1]
			c.Webhooks = c.Webhooks[:len(c.Webhooks)-1]
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return &SimpleResponse{Result: "failed"}, nil
}

// ListPermissions implemented in a fake way for automated tests
func (c *FakeClient) ListPermissions() ([]Permission, error) {
	return []Permission{
		{
			Name:        "instance.create",
			Description: "Create Compute instances",
		},
		{
			Name:        "kubernetes.*",
			Description: "Manage Civo Kubernetes clusters",
		},
	}, nil
}

// GetOrganisation implemented in a fake way for automated tests
func (c *FakeClient) GetOrganisation() (*Organisation, error) {
	return &c.Organisation, nil
}

// CreateOrganisation implemented in a fake way for automated tests
func (c *FakeClient) CreateOrganisation(name string) (*Organisation, error) {
	c.Organisation.ID = c.generateID()
	c.Organisation.Name = name
	return &c.Organisation, nil
}

// RenameOrganisation implemented in a fake way for automated tests
func (c *FakeClient) RenameOrganisation(name string) (*Organisation, error) {
	c.Organisation.Name = name
	return &c.Organisation, nil
}

// AddAccountToOrganisation implemented in a fake way for automated tests
func (c *FakeClient) AddAccountToOrganisation(accountID string) ([]Account, error) {
	c.OrganisationAccounts = append(c.OrganisationAccounts, Account{
		ID:        accountID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return c.ListAccountsInOrganisation()
}

// ListAccountsInOrganisation implemented in a fake way for automated tests
func (c *FakeClient) ListAccountsInOrganisation() ([]Account, error) {
	return c.OrganisationAccounts, nil
}

// ListRoles implemented in a fake way for automated tests
func (c *FakeClient) ListRoles() ([]Role, error) {
	return c.OrganisationRoles, nil
}

// CreateRole implemented in a fake way for automated tests
func (c *FakeClient) CreateRole(name, permissions string) (*Role, error) {
	role := Role{
		ID:          c.generateID(),
		Name:        name,
		Permissions: permissions,
	}
	c.OrganisationRoles = append(c.OrganisationRoles, role)
	return &role, nil
}

// DeleteRole implemented in a fake way for automated tests
func (c *FakeClient) DeleteRole(id string) (*SimpleResponse, error) {
	for i, role := range c.OrganisationRoles {
		if role.ID == id {
			c.OrganisationRoles[len(c.OrganisationRoles)-1], c.OrganisationRoles[i] = c.OrganisationRoles[i], c.OrganisationRoles[len(c.OrganisationRoles)-1]
			c.OrganisationRoles = c.OrganisationRoles[:len(c.OrganisationRoles)-1]
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return &SimpleResponse{Result: "failed"}, fmt.Errorf("unable to find that role")
}

// ListTeams implemented in a fake way for automated tests
func (c *FakeClient) ListTeams() ([]Team, error) {
	return c.OrganisationTeams, nil
}

// CreateTeam implemented in a fake way for automated tests
func (c *FakeClient) CreateTeam(name string) (*Team, error) {
	team := Team{
		ID:        c.generateID(),
		Name:      name,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	c.OrganisationTeams = append(c.OrganisationTeams, team)
	return &team, nil
}

// RenameTeam implemented in a fake way for automated tests
func (c *FakeClient) RenameTeam(teamID, name string) (*Team, error) {
	for _, team := range c.OrganisationTeams {
		if team.ID == teamID {
			team.Name = name
			return &team, nil
		}
	}

	return nil, fmt.Errorf("unable to find that role")
}

// DeleteTeam implemented in a fake way for automated tests
func (c *FakeClient) DeleteTeam(id string) (*SimpleResponse, error) {
	for i, team := range c.OrganisationTeams {
		if team.ID == id {
			c.OrganisationTeams[len(c.OrganisationTeams)-1], c.OrganisationTeams[i] = c.OrganisationTeams[i], c.OrganisationTeams[len(c.OrganisationTeams)-1]
			c.OrganisationTeams = c.OrganisationTeams[:len(c.OrganisationTeams)-1]
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return &SimpleResponse{Result: "failure"}, fmt.Errorf("unable to find that team")
}

// ListTeamMembers implemented in a fake way for automated tests
func (c *FakeClient) ListTeamMembers(teamID string) ([]TeamMember, error) {
	return c.OrganisationTeamMembers[teamID], nil
}

// AddTeamMember implemented in a fake way for automated tests
func (c *FakeClient) AddTeamMember(teamID, userID, permissions, roles string) ([]TeamMember, error) {
	c.OrganisationTeamMembers[teamID] = append(c.OrganisationTeamMembers[teamID], TeamMember{
		ID:          c.generateID(),
		TeamID:      teamID,
		UserID:      userID,
		Permissions: permissions,
		Roles:       roles,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	return c.ListTeamMembers(teamID)
}

// UpdateTeamMember implemented in a fake way for automated tests
func (c *FakeClient) UpdateTeamMember(teamID, teamMemberID, permissions, roles string) (*TeamMember, error) {
	for _, teamMember := range c.OrganisationTeamMembers[teamID] {
		if teamMember.ID == teamMemberID {
			teamMember.Permissions = permissions
			teamMember.Roles = roles
			return &teamMember, nil
		}
	}

	return nil, fmt.Errorf("unable to find that role")
}

// RemoveTeamMember implemented in a fake way for automated tests
func (c *FakeClient) RemoveTeamMember(teamID, teamMemberID string) (*SimpleResponse, error) {
	for i, teamMember := range c.OrganisationTeamMembers[teamID] {
		if teamMember.ID == teamMemberID {
			c.OrganisationTeamMembers[teamID][len(c.OrganisationTeamMembers[teamID])-1], c.OrganisationTeamMembers[teamID][i] = c.OrganisationTeamMembers[teamID][i], c.OrganisationTeamMembers[teamID][len(c.OrganisationTeamMembers[teamID])-1]
			c.OrganisationTeamMembers[teamID] = c.OrganisationTeamMembers[teamID][:len(c.OrganisationTeamMembers[teamID])-1]
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return &SimpleResponse{Result: "failure"}, fmt.Errorf("unable to find that team member")
}

// ListLoadBalancers implemented in a fake way for automated tests
func (c *FakeClient) ListLoadBalancers() ([]LoadBalancer, error) {
	return c.LoadBalancers, nil
}

// GetLoadBalancer implemented in a fake way for automated tests
func (c *FakeClient) GetLoadBalancer(id string) (*LoadBalancer, error) {
	for _, lb := range c.LoadBalancers {
		if lb.ID == id {
			return &lb, nil
		}
	}

	err := fmt.Errorf("unable to get load balancer %s", id)
	return nil, DatabaseLoadBalancerNotFoundError.wrap(err)
}

// FindLoadBalancer implemented in a fake way for automated tests
func (c *FakeClient) FindLoadBalancer(search string) (*LoadBalancer, error) {
	exactMatch := false
	partialMatchesCount := 0
	result := LoadBalancer{}

	for _, lb := range c.LoadBalancers {
		if lb.ID == search || lb.Name == search {
			exactMatch = true
			result = lb
		} else if strings.Contains(lb.Name, search) || strings.Contains(lb.ID, search) {
			if !exactMatch {
				result = lb
				partialMatchesCount++
			}
		}
	}

	if exactMatch || partialMatchesCount == 1 {
		return &result, nil
	} else if partialMatchesCount > 1 {
		err := fmt.Errorf("unable to find %s because there were multiple matches", search)
		return nil, MultipleMatchesError.wrap(err)
	} else {
		err := fmt.Errorf("unable to find %s, zero matches", search)
		return nil, ZeroMatchesError.wrap(err)
	}
}

// CreateLoadBalancer implemented in a fake way for automated tests
func (c *FakeClient) CreateLoadBalancer(r *LoadBalancerConfig) (*LoadBalancer, error) {
	loadbalancer := LoadBalancer{
		ID:                           c.generateID(),
		Name:                         r.Name,
		Algorithm:                    r.Algorithm,
		ExternalTrafficPolicy:        r.ExternalTrafficPolicy,
		SessionAffinityConfigTimeout: r.SessionAffinityConfigTimeout,
		SessionAffinity:              r.SessionAffinity,
		EnableProxyProtocol:          r.EnableProxyProtocol,
		FirewallID:                   r.FirewallID,
		ClusterID:                    r.ClusterID,
	}

	if r.Algorithm == "" {
		loadbalancer.Algorithm = "round_robin"
	}
	if r.FirewallID == "" {
		loadbalancer.FirewallID = c.generateID()
	}
	if r.ExternalTrafficPolicy == "" {
		loadbalancer.ExternalTrafficPolicy = "Cluster"
	}

	backends := make([]LoadBalancerBackend, 0)
	for _, b := range r.Backends {
		backend := LoadBalancerBackend{
			IP:         b.IP,
			Protocol:   b.Protocol,
			SourcePort: b.SourcePort,
			TargetPort: b.TargetPort,
		}
		backends = append(backends, backend)
	}
	loadbalancer.Backends = backends
	loadbalancer.PublicIP = c.generatePublicIP()
	loadbalancer.State = "available"

	c.LoadBalancers = append(c.LoadBalancers, loadbalancer)
	return &loadbalancer, nil
}

// UpdateLoadBalancer implemented in a fake way for automated tests
func (c *FakeClient) UpdateLoadBalancer(id string, r *LoadBalancerUpdateConfig) (*LoadBalancer, error) {
	for _, lb := range c.LoadBalancers {
		if lb.ID == id {
			lb.Name = r.Name
			lb.Algorithm = r.Algorithm
			lb.EnableProxyProtocol = r.EnableProxyProtocol
			lb.ExternalTrafficPolicy = r.ExternalTrafficPolicy
			lb.SessionAffinity = r.SessionAffinity
			lb.SessionAffinityConfigTimeout = r.SessionAffinityConfigTimeout

			backends := make([]LoadBalancerBackend, len(r.Backends))
			for i, b := range r.Backends {
				backends[i].IP = b.IP
				backends[i].Protocol = b.Protocol
				backends[i].SourcePort = b.SourcePort
				backends[i].TargetPort = b.TargetPort
			}

			if r.ExternalTrafficPolicy == "" {
				lb.ExternalTrafficPolicy = "Cluster"
			}

			return &lb, nil
		}
	}

	err := fmt.Errorf("unable to find load balancer %s", id)
	return nil, DatabaseLoadBalancerNotFoundError.wrap(err)
}

// DeleteLoadBalancer implemented in a fake way for automated tests
func (c *FakeClient) DeleteLoadBalancer(id string) (*SimpleResponse, error) {
	for i, lb := range c.LoadBalancers {
		if lb.ID == id {
			c.LoadBalancers[len(c.LoadBalancers)-1], c.LoadBalancers[i] = c.LoadBalancers[i], c.LoadBalancers[len(c.LoadBalancers)-1]
			c.LoadBalancers = c.LoadBalancers[:len(c.LoadBalancers)-1]
			return &SimpleResponse{Result: "success"}, nil
		}
	}

	return &SimpleResponse{Result: "failed"}, nil
}

// ListKubernetesClusterPools implemented in a fake way for automated tests
func (c *FakeClient) ListKubernetesClusterPools(cid string) ([]KubernetesPool, error) {
	pools := []KubernetesPool{}
	found := false

	for _, cs := range c.Clusters {
		if cs.ID == cid {
			found = true
			pools = cs.Pools
			break
		}
	}

	if found {
		return pools, nil
	}

	err := fmt.Errorf("unable to get kubernetes cluster %s", cid)
	return nil, DatabaseKubernetesClusterNotFoundError.wrap(err)
}

// GetKubernetesClusterPool implemented in a fake way for automated tests
func (c *FakeClient) GetKubernetesClusterPool(cid, pid string) (*KubernetesPool, error) {
	pool := &KubernetesPool{}
	clusterFound := false
	poolFound := false

	for _, cs := range c.Clusters {
		if cs.ID == cid {
			clusterFound = true
			for _, p := range cs.Pools {
				if p.ID == pid {
					poolFound = true
					pool = &p
					break
				}
			}
		}
	}

	if !clusterFound {
		err := fmt.Errorf("unable to get kubernetes cluster %s", cid)
		return nil, DatabaseKubernetesClusterNotFoundError.wrap(err)
	}

	if !poolFound {
		err := fmt.Errorf("unable to get kubernetes pool %s", pid)
		return nil, DatabaseKubernetesClusterNotFoundError.wrap(err)
	}

	return pool, nil
}

// FindKubernetesClusterPool implemented in a fake way for automated tests
func (c *FakeClient) FindKubernetesClusterPool(cid, search string) (*KubernetesPool, error) {
	pool := &KubernetesPool{}
	clusterFound := false
	poolFound := false

	for _, cs := range c.Clusters {
		if cs.ID == cid {
			clusterFound = true
			for _, p := range cs.Pools {
				if p.ID == search || strings.Contains(p.ID, search) {
					poolFound = true
					pool = &p
					break
				}
			}
		}
	}

	if !clusterFound {
		err := fmt.Errorf("unable to get kubernetes cluster %s", cid)
		return nil, DatabaseKubernetesClusterNotFoundError.wrap(err)
	}

	if !poolFound {
		err := fmt.Errorf("unable to get kubernetes pool %s", search)
		return nil, DatabaseKubernetesClusterNotFoundError.wrap(err)
	}

	return pool, nil
}

// DeleteKubernetesClusterPoolInstance implemented in a fake way for automated tests
func (c *FakeClient) DeleteKubernetesClusterPoolInstance(cid, pid, id string) (*SimpleResponse, error) {
	clusterFound := false
	poolFound := false
	instanceFound := false

	for ci, cs := range c.Clusters {
		if cs.ID == cid {
			clusterFound = true
			for pi, p := range cs.Pools {
				if p.ID == pid {
					poolFound = true
					for i, in := range p.Instances {
						if in.ID == id {
							instanceFound = true
							p.Instances = append(p.Instances[:i], p.Instances[i+1:]...)

							instanceNames := []string{}
							for _, in := range p.Instances {
								instanceNames = append(instanceNames, in.Hostname)
							}
							p.InstanceNames = instanceNames
							c.Clusters[ci].Pools[pi] = p
							break
						}
					}
				}
			}
		}
	}

	if !clusterFound {
		err := fmt.Errorf("unable to get kubernetes cluster %s", cid)
		return nil, DatabaseKubernetesClusterNotFoundError.wrap(err)
	}

	if !poolFound {
		err := fmt.Errorf("unable to get kubernetes pool %s", pid)
		return nil, DatabaseKubernetesClusterNotFoundError.wrap(err)
	}

	if !instanceFound {
		err := fmt.Errorf("unable to get kubernetes pool instance %s", id)
		return nil, DatabaseKubernetesClusterNotFoundError.wrap(err)
	}

	return &SimpleResponse{
		Result: "success",
	}, nil
}

// UpdateKubernetesClusterPool implemented in a fake way for automated tests
func (c *FakeClient) UpdateKubernetesClusterPool(cid, pid string, config *KubernetesClusterPoolUpdateConfig) (*KubernetesPool, error) {
	clusterFound := false
	poolFound := false

	pool := KubernetesPool{}
	for _, cs := range c.Clusters {
		if cs.ID == cid {
			clusterFound = true
			for _, p := range cs.Pools {
				if p.ID == pid {
					poolFound = true
					p.Count = config.Count
					pool = p
				}
			}
		}
	}

	if !clusterFound {
		err := fmt.Errorf("unable to get kubernetes cluster %s", cid)
		return nil, DatabaseKubernetesClusterNotFoundError.wrap(err)
	}

	if !poolFound {
		err := fmt.Errorf("unable to get kubernetes pool %s", pid)
		return nil, DatabaseKubernetesClusterNotFoundError.wrap(err)
	}

	return &pool, nil
}

// ListIPs returns a list of fake IPs
func (c *FakeClient) ListIPs() (*PaginatedIPs, error) {
	return &PaginatedIPs{
		Page:    1,
		PerPage: 20,
		Pages:   100,
		Items: []IP{
			{
				ID:   c.generateID(),
				Name: "test-ip",
				IP:   c.generatePublicIP(),
			},
		},
	}, nil
}

// GetIP returns a fake IP
func (c *FakeClient) GetIP(id string) (*IP, error) {
	return &IP{
		ID:   c.generateID(),
		Name: "test-ip",
		IP:   c.generatePublicIP(),
	}, nil
}

// FindIP finds a fake IP
func (c *FakeClient) FindIP(search string) (*IP, error) {
	return &IP{
		ID:   c.generateID(),
		Name: "test-ip",
		IP:   c.generatePublicIP(),
	}, nil
}

// NewIP creates a fake IP
func (c *FakeClient) NewIP(v *CreateIPRequest) (*IP, error) {
	return &IP{
		ID:   c.generateID(),
		Name: "test-ip",
		IP:   c.generatePublicIP(),
	}, nil
}

// UpdateIP updates a fake IP
func (c *FakeClient) UpdateIP(id string, v *UpdateIPRequest) (*IP, error) {
	return &IP{
		ID:   c.generateID(),
		Name: v.Name,
		IP:   c.generatePublicIP(),
	}, nil
}

// DeleteIP deletes a fake IP
func (c *FakeClient) DeleteIP(id string) (*SimpleResponse, error) {
	return &SimpleResponse{
		Result: "success",
	}, nil
}

// AssignIP assigns a fake IP
func (c *FakeClient) AssignIP(id, resourceID, resourceType, region string) (*SimpleResponse, error) {
	return &SimpleResponse{
		Result: "success",
	}, nil
}

// UnassignIP unassigns a fake IP
func (c *FakeClient) UnassignIP(id, region string) (*SimpleResponse, error) {
	return &SimpleResponse{
		Result: "success",
	}, nil
}
