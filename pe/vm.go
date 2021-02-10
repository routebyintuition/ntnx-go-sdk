package pe

import (
	"fmt"
	"net/http"
)

// VMService handles communication to CVM/PE for VM interactions
type VMService Service

type VMListRequest struct{}

type VMListResponse struct{}

// VMGetRequest provides a UUID URL parameter to get details on one VM
// provides path query parameters to request
type VMGetRequest struct {
	Params *VMGetRequestParams
	Query  *VMGetRequestQuery
}

// VMGetRequestParams provides vm search criteria - VM UUID
type VMGetRequestParams struct {
	UUID string
}

// VMGetRequestQuery provides vm search criteria - VM UUID
type VMGetRequestQuery struct {
	IncludeVMDiskConfig bool `json:"include_vm_disk_config,omitempty"`
	IncludeVMNICConfig  bool `json:"include_vm_nic_config,omitempty"`
}

// VMGetResponse gets the VM details
type VMGetResponse struct{}

// List makes the call to list virtual machines
func (vm *VMService) List(reqdata *VMListRequest) (*VMListResponse, *http.Response, error) {

	u := "vms"

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := vm.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var data *VMListResponse
	resp, err := vm.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}

// Get makes the call to get one virtual machine by UUID
func (vm *VMService) Get(reqdata *VMGetRequest) (*VMGetResponse, *http.Response, error) {

	u := fmt.Sprintf("vms/%s", reqdata.Params.UUID)

	u, err := addOptions(u, reqdata.Query)
	if err != nil {
		return nil, nil, err
	}

	req, err := vm.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var data *VMGetResponse
	resp, err := vm.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}
