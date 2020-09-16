package karbon

import (
	"errors"
	"fmt"
	"net/http"
)

// ClusterService handles communication to the clusters REST API endpoint
type ClusterService Service

type ClusterUpdateSecretRequest struct {
	Password                string `json:"password"`
	Username                string `json:"username"`
	PrismElementClusterUUID string `json:"prism_element_cluster_uuid"`
}
type ClusterUpdateSecretResponse struct {
	TaskUUID string `json:"task_uuid"`
}

// ClusterListRegistriesRequest requests a list of registries by cluster name
type ClusterListRegistriesRequest struct {
	ClusterName string `json:"cluster_name"`
}

// ClusterListRegistriesResponse provides the list of cluster registries by cluster name
type ClusterListRegistriesResponse []struct {
	Endpoint string `json:"endpoint"`
	Name     string `json:"name"`
	UUID     string `json:"uuid"`
}

type ClusterAddRegistryRequest struct{}
type ClusterAddRegistryResponse struct{}

type ClusterRemoveRegistryRequest struct{}
type ClusterRemoveRegistryResponse struct{}

type ClusterGetTasksRequest struct{}
type ClusterGetTasksResponse struct{}

// ClusterListRequest requests a full list of k8s clusters
type ClusterListRequest struct{}

// ClusterListResponse returns a full list of the clusters along with cluster details in a slice
type ClusterListResponse []struct {
	EtcdConfig               EtcdConfig   `json:"etcd_config"`
	KubeapiServerIpv4Address string       `json:"kubeapi_server_ipv4_address"`
	MasterConfig             MasterConfig `json:"master_config"`
	Name                     string       `json:"name"`
	Status                   string       `json:"status"`
	UUID                     string       `json:"uuid"`
	Version                  string       `json:"version"`
	WorkerConfig             WorkerConfig `json:"worker_config"`
}

type ClusterGetNodePoolRequest struct{}
type ClusterGetNodePoolResponse struct{}

// ClusterCreateRequest Create a k8s cluster with the provided configuration.
type ClusterCreateRequest struct {
	CniConfig          CniConfig          `json:"cni_config"`
	EtcdConfig         CreateEtcdConfig   `json:"etcd_config"`
	MastersConfig      MastersConfig      `json:"masters_config"`
	Metadata           Metadata           `json:"metadata"`
	Name               string             `json:"name"`
	StorageClassConfig StorageClassConfig `json:"storage_class_config"`
	Version            string             `json:"version"`
	WorkersConfig      WorkersConfig      `json:"workers_config"`
}

// ClusterCreateResponse cluster create response of new cluster
type ClusterCreateResponse struct {
	ClusterName string `json:"cluster_name"`
	ClusterUUID string `json:"cluster_uuid"`
	TaskUUID    string `json:"task_uuid"`
}

// ClusterGetRequest provides a cluster name to get details on the cluster
type ClusterGetRequest struct {
	ClusterName string `json:"cluster_name"`
}

// ClusterGetResponse returns details about a k8s cluster based on provided cluster name
type ClusterGetResponse struct {
	EtcdConfig               EtcdConfig   `json:"etcd_config"`
	KubeapiServerIpv4Address string       `json:"kubeapi_server_ipv4_address"`
	MasterConfig             MasterConfig `json:"master_config"`
	Name                     string       `json:"name"`
	Status                   string       `json:"status"`
	UUID                     string       `json:"uuid"`
	Version                  string       `json:"version"`
	WorkerConfig             WorkerConfig `json:"worker_config"`
}

// ClusterDeleteRequest deletes the identified cluster by name with or without skipping pre-checks
type ClusterDeleteRequest struct {
	ClusterName string `json:"cluster_name"`
	SkipChecks  bool   `json:"skip_checks"`
}

// ClusterDeleteResponse is the task and cluster information for the deleted resource
type ClusterDeleteResponse struct {
	ClusterName string `json:"cluster_name"`
	ClusterUUID string `json:"cluster_uuid"`
	TaskUUID    string `json:"task_uuid"`
}

// ClusterGetHealthRequest requests cluster health by cluster name for k8s
type ClusterGetHealthRequest struct {
	ClusterName string `json:"cluster_name"`
}

// ClusterGetHealthResponse provides the cluster health retrieved by cluster name
type ClusterGetHealthResponse struct {
	Messages []string `json:"messages"`
	Status   bool     `json:"status"`
}

// ClusterGetKubeConfigRequest requests KubeConfig for CLUSTER_NAME
type ClusterGetKubeConfigRequest struct {
	ClusterName string
}

// ClusterGetKubeConfigResponse returns the KubeConfig data to be used directly or placed in a file
type ClusterGetKubeConfigResponse struct {
	KubeConfig string `json:"kube_config"`
}

// ClusterGetSSHRequest Get SSH credentials to remotely access nodes belonging to the k8s cluster. The credentials have an expiry time of 24 hours.
type ClusterGetSSHRequest struct {
	ClusterName string `json:"cluster_name"`
}

// ClusterGetSSHResponse SSH cluster credentials with 24 hour expiration time
type ClusterGetSSHResponse struct {
	Certificate string `json:"certificate"`
	ExpiryTime  string `json:"expiry_time"`
	PrivateKey  string `json:"private_key"`
	Username    string `json:"username"`
}

