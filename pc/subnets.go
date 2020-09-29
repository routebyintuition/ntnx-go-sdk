package pc

import (
	"fmt"
	"net/http"
)

// SubnetService service for operations against the subnet prism central API
type SubnetService Service

// SubnetCreateRequest is the definition of the subnet to create
type SubnetCreateRequest struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
}

// SubnetCreateResponse is the response with definition for creating subnets
type SubnetCreateResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// SubnetListRequest contains parameters for the list request
type SubnetListRequest struct {
	Kind          string `json:"kind,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Length        int    `json:"length,omitempty"`
	Filter        string `json:"filter,omitempty"`
	SortOrder     string `json:"sort_order,omitempty"`
	SortAttribute string `json:"sort_attribute,omitempty"`
}

// SubnetListResponse contains a list and details about each subnet
type SubnetListResponse struct {
	APIVersion string     `json:"api_version,omitempty"`
	Metadata   Metadata   `json:"metadata,omitempty"`
	Entities   []Entities `json:"entities,omitempty"`
}

// SubnetGetRequest provides UUID of subnet to get
type SubnetGetRequest struct {
	UUID string
}

// SubnetGetResponse details on subnet identified by UUID
type SubnetGetResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// SubnetUpdateRequest provides UUID and subnet data to update
type SubnetUpdateRequest struct {
	UUID string
	Data SubnetUpdateRequestData
}

// SubnetUpdateRequestData new details about subnet to provide for update
type SubnetUpdateRequestData struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
}

// SubnetUpdateResponse provides details about subnet update
type SubnetUpdateResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// SubnetDeleteRequest provides UUID of subnet to delete
type SubnetDeleteRequest struct {
	UUID string
}

// SubnetDeleteResponse provides details from delete subnet request
type SubnetDeleteResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// Delete - This operation submits a request to delete a existing subnet.
func (ss *SubnetService) Delete(reqdata *SubnetDeleteRequest) (*SubnetDeleteResponse, *http.Response, error) {

	u := fmt.Sprintf("subnets/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ss.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *SubnetDeleteResponse
	resp, err := ss.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Update - This operation submits a request to update a existing subnet based on the input parameters.
func (ss *SubnetService) Update(reqdata *SubnetUpdateRequest) (*SubnetUpdateResponse, *http.Response, error) {

	u := fmt.Sprintf("subnets/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ss.client.NewRequest("PUT", u, reqdata.Data)
	if err != nil {
		return nil, nil, err
	}

	var result *SubnetUpdateResponse
	resp, err := ss.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Get - This operation gets a existing subnet.
func (ss *SubnetService) Get(reqdata *SubnetGetRequest) (*SubnetGetResponse, *http.Response, error) {

	u := fmt.Sprintf("subnets/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ss.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *SubnetGetResponse
	resp, err := ss.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// List - This operation gets a list of subnets, allowing for sorting and pagination. Note: Entities that have not been created successfully are not listed.
func (ss *SubnetService) List(reqdata *SubnetListRequest) (*SubnetListResponse, *http.Response, error) {

	u := "subnets/list"

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ss.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var result *SubnetListResponse
	resp, err := ss.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Create - This operation submits a request to create a new subnet based on the input parameters. A subnet is a block of IP addresses.
func (ss *SubnetService) Create(reqdata *SubnetCreateRequest) (*SubnetCreateResponse, *http.Response, error) {

	u := "subnets"

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ss.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var result *SubnetCreateResponse
	resp, err := ss.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
