package cs

import (
	"net/http"
	"net/url"

	"math"
	"strconv"
	"time"

	"fmt"

	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/ecs"
	"github.com/denverdino/aliyungo/util"
)

type ClusterState string

const (
	Initial      = ClusterState("initial")
	Running      = ClusterState("running")
	Updating     = ClusterState("updating")
	Scaling      = ClusterState("scaling")
	Failed       = ClusterState("failed")
	Deleting     = ClusterState("deleting")
	DeleteFailed = ClusterState("deleteFailed")
	Deleted      = ClusterState("deleted")
	InActive     = ClusterState("inactive")
)

type NodeStatus struct {
	Health   int64 `json:"health"`
	Unhealth int64 `json:"unhealth"`
}

type NetworkModeType string

const (
	ClassicNetwork = NetworkModeType("classic")
	VPCNetwork     = NetworkModeType("vpc")
)

// https://help.aliyun.com/document_detail/26053.html
type ClusterType struct {
	AgentVersion           string           `json:"agent_version"`
	ClusterID              string           `json:"cluster_id"`
	Name                   string           `json:"name"`
	Created                util.ISO6801Time `json:"created"`
	ExternalLoadbalancerID string           `json:"external_loadbalancer_id"`
	MasterURL              string           `json:"master_url"`
	NetworkMode            NetworkModeType  `json:"network_mode"`
	RegionID               common.Region    `json:"region_id"`
	SecurityGroupID        string           `json:"security_group_id"`
	Size                   int64            `json:"size"`
	State                  ClusterState     `json:"state"`
	Updated                util.ISO6801Time `json:"updated"`
	VPCID                  string           `json:"vpc_id"`
	VSwitchID              string           `json:"vswitch_id"`
	NodeStatus             string           `json:"node_status"`
	DockerVersion          string           `json:"docker_version"`
}

func (client *Client) DescribeClusters(nameFilter string) (clusters []ClusterType, err error) {
	query := make(url.Values)

	if nameFilter != "" {
		query.Add("name", nameFilter)
	}

	err = client.Invoke("", http.MethodGet, "/clusters", query, nil, &clusters)
	return
}

func (client *Client) DescribeCluster(id string) (cluster ClusterType, err error) {
	err = client.Invoke("", http.MethodGet, "/clusters/"+id, nil, nil, &cluster)
	return
}

type ClusterCreationArgs struct {
	Name             string           `json:"name"`
	Size             int64            `json:"size"`
	NetworkMode      NetworkModeType  `json:"network_mode"`
	SubnetCIDR       string           `json:"subnet_cidr,omitempty"`
	InstanceType     string           `json:"instance_type"`
	VPCID            string           `json:"vpc_id,omitempty"`
	VSwitchID        string           `json:"vswitch_id,omitempty"`
	Password         string           `json:"password"`
	DataDiskSize     int64            `json:"data_disk_size"`
	DataDiskCategory ecs.DiskCategory `json:"data_disk_category"`
	ECSImageID       string           `json:"ecs_image_id,omitempty"`
	IOOptimized      ecs.IoOptimized  `json:"io_optimized"`
	ReleaseEipFlag   bool             `json:"release_eip_flag"`
}

type ClusterCreationResponse struct {
	Response
	ClusterID string `json:"cluster_id"`
}

func (client *Client) CreateCluster(region common.Region, args *ClusterCreationArgs) (cluster ClusterCreationResponse, err error) {
	err = client.Invoke(region, http.MethodPost, "/clusters", nil, args, &cluster)
	return
}

type KubernetesStackArgs struct {
	VPCID                    string           `json:"VpcId,omitempty"`
	VSwitchID                string           `json:"VSwitchId,omitempty"`
	MasterInstanceType       string           `json:"MasterInstanceType"`
	WorkerInstanceType       string           `json:"WorkerInstanceType"`
	NumOfNodes               int64            `json:"NumOfNodes"`
	Password                 string           `json:"LoginPassword"`
	DockerVersion            string           `json:"DockerVersion"`
	KubernetesVersion        string           `json:"KubernetesVersion"`
	ZoneId                   string           `json:"ZoneId"`
	ContainerCIDR            string           `json:"ContainerCIDR"`
	ServiceCIDR              string           `json:"ServiceCIDR"`
	SSHFlags                 bool             `json:"SSHFlags"`
	MasterSystemDiskSize     int64            `json:"MasterSystemDiskSize"`
	MasterSystemDiskCategory ecs.DiskCategory `json:"MasterSystemDiskCategory"`
	WorkerSystemDiskSize     int64            `json:"WorkerSystemDiskSize"`
	WorkerSystemDiskCategory ecs.DiskCategory `json:"WorkerSystemDiskCategory"`
	ImageID                  string           `json:"ImageId,omitempty"`
	CloudMonitorFlags        bool             `json:"CloudMonitorFlags"`
	SNatEntry                bool             `json:"SNatEntry"`
}

