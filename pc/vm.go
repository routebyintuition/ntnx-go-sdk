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

// ClusterReference is the ClusterReference
type ClusterReference struct {
	Kind string `json:"kind,omitempty"`
	UUID string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
}

// DiskAddress is the DiskAddress
type DiskAddress struct {
	DeviceIndex int    `json:"device_index,omitempty"`
	AdapterType string `json:"adapter_type,omitempty"`
}

// DeviceProperties is the DeviceProperties
type DeviceProperties struct {
	DiskAddress DiskAddress `json:"disk_address,omitempty"`
	DeviceType  string      `json:"device_type,omitempty"`
}

// VnumaConfig is the VnumaConfig
type VnumaConfig struct {
	NumVnumaNodes int `json:"num_vnuma_nodes,omitempty"`
}

// IPEndpointList is the IPEndpointList
type IPEndpointList struct {
	IP   string `json:"ip,omitempty"`
	Type string `json:"type,omitempty"`
}

// SubnetReference is the SubnetReference
type SubnetReference struct {
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

// NicList is the NicList
type NicList struct {
	NicType         string           `json:"nic_type,omitempty"`
	UUID            string           `json:"uuid,omitempty"`
	IPEndpointList  []IPEndpointList `json:"ip_endpoint_list,omitempty"`
	MacAddress      string           `json:"mac_address,omitempty"`
	SubnetReference SubnetReference  `json:"subnet_reference,omitempty"`
	IsConnected     bool             `json:"is_connected,omitempty"`
}

// HostReference is the HostReference
type HostReference struct {
	Kind string `json:"kind,omitempty"`
	UUID string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
}

// DataSourceReference is the DataSourceReference
type DataSourceReference struct {
	Kind string `json:"kind,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

// DiskList is the DiskList
type DiskList struct {
	DataSourceReference DataSourceReference `json:"data_source_reference,omitempty"`
	DeviceProperties    DeviceProperties    `json:"device_properties,omitempty"`
	UUID                string              `json:"uuid,omitempty"`
	DiskSizeBytes       int                 `json:"disk_size_bytes,omitempty"`
	DiskSizeMib         int                 `json:"disk_size_mib,omitempty"`
}

// ExecutionContext is the ExecutionContext
type ExecutionContext struct {
	TaskUuids []string `json:"task_uuids,omitempty"`
}

// BootDevice is the BootDevice
type BootDevice struct {
	DiskAddress DiskAddress `json:"disk_address,omitempty"`
}

// BootConfig is the BootConfig
type BootConfig struct {
	BootDevice BootDevice `json:"boot_device,omitempty"`
}

// CloudInit is the CloudInit
type CloudInit struct {
	UserData string `json:"user_data,omitempty"`
}

// GuestCustomization is the GuestCustomization
type GuestCustomization struct {
	CloudInit     CloudInit `json:"cloud_init,omitempty"`
	IsOverridable bool      `json:"is_overridable,omitempty"`
}

// GuestTransitionConfig is the GuestTransitionConfig
type GuestTransitionConfig struct {
	ShouldFailOnScriptFailure bool `json:"should_fail_on_script_failure,omitempty"`
	EnableScriptExec          bool `json:"enable_script_exec,omitempty"`
}

// PowerStateMechanism is the PowerStateMechanism
type PowerStateMechanism struct {
	GuestTransitionConfig GuestTransitionConfig `json:"guest_transition_config,omitempty"`
	Mechanism             string                `json:"mechanism,omitempty"`
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
