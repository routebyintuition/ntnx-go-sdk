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

// VMCreateRequest is the JSON sent to the REST API
// creates a new VM with provided configuration
type VMCreateRequest struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
}

// VMCreateResponse is the response recieved from the REST API request
type VMCreateResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// VMUpdateRequest is an update request
type VMUpdateRequest struct {
	UUID string
	Data VMUpdateRequestData
}

// VMUpdateRequestData is the configuration of the VM being updated
type VMUpdateRequestData struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
}

// VMUpdateResponse is an update response
type VMUpdateResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// VMGetRequest is a get request - This operation gets an existing VM by UUID
type VMGetRequest struct {
	UUID string
}

// VMGetResponse is the get response
type VMGetResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// VMDeleteRequest is the delete request - This operation submits a request to delete an existing VM.
type VMDeleteRequest struct {
	UUID string
}

// VMDeleteResponse is the delete response
type VMDeleteResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// VMCloneRequest is the clone request
type VMCloneRequest struct {
	UUID string
	Data VMCloneRequestData
}

// VMCloneRequestData is the data passed in the request body
type VMCloneRequestData struct {
	Metadata VMCloneRequestMetadata `json:"metadata,omitempty"`
}

// VMCloneRequestMetadata is the metadata passed to the clone request in the request body
type VMCloneRequestMetadata struct {
	UUID          string `json:"uuid,omitempty"`
	EntityVersion string `json:"entity_version,omitempty"`
}

// VMCloneResponse is the clone response
type VMCloneResponse struct {
	TaskUUID string `json:"task_uuid,omitempty"`
}

// VMResumeReplicationRequest is the request to API to resume replication for a given VM protected using sync protection policy.
type VMResumeReplicationRequest struct {
	UUID string
}

// VMResumeReplicationResponse response task UUID to resume replication
type VMResumeReplicationResponse struct {
	TaskUUID string `json:"task_uuid,omitempty"`
}

// VMPauseReplicationRequest is the request to API to pause replication for a given VM protected using sync protection policy.
type VMPauseReplicationRequest struct {
	UUID string
}

// VMPauseReplicationResponse is the response to a pause request
type VMPauseReplicationResponse struct {
	TaskUUID string `json:"task_uuid,omitempty"`
}

// VMUpdateIPsRequest is the UUID and associated data to Request a new IP address the currently allocated IP address.
type VMUpdateIPsRequest struct {
	UUID string
	Data VMUpdateIPsRequestData
}

// VMUpdateIPsRequestData is the slice of updates requested
type VMUpdateIPsRequestData struct {
	UpdateList []VMUpdateIPsRequestUpdateList `json:"update_list,omitempty"`
}

// VMUpdateIPsRequestUpdateList is each updated for an address requested
type VMUpdateIPsRequestUpdateList struct {
	UUID string `json:"uuid,omitempty"`
	IP   string `json:"ip,omitempty"`
}

// VMUpdateIPsResponse task id being processed for the update request
type VMUpdateIPsResponse struct {
	TaskUUID string `json:"task_uuid,omitempty"`
}

// Create This operation submits a request to create a new VM based on the input parameters.
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

	req, err := vms.client.NewRequest("PUT", u, reqdata)
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
// This operation submits a request to delete an existing VM.
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

	req, err := vms.client.NewRequest("POST", u, reqdata.Data)
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

// ResumeReplication API to resume replication for a given VM protected using sync protection policy.
func (vms *VMService) ResumeReplication(reqdata *VMResumeReplicationRequest) (*VMResumeReplicationResponse, *http.Response, error) {

	u := fmt.Sprintf("vms/%s/resume_replication", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := vms.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var vmdata *VMResumeReplicationResponse
	resp, err := vms.client.Do(req, &vmdata)
	if err != nil {
		return nil, resp, err
	}

	return vmdata, resp, nil
}

// PauseReplication API to pause replication for a given VM protected using sync protection policy.
func (vms *VMService) PauseReplication(reqdata *VMPauseReplicationRequest) (*VMPauseReplicationResponse, *http.Response, error) {

	u := fmt.Sprintf("vms/%s/pause_replication", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := vms.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var vmdata *VMPauseReplicationResponse
	resp, err := vms.client.Do(req, &vmdata)
	if err != nil {
		return nil, resp, err
	}

	return vmdata, resp, nil
}

// UpdateIPs Request a new IP address the currently allocated IP address.
func (vms *VMService) UpdateIPs(reqdata *VMUpdateIPsRequest) (*VMUpdateIPsResponse, *http.Response, error) {

	u := fmt.Sprintf("vms/%s/pause_replication", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := vms.client.NewRequest("PUT", u, reqdata.Data)
	if err != nil {
		return nil, nil, err
	}

	var vmdata *VMUpdateIPsResponse
	resp, err := vms.client.Do(req, &vmdata)
	if err != nil {
		return nil, resp, err
	}

	return vmdata, resp, nil
}