// Create makes the call to cluster create
func (cs *ClusterService) Create(reqdata *ClusterCreateRequest) (*ClusterCreateResponse, *http.Response, error) {
	u := "v1/k8s/clusters"

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var cluster *ClusterCreateResponse
	resp, err := cs.client.Do(req, &cluster)
	if err != nil {
		return nil, resp, err
	}

	return cluster, resp, nil
}

// UpdateSecret the k8s secret and docker volume plugin password across k8s clusters deployed in Karbon. The update operation is performed only for the clusters where the provided Prism Element UUID and the username match.
func (cs *ClusterService) UpdateSecret(reqdata *ClusterUpdateSecretRequest) (*ClusterUpdateSecretResponse, *http.Response, error) {

	u := "v1-alpha.1/k8s/clusters/storage-auth"

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest("PUT", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var cluster *ClusterUpdateSecretResponse
	resp, err := cs.client.Do(req, &cluster)
	if err != nil {
		return nil, resp, err
	}

	return cluster, resp, nil
}

// GetKubeConfig Get the kubeconfig to access the K8s cluster.
func (cs *ClusterService) GetKubeConfig(reqdata *ClusterGetKubeConfigRequest) (*ClusterGetKubeConfigResponse, *http.Response, error) {
	if reqdata == nil {
		return nil, nil, errors.New("ClusterGetKubeConfigRequest cannot be nil")
	}

	if len(reqdata.ClusterName) < 3 {
		return nil, nil, errors.New("invalid cluster_name format in ClusterGetKubeConfigRequest")
	}

	u := fmt.Sprintf("v1/k8s/clusters/%s/kubeconfig", reqdata.ClusterName)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cluster *ClusterGetKubeConfigResponse
	resp, err := cs.client.Do(req, &cluster)
	if err != nil {
		return nil, resp, err
	}

	return cluster, resp, nil
}

// GetHealth will return cluster health for cluster name
func (cs *ClusterService) GetHealth(reqdata *ClusterGetHealthRequest) (*ClusterGetHealthResponse, *http.Response, error) {
	if reqdata == nil {
		return nil, nil, errors.New("GetHealth cannot be nil")
	}

	if len(reqdata.ClusterName) < 3 {
		return nil, nil, errors.New("invalid cluster_name format in GetHealth")
	}

	u := fmt.Sprintf("v1/k8s/clusters/%s/health", reqdata.ClusterName)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cluster *ClusterGetHealthResponse
	resp, err := cs.client.Do(req, &cluster)
	if err != nil {
		return nil, resp, err
	}

	return cluster, resp, nil
}

// ListRegistries will list the private registries associated with a cluster
func (cs *ClusterService) ListRegistries(reqdata *ClusterListRegistriesRequest) (*ClusterListRegistriesResponse, *http.Response, error) {
	if reqdata == nil {
		return nil, nil, errors.New("ClusterListRegistriesRequest cannot be nil")
	}

	if len(reqdata.ClusterName) < 3 {
		return nil, nil, errors.New("invalid cluster_name format in ClusterListRegistriesRequest")
	}

	u := fmt.Sprintf("v1-alpha.1/k8s/clusters/%s/registries", reqdata.ClusterName)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cluster *ClusterListRegistriesResponse
	resp, err := cs.client.Do(req, &cluster)
	if err != nil {
		return nil, resp, err
	}

	return cluster, resp, nil
}

// List Get all the k8s cluster objects, which include worker, etcd and master configuration details.
func (cs *ClusterService) List(reqdata *ClusterListRequest) (*ClusterListResponse, *http.Response, error) {

	u := "v1-beta.1/k8s/clusters"

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cluster *ClusterListResponse
	resp, err := cs.client.Do(req, &cluster)
	if err != nil {
		return nil, resp, err
	}

	return cluster, resp, nil
}

// Get Get all the k8s cluster objects, which include worker, etcd and master configuration details.
func (cs *ClusterService) Get(reqdata *ClusterGetRequest) (*ClusterGetResponse, *http.Response, error) {

	u := fmt.Sprintf("v1/k8s/clusters/%s", reqdata.ClusterName)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cluster *ClusterGetResponse
	resp, err := cs.client.Do(req, &cluster)
	if err != nil {
		return nil, resp, err
	}

	return cluster, resp, nil
}

// GetSSH Gets the associated k8s SSH credentials to remotely access the k8s cluster. Credentials expire in 24 hours.
func (cs *ClusterService) GetSSH(reqdata *ClusterGetSSHRequest) (*ClusterGetSSHResponse, *http.Response, error) {

	u := fmt.Sprintf("v1/k8s/clusters/%s/ssh", reqdata.ClusterName)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cluster *ClusterGetSSHResponse
	resp, err := cs.client.Do(req, &cluster)
	if err != nil {
		return nil, resp, err
	}

	return cluster, resp, nil
}

// Delete identified cluster by name defining whether to allow cluster deletion during cluster deployment with skip-pre-check
func (cs *ClusterService) Delete(reqdata *ClusterDeleteRequest) (*ClusterDeleteResponse, *http.Response, error) {

	u := fmt.Sprintf("v1/k8s/clusters/%s", reqdata.ClusterName)

	skip := DeleteSkipCheck{SkipCheck: reqdata.SkipChecks}

	u, err := addOptions(u, skip)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cluster *ClusterDeleteResponse
	resp, err := cs.client.Do(req, &cluster)
	if err != nil {
		return nil, resp, err
	}

	return cluster, resp, nil
}
