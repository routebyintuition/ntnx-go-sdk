package pc

import (
	"fmt"
	"net/http"
)

// VMService handles communication to the VM REST API endpoint
type VMService Service

// VMListRequest is the request sent to the REST API
// Filter must use the FIQL notation, so like "filter": "power_state==on"
type VMListRequest struct {
	Filter        string `json:"filter,omitempty"`
	Kind          string `json:"kind,omitempty"`
	SortOrder     string `json:"sort_order,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Length        int    `json:"length,omitempty"`
	SortAttribute string `json:"sort_attribute,omitempty"`
}

// VMListResponse is the response recieved from the REST API
type VMListResponse struct {
	APIVersion string     `json:"api_version,omitempty"`
	Metadata   Metadata   `json:"metadata,omitempty"`
	Entities   []Entities `json:"entities,omitempty"`
}

// VMCreateResponse is the response recieved from the REST API request
type VMCreateResponse struct {
}

// VMCreateRequest is the JSON sent to the REST API
type VMCreateRequest struct {
}

// VMUpdateRequest is an update request
type VMUpdateRequest struct {
	UUID string
}

// VMUpdateResponse is an update response
type VMUpdateResponse struct {
}

// VMGetRequest is a get request
type VMGetRequest struct {
	UUID string
}

// VMGetResponse is the get response
type VMGetResponse struct {
}

// VMDeleteRequest is the delete request
type VMDeleteRequest struct {
	UUID string
}

// VMDeleteResponse is the delete response
type VMDeleteResponse struct {
}

// VMCloneRequest is the clone request
type VMCloneRequest struct {
	Metadata `json:"metadata,omitempty"`
}

// VMCloneResponse is the clone response
type VMCloneResponse struct {
}

// List makes the call to cluster list
func (vms *VMService) List(reqdata *VMListRequest) (*VMListResponse, *http.Response, error) {

	u := "vms/list"

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := vms.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var vmdata *VMListResponse
	resp, err := vms.client.Do(req, &vmdata)
	if err != nil {
		return nil, resp, err
	}

	return vmdata, resp, nil
}

// Update makes the call to cluster list
func (vms *VMService) Update(reqdata *VMUpdateRequest) (*VMUpdateResponse, *http.Response, error) {

	u := fmt.Sprintf("vms/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := vms.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var vmdata *VMUpdateResponse
	resp, err := vms.client.Do(req, &vmdata)
	if err != nil {
		return nil, resp, err
	}

	return vmdata, resp, nil
}

// Create makes the call to create a new virtual machine
func (vms *VMService) Create(reqdata *VMCreateRequest) (*VMCreateResponse, *http.Response, error) {

	u := "vms"

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := vms.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var vmdata *VMCreateResponse
	resp, err := vms.client.Do(req, &vmdata)
	if err != nil {
		return nil, resp, err
	}

	return vmdata, resp, nil
}

// Get retreives information about a specific virtual machine
func (vms *VMService) Get(reqdata *VMGetRequest) (*VMGetResponse, *http.Response, error) {

	u := fmt.Sprintf("vms/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := vms.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var vmdata *VMGetResponse
	resp, err := vms.client.Do(req, &vmdata)
	if err != nil {
		return nil, resp, err
	}

	return vmdata, resp, nil
}

// Delete will delete the virtual machine associated with the provided UUID
func (vms *VMService) Delete(reqdata *VMDeleteRequest) (*VMDeleteResponse, *http.Response, error) {

	u := fmt.Sprintf("vms/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := vms.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var vmdata *VMDeleteResponse
	resp, err := vms.client.Do(req, &vmdata)
	if err != nil {
		return nil, resp, err
	}

	return vmdata, resp, nil
}

// Clone will clone a virtual machine identified by UUID
func (vms *VMService) Clone(reqdata *VMCloneRequest) (*VMCloneResponse, *http.Response, error) {

	u := fmt.Sprintf("vms/%s/clone", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := vms.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var vmdata *VMCloneResponse
	resp, err := vms.client.Do(req, &vmdata)
	if err != nil {
		return nil, resp, err
	}

	return vmdata, resp, nil
}
