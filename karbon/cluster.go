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

type ClusterListRegistriesRequest struct {
	ClusterName string `json:"cluster_name"`
}
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

type ClusterListRequest struct{}
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

type ClusterGetRequest struct{}
type ClusterGetResponse struct{}

type ClusterDeleteRequest struct{}
type ClusterDeleteResponse struct{}

type ClusterGetHealthRequest struct{}
type ClusterGetHealthResponse struct{}

type ClusterGetKubeConfigRequest struct {
	ClusterName string
}
type ClusterGetKubeConfigResponse struct {
	KubeConfig string `json:"kube_config"`
}

type ClusterGetSSHRequest struct{}
type ClusterGetSSHResponse struct{}

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
