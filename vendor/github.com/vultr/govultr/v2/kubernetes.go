package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const vkePath = "/v2/kubernetes/clusters"

// KubernetesService is the interface to interact with kubernetes endpoint on the Vultr API
// Link : https://www.vultr.com/api/#tag/kubernetes
type KubernetesService interface {
	CreateCluster(ctx context.Context, createReq *ClusterReq) (*Cluster, error)
	GetCluster(ctx context.Context, id string) (*Cluster, error)
	ListClusters(ctx context.Context, options *ListOptions) ([]Cluster, *Meta, error)
	UpdateCluster(ctx context.Context, vkeID string, updateReq *ClusterReqUpdate) error
	DeleteCluster(ctx context.Context, id string) error
	DeleteClusterWithResources(ctx context.Context, id string) error

	CreateNodePool(ctx context.Context, vkeID string, nodePoolReq *NodePoolReq) (*NodePool, error)
	ListNodePools(ctx context.Context, vkeID string, options *ListOptions) ([]NodePool, *Meta, error)
	GetNodePool(ctx context.Context, vkeID, nodePoolID string) (*NodePool, error)
	UpdateNodePool(ctx context.Context, vkeID, nodePoolID string, updateReq *NodePoolReqUpdate) (*NodePool, error)
	DeleteNodePool(ctx context.Context, vkeID, nodePoolID string) error

	DeleteNodePoolInstance(ctx context.Context, vkeID, nodePoolID, nodeID string) error
	RecycleNodePoolInstance(ctx context.Context, vkeID, nodePoolID, nodeID string) error

	GetKubeConfig(ctx context.Context, vkeID string) (*KubeConfig, error)
	GetVersions(ctx context.Context) (*Versions, error)
}

// KubernetesHandler handles interaction with the kubernetes methods for the Vultr API
type KubernetesHandler struct {
	client *Client
}

// Cluster represents a full VKE cluster
type Cluster struct {
	ID            string     `json:"id"`
	Label         string     `json:"label"`
	DateCreated   string     `json:"date_created"`
	ClusterSubnet string     `json:"cluster_subnet"`
	ServiceSubnet string     `json:"service_subnet"`
	IP            string     `json:"ip"`
	Endpoint      string     `json:"endpoint"`
	Version       string     `json:"version"`
	Region        string     `json:"region"`
	Status        string     `json:"status"`
	NodePools     []NodePool `json:"node_pools"`
}

// NodePool represents a pool of nodes that are grouped by their label and plan type
type NodePool struct {
	ID           string `json:"id"`
	DateCreated  string `json:"date_created"`
	DateUpdated  string `json:"date_updated"`
	Label        string `json:"label"`
	Plan         string `json:"plan"`
	Status       string `json:"status"`
	NodeQuantity int    `json:"node_quantity"`
	Nodes        []Node `json:"nodes"`
}

// Node represents a node that will live within a nodepool
type Node struct {
	ID          string `json:"id"`
	DateCreated string `json:"date_created"`
	Label       string `json:"label"`
	Status      string `json:"status"`
}

// KubeConfig will contain the kubeconfig b64 encoded
type KubeConfig struct {
	KubeConfig string `json:"kube_config"`
}

// ClusterReq struct used to create a cluster
type ClusterReq struct {
	Label     string        `json:"label"`
	Region    string        `json:"region"`
	Version   string        `json:"version"`
	NodePools []NodePoolReq `json:"node_pools"`
}

// ClusterReqUpdate struct used to update update a cluster
type ClusterReqUpdate struct {
	Label string `json:"label"`
}

// NodePoolReq struct used to create a node pool
type NodePoolReq struct {
	NodeQuantity int    `json:"node_quantity"`
	Label        string `json:"label"`
	Plan         string `json:"plan"`
}

// NodePoolReqUpdate struct used to update a node pool
type NodePoolReqUpdate struct {
	NodeQuantity int `json:"node_quantity"`
}

type vkeClustersBase struct {
	VKEClusters []Cluster `json:"vke_clusters"`
	Meta        *Meta     `json:"meta"`
}

type vkeClusterBase struct {
	VKECluster *Cluster `json:"vke_cluster"`
}

type vkeNodePoolsBase struct {
	NodePools []NodePool `json:"node_pools"`
	Meta      *Meta      `json:"meta"`
}

type vkeNodePoolBase struct {
	NodePool *NodePool `json:"node_pool"`
}

// Versions that are supported for VKE
type Versions struct {
	Versions []string `json:"versions"`
}

// CreateCluster will create a Kubernetes cluster.
func (k *KubernetesHandler) CreateCluster(ctx context.Context, createReq *ClusterReq) (*Cluster, error) {
	req, err := k.client.NewRequest(ctx, http.MethodPost, vkePath, createReq)
	if err != nil {
		return nil, err
	}

	var k8 = new(vkeClusterBase)
	if err = k.client.DoWithContext(ctx, req, &k8); err != nil {
		return nil, err
	}

	return k8.VKECluster, nil
}

// GetCluster will return a Kubernetes cluster.
func (k *KubernetesHandler) GetCluster(ctx context.Context, id string) (*Cluster, error) {
	req, err := k.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s", vkePath, id), nil)
	if err != nil {
		return nil, err
	}

	k8 := new(vkeClusterBase)
	if err = k.client.DoWithContext(ctx, req, &k8); err != nil {
		return nil, err
	}

	return k8.VKECluster, nil
}