type KubernetesCreationArgs struct {
	ClusterType       string              `json:"cluster_type"`
	Name              string              `json:"name"`
	DisableRollback   bool                `json:"disable_rollback"`
	TimeoutMins       int64               `json:"timeout_mins"`
	KubernetesVersion string              `json:"kubernetes_version"`
	StackParams       KubernetesStackArgs `json:"stack_params"`
}

type KubernetesMultiAZCreationArgs struct {
	DisableRollback          bool             `json:"disable_rollback"`
	Name                     string           `json:"name"`
	TimeoutMins              int64            `json:"timeout_mins"`
	ClusterType              string           `json:"cluster_type"`
	MultiAZ                  bool             `json:"multi_az,omitempty"`
	VPCID                    string           `json:"vpcid,omitempty"`
	ContainerCIDR            string           `json:"container_cidr"`
	ServiceCIDR              string           `json:"service_cidr"`
	VSwitchIdA               string           `json:"vswitch_id_a,omitempty"`
	VSwitchIdB               string           `json:"vswitch_id_b,omitempty"`
	VSwitchIdC               string           `json:"vswitch_id_c,omitempty"`
	MasterInstanceTypeA      string           `json:"master_instance_type_a,omitempty"`
	MasterInstanceTypeB      string           `json:"master_instance_type_b,omitempty"`
	MasterInstanceTypeC      string           `json:"master_instance_type_c,omitempty"`
	MasterSystemDiskSize     int64            `json:"master_system_disk_size"`
	MasterSystemDiskCategory ecs.DiskCategory `json:"master_system_disk_category"`
	WorkerInstanceTypeA      string           `json:"worker_instance_type_a,omitempty"`
	WorkerInstanceTypeB      string           `json:"worker_instance_type_b,omitempty"`
	WorkerInstanceTypeC      string           `json:"worker_instance_type_c,omitempty"`
	WorkerSystemDiskSize     int64            `json:"worker_system_disk_size"`
	WorkerSystemDiskCategory ecs.DiskCategory `json:"worker_system_disk_category"`
	NumOfNodesA              int64            `json:"num_of_nodes_a"`
	NumOfNodesB              int64            `json:"num_of_nodes_b"`
	NumOfNodesC              int64            `json:"num_of_nodes_c"`
	SSHFlags                 bool             `json:"ssh_flags"`
	CloudMonitorFlags        bool             `json:"cloud_monitor_flags"`
	LoginPassword            string           `json:"login_password,omitempty"`
	ImageId                  string           `json:"image_id,omitempty"`
	KeyPair                  string           `json:"key_pair,omitempty"`
}

func (client *Client) CreateKubernetesMultiAZCluster(region common.Region, args *KubernetesMultiAZCreationArgs) (cluster ClusterCreationResponse, err error) {
	err = client.Invoke(region, http.MethodPost, "/clusters", nil, args, &cluster)
	return
}

func (client *Client) CreateKubernetesCluster(region common.Region, args *KubernetesCreationArgs) (cluster ClusterCreationResponse, err error) {
	err = client.Invoke(region, http.MethodPost, "/clusters", nil, args, &cluster)
	return
}

type ClusterResizeArgs struct {
	Size             int64            `json:"size"`
	InstanceType     string           `json:"instance_type"`
	Password         string           `json:"password"`
	DataDiskSize     int64            `json:"data_disk_size"`
	DataDiskCategory ecs.DiskCategory `json:"data_disk_category"`
	ECSImageID       string           `json:"ecs_image_id,omitempty"`
	IOOptimized      ecs.IoOptimized  `json:"io_optimized"`
}

type ModifyClusterNameArgs struct {
	Name string `json:"name"`
}

func (client *Client) ResizeCluster(clusterID string, args *ClusterResizeArgs) error {
	return client.Invoke("", http.MethodPut, "/clusters/"+clusterID, nil, args, nil)
}

