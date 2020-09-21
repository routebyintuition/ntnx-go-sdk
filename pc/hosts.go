package pc

import (
	"errors"
	"fmt"
	"net/http"
)

// HostService handles the API communications to the hosts service
type HostService Service

// HostListRequest defintes the host list request
type HostListRequest struct {
	Kind          string `json:"kind,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Length        int    `json:"length,omitempty"`
	Filter        string `json:"filter,omitempty"`
	SortOrder     string `json:"sort_order,omitempty"`
	SortAttribute string `json:"sort_attribute,omitempty"`
}

// HostListResponse full list of hosts with details
type HostListResponse struct {
	APIVersion string     `json:"api_version,omitempty"`
	Metadata   Metadata   `json:"metadata,omitempty"`
	Entities   []Entities `json:"entities,omitempty"`
}

// HostGetRequest gets a host based upon
type HostGetRequest struct {
	UUID string
}

// HostGetResponse returns host details
type HostGetResponse struct {
	APIVersion string   `json:"api_version"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
	Status     Status   `json:"status"`
}

// Get passes as UUID of a specific cluster to get information about that cluster
func (hs *HostService) Get(reqdata *HostGetRequest) (*HostGetResponse, *http.Response, error) {
	if reqdata == nil {
		return nil, nil, errors.New("HostGetRequest cannot be nil")
	}

	if len(reqdata.UUID) < 8 {
		return nil, nil, errors.New("invalid UUID format in HostGetRequest")
	}

	u := fmt.Sprintf("hosts/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := hs.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var host *HostGetResponse
	resp, err := hs.client.Do(req, &host)
	if err != nil {
		return nil, resp, err
	}

	return host, resp, nil
}

// List returns all of the cluster hosts
func (hs *HostService) List(reqdata *HostListRequest) (*HostListResponse, *http.Response, error) {
	u := fmt.Sprintf("hosts/list")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := hs.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var host *HostListResponse

	resp, err := hs.client.Do(req, &host)
	if err != nil {
		return nil, resp, err
	}

	return host, resp, nil
}