// ListClusters will return all kubernetes clusters.
func (k *KubernetesHandler) ListClusters(ctx context.Context, options *ListOptions) ([]Cluster, *Meta, error) {
	req, err := k.client.NewRequest(ctx, http.MethodGet, vkePath, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	k8s := new(vkeClustersBase)
	if err = k.client.DoWithContext(ctx, req, &k8s); err != nil {
		return nil, nil, err
	}

	return k8s.VKEClusters, k8s.Meta, nil
}

// UpdateCluster updates label on VKE cluster
func (k *KubernetesHandler) UpdateCluster(ctx context.Context, vkeID string, updateReq *ClusterReqUpdate) error {
	req, err := k.client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("%s/%s", vkePath, vkeID), updateReq)
	if err != nil {
		return err
	}

	return k.client.DoWithContext(ctx, req, nil)
}

// DeleteCluster will delete a Kubernetes cluster.
func (k *KubernetesHandler) DeleteCluster(ctx context.Context, id string) error {
	req, err := k.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s", vkePath, id), nil)
	if err != nil {
		return err
	}

	return k.client.DoWithContext(ctx, req, nil)
}

// DeleteClusterWithResources will delete a Kubernetes cluster and all related resources.
func (k *KubernetesHandler) DeleteClusterWithResources(ctx context.Context, id string) error {
	req, err := k.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s/delete-with-linked-resources", vkePath, id), nil)
	if err != nil {
		return err
	}

	return k.client.DoWithContext(ctx, req, nil)
}

// CreateNodePool creates a nodepool on a VKE cluster
func (k *KubernetesHandler) CreateNodePool(ctx context.Context, vkeID string, nodePoolReq *NodePoolReq) (*NodePool, error) {
	req, err := k.client.NewRequest(ctx, http.MethodPost, fmt.Sprintf("%s/%s/node-pools", vkePath, vkeID), nodePoolReq)
	if err != nil {
		return nil, err
	}

	n := new(vkeNodePoolBase)
	err = k.client.DoWithContext(ctx, req, n)
	if err != nil {
		return nil, err
	}

	return n.NodePool, nil
}

// ListNodePools will return all nodepools for a given VKE cluster
func (k *KubernetesHandler) ListNodePools(ctx context.Context, vkeID string, options *ListOptions) ([]NodePool, *Meta, error) {
	req, err := k.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s/node-pools", vkePath, vkeID), nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	n := new(vkeNodePoolsBase)
	if err = k.client.DoWithContext(ctx, req, &n); err != nil {
		return nil, nil, err
	}

	return n.NodePools, n.Meta, nil
}

// GetNodePool will return a single nodepool
func (k *KubernetesHandler) GetNodePool(ctx context.Context, vkeID, nodePoolID string) (*NodePool, error) {
	req, err := k.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s/node-pools/%s", vkePath, vkeID, nodePoolID), nil)
	if err != nil {
		return nil, err
	}

	n := new(vkeNodePoolBase)
	if err = k.client.DoWithContext(ctx, req, &n); err != nil {
		return nil, err
	}

	return n.NodePool, nil
}

// UpdateNodePool will allow you change the quantity of nodes within a nodepool
func (k *KubernetesHandler) UpdateNodePool(ctx context.Context, vkeID, nodePoolID string, updateReq *NodePoolReqUpdate) (*NodePool, error) {
	req, err := k.client.NewRequest(ctx, http.MethodPatch, fmt.Sprintf("%s/%s/node-pools/%s", vkePath, vkeID, nodePoolID), updateReq)
	if err != nil {
		return nil, err
	}

	np := new(vkeNodePoolBase)
	if err = k.client.DoWithContext(ctx, req, np); err != nil {
		return nil, err
	}

	return np.NodePool, nil
}

// DeleteNodePool will remove a nodepool from a VKE cluster
func (k *KubernetesHandler) DeleteNodePool(ctx context.Context, vkeID, nodePoolID string) error {
	req, err := k.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s/node-pools/%s", vkePath, vkeID, nodePoolID), nil)
	if err != nil {
		return err
	}

	return k.client.DoWithContext(ctx, req, nil)
}

// DeleteNodePoolInstance will remove a specified node from a nodepool
func (k *KubernetesHandler) DeleteNodePoolInstance(ctx context.Context, vkeID, nodePoolID, nodeID string) error {
	req, err := k.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s/node-pools/%s/nodes/%s", vkePath, vkeID, nodePoolID, nodeID), nil)
	if err != nil {
		return err
	}

	return k.client.DoWithContext(ctx, req, nil)
}

// RecycleNodePoolInstance will recycle (destroy + redeploy) a given node on a nodepool
func (k *KubernetesHandler) RecycleNodePoolInstance(ctx context.Context, vkeID, nodePoolID, nodeID string) error {
	req, err := k.client.NewRequest(ctx, http.MethodPost, fmt.Sprintf("%s/%s/node-pools/%s/nodes/%s/recycle", vkePath, vkeID, nodePoolID, nodeID), nil)
	if err != nil {
		return err
	}

	return k.client.DoWithContext(ctx, req, nil)
}

// GetKubeConfig returns the kubeconfig for the specified VKE cluster
func (k *KubernetesHandler) GetKubeConfig(ctx context.Context, vkeID string) (*KubeConfig, error) {
	req, err := k.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s/config", vkePath, vkeID), nil)
	if err != nil {
		return nil, err
	}

	kc := new(KubeConfig)
	if err = k.client.DoWithContext(ctx, req, &kc); err != nil {
		return nil, err
	}

	return kc, nil
}

// GetVersions returns the supported kubernetes versions
func (k *KubernetesHandler) GetVersions(ctx context.Context) (*Versions, error) {
	uri := "/v2/kubernetes/versions"
	req, err := k.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	versions := new(Versions)
	if err = k.client.DoWithContext(ctx, req, &versions); err != nil {
		return nil, err
	}

	return versions, nil
}
