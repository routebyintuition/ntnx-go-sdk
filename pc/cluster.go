package pc

import (
	"errors"
	"fmt"
	"net/http"
)

// ClusterService handles communication to the clusters REST API endpoint
type ClusterService Service

// ClusterListRequest is the request JSON sent to the API
type ClusterListRequest struct {
	Filter        string `json:"filter,omitempty"`
	Kind          string `json:"kind,omitempty"`
	SortOrder     string `json:"sort_order,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Length        int    `json:"length,omitempty"`
	SortAttribute string `json:"sort_attribute,omitempty"`
}

// ClusterListResponse is the response received from the REST API
type ClusterListResponse struct {
	Entities   []Entities `json:"entities,omitempty"`
	APIVersion string     `json:"api_version,omitempty"`
	Metadata   Metadata   `json:"metadata,omitempty"`
}

// ClusterGetRequest is the request used to identify a specific cluster to get...really just need the cluster UUID
type ClusterGetRequest struct {
	UUID string `json:"uuid,omitempty"`
}

// ClusterGetResponse are the details associated with a cluster
type ClusterGetResponse struct {
	Status     Status   `json:"status,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
}

// ClusterUpdateRequest is used to update a cluster. The UUID of the cluster is needed along with
// the other data elements related to the update
type ClusterUpdateRequest struct {
	UUID       string   `json:"uuid,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
}

// ClusterUpdateResponse is the response from the update request
type ClusterUpdateResponse struct {
	Status     Status   `json:"status,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
}

// List makes the call to cluster list
func (cs *ClusterService) List(reqdata *ClusterListRequest) (*ClusterListResponse, *http.Response, error) {
	u := "clusters/list"

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var clusters *ClusterListResponse
	resp, err := cs.client.Do(req, &clusters)
	if err != nil {
		return nil, resp, err
	}

	return clusters, resp, nil
}

// Get passes as UUID of a specific cluster to get information about that cluster
func (cs *ClusterService) Get(reqdata *ClusterGetRequest) (*ClusterGetResponse, *http.Response, error) {
	if reqdata == nil {
		return nil, nil, errors.New("ClusterGetRequest cannot be nil")
	}

	if len(reqdata.UUID) < 8 {
		return nil, nil, errors.New("invalid UUID format in ClusterGetRequest")
	}

	u := fmt.Sprintf("clusters/%s", reqdata.UUID)

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

// Update will update the cluster defined by the UUID with the provided information
func (cs *ClusterService) Update(reqdata *ClusterUpdateRequest) (*ClusterUpdateResponse, *http.Response, error) {

	u := fmt.Sprintf("clusters/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest("PUT", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var cluster *ClusterUpdateResponse
	resp, err := cs.client.Do(req, &cluster)
	if err != nil {
		return nil, resp, err
	}

	return cluster, resp, nil
}
