package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KubernetesInstance represents a single node/master within a Kubernetes cluster
type KubernetesInstance struct {
	ID              string    `json:"id"`
	Hostname        string    `json:"hostname,omitempty"`
	Size            string    `json:"size,omitempty"`
	Region          string    `json:"region,omitempty"`
	SourceType      string    `json:"source_type,omitempty"`
	SourceID        string    `json:"source_id,omitempty"`
	InitialUser     string    `json:"initial_user,omitempty"`
	InitialPassword string    `json:"initial_password,omitempty"`
	Status          string    `json:"status,omitempty"`
	FirewallID      string    `json:"firewall_id,omitempty"`
	PublicIP        string    `json:"public_ip,omitempty"`
	CPUCores        int       `json:"cpu_cores,omitempty"`
	RAMMegabytes    int       `json:"ram_mb,omitempty"`
	DiskGigabytes   int       `json:"disk_gb,omitempty"`
	Tags            []string  `json:"tags,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	CivoStatsdToken string    `json:"civostatsd_token,omitempty"`
}

// KubernetesPool represents a single pool within a Kubernetes cluster
type KubernetesPool struct {
	ID               string               `json:"id"`
	Count            int                  `json:"count,omitempty"`
	Size             string               `json:"size,omitempty"`
	InstanceNames    []string             `json:"instance_names,omitempty"`
	Instances        []KubernetesInstance `json:"instances,omitempty"`
	Labels           map[string]string    `json:"labels,omitempty"`
	Annotations      map[string]string    `json:"annotations,omitempty"`
	Taints           []corev1.Taint       `json:"taints,omitempty"`
	PublicIPNodePool bool                 `json:"public_ip_node_pool,omitempty"`
}

// KubernetesInstalledApplication is an application within our marketplace available for
// installation
type KubernetesInstalledApplication struct {
	Application   string                              `json:"application,omitempty"`
	Name          string                              `json:"name,omitempty"`
	Version       string                              `json:"version,omitempty"`
	Dependencies  []string                            `json:"dependencies,omitempty"`
	Maintainer    string                              `json:"maintainer,omitempty"`
	Description   string                              `json:"description,omitempty"`
	PostInstall   string                              `json:"post_install,omitempty"`
	Installed     bool                                `json:"installed,omitempty"`
	URL           string                              `json:"url,omitempty"`
	Category      string                              `json:"category,omitempty"`
	UpdatedAt     time.Time                           `json:"updated_at,omitempty"`
	ImageURL      string                              `json:"image_url,omitempty"`
	Plan          string                              `json:"plan,omitempty"`
	Configuration map[string]ApplicationConfiguration `json:"configuration,omitempty"`
}

// ApplicationConfiguration is a configuration for installed application
type ApplicationConfiguration map[string]string

// Condition is a condition for a Kubernetes cluster
type Condition struct {
	Type               string                 `json:"type"`
	Status             metav1.ConditionStatus `json:"status"`
	Synced             bool                   `json:"synced"`
	LastTransitionTime metav1.Time            `json:"last_transition_time"`
	Reason             string                 `json:"reason,omitempty"`
	Message            string                 `json:"message,omitempty"`
}

// KubernetesCluster is a Kubernetes item inside the cluster
type KubernetesCluster struct {
	ID                    string                           `json:"id"`
	Name                  string                           `json:"name,omitempty"`
	GeneratedName         string                           `json:"generated_name,omitempty"`
	Version               string                           `json:"version,omitempty"`
	Status                string                           `json:"status,omitempty"`
	Ready                 bool                             `json:"ready,omitempty"`
	ClusterType           string                           `json:"cluster_type,omitempty"`
	NumTargetNode         int                              `json:"num_target_nodes,omitempty"`
	TargetNodeSize        string                           `json:"target_nodes_size,omitempty"`
	BuiltAt               time.Time                        `json:"built_at,omitempty"`
	KubeConfig            string                           `json:"kubeconfig,omitempty"`
	KubernetesVersion     string                           `json:"kubernetes_version,omitempty"`
	APIEndPoint           string                           `json:"api_endpoint,omitempty"`
	MasterIP              string                           `json:"master_ip,omitempty"`
	DNSEntry              string                           `json:"dns_entry,omitempty"`
	UpgradeAvailableTo    string                           `json:"upgrade_available_to,omitempty"`
	Legacy                bool                             `json:"legacy,omitempty"`
	NetworkID             string                           `json:"network_id,omitempty"`
	NameSpace             string                           `json:"namespace,omitempty"`
	Tags                  []string                         `json:"tags,omitempty"`
	CreatedAt             time.Time                        `json:"created_at,omitempty"`
	Instances             []KubernetesInstance             `json:"instances,omitempty"`
	Pools                 []KubernetesPool                 `json:"pools,omitempty"`
	RequiredPools         []RequiredPools                  `json:"required_pools,omitempty"`
	InstalledApplications []KubernetesInstalledApplication `json:"installed_applications,omitempty"`
	FirewallID            string                           `json:"firewall_id,omitempty"`
	CNIPlugin             string                           `json:"cni_plugin,omitempty"`
	CCMInstalled          string                           `json:"ccm_installed,omitempty"`
	Conditions            []Condition                      `json:"conditions"`
	VolumeType            string                           `json:"volume_type,omitempty"`
}

// RequiredPools returns the required pools for a given Kubernetes cluster
type RequiredPools struct {
	ID               string            `json:"id"`
	Size             string            `json:"size"`
	Count            int               `json:"count"`
	Labels           map[string]string `json:"labels,omitempty"`
	Annotations      map[string]string `json:"annotations,omitempty"`
	Taints           []corev1.Taint    `json:"taints,omitempty"`
	PublicIPNodePool bool              `json:"public_ip_node_pool,omitempty"`
}

// PaginatedKubernetesClusters is a Kubernetes k3s cluster
type PaginatedKubernetesClusters struct {
	Page    int                 `json:"page"`
	PerPage int                 `json:"per_page"`
	Pages   int                 `json:"pages"`
	Items   []KubernetesCluster `json:"items"`
}

// KubernetesClusterConfig is used to create a new cluster
type KubernetesClusterConfig struct {
	Name              string                        `json:"name,omitempty"`
	Region            string                        `json:"region,omitempty"`
	ClusterType       string                        `json:"cluster_type,omitempty"`
	NumTargetNodes    int                           `json:"num_target_nodes,omitempty"`
	TargetNodesSize   string                        `json:"target_nodes_size,omitempty"`
	KubernetesVersion string                        `json:"kubernetes_version,omitempty"`
	NodeDestroy       string                        `json:"node_destroy,omitempty"`
	NetworkID         string                        `json:"network_id,omitempty"`
	Tags              string                        `json:"tags,omitempty"`
	Pools             []KubernetesClusterPoolConfig `json:"pools,omitempty"`
	Applications      string                        `json:"applications,omitempty"`
	InstanceFirewall  string                        `json:"instance_firewall,omitempty"`
	FirewallRule      string                        `json:"firewall_rule,omitempty"`
	FirewallID        string                        `json:"firewall_id,omitempty"`
	CNIPlugin         string                        `json:"cni_plugin,omitempty"`
	VolumeType        string                        `json:"volume_type,omitempty"`
}

// KubernetesClusterPoolConfig is used to create a new cluster pool
type KubernetesClusterPoolConfig struct {
	Region           string            `json:"region,omitempty"`
	ID               string            `json:"id,omitempty"`
	Count            int               `json:"count,omitempty"`
	Size             string            `json:"size,omitempty"`
	Labels           map[string]string `json:"labels,omitempty"`
	Taints           []corev1.Taint    `json:"taints"`
	PublicIPNodePool bool              `json:"public_ip_node_pool,omitempty"`
}

// KubernetesPlanConfiguration is a value within a configuration for
// an application's plan
type KubernetesPlanConfiguration struct {
	Value string `json:"value"`
}

// KubernetesMarketplacePlan is a plan for
type KubernetesMarketplacePlan struct {
	Label         string                                 `json:"label"`
	Configuration map[string]KubernetesPlanConfiguration `json:"configuration"`
}

// KubernetesMarketplaceApplication is an application within our marketplace
// available for installation
type KubernetesMarketplaceApplication struct {
	Name         string                      `json:"name"`
	Title        string                      `json:"title,omitempty"`
	Version      string                      `json:"version"`
	Default      bool                        `json:"default,omitempty"`
	Dependencies []string                    `json:"dependencies,omitempty"`
	Maintainer   string                      `json:"maintainer"`
	Description  string                      `json:"description"`
	PostInstall  string                      `json:"post_install"`
	URL          string                      `json:"url"`
	Category     string                      `json:"category"`
	Plans        []KubernetesMarketplacePlan `json:"plans"`
}

// KubernetesVersion represents an available version of k3s to install
type KubernetesVersion struct {
	Label       string `json:"label"`
	Version     string `json:"version"`
	Type        string `json:"type"`
	Default     bool   `json:"default,omitempty"`
	ClusterType string `json:"clusterType,omitempty"`
}

// ListKubernetesClusters returns all cluster of kubernetes in the account
func (c *Client) ListKubernetesClusters() (*PaginatedKubernetesClusters, error) {
	resp, err := c.SendGetRequest("/v2/kubernetes/clusters")
	if err != nil {
		return nil, decodeError(err)
	}

	kubernetes := &PaginatedKubernetesClusters{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&kubernetes); err != nil {
		return nil, err
	}

	return kubernetes, nil
}

// FindKubernetesCluster finds a Kubernetes cluster by either part of the ID or part of the name
func (c *Client) FindKubernetesCluster(search string) (*KubernetesCluster, error) {
	clusters, err := c.ListKubernetesClusters()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := KubernetesCluster{}

	for _, value := range clusters.Items {
		if strings.EqualFold(value.Name, search) || value.ID == search {
			exactMatch = true
			result = value
		} else if strings.Contains(strings.ToUpper(value.Name), strings.ToUpper(search)) || strings.Contains(value.ID, search) {
			if !exactMatch {
				result = value
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

// NewKubernetesClusters create a new cluster of kubernetes
func (c *Client) NewKubernetesClusters(kc *KubernetesClusterConfig) (*KubernetesCluster, error) {
	kc.Region = c.Region
	body, err := c.SendPostRequest("/v2/kubernetes/clusters", kc)
	if err != nil {
		return nil, decodeError(err)
	}

	kubernetes := &KubernetesCluster{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(kubernetes); err != nil {
		return nil, err
	}

	return kubernetes, nil
}

// GetKubernetesCluster returns a single kubernetes cluster by its full ID
func (c *Client) GetKubernetesCluster(id string) (*KubernetesCluster, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/kubernetes/clusters/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	kubernetes := &KubernetesCluster{}
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(kubernetes); err != nil {
		return nil, err
	}
	return kubernetes, nil
}

// UpdateKubernetesCluster update a single kubernetes cluster by its full ID
func (c *Client) UpdateKubernetesCluster(id string, i *KubernetesClusterConfig) (*KubernetesCluster, error) {
	i.Region = c.Region
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/kubernetes/clusters/%s", id), i)
	if err != nil {
		return nil, decodeError(err)
	}

	kubernetes := &KubernetesCluster{}
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(kubernetes); err != nil {
		return nil, err
	}
	return kubernetes, nil
}

// ListKubernetesMarketplaceApplications returns all application inside marketplace
func (c *Client) ListKubernetesMarketplaceApplications() ([]KubernetesMarketplaceApplication, error) {
	resp, err := c.SendGetRequest("/v2/kubernetes/applications")
	if err != nil {
		return nil, decodeError(err)
	}

	kubernetes := make([]KubernetesMarketplaceApplication, 0)
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&kubernetes); err != nil {
		return nil, err
	}

	return kubernetes, nil
}

// DeleteKubernetesCluster deletes a cluster
func (c *Client) DeleteKubernetesCluster(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/kubernetes/clusters/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// RecycleKubernetesCluster create a new cluster of kubernetes
func (c *Client) RecycleKubernetesCluster(id string, hostname string) (*SimpleResponse, error) {
	body, err := c.SendPostRequest(fmt.Sprintf("/v2/kubernetes/clusters/%s/recycle", id), map[string]string{
		"hostname": hostname,
		"region":   c.Region,
	})
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(body)
}

// ListAvailableKubernetesVersions returns all version of kubernetes available
func (c *Client) ListAvailableKubernetesVersions() ([]KubernetesVersion, error) {
	resp, err := c.SendGetRequest("/v2/kubernetes/versions")
	if err != nil {
		return nil, decodeError(err)
	}

	kubernetes := make([]KubernetesVersion, 0)
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&kubernetes); err != nil {
		return nil, err
	}

	return kubernetes, nil
}

// ListKubernetesClusterInstances returns all cluster instances
func (c *Client) ListKubernetesClusterInstances(id string) ([]Instance, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/kubernetes/clusters/%s/instances", id))
	if err != nil {
		return nil, decodeError(err)
	}

	instances := make([]Instance, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&instances); err != nil {
		return nil, err
	}

	return instances, nil
}

// FindKubernetesClusterInstance finds a Kubernetes cluster instance by either part of the ID or part of the name
func (c *Client) FindKubernetesClusterInstance(clusterID, search string) (*Instance, error) {
	instances, err := c.ListKubernetesClusterInstances(clusterID)
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := Instance{}

	for _, value := range instances {
		if strings.EqualFold(value.Hostname, search) || value.ID == search {
			exactMatch = true
			result = value
		} else if strings.Contains(strings.ToUpper(value.Hostname), strings.ToUpper(search)) || strings.Contains(value.ID, search) {
			if !exactMatch {
				result = value
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
