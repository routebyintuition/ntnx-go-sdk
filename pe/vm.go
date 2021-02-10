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
	IncludeVMDiskConfig bool `url:"include_vm_disk_config,omitempty"`
	IncludeVMNICConfig  bool `url:"include_vm_nic_config,omitempty"`
}

// VMGetResponse gets the VM details
type VMGetResponse struct {
	AllowLiveMigrate      *bool                  `json:"allow_live_migrate,omitempty"`
	Boot                  *Boot                  `json:"boot,omitempty"`
	CbrNotCapableReason   *string                `json:"cbr_not_capable_reason,omitempty"`
	ClearAffinity         *bool                  `json:"clear_affinity,omitempty"`
	Description           *string                `json:"description,omitempty"`
	GpusAssigned          *bool                  `json:"gpus_assigned,omitempty"`
	GuestDriverVersion    *string                `json:"guest_driver_version,omitempty"`
	GuestOs               *string                `json:"guest_os,omitempty"`
	HaPriority            *int                   `json:"ha_priority,omitempty"`
	HostUUID              *string                `json:"host_uuid,omitempty"`
	MachineType           *string                `json:"machine_type,omitempty"`
	MemoryMb              *int                   `json:"memory_mb,omitempty"`
	MemoryReservationMb   *int                   `json:"memory_reservation_mb,omitempty"`
	Name                  *string                `json:"name,omitempty"`
	NumCoresPerVcpu       *int                   `json:"num_cores_per_vcpu,omitempty"`
	NumVcpus              *int                   `json:"num_vcpus,omitempty"`
	PowerState            *string                `json:"power_state,omitempty"`
	SerialPorts           *[]SerialPorts         `json:"serial_ports,omitempty"`
	StorageContainerUUID  *string                `json:"storage_container_uuid,omitempty"`
	Timezone              *string                `json:"timezone,omitempty"`
	ToolsInstallerMounted *bool                  `json:"tools_installer_mounted,omitempty"`
	ToolsRunningStatus    *string                `json:"tools_running_status,omitempty"`
	UUID                  *string                `json:"uuid,omitempty"`
	VcpuReservationHz     *int                   `json:"vcpu_reservation_hz,omitempty"`
	VMCustomizationConfig *VMCustomizationConfig `json:"vm_customization_config,omitempty"`
	VMDiskInfo            *[]VMDiskInfo          `json:"vm_disk_info,omitempty"`
	VMDisks               *[]VMDisks             `json:"vm_disks,omitempty"`
	VMFeatures            *VMFeatures            `json:"vm_features,omitempty"`
	VMGpus                *[]VMGpus              `json:"vm_gpus,omitempty"`
	VMLogicalTimestamp    *int                   `json:"vm_logical_timestamp,omitempty"`
	VMNics                *[]VMNics              `json:"vm_nics,omitempty"`
}

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