func (client *Client) ResizeKubernetes(clusterID string, args *KubernetesCreationArgs) error {
	return client.Invoke("", http.MethodPut, "/clusters/"+clusterID, nil, args, nil)
}

func (client *Client) ModifyClusterName(clusterID, clusterName string) error {
	return client.Invoke("", http.MethodPost, "/clusters/"+clusterID+"/name/"+clusterName, nil, nil, nil)
}

func (client *Client) DeleteCluster(clusterID string) error {
	return client.Invoke("", http.MethodDelete, "/clusters/"+clusterID, nil, nil, nil)
}

type ClusterCerts struct {
	CA   string `json:"ca,omitempty"`
	Key  string `json:"key,omitempty"`
	Cert string `json:"cert,omitempty"`
}

func (client *Client) GetClusterCerts(id string) (certs ClusterCerts, err error) {
	err = client.Invoke("", http.MethodGet, "/clusters/"+id+"/certs", nil, nil, &certs)
	return
}

type ClusterConfig struct {
	Config string `json:"config"`
}

func (client *Client) GetClusterConfig(id string) (config ClusterConfig, err error) {
	err = client.Invoke("", http.MethodGet, "/k8s/"+id+"/user_config", nil, nil, &config)
	return
}

type KubernetesNodeType struct {
	InstanceType       string   `json:"instance_type"`
	IpAddress          []string `json:"ip_address"`
	InstanceChargeType string   `json:"instance_charge_type"`
	InstanceRole       string   `json:"instance_role"`
	CreationTime       string   `json:"creation_time"`
	InstanceName       string   `json:"instance_name"`
	InstanceTypeFamily string   `json:"instance_type_family"`
	HostName           string   `json:"host_name"`
	ImageId            string   `json:"image_id"`
	InstanceId         string   `json:"instance_id"`
}

type GetKubernetesClusterNodesResponse struct {
	Response
	Page  PaginationResult     `json:"page"`
	Nodes []KubernetesNodeType `json:"nodes"`
}

func (client *Client) GetKubernetesClusterNodes(id string, pagination common.Pagination) (nodes []KubernetesNodeType, paginationResult *PaginationResult, err error) {
	response := &GetKubernetesClusterNodesResponse{}
	err = client.Invoke("", http.MethodGet, "/clusters/"+id+"/nodes?pageNumber="+strconv.Itoa(pagination.PageNumber)+"&pageSize="+strconv.Itoa(pagination.PageSize), nil, nil, &response)
	if err != nil {
		return nil, nil, err
	}

	return response.Nodes, &response.Page, nil
}

const ClusterDefaultTimeout = 300
const DefaultWaitForInterval = 10
const DefaultPreSleepTime = 240

// WaitForCluster waits for instance to given status
// when instance.NotFound wait until timeout
func (client *Client) WaitForClusterAsyn(clusterId string, status ClusterState, timeout int) error {
	if timeout <= 0 {
		timeout = ClusterDefaultTimeout
	}
	cluster, err := client.DescribeCluster(clusterId)
	if err != nil {
		return err
	} else if cluster.State == status {
		//TODO
		return nil
	}
	// Create or Reset cluster usually cost at least 4 min, so there will sleep a long time before polling
	sleep := math.Min(float64(timeout), float64(DefaultPreSleepTime))
	time.Sleep(time.Duration(sleep) * time.Second)

	for {
		cluster, err := client.DescribeCluster(clusterId)
		if err != nil {
			return err
		} else if cluster.State == Failed {
			return fmt.Errorf("Waitting for cluster %s %s failed. Looking the specified reason in the web console.", clusterId, status)
		} else if cluster.State == status {
			//TODO
			break
		}
		timeout = timeout - DefaultWaitForInterval
		if timeout <= 0 {
			return common.GetClientErrorFromString("Timeout")
		}
		time.Sleep(DefaultWaitForInterval * time.Second)
	}
	return nil
}

func (client *Client) GetProjectClient(clusterId string) (projectClient *ProjectClient, err error) {
	cluster, err := client.DescribeCluster(clusterId)
	if err != nil {
		return
	}

	certs, err := client.GetClusterCerts(clusterId)
	if err != nil {
		return
	}

	projectClient, err = NewProjectClient(clusterId, cluster.MasterURL, certs)

	if err != nil {
		return
	}

	projectClient.SetDebug(client.debug)

	return
}
