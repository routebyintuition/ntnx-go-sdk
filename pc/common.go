package pc

import "time"

// Metadata is the Metadata associated with an API responsee
// this same type can be used across the calls as it includes
// type entries for other API calls
type Metadata struct {
	TotalMatches         *int               `json:"total_matches,omitempty" yaml:"total_matches,omitempty"`
	Kind                 *string            `json:"kind,omitempty" yaml:"kind,omitempty"`
	Length               *int               `json:"length,omitempty" yaml:"length,omitempty"`
	Offset               *int               `json:"offset,omitempty" yaml:"offset,omitempty"`
	LastUpdateTime       *time.Time         `json:"last_update_time,omitempty" yaml:"last_update_time,omitempty"`
	UUID                 *string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	ProjectReference     *ProjectReference  `json:"project_reference,omitempty" yaml:"project_reference,omitempty"`
	SpecVersion          *int               `json:"spec_version,omitempty" yaml:"spec_version,omitempty"`
	CreationTime         *time.Time         `json:"creation_time,omitempty" yaml:"creation_time,omitempty"`
	OwnerReference       *OwnerReference    `json:"owner_reference,omitempty" yaml:"owner_reference,omitempty"`
	Categories           *Categories        `json:"categories,omitempty" yaml:"categories,omitempty"`
	SpecHash             *string            `json:"spec_hash,omitempty" yaml:"spec_hash,omitempty"`
	ShouldForceTranslate *bool              `json:"should_force_translate,omitempty" yaml:"should_force_translate,omitempty"`
	Name                 *string            `json:"name,omitempty" yaml:"name,omitempty"`
	Filter               *string            `json:"filter,omitempty" yaml:"filter,omitempty"`
	SortOrder            *string            `json:"sort_order,omitempty" yaml:"sort_order,omitempty"`
	SortAttribute        *string            `json:"sort_attribute,omitempty" yaml:"sort_attribute,omitempty"`
	CategoriesMapping    *CategoriesMapping `json:"categories_mapping,omitempty" yaml:"categories_mapping,omitempty"`
	UseCategoriesMapping *bool              `json:"use_categories_mapping,omitempty" yaml:"use_categories_mapping,omitempty"`
	EntityVersion        *string            `json:"entity_version,omitempty" yaml:"entity_version,omitempty"`
	GroupMemberCount     *int               `json:"group_member_count,omitempty" yaml:"group_member_count,omitempty"`
	GroupMemberOffset    *int               `json:"group_member_offset,omitempty" yaml:"group_member_offset,omitempty"`
	UsageType            *string            `json:"usage_type,omitempty" yaml:"usage_type,omitempty"`
}

// Config is the API response for Config used across all API requests
type Config struct {
	AuthorizedPublicKeyList       []AuthorizedPublicKeyList `json:"authorized_public_key_list,omitempty" yaml:"authorized_public_key_list,omitempty"`
	Build                         Build                     `json:"build,omitempty" yaml:"build,omitempty"`
	CaCertificateList             []CaCertificateList       `json:"ca_certificate_list,omitempty" yaml:"ca_certificate_list,omitempty"`
	CertificationSigningInfo      CertificationSigningInfo  `json:"certification_signing_info,omitempty" yaml:"certification_signing_info,omitempty"`
	ClientAuth                    ClientAuth                `json:"client_auth,omitempty" yaml:"client_auth,omitempty"`
	ClusterArch                   string                    `json:"cluster_arch,omitempty" yaml:"cluster_arch,omitempty"`
	DomainAwarenessLevel          string                    `json:"domain_awareness_level,omitempty" yaml:"domain_awareness_level,omitempty"`
	EnabledFeatureList            []string                  `json:"enabled_feature_list,omitempty" yaml:"enabled_feature_list,omitempty"`
	EncryptionStatus              string                    `json:"encryption_status,omitempty" yaml:"encryption_status,omitempty"`
	ExternalConfigurations        ExternalConfigurations    `json:"external_configurations,omitempty" yaml:"external_configurations,omitempty"`
	GpuDriverVersion              string                    `json:"gpu_driver_version,omitempty" yaml:"gpu_driver_version,omitempty"`
	IsAvailable                   bool                      `json:"is_available,omitempty" yaml:"is_available,omitempty"`
	ManagementServerList          []ManagementServerList    `json:"management_server_list,omitempty" yaml:"management_server_list,omitempty"`
	OperationMode                 string                    `json:"operation_mode,omitempty" yaml:"operation_mode,omitempty"`
	RedundancyFactor              int                       `json:"redundancy_factor,omitempty" yaml:"redundancy_factor,omitempty"`
	ServiceList                   []string                  `json:"service_list,omitempty" yaml:"service_list,omitempty"`
	SoftwareMap                   SoftwareMap               `json:"software_map,omitempty" yaml:"software_map,omitempty"`
	SslKey                        SslKey                    `json:"ssl_key,omitempty" yaml:"ssl_key,omitempty"`
	SupportedInformationVerbosity string                    `json:"supported_information_verbosity,omitempty" yaml:"supported_information_verbosity,omitempty"`
	Timezone                      string                    `json:"timezone,omitempty" yaml:"timezone,omitempty"`
}

// Resources is the main Resources type across all Response calls
// used for JSON and YAML decoder
type Resources struct {
	RetrievalURIList                 *[]string                           `json:"retrieval_uri_list,omitempty" yaml:"retrieval_uri_list,omitempty"`
	SizeBytes                        *int                                `json:"size_bytes,omitempty" yaml:"size_bytes,omitempty"`
	CurrentClusterReferenceList      *[]CurrentClusterReferenceList      `json:"current_cluster_reference_list,omitempty" yaml:"current_cluster_reference_list,omitempty"`
	ImagePlacementPolicyStateList    *[]ImagePlacementPolicyStateList    `json:"image_placement_policy_state_list,omitempty" yaml:"image_placement_policy_state_list,omitempty"`
	ImageType                        *string                             `json:"image_type,omitempty" yaml:"image_type,omitempty"`
	SourceURI                        *string                             `json:"source_uri,omitempty" yaml:"source_uri,omitempty"`
	SourceOptions                    *SourceOptions                      `json:"source_options,omitempty" yaml:"source_options,omitempty"`
	InitialPlacementRefList          *[]InitialPlacementRefList          `json:"initial_placement_ref_list,omitempty" yaml:"initial_placement_ref_list,omitempty"`
	DataSourceReference              *DataSourceReference                `json:"data_source_reference,omitempty" yaml:"data_source_reference,omitempty"`
	Version                          *Version                            `json:"version,omitempty" yaml:"version,omitempty"`
	Architecture                     *string                             `json:"architecture,omitempty" yaml:"architecture,omitempty"`
	BootConfig                       *BootConfig                         `json:"boot_config,omitempty" yaml:"boot_config,omitempty"`
	DiskList                         *[]DiskList                         `json:"disk_list,omitempty" yaml:"disk_list,omitempty"`
	GuestCustomization               *GuestCustomization                 `json:"guest_customization,omitempty" yaml:"guest_customization,omitempty"`
	HardwareClockTimezone            *string                             `json:"hardware_clock_timezone,omitempty" yaml:"hardware_clock_timezone,omitempty"`
	HostReference                    *HostReference                      `json:"host_reference,omitempty" yaml:"host_reference,omitempty"`
	HypervisorType                   *string                             `json:"hypervisor_type,omitempty" yaml:"hypervisor_type,omitempty"`
	MemorySizeMib                    *int                                `json:"memory_size_mib,omitempty" yaml:"memory_size_mib,omitempty"`
	NicList                          *[]NicList                          `json:"nic_list,omitempty" yaml:"nic_list,omitempty"`
	NumSockets                       *int                                `json:"num_sockets,omitempty" yaml:"num_sockets,omitempty"`
	NumThreadsPerCore                *int                                `json:"num_threads_per_core,omitempty" yaml:"num_threads_per_core,omitempty"`
	NumVcpusPerSocket                *int                                `json:"num_vcpus_per_socket,omitempty" yaml:"num_vcpus_per_socket,omitempty"`
	PowerState                       *string                             `json:"power_state,omitempty" yaml:"power_state,omitempty"`
	PowerStateMechanism              *PowerStateMechanism                `json:"power_state_mechanism,omitempty" yaml:"power_state_mechanism,omitempty"`
	ProtectionType                   *string                             `json:"protection_type,omitempty" yaml:"protection_type,omitempty"`
	SerialPortList                   *[]interface{}                      `json:"serial_port_list,omitempty" yaml:"serial_port_list,omitempty"`
	VgaConsoleEnabled                *bool                               `json:"vga_console_enabled,omitempty" yaml:"vga_console_enabled,omitempty"`
	VnumaConfig                      *VnumaConfig                        `json:"vnuma_config,omitempty" yaml:"vnuma_config,omitempty"`
	Nodes                            *Nodes                              `json:"nodes,omitempty" yaml:"nodes,omitempty"`
	Config                           *Config                             `json:"config,omitempty" yaml:"config,omitempty"`
	Network                          *Network                            `json:"network,omitempty" yaml:"network,omitempty"`
	Analysis                         *Analysis                           `json:"analysis,omitempty" yaml:"analysis,omitempty"`
	RuntimeStatusList                *[]string                           `json:"runtime_status_list,omitempty" yaml:"runtime_status_list,omitempty"`
	CPUModel                         *string                             `json:"cpu_model,omitempty" yaml:"cpu_model,omitempty"`
	CPUCapacityHz                    *int                                `json:"cpu_capacity_hz,omitempty" yaml:"cpu_capacity_hz,omitempty"`
	NumCPUCores                      *int                                `json:"num_cpu_cores,omitempty" yaml:"num_cpu_cores,omitempty"`
	NumCPUSockets                    *int                                `json:"num_cpu_sockets,omitempty" yaml:"num_cpu_sockets,omitempty"`
	MemoryCapacityMib                *int                                `json:"memory_capacity_mib,omitempty" yaml:"memory_capacity_mib,omitempty"`
	SerialNumber                     *string                             `json:"serial_number,omitempty" yaml:"serial_number,omitempty"`
	MonitoringState                  *string                             `json:"monitoring_state,omitempty" yaml:"monitoring_state,omitempty"`
	HostNicsIDList                   *[]string                           `json:"host_nics_id_list,omitempty" yaml:"host_nics_id_list,omitempty"`
	WindowsDomain                    *WindowsDomain                      `json:"windows_domain,omitempty" yaml:"windows_domain,omitempty"`
	FailoverCluster                  *string                             `json:"failover_cluster,omitempty" yaml:"failover_cluster,omitempty"`
	Ipmi                             *Ipmi                               `json:"ipmi,omitempty" yaml:"ipmi,omitempty"`
	ControllerVM                     *ControllerVM                       `json:"controller_vm,omitempty" yaml:"controller_vm,omitempty"`
	Hypervisor                       *Hypervisor                         `json:"hypervisor,omitempty" yaml:"hypervisor,omitempty"`
	Block                            *Block                              `json:"block,omitempty" yaml:"block,omitempty"`
	RackableUnitReference            *RackableUnitReference              `json:"rackable_unit_reference,omitempty" yaml:"rackable_unit_reference,omitempty"`
	HostDisksReferenceList           *[]HostDisksReferenceList           `json:"host_disks_reference_list,omitempty" yaml:"host_disks_reference_list,omitempty"`
	GpuList                          *[]GpuList                          `json:"gpu_list,omitempty" yaml:"gpu_list,omitempty"`
	GpuDriverVersion                 *string                             `json:"gpu_driver_version,omitempty" yaml:"gpu_driver_version,omitempty"`
	HostType                         *string                             `json:"host_type,omitempty" yaml:"host_type,omitempty"`
	GuestOsID                        *string                             `json:"guest_os_id,omitempty" yaml:"guest_os_id,omitempty"`
	IsVcpuHardPinned                 *bool                               `json:"is_vcpu_hard_pinned,omitempty" yaml:"is_vcpu_hard_pinned,omitempty"`
	IsAgentVM                        *bool                               `json:"is_agent_vm,omitempty" yaml:"is_agent_vm,omitempty"`
	DisableBranding                  *bool                               `json:"disable_branding,omitempty" yaml:"disable_branding,omitempty"`
	EnableCPUPassthrough             *bool                               `json:"enable_cpu_passthrough,omitempty" yaml:"enable_cpu_passthrough,omitempty"`
	MachineType                      *string                             `json:"machine_type,omitempty" yaml:"machine_type,omitempty"`
	HardwareVirtualizationEnabled    *bool                               `json:"hardware_virtualization_enabled,omitempty" yaml:"hardware_virtualization_enabled,omitempty"`
	ParentReference                  *ParentReference                    `json:"parent_reference,omitempty" yaml:"parent_reference,omitempty"`
	GuestTools                       *GuestTools                         `json:"guest_tools,omitempty" yaml:"guest_tools,omitempty"`
	StorageConfig                    *StorageConfig                      `json:"storage_config,omitempty" yaml:"storage_config,omitempty"`
	DirectoryServiceUser             *DirectoryServiceUser               `json:"directory_service_user,omitempty" yaml:"directory_service_user,omitempty"`
	IdentityProviderUser             *IdentityProviderUser               `json:"identity_provider_user,omitempty" yaml:"identity_provider_user,omitempty"`
	UserType                         *string                             `json:"user_type,omitempty" yaml:"user_type,omitempty"`
	DisplayName                      *string                             `json:"display_name,omitempty" yaml:"display_name,omitempty"`
	ProjectsReferenceList            *[]ProjectsReferenceList            `json:"projects_reference_list,omitempty" yaml:"projects_reference_list,omitempty"`
	AccessControlPolicyReferenceList *[]AccessControlPolicyReferenceList `json:"access_control_policy_reference_list,omitempty" yaml:"access_control_policy_reference_list,omitempty"`
	ResourceUsageSummary             *ResourceUsageSummary               `json:"resource_usage_summary,omitempty" yaml:"resource_usage_summary,omitempty"`
	SubnetType                       *string                             `json:"subnet_type,omitempty" yaml:"subnet_type,omitempty"`
	VlanID                           *int                                `json:"vlan_id,omitempty" yaml:"vlan_id,omitempty"`
	IPConfig                         *IPConfig                           `json:"ip_config,omitempty" yaml:"ip_config,omitempty"`
	VswitchName                      *string                             `json:"vswitch_name,omitempty" yaml:"vswitch_name,omitempty"`
	NetworkFunctionChainReference    *NetworkFunctionChainReference      `json:"network_function_chain_reference,omitempty" yaml:"network_function_chain_reference,omitempty"`
	VirtualNetworkReference          *VirtualNetworkReference            `json:"virtual_network_reference,omitempty" yaml:"virtual_network_reference,omitempty"`
	VpcReference                     *VpcReference                       `json:"vpc_reference,omitempty" yaml:"vpc_reference,omitempty"`
	AvailabilityZoneReferenceList    *[]AvailabilityZoneReferenceList    `json:"availability_zone_reference_list,omitempty" yaml:"availability_zone_reference_list,omitempty"`
	ExternalConnectivityState        *string                             `json:"external_connectivity_state,omitempty" yaml:"external_connectivity_state,omitempty"`
	ResourceDomain                   *string                             `json:"resource_domain,omitempty" yaml:"resource_domain,omitempty"`
	DefaultSubnetReference           *DefaultSubnetReference             `json:"default_subnet_reference,omitempty" yaml:"default_subnet_reference,omitempty"`
	SubnetReferenceList              *[]SubnetReferenceList              `json:"subnet_reference_list,omitempty" yaml:"subnet_reference_list,omitempty"`
	UserReferenceList                *[]UserReferenceList                `json:"user_reference_list,omitempty" yaml:"user_reference_list,omitempty"`
	ExternalUserGroupReferenceList   *[]ExternalUserGroupReferenceList   `json:"external_user_group_reference_list,omitempty" yaml:"external_user_group_reference_list,omitempty"`
	AccountReferenceList             *[]AccountReferenceList             `json:"account_reference_list,omitempty" yaml:"account_reference_list,omitempty"`
	TunnelReferenceList              *[]TunnelReferenceList              `json:"tunnel_reference_list,omitempty" yaml:"tunnel_reference_list,omitempty"`
	ExternalNetworkList              *[]ExternalNetworkList              `json:"external_network_list,omitempty" yaml:"external_network_list,omitempty"`
	EnvironmentReferenceList         *[]EnvironmentReferenceList         `json:"environment_reference_list,omitempty" yaml:"environment_reference_list,omitempty"`
	DefaultEnvironmentReference      *DefaultEnvironmentReference        `json:"default_environment_reference,omitempty" yaml:"default_environment_reference,omitempty"`
	PermissionReferenceList          *[]PermissionReferenceList          `json:"permission_reference_list,omitempty" yaml:"permission_reference_list,omitempty"`
	Operation                        *string                             `json:"operation,omitempty" yaml:"operation,omitempty"`
	Kind                             *string                             `json:"kind,omitempty" yaml:"kind,omitempty"`
	Fields                           *Fields                             `json:"fields,omitempty" yaml:"fields,omitempty"`
	AllowIpv6Traffic                 *bool                               `json:"allow_ipv6_traffic,omitempty" yaml:"allow_ipv_6_traffic,omitempty"`
	IsPolicyHitlogEnabled            *bool                               `json:"is_policy_hitlog_enabled,omitempty" yaml:"is_policy_hitlog_enabled,omitempty"`
	AdRule                           *AdRule                             `json:"ad_rule,omitempty" yaml:"ad_rule,omitempty"`
	AppRule                          *AppRule                            `json:"app_rule,omitempty" yaml:"app_rule,omitempty"`
	IsolationRule                    *IsolationRule                      `json:"isolation_rule,omitempty" yaml:"isolation_rule,omitempty"`
	QuarantineRule                   *QuarantineRule                     `json:"quarantine_rule,omitempty" yaml:"quarantine_rule,omitempty"`
}

// Spec is the main Spec type across all Response calls
type Spec struct {
	ClusterReference          *ClusterReference          `json:"cluster_reference,omitempty" yaml:"cluster_reference,omitempty"`
	Description               *string                    `json:"description,omitempty" yaml:"description,omitempty"`
	Resources                 *Resources                 `json:"resources,omitempty" yaml:"resources,omitempty"`
	Name                      *string                    `json:"name,omitempty" yaml:"name,omitempty"`
	APIVersion                *string                    `json:"api_version,omitempty" yaml:"api_version,omitempty"`
	AvailabilityZoneReference *AvailabilityZoneReference `json:"availability_zone_reference,omitempty" yaml:"availability_zone_reference,omitempty"`
}

// PermissionReferenceList is used for role creation
type PermissionReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// ImageReferenceList is used for image service details
type ImageReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ProjectsReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type AccessControlPolicyReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// SourceOptions used for Image Update
type SourceOptions struct {
	AllowInsecureConnection bool `json:"allow_insecure_connection,omitempty" yaml:"allow_insecure_connection,omitempty"`
}

type ResourceUsageSummary struct {
	ResourceDomain string `json:"resource_domain,omitempty" yaml:"resource_domain,omitempty"`
}

// InitialPlacementRefList used or image update and placement
type InitialPlacementRefList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// DataSourceReference used for image update
type DataSourceReference struct {
	Kind           string `json:"kind,omitempty" yaml:"kind,omitempty"`
	UUID           string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name           string `json:"name,omitempty" yaml:"name,omitempty"`
	URL            string `json:"url,omitempty" yaml:"url,omitempty"`
	IsDirectAttach bool   `json:"is_direct_attach,omitempty" yaml:"is_direct_attach,omitempty"`
}

type VolumeGroupReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}
type StorageContainerReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

// StorageConfig VM storage definitions
type StorageConfig struct {
	StorageContainerReference *StorageContainerReference `json:"storage_container_reference,omitempty" yaml:"storage_container_reference,omitempty"`
	FlashMode                 string                     `json:"flash_mode,omitempty" yaml:"flash_mode,omitempty"`
	QosPolicy                 *QosPolicy                 `json:"qos_policy,omitempty" yaml:"qos_policy,omitempty"`
}

type DirectoryServiceReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type DirectoryServiceUser struct {
	UserPrincipalName         string                     `json:"user_principal_name,omitempty" yaml:"user_principal_name,omitempty"`
	DefaultUserPrincipalName  string                     `json:"default_user_principal_name,omitempty" yaml:"default_user_principal_name,omitempty"`
	DirectoryServiceReference *DirectoryServiceReference `json:"directory_service_reference,omitempty" yaml:"directory_service_reference,omitempty"`
}
type IdentityProviderReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type IdentityProviderUser struct {
	Username                  string                     `json:"username,omitempty" yaml:"username,omitempty"`
	IdentityProviderReference *IdentityProviderReference `json:"identity_provider_reference,omitempty" yaml:"identity_provider_reference,omitempty"`
}

// ProjectReference is the main ProjectReference type across all Response calls
type ProjectReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// ProjectResourceDomainList used for user resource/domain get information
type ProjectResourceDomainList struct {
	ProjectReference ProjectReference `json:"project_reference,omitempty" yaml:"project_reference,omitempty"`
	ResourceDomain   string           `json:"resource_domain,omitempty" yaml:"resource_domain,omitempty"`
}

// OwnerReference is the main OwnerReference type across all Response calls
type OwnerReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Categories is the main Categories type across all Response calls
type Categories struct{}

// State is the various states associated with a cluster
type State struct {
	ReadOnly    bool   `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	Type        string `json:"type,omitempty" yaml:"type,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

type AvailabilityZoneReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ImagePlacementPolicyStateList struct {
	ComplianceStatus string             `json:"compliance_status,omitempty" yaml:"compliance_status,omitempty"`
	EnforcementMode  string             `json:"enforcement_mode,omitempty" yaml:"enforcement_mode,omitempty"`
	PolicyReference  PolicyReference    `json:"policy_reference,omitempty" yaml:"policy_reference,omitempty"`
	PolicyInfo       PolicyInfo         `json:"policy_info,omitempty" yaml:"policy_info,omitempty"`
	ErrorMessageList []ErrorMessageList `json:"error_message_list,omitempty" yaml:"error_message_list,omitempty"`
}

type CurrentClusterReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type PolicyReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}
type EnforcedPolicyClusterReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type CompletePolicyClusterReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type ConflictingImagePlacementPolicyReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type PolicyInfo struct {
	EnforcedPolicyClusterReferenceList           []EnforcedPolicyClusterReferenceList           `json:"enforced_policy_cluster_reference_list,omitempty" yaml:"enforced_policy_cluster_reference_list,omitempty"`
	CompletePolicyClusterReferenceList           []CompletePolicyClusterReferenceList           `json:"complete_policy_cluster_reference_list,omitempty" yaml:"complete_policy_cluster_reference_list,omitempty"`
	ConflictingImagePlacementPolicyReferenceList []ConflictingImagePlacementPolicyReferenceList `json:"conflicting_image_placement_policy_reference_list,omitempty" yaml:"conflicting_image_placement_policy_reference_list,omitempty"`
}

type ErrorMessageList struct {
	Message string  `json:"message,omitempty" yaml:"message,omitempty"`
	Reason  string  `json:"reason,omitempty" yaml:"reason,omitempty"`
	Details Details `json:"details,omitempty" yaml:"details,omitempty"`
}

// Status is the main Status type across all Response calls
type Status struct {
	State                     string                     `json:"state,omitempty" yaml:"state,omitempty"`
	ClusterReference          *ClusterReference          `json:"cluster_reference,omitempty" yaml:"cluster_reference,omitempty"`
	ExecutionContext          *ExecutionContext          `json:"execution_context,omitempty" yaml:"execution_context,omitempty"`
	Description               string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Resources                 *Resources                 `json:"resources,omitempty" yaml:"resources,omitempty"`
	Name                      string                     `json:"name,omitempty" yaml:"name,omitempty"`
	MessageList               []MessageList              `json:"message_list,omitempty" yaml:"message_list,omitempty"`
	AvailabilityZoneReference *AvailabilityZoneReference `json:"availability_zone_reference,omitempty" yaml:"availability_zone_reference,omitempty"`
}

// Entities is the main Entities type across all Response calls
type Entities struct {
	Status               Status                 `json:"status,omitempty" yaml:"status,omitempty"`
	Spec                 Spec                   `json:"spec,omitempty" yaml:"spec,omitempty"`
	APIVersion           string                 `json:"api_version,omitempty" yaml:"api_version,omitempty"`
	Metadata             Metadata               `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Name                 string                 `json:"name,omitempty" yaml:"name,omitempty"`
	Capabilities         Capabilities           `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
	Description          string                 `json:"description,omitempty" yaml:"description,omitempty"`
	SystemDefined        bool                   `json:"system_defined,omitempty" yaml:"system_defined,omitempty"`
	UUID                 string                 `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	LogicalTimestamp     int                    `json:"logical_timestamp,omitempty" yaml:"logical_timestamp,omitempty"`
	ProgressMessage      string                 `json:"progress_message,omitempty" yaml:"progress_message,omitempty"`
	OperationType        string                 `json:"operation_type,omitempty" yaml:"operation_type,omitempty"`
	PercentageComplete   int                    `json:"percentage_complete,omitempty" yaml:"percentage_complete,omitempty"`
	ParentTaskReference  ParentTaskReference    `json:"parent_task_reference,omitempty" yaml:"parent_task_reference,omitempty"`
	SubtaskReferenceList []SubtaskReferenceList `json:"subtask_reference_list,omitempty" yaml:"subtask_reference_list,omitempty"`
	ClusterReference     *ClusterReference      `json:"cluster_reference,omitempty" yaml:"cluster_reference,omitempty"`
	CreationTime         string                 `json:"creation_time,omitempty" yaml:"creation_time,omitempty"`
	CreationTimeUsecs    int                    `json:"creation_time_usecs,omitempty" yaml:"creation_time_usecs,omitempty"`
	LastUpdateTime       string                 `json:"last_update_time,omitempty" yaml:"last_update_time,omitempty"`
	StartTime            string                 `json:"start_time,omitempty" yaml:"start_time,omitempty"`
	StartTimeUsecs       int                    `json:"start_time_usecs,omitempty" yaml:"start_time_usecs,omitempty"`
	CompletionTime       string                 `json:"completion_time,omitempty" yaml:"completion_time,omitempty"`
	CompletionTimeUsecs  int                    `json:"completion_time_usecs,omitempty" yaml:"completion_time_usecs,omitempty"`
	EntityReferenceList  []EntityReferenceList  `json:"entity_reference_list,omitempty" yaml:"entity_reference_list,omitempty"`
	ErrorCode            string                 `json:"error_code,omitempty" yaml:"error_code,omitempty"`
	ErrorDetail          string                 `json:"error_detail,omitempty" yaml:"error_detail,omitempty"`
}

// MessageList is the list of messages
type MessageList struct {
	Message string  `json:"message,omitempty" yaml:"message,omitempty"`
	Reason  string  `json:"reason,omitempty" yaml:"reason,omitempty"`
	Details Details `json:"details,omitempty" yaml:"details,omitempty"`
}

// Properties are tied to response properties
type Properties struct {
	Kind                 Kind                 `json:"kind,omitempty" yaml:"kind,omitempty"`
	ProjectReference     ProjectReference     `json:"project_reference,omitempty" yaml:"project_reference,omitempty"`
	UUID                 UUID                 `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name                 Name                 `json:"name,omitempty" yaml:"name,omitempty"`
	SpecVersion          SpecVersion          `json:"spec_version,omitempty" yaml:"spec_version,omitempty"`
	SpecHash             SpecHash             `json:"spec_hash,omitempty" yaml:"spec_hash,omitempty"`
	Categories           Categories           `json:"categories,omitempty" yaml:"categories,omitempty"`
	CategoriesMapping    CategoriesMapping    `json:"categories_mapping,omitempty" yaml:"categories_mapping,omitempty"`
	UseCategoriesMapping UseCategoriesMapping `json:"use_categories_mapping,omitempty" yaml:"use_categories_mapping,omitempty"`
	CreationTime         CreationTime         `json:"creation_time,omitempty" yaml:"creation_time,omitempty"`
	LastUpdateTime       LastUpdateTime       `json:"last_update_time,omitempty" yaml:"last_update_time,omitempty"`
	OwnerReference       OwnerReference       `json:"owner_reference,omitempty" yaml:"owner_reference,omitempty"`
	ShouldForceTranslate ShouldForceTranslate `json:"should_force_translate,omitempty" yaml:"should_force_translate,omitempty"`
	EntityVersion        EntityVersion        `json:"entity_version,omitempty" yaml:"entity_version,omitempty"`
	Offset               Offset               `json:"offset,omitempty" yaml:"offset,omitempty"`
	Length               Length               `json:"length,omitempty" yaml:"length,omitempty"`
	TotalMatches         TotalMatches         `json:"total_matches,omitempty" yaml:"total_matches,omitempty"`
	Filter               Filter               `json:"filter,omitempty" yaml:"filter,omitempty"`
	SortOrder            SortOrder            `json:"sort_order,omitempty" yaml:"sort_order,omitempty"`
	SortAttribute        SortAttribute        `json:"sort_attribute,omitempty" yaml:"sort_attribute,omitempty"`
	MessageList          MessageList          `json:"message_list,omitempty" yaml:"message_list,omitempty"`
}

type Kind struct {
	Type        string   `json:"type,omitempty" yaml:"type,omitempty"`
	ReadOnly    bool     `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
	Default     string   `json:"default,omitempty" yaml:"default,omitempty"`
	XNtnxEnum   []string `json:"x-ntnx-enum,omitempty" yaml:"x_ntnx_enum,omitempty"`
}

type ShouldForceTranslate struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Type        string `json:"type,omitempty" yaml:"type,omitempty"`
}

type EntityVersion struct {
	Type        string `json:"type,omitempty" yaml:"type,omitempty"`
	ReadOnly    bool   `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

type UUID struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Type        string `json:"type,omitempty" yaml:"type,omitempty"`
	Format      string `json:"format,omitempty" yaml:"format,omitempty"`
	Pattern     string `json:"pattern,omitempty" yaml:"pattern,omitempty"`
}

type Offset struct {
	Type        string `json:"type,omitempty" yaml:"type,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
type Length struct {
	Type        string `json:"type,omitempty" yaml:"type,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
type TotalMatches struct {
	Type        string `json:"type,omitempty" yaml:"type,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

// Filter is used in NSR and other areas of List
type Filter struct {
	Type        string   `json:"type,omitempty" yaml:"type,omitempty"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
	KindList    []string `json:"kind_list,omitempty" yaml:"kind_list,omitempty"`
	Params      Params   `json:"params,omitempty" yaml:"params,omitempty"`
}
type SortOrder struct {
	Title       string   `json:"title,omitempty" yaml:"title,omitempty"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
	Type        string   `json:"type,omitempty" yaml:"type,omitempty"`
	XNtnxEnum   []string `json:"x-ntnx-enum,omitempty" yaml:"x_ntnx_enum,omitempty"`
}
type SortAttribute struct {
	Type        string `json:"type,omitempty" yaml:"type,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
type LastUpdateTime struct {
	Type        string `json:"type,omitempty" yaml:"type,omitempty"`
	ReadOnly    bool   `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	Format      string `json:"format,omitempty" yaml:"format,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

type CreationTime struct {
	Type        string `json:"type,omitempty" yaml:"type,omitempty"`
	Format      string `json:"format,omitempty" yaml:"format,omitempty"`
	ReadOnly    bool   `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

type Name struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Type        string `json:"type,omitempty" yaml:"type,omitempty"`
	ReadOnly    bool   `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	MaxLength   int    `json:"maxLength,omitempty" yaml:"max_length,omitempty"`
}

type SpecVersion struct {
	Type        string `json:"type,omitempty" yaml:"type,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

type SpecHash struct {
	Type        string `json:"type,omitempty" yaml:"type,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

type UseCategoriesMapping struct {
	Type        string `json:"type,omitempty" yaml:"type,omitempty"`
	Default     bool   `json:"default,omitempty" yaml:"default,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

// Results is used for category query results
type Results struct {
	Kind                string              `json:"kind,omitempty" yaml:"kind,omitempty"`
	TotalEntityCount    int                 `json:"total_entity_count,omitempty" yaml:"total_entity_count,omitempty"`
	FilteredEntityCount int                 `json:"filtered_entity_count,omitempty" yaml:"filtered_entity_count,omitempty"`
	KindReferenceList   []KindReferenceList `json:"kind_reference_list,omitempty" yaml:"kind_reference_list,omitempty"`
}

// KindReferenceList is used for category query results
type KindReferenceList struct {
	Kind       string     `json:"kind,omitempty" yaml:"kind,omitempty"`
	UUID       string     `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name       string     `json:"name,omitempty" yaml:"name,omitempty"`
	Type       string     `json:"type,omitempty" yaml:"type,omitempty"`
	Categories Categories `json:"categories,omitempty" yaml:"categories,omitempty"`
}

type CategoryFilter struct {
	Type     string   `json:"type,omitempty" yaml:"type,omitempty"`
	KindList []string `json:"kind_list,omitempty" yaml:"kind_list,omitempty"`
	Params   Params   `json:"params,omitempty" yaml:"params,omitempty"`
}

type Params struct {
}

// HypervisorServerList is the hyprevisor list of servers
type HypervisorServerList struct {
	IP      string `json:"ip,omitempty" yaml:"ip,omitempty"`
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
	Type    string `json:"type,omitempty" yaml:"type,omitempty"`
}

// Nodes are the nodes
type Nodes struct {
	HypervisorServerList []HypervisorServerList `json:"hypervisor_server_list,omitempty" yaml:"hypervisor_server_list,omitempty"`
}

// ClientAuth is the PKE auth info
type ClientAuth struct {
	Status  string `json:"status,omitempty" yaml:"status,omitempty"`
	CaChain string `json:"ca_chain,omitempty" yaml:"ca_chain,omitempty"`
	Name    string `json:"name,omitempty" yaml:"name,omitempty"`
}

// SoftwareMap is the software map
type SoftwareMap struct {
}

// SigningInfo is the SigningInfo
type SigningInfo struct {
	City             string `json:"city,omitempty" yaml:"city,omitempty"`
	State            string `json:"state,omitempty" yaml:"state,omitempty"`
	CountryCode      string `json:"country_code,omitempty" yaml:"country_code,omitempty"`
	CommonName       string `json:"common_name,omitempty" yaml:"common_name,omitempty"`
	Organization     string `json:"organization,omitempty" yaml:"organization,omitempty"`
	EmailAddress     string `json:"email_address,omitempty" yaml:"email_address,omitempty"`
	CommonNameSuffix string `json:"common_name_suffix,omitempty" yaml:"common_name_suffix,omitempty"`
}

// Details are the details
type Details struct {
}

// SslKey is the SSL Key
type SslKey struct {
	KeyType        string      `json:"key_type,omitempty" yaml:"key_type,omitempty"`
	KeyName        string      `json:"key_name,omitempty" yaml:"key_name,omitempty"`
	SigningInfo    SigningInfo `json:"signing_info,omitempty" yaml:"signing_info,omitempty"`
	ExpireDatetime time.Time   `json:"expire_datetime,omitempty" yaml:"expire_datetime,omitempty"`
}

// CertificationSigningInfo is the signing info PKE
type CertificationSigningInfo struct {
	City             string `json:"city,omitempty" yaml:"city,omitempty"`
	State            string `json:"state,omitempty" yaml:"state,omitempty"`
	CountryCode      string `json:"country_code,omitempty" yaml:"country_code,omitempty"`
	CommonName       string `json:"common_name,omitempty" yaml:"common_name,omitempty"`
	Organization     string `json:"organization,omitempty" yaml:"organization,omitempty"`
	EmailAddress     string `json:"email_address,omitempty" yaml:"email_address,omitempty"`
	CommonNameSuffix string `json:"common_name_suffix,omitempty" yaml:"common_name_suffix,omitempty"`
}

// CitrixVMReferenceList is the CitrixVMReferenceList
type CitrixVMReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// ResourceLocation is the ResourceLocation
type ResourceLocation struct {
	ID   string `json:"id,omitempty" yaml:"id,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// CitrixConnectorConfig is the CitrixConnectorConfig
type CitrixConnectorConfig struct {
	CitrixVMReferenceList []CitrixVMReferenceList `json:"citrix_vm_reference_list,omitempty" yaml:"citrix_vm_reference_list,omitempty"`
	ClientSecret          string                  `json:"client_secret,omitempty" yaml:"client_secret,omitempty"`
	CustomerID            string                  `json:"customer_id,omitempty" yaml:"customer_id,omitempty"`
	ClientID              string                  `json:"client_id,omitempty" yaml:"client_id,omitempty"`
	ResourceLocation      ResourceLocation        `json:"resource_location,omitempty" yaml:"resource_location,omitempty"`
}

// ExternalConfigurations is the ExternalConfigurations
type ExternalConfigurations struct {
	CitrixConnectorConfig CitrixConnectorConfig `json:"citrix_connector_config,omitempty" yaml:"citrix_connector_config,omitempty"`
}

// CaCertificateList is the CaCertificateList
type CaCertificateList struct {
	CaName      string `json:"ca_name,omitempty" yaml:"ca_name,omitempty"`
	Certificate string `json:"certificate,omitempty" yaml:"certificate,omitempty"`
}

// ManagementServerList is the ManagementServerList
type ManagementServerList struct {
	DrsEnabled bool     `json:"drs_enabled,omitempty" yaml:"drs_enabled,omitempty"`
	IP         string   `json:"ip,omitempty" yaml:"ip,omitempty"`
	Type       string   `json:"type,omitempty" yaml:"type,omitempty"`
	StatusList []string `json:"status_list,omitempty" yaml:"status_list,omitempty"`
}

// AuthorizedPublicKeyList is the AuthorizedPublicKeyList
type AuthorizedPublicKeyList struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	Key  string `json:"key,omitempty" yaml:"key,omitempty"`
}

// Build is the Build
type Build struct {
	CommitID      string    `json:"commit_id,omitempty" yaml:"commit_id,omitempty"`
	Version       string    `json:"version,omitempty" yaml:"version,omitempty"`
	ShortCommitID string    `json:"short_commit_id,omitempty" yaml:"short_commit_id,omitempty"`
	FullVersion   string    `json:"full_version,omitempty" yaml:"full_version,omitempty"`
	CommitDate    time.Time `json:"commit_date,omitempty" yaml:"commit_date,omitempty"`
	BuildType     string    `json:"build_type,omitempty" yaml:"build_type,omitempty"`
}

// HTTPProxyWhitelist is the HTTPProxyWhitelist
type HTTPProxyWhitelist struct {
	Target     string `json:"target,omitempty" yaml:"target,omitempty"`
	TargetType string `json:"target_type,omitempty" yaml:"target_type,omitempty"`
}

// Credentials are the Credentials
type Credentials struct {
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
}

type NutanixGuestTools struct {
	NgtState              string      `json:"ngt_state,omitempty" yaml:"ngt_state,omitempty"`
	Credentials           Credentials `json:"credentials,omitempty" yaml:"credentials,omitempty"`
	State                 string      `json:"state,omitempty" yaml:"state,omitempty"`
	IsoMountState         string      `json:"iso_mount_state,omitempty" yaml:"iso_mount_state,omitempty"`
	EnabledCapabilityList []string    `json:"enabled_capability_list,omitempty" yaml:"enabled_capability_list,omitempty"`
	Version               string      `json:"version,omitempty" yaml:"version,omitempty"`
}

type GuestTools struct {
	NutanixGuestTools NutanixGuestTools `json:"nutanix_guest_tools,omitempty" yaml:"nutanix_guest_tools,omitempty"`
}

// Address is the address
type Address struct {
	IP       string `json:"ip,omitempty" yaml:"ip,omitempty"`
	Ipv6     string `json:"ipv6,omitempty" yaml:"ipv_6,omitempty"`
	IsBackup bool   `json:"is_backup,omitempty" yaml:"is_backup,omitempty"`
	Port     int    `json:"port,omitempty" yaml:"port,omitempty"`
	Fqdn     string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`
}

// HTTPProxyList is the HTTPProxyList
type HTTPProxyList struct {
	Credentials   Credentials `json:"credentials,omitempty" yaml:"credentials,omitempty"`
	Address       Address     `json:"address,omitempty" yaml:"address,omitempty"`
	ProxyTypeList []string    `json:"proxy_type_list,omitempty" yaml:"proxy_type_list,omitempty"`
}

// Server is the server
type Server struct {
	Credentials   Credentials `json:"credentials,omitempty" yaml:"credentials,omitempty"`
	Address       Address     `json:"address,omitempty" yaml:"address,omitempty"`
	ProxyTypeList []string    `json:"proxy_type_list,omitempty" yaml:"proxy_type_list,omitempty"`
}

// SMTPServer is the SMTP server
type SMTPServer struct {
	EmailAddress string `json:"email_address,omitempty" yaml:"email_address,omitempty"`
	Type         string `json:"type,omitempty" yaml:"type,omitempty"`
	Server       Server `json:"server,omitempty" yaml:"server,omitempty"`
}

// DomainCredentials are the domain credentials
type DomainCredentials struct {
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
}

// DomainServer is the domain server
type DomainServer struct {
	Nameserver        string            `json:"nameserver,omitempty" yaml:"nameserver,omitempty"`
	Name              string            `json:"name,omitempty" yaml:"name,omitempty"`
	DomainCredentials DomainCredentials `json:"domain_credentials,omitempty" yaml:"domain_credentials,omitempty"`
}

// Network is the network
type Network struct {
	HTTPProxyWhitelist     []HTTPProxyWhitelist `json:"http_proxy_whitelist,omitempty" yaml:"http_proxy_whitelist,omitempty"`
	MasqueradingPort       int                  `json:"masquerading_port,omitempty" yaml:"masquerading_port,omitempty"`
	MasqueradingIP         string               `json:"masquerading_ip,omitempty" yaml:"masquerading_ip,omitempty"`
	ExternalIP             string               `json:"external_ip,omitempty" yaml:"external_ip,omitempty"`
	HTTPProxyList          []HTTPProxyList      `json:"http_proxy_list,omitempty" yaml:"http_proxy_list,omitempty"`
	SMTPServer             SMTPServer           `json:"smtp_server,omitempty" yaml:"smtp_server,omitempty"`
	NtpServerIPList        []string             `json:"ntp_server_ip_list,omitempty" yaml:"ntp_server_ip_list,omitempty"`
	ExternalSubnet         string               `json:"external_subnet,omitempty" yaml:"external_subnet,omitempty"`
	ExternalDataServicesIP string               `json:"external_data_services_ip,omitempty" yaml:"external_data_services_ip,omitempty"`
	DomainServer           DomainServer         `json:"domain_server,omitempty" yaml:"domain_server,omitempty"`
	NameServerIPList       []string             `json:"name_server_ip_list,omitempty" yaml:"name_server_ip_list,omitempty"`
	NfsSubnetWhitelist     []string             `json:"nfs_subnet_whitelist,omitempty" yaml:"nfs_subnet_whitelist,omitempty"`
	InternalSubnet         string               `json:"internal_subnet,omitempty" yaml:"internal_subnet,omitempty"`
}

// VMEfficiencyMap is the VMEfficiencyMap
type VMEfficiencyMap struct {
}

// Analysis is the Analysis
type Analysis struct {
	VMEfficiencyMap VMEfficiencyMap `json:"vm_efficiency_map,omitempty" yaml:"vm_efficiency_map,omitempty"`
}

// CategoriesMapping is mapping for categories
type CategoriesMapping struct {
}

type ACPMetadata struct {
	Kind                 string            `json:"kind,omitempty" yaml:"kind,omitempty"`
	ProjectReference     ProjectReference  `json:"project_reference,omitempty" yaml:"project_reference,omitempty"`
	UUID                 string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name                 string            `json:"name,omitempty" yaml:"name,omitempty"`
	SpecVersion          int               `json:"spec_version,omitempty" yaml:"spec_version,omitempty"`
	SpecHash             string            `json:"spec_hash,omitempty" yaml:"spec_hash,omitempty"`
	Categories           Categories        `json:"categories,omitempty" yaml:"categories,omitempty"`
	CategoriesMapping    CategoriesMapping `json:"categories_mapping,omitempty" yaml:"categories_mapping,omitempty"`
	UseCategoriesMapping bool              `json:"use_categories_mapping,omitempty" yaml:"use_categories_mapping,omitempty"`
	CreationTime         string            `json:"creation_time,omitempty" yaml:"creation_time,omitempty"`
	LastUpdateTime       string            `json:"last_update_time,omitempty" yaml:"last_update_time,omitempty"`
	OwnerReference       OwnerReference    `json:"owner_reference,omitempty" yaml:"owner_reference,omitempty"`
	ShouldForceTranslate bool              `json:"should_force_translate,omitempty" yaml:"should_force_translate,omitempty"`
	EntityVersion        string            `json:"entity_version,omitempty" yaml:"entity_version,omitempty"`
}
type UserReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type UserGroupReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type RoleReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type RightHandSide struct {
	Collection string     `json:"collection,omitempty" yaml:"collection,omitempty"`
	Categories Categories `json:"categories,omitempty" yaml:"categories,omitempty"`
	UUIDList   []string   `json:"uuid_list,omitempty" yaml:"uuid_list,omitempty"`
}
type ScopeFilterExpressionList struct {
	LeftHandSide  string        `json:"left_hand_side,omitempty" yaml:"left_hand_side,omitempty"`
	Operator      string        `json:"operator,omitempty" yaml:"operator,omitempty"`
	RightHandSide RightHandSide `json:"right_hand_side,omitempty" yaml:"right_hand_side,omitempty"`
}
type LeftHandSide struct {
	EntityType string `json:"entity_type,omitempty" yaml:"entity_type,omitempty"`
}
type EntityFilterExpressionList struct {
	LeftHandSide  LeftHandSide  `json:"left_hand_side,omitempty" yaml:"left_hand_side,omitempty"`
	Operator      string        `json:"operator,omitempty" yaml:"operator,omitempty"`
	RightHandSide RightHandSide `json:"right_hand_side,omitempty" yaml:"right_hand_side,omitempty"`
}
type ContextList struct {
	ScopeFilterExpressionList  []ScopeFilterExpressionList  `json:"scope_filter_expression_list,omitempty" yaml:"scope_filter_expression_list,omitempty"`
	EntityFilterExpressionList []EntityFilterExpressionList `json:"entity_filter_expression_list,omitempty" yaml:"entity_filter_expression_list,omitempty"`
}
type FilterList struct {
	ContextList []ContextList `json:"context_list,omitempty" yaml:"context_list,omitempty"`
}
type ACPResources struct {
	UserReferenceList      []UserReferenceList      `json:"user_reference_list,omitempty" yaml:"user_reference_list,omitempty"`
	UserGroupReferenceList []UserGroupReferenceList `json:"user_group_reference_list,omitempty" yaml:"user_group_reference_list,omitempty"`
	RoleReference          RoleReference            `json:"role_reference,omitempty" yaml:"role_reference,omitempty"`
	FilterList             FilterList               `json:"filter_list,omitempty" yaml:"filter_list,omitempty"`
}
type ACPSpec struct {
	Name        string       `json:"name,omitempty" yaml:"name,omitempty"`
	Description string       `json:"description,omitempty" yaml:"description,omitempty"`
	Resources   ACPResources `json:"resources,omitempty" yaml:"resources,omitempty"`
}

type AutoGenerated struct {
	APIVersion string     `json:"api_version,omitempty" yaml:"api_version,omitempty"`
	Metadata   Metadata   `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Entities   []Entities `json:"entities,omitempty" yaml:"entities,omitempty"`
}

// OplogUsage produces the CVM OpLog utilization with a very precise result
type OplogUsage struct {
	OplogDiskPct  float64 `json:"oplog_disk_pct,omitempty" yaml:"oplog_disk_pct,omitempty"`
	OplogDiskSize int64   `json:"oplog_disk_size,omitempty" yaml:"oplog_disk_size,omitempty"`
}

// ControllerVM returns information about CVM and utilization
type ControllerVM struct {
	IP         string     `json:"ip,omitempty" yaml:"ip,omitempty"`
	NatIP      string     `json:"nat_ip,omitempty" yaml:"nat_ip,omitempty"`
	NatPort    int        `json:"nat_port,omitempty" yaml:"nat_port,omitempty"`
	OplogUsage OplogUsage `json:"oplog_usage,omitempty" yaml:"oplog_usage,omitempty"`
}
type DomainCredential struct {
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
}
type WindowsDomain struct {
	Name                 string           `json:"name,omitempty" yaml:"name,omitempty"`
	DomainName           string           `json:"domain_name,omitempty" yaml:"domain_name,omitempty"`
	OrganizationUnitPath string           `json:"organization_unit_path,omitempty" yaml:"organization_unit_path,omitempty"`
	NamePrefix           string           `json:"name_prefix,omitempty" yaml:"name_prefix,omitempty"`
	NameServerIP         string           `json:"name_server_ip,omitempty" yaml:"name_server_ip,omitempty"`
	DomainCredential     DomainCredential `json:"domain_credential,omitempty" yaml:"domain_credential,omitempty"`
}

type ClusterReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

type Ipmi struct {
	IP string `json:"ip,omitempty" yaml:"ip,omitempty"`
}
type Hypervisor struct {
	IP                 string `json:"ip,omitempty" yaml:"ip,omitempty"`
	HypervisorFullName string `json:"hypervisor_full_name,omitempty" yaml:"hypervisor_full_name,omitempty"`
	NumVms             int    `json:"num_vms,omitempty" yaml:"num_vms,omitempty"`
}
type Block struct {
	BlockSerialNumber string `json:"block_serial_number,omitempty" yaml:"block_serial_number,omitempty"`
	BlockModel        string `json:"block_model,omitempty" yaml:"block_model,omitempty"`
}
type RackableUnitReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type HostDisksReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type ConsumerReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}
type GpuList struct {
	Name                   string            `json:"name,omitempty" yaml:"name,omitempty"`
	Mode                   string            `json:"mode,omitempty" yaml:"mode,omitempty"`
	DeviceID               int               `json:"device_id,omitempty" yaml:"device_id,omitempty"`
	Vendor                 string            `json:"vendor,omitempty" yaml:"vendor,omitempty"`
	Status                 string            `json:"status,omitempty" yaml:"status,omitempty"`
	Index                  int               `json:"index,omitempty" yaml:"index,omitempty"`
	NumVgpusAllocated      int               `json:"num_vgpus_allocated,omitempty" yaml:"num_vgpus_allocated,omitempty"`
	Assignable             bool              `json:"assignable,omitempty" yaml:"assignable,omitempty"`
	PciAddress             string            `json:"pci_address,omitempty" yaml:"pci_address,omitempty"`
	NumaNode               int               `json:"numa_node,omitempty" yaml:"numa_node,omitempty"`
	GuestDriverVersion     string            `json:"guest_driver_version,omitempty" yaml:"guest_driver_version,omitempty"`
	FrameBufferSizeMib     int               `json:"frame_buffer_size_mib,omitempty" yaml:"frame_buffer_size_mib,omitempty"`
	MaxInstancesPerVM      int               `json:"max_instances_per_vm,omitempty" yaml:"max_instances_per_vm,omitempty"`
	NumVirtualDisplayHeads int               `json:"num_virtual_display_heads,omitempty" yaml:"num_virtual_display_heads,omitempty"`
	MaxResolution          string            `json:"max_resolution,omitempty" yaml:"max_resolution,omitempty"`
	Fraction               int               `json:"fraction,omitempty" yaml:"fraction,omitempty"`
	UUID                   string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	LicenseList            []string          `json:"license_list,omitempty" yaml:"license_list,omitempty"`
	ConsumerReference      ConsumerReference `json:"consumer_reference,omitempty" yaml:"consumer_reference,omitempty"`
}

type ParentReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

// DiskAddress is the DiskAddress
type DiskAddress struct {
	DeviceIndex *int   `json:"device_index,omitempty" yaml:"device_index,omitempty"`
	AdapterType string `json:"adapter_type,omitempty" yaml:"adapter_type,omitempty"`
}

// DeviceProperties is the DeviceProperties
type DeviceProperties struct {
	DiskAddress *DiskAddress `json:"disk_address,omitempty" yaml:"disk_address,omitempty"`
	DeviceType  string       `json:"device_type,omitempty" yaml:"device_type,omitempty"`
}

// VnumaConfig is the VnumaConfig
type VnumaConfig struct {
	NumVnumaNodes int `json:"num_vnuma_nodes,omitempty" yaml:"num_vnuma_nodes,omitempty"`
}

// IPEndpointList is the IPEndpointList
type IPEndpointList struct {
	IP                 string   `json:"ip,omitempty" yaml:"ip,omitempty"`
	Type               string   `json:"type,omitempty" yaml:"type,omitempty"`
	IPType             string   `json:"ip_type,omitempty" yaml:"ip_type,omitempty"`
	PrefixLength       int      `json:"prefix_length,omitempty" yaml:"prefix_length,omitempty"`
	GatewayAddressList []string `json:"gateway_address_list,omitempty" yaml:"gateway_address_list,omitempty"`
}

// SubnetReference is the SubnetReference
type SubnetReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type SerialPortList struct {
	IsConnected bool `json:"is_connected,omitempty" yaml:"is_connected,omitempty"`
	Index       int  `json:"index,omitempty" yaml:"index,omitempty"`
}
type QosPolicy struct {
	ThrottledIops int `json:"throttled_iops,omitempty" yaml:"throttled_iops,omitempty"`
}

type NetworkFunctionChainReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// NicList is the NicList
type NicList struct {
	NicType                       string                         `json:"nic_type,omitempty" yaml:"nic_type,omitempty"`
	UUID                          string                         `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	IPEndpointList                []IPEndpointList               `json:"ip_endpoint_list,omitempty" yaml:"ip_endpoint_list,omitempty"`
	MacAddress                    string                         `json:"mac_address,omitempty" yaml:"mac_address,omitempty"`
	SubnetReference               SubnetReference                `json:"subnet_reference,omitempty" yaml:"subnet_reference,omitempty"`
	IsConnected                   bool                           `json:"is_connected,omitempty" yaml:"is_connected,omitempty"`
	VlanMode                      string                         `json:"vlan_mode,omitempty" yaml:"vlan_mode,omitempty"`
	TrunkedVlanList               []int                          `json:"trunked_vlan_list,omitempty" yaml:"trunked_vlan_list,omitempty"`
	NetworkFunctionChainReference *NetworkFunctionChainReference `json:"network_function_chain_reference,omitempty" yaml:"network_function_chain_reference,omitempty"`
	NetworkFunctionNicType        string                         `json:"network_function_nic_type,omitempty" yaml:"network_function_nic_type,omitempty"`
	Model                         string                         `json:"model,omitempty" yaml:"model,omitempty"`
}

// HostReference is the HostReference
type HostReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// DiskList is the DiskList
type DiskList struct {
	DataSourceReference  *DataSourceReference  `json:"data_source_reference,omitempty" yaml:"data_source_reference,omitempty"`
	DeviceProperties     *DeviceProperties     `json:"device_properties,omitempty" yaml:"device_properties,omitempty"`
	UUID                 string                `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	DiskSizeBytes        int                   `json:"disk_size_bytes,omitempty" yaml:"disk_size_bytes,omitempty"`
	DiskSizeMib          int                   `json:"disk_size_mib,omitempty" yaml:"disk_size_mib,omitempty"`
	VolumeGroupReference *VolumeGroupReference `json:"volume_group_reference,omitempty" yaml:"volume_group_reference,omitempty"`
	StorageConfig        *StorageConfig        `json:"storage_config,omitempty" yaml:"storage_config,omitempty"`
}

// ExecutionContext is the ExecutionContext
type ExecutionContext struct {
	TaskUuids []string `json:"task_uuids,omitempty" yaml:"task_uuids,omitempty"`
}

// Version is used for image version
type Version struct {
	ProductName    string `json:"product_name,omitempty" yaml:"product_name,omitempty"`
	ProductVersion string `json:"product_version,omitempty" yaml:"product_version,omitempty"`
}

// BootDevice is the BootDevice
type BootDevice struct {
	DiskAddress *DiskAddress `json:"disk_address,omitempty" yaml:"disk_address,omitempty"`
	MacAddress  string       `json:"mac_address,omitempty" yaml:"mac_address,omitempty"`
}
type CustomKeyValues struct {
}
type Sysprep struct {
	InstallType     string           `json:"install_type,omitempty" yaml:"install_type,omitempty"`
	UnattendXML     string           `json:"unattend_xml,omitempty" yaml:"unattend_xml,omitempty"`
	CustomKeyValues *CustomKeyValues `json:"custom_key_values,omitempty" yaml:"custom_key_values,omitempty"`
}
type CloudInit struct {
	MetaData        string           `json:"meta_data,omitempty" yaml:"meta_data,omitempty"`
	UserData        string           `json:"user_data,omitempty" yaml:"user_data,omitempty"`
	CustomKeyValues *CustomKeyValues `json:"custom_key_values,omitempty" yaml:"custom_key_values,omitempty"`
}
type GuestCustomization struct {
	IsOverridable bool       `json:"is_overridable,omitempty" yaml:"is_overridable,omitempty"`
	Sysprep       *Sysprep   `json:"sysprep,omitempty" yaml:"sysprep,omitempty"`
	CloudInit     *CloudInit `json:"cloud_init,omitempty" yaml:"cloud_init,omitempty"`
}

// BootConfig is the BootConfig
type BootConfig struct {
	BootDevice          *BootDevice          `json:"boot_device,omitempty" yaml:"boot_device,omitempty"`
	BootDeviceOrderList []string             `json:"boot_device_order_list,omitempty" yaml:"boot_device_order_list,omitempty"`
	BootType            string               `json:"boot_type,omitempty" yaml:"boot_type,omitempty"`
	DataSourceReference *DataSourceReference `json:"data_source_reference,omitempty" yaml:"data_source_reference,omitempty"`
}

// GuestTransitionConfig is the GuestTransitionConfig
type GuestTransitionConfig struct {
	ShouldFailOnScriptFailure bool `json:"should_fail_on_script_failure,omitempty" yaml:"should_fail_on_script_failure,omitempty"`
	EnableScriptExec          bool `json:"enable_script_exec,omitempty" yaml:"enable_script_exec,omitempty"`
}

// PowerStateMechanism is the PowerStateMechanism
type PowerStateMechanism struct {
	GuestTransitionConfig *GuestTransitionConfig `json:"guest_transition_config,omitempty" yaml:"guest_transition_config,omitempty"`
	Mechanism             string                 `json:"mechanism,omitempty" yaml:"mechanism,omitempty"`
}

// Capabilities are used for categories
type Capabilities struct {
	Cardinality int `json:"cardinality,omitempty" yaml:"cardinality,omitempty"`
}

// ExpressionList used for category get value results
type ExpressionList struct {
	PropertyName       string `json:"property_name,omitempty" yaml:"property_name,omitempty"`
	Operator           string `json:"operator,omitempty" yaml:"operator,omitempty"`
	Value              string `json:"value,omitempty" yaml:"value,omitempty"`
	DisplayForOperator string `json:"display_for_operator,omitempty" yaml:"display_for_operator,omitempty"`
	DisplayForValue    string `json:"display_for_value,omitempty" yaml:"display_for_value,omitempty"`
}

// SelectionCriteriaList used for category get value results
type SelectionCriteriaList struct {
	EntityType     string           `json:"entity_type,omitempty" yaml:"entity_type,omitempty"`
	ExpressionList []ExpressionList `json:"expression_list,omitempty" yaml:"expression_list,omitempty"`
}

// ExclusionList used for category get value results
type ExclusionList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

// InclusionList used for category get value results
type InclusionList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

// AssignmentRule used for category get value results
type AssignmentRule struct {
	Name                  string                  `json:"name,omitempty" yaml:"name,omitempty"`
	Description           string                  `json:"description,omitempty" yaml:"description,omitempty"`
	SelectionCriteriaList []SelectionCriteriaList `json:"selection_criteria_list,omitempty" yaml:"selection_criteria_list,omitempty"`
	ExclusionList         []ExclusionList         `json:"exclusion_list,omitempty" yaml:"exclusion_list,omitempty"`
	InclusionList         []InclusionList         `json:"inclusion_list,omitempty" yaml:"inclusion_list,omitempty"`
}

type DhcpOptions struct {
	BootFileName         string   `json:"boot_file_name,omitempty" yaml:"boot_file_name,omitempty"`
	DomainName           string   `json:"domain_name,omitempty" yaml:"domain_name,omitempty"`
	DomainNameServerList []string `json:"domain_name_server_list,omitempty" yaml:"domain_name_server_list,omitempty"`
	DomainSearchList     []string `json:"domain_search_list,omitempty" yaml:"domain_search_list,omitempty"`
	TftpServerName       string   `json:"tftp_server_name,omitempty" yaml:"tftp_server_name,omitempty"`
}
type DhcpServerAddress struct {
	IP       string `json:"ip,omitempty" yaml:"ip,omitempty"`
	Ipv6     string `json:"ipv6,omitempty" yaml:"ipv_6,omitempty"`
	Fqdn     string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`
	Port     int    `json:"port,omitempty" yaml:"port,omitempty"`
	IsBackup bool   `json:"is_backup,omitempty" yaml:"is_backup,omitempty"`
}

type PoolList struct {
	Range string `json:"range,omitempty" yaml:"range,omitempty"`
}
type IPConfig struct {
	SubnetIP          string            `json:"subnet_ip,omitempty" yaml:"subnet_ip,omitempty"`
	PrefixLength      int               `json:"prefix_length,omitempty" yaml:"prefix_length,omitempty"`
	DefaultGatewayIP  string            `json:"default_gateway_ip,omitempty" yaml:"default_gateway_ip,omitempty"`
	DhcpOptions       DhcpOptions       `json:"dhcp_options,omitempty" yaml:"dhcp_options,omitempty"`
	DhcpServerAddress DhcpServerAddress `json:"dhcp_server_address,omitempty" yaml:"dhcp_server_address,omitempty"`
	PoolList          []PoolList        `json:"pool_list,omitempty" yaml:"pool_list,omitempty"`
}

type VirtualNetworkReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type VpcReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type AvailabilityZoneReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type DefaultSubnetReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type SubnetReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ExternalUserGroupReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type AccountReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type TunnelReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type ExternalNetworkList struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type EnvironmentReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type DefaultEnvironmentReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// ParentTaskReference used for task get details
type ParentTaskReference struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// SubtaskReferenceList used for task get details
type SubtaskReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// EntityReferenceList used for task get details
type EntityReferenceList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

// Fields is used for permission creation
type Fields struct {
	FieldMode     string   `json:"field_mode,omitempty" yaml:"field_mode,omitempty"`
	FieldNameList []string `json:"field_name_list,omitempty" yaml:"field_name_list,omitempty"`
}

// TargetGroup used in NSR
type TargetGroup struct {
	PeerSpecificationType string `json:"peer_specification_type,omitempty" yaml:"peer_specification_type,omitempty"`
	Filter                Filter `json:"filter,omitempty" yaml:"filter,omitempty"`
	DefaultInternalPolicy string `json:"default_internal_policy,omitempty" yaml:"default_internal_policy,omitempty"`
}

// AddressGroupInclusionList used for NSR
type AddressGroupInclusionList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// IPSubnet used for NSR
type IPSubnet struct {
	IP           string `json:"ip,omitempty" yaml:"ip,omitempty"`
	PrefixLength int    `json:"prefix_length,omitempty" yaml:"prefix_length,omitempty"`
}

// ServiceGroupList used for NSR
type ServiceGroupList struct {
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// TCPPortRangeList used for NSR
type TCPPortRangeList struct {
	StartPort int `json:"start_port,omitempty" yaml:"start_port,omitempty"`
	EndPort   int `json:"end_port,omitempty" yaml:"end_port,omitempty"`
}

// UDPPortRangeList used for NSR
type UDPPortRangeList struct {
	StartPort int `json:"start_port,omitempty" yaml:"start_port,omitempty"`
	EndPort   int `json:"end_port,omitempty" yaml:"end_port,omitempty"`
}

// IcmpTypeCodeList used for NSR
type IcmpTypeCodeList struct {
	Type int `json:"type,omitempty" yaml:"type,omitempty"`
	Code int `json:"code,omitempty" yaml:"code,omitempty"`
}

// InboundAllowList used for NSR
type InboundAllowList struct {
	PeerSpecificationType         string                         `json:"peer_specification_type,omitempty" yaml:"peer_specification_type,omitempty"`
	Filter                        *Filter                        `json:"filter,omitempty" yaml:"filter,omitempty"`
	AddressGroupInclusionList     []AddressGroupInclusionList    `json:"address_group_inclusion_list,omitempty" yaml:"address_group_inclusion_list,omitempty"`
	IPSubnet                      *IPSubnet                      `json:"ip_subnet,omitempty" yaml:"ip_subnet,omitempty"`
	ServiceGroupList              []ServiceGroupList             `json:"service_group_list,omitempty" yaml:"service_group_list,omitempty"`
	Protocol                      string                         `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	TCPPortRangeList              []TCPPortRangeList             `json:"tcp_port_range_list,omitempty" yaml:"tcp_port_range_list,omitempty"`
	UDPPortRangeList              []UDPPortRangeList             `json:"udp_port_range_list,omitempty" yaml:"udp_port_range_list,omitempty"`
	IcmpTypeCodeList              []IcmpTypeCodeList             `json:"icmp_type_code_list,omitempty" yaml:"icmp_type_code_list,omitempty"`
	NetworkFunctionChainReference *NetworkFunctionChainReference `json:"network_function_chain_reference,omitempty" yaml:"network_function_chain_reference,omitempty"`
	ExpirationTime                string                         `json:"expiration_time,omitempty" yaml:"expiration_time,omitempty"`
	Description                   string                         `json:"description,omitempty" yaml:"description,omitempty"`
}

// OutboundAllowList used for NSR
type OutboundAllowList struct {
	PeerSpecificationType         string                         `json:"peer_specification_type,omitempty" yaml:"peer_specification_type,omitempty"`
	Filter                        *Filter                        `json:"filter,omitempty" yaml:"filter,omitempty"`
	AddressGroupInclusionList     []AddressGroupInclusionList    `json:"address_group_inclusion_list,omitempty" yaml:"address_group_inclusion_list,omitempty"`
	IPSubnet                      *IPSubnet                      `json:"ip_subnet,omitempty" yaml:"ip_subnet,omitempty"`
	ServiceGroupList              []ServiceGroupList             `json:"service_group_list,omitempty" yaml:"service_group_list,omitempty"`
	Protocol                      string                         `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	TCPPortRangeList              []TCPPortRangeList             `json:"tcp_port_range_list,omitempty" yaml:"tcp_port_range_list,omitempty"`
	UDPPortRangeList              []UDPPortRangeList             `json:"udp_port_range_list,omitempty" yaml:"udp_port_range_list,omitempty"`
	IcmpTypeCodeList              []IcmpTypeCodeList             `json:"icmp_type_code_list,omitempty" yaml:"icmp_type_code_list,omitempty"`
	NetworkFunctionChainReference *NetworkFunctionChainReference `json:"network_function_chain_reference,omitempty" yaml:"network_function_chain_reference,omitempty"`
	ExpirationTime                string                         `json:"expiration_time,omitempty" yaml:"expiration_time,omitempty"`
	Description                   string                         `json:"description,omitempty" yaml:"description,omitempty"`
}

// AdRule used for NSR
type AdRule struct {
	TargetGroup       TargetGroup         `json:"target_group,omitempty" yaml:"target_group,omitempty"`
	InboundAllowList  []InboundAllowList  `json:"inbound_allow_list,omitempty" yaml:"inbound_allow_list,omitempty"`
	OutboundAllowList []OutboundAllowList `json:"outbound_allow_list,omitempty" yaml:"outbound_allow_list,omitempty"`
	Action            string              `json:"action,omitempty" yaml:"action,omitempty"`
}

// AppRule used for NSR
type AppRule struct {
	TargetGroup       TargetGroup         `json:"target_group,omitempty" yaml:"target_group,omitempty"`
	InboundAllowList  []InboundAllowList  `json:"inbound_allow_list,omitempty" yaml:"inbound_allow_list,omitempty"`
	OutboundAllowList []OutboundAllowList `json:"outbound_allow_list,omitempty" yaml:"outbound_allow_list,omitempty"`
	Action            string              `json:"action,omitempty" yaml:"action,omitempty"`
}

// FirstEntityFilter used for NSR
type FirstEntityFilter struct {
	Type     string   `json:"type,omitempty" yaml:"type,omitempty"`
	KindList []string `json:"kind_list,omitempty" yaml:"kind_list,omitempty"`
	Params   Params   `json:"params,omitempty" yaml:"params,omitempty"`
}

// SecondEntityFilter used for NSR
type SecondEntityFilter struct {
	Type     string   `json:"type,omitempty" yaml:"type,omitempty"`
	KindList []string `json:"kind_list,omitempty" yaml:"kind_list,omitempty"`
	Params   Params   `json:"params,omitempty" yaml:"params,omitempty"`
}

// IsolationRule used for NSR
type IsolationRule struct {
	FirstEntityFilter  FirstEntityFilter  `json:"first_entity_filter,omitempty" yaml:"first_entity_filter,omitempty"`
	SecondEntityFilter SecondEntityFilter `json:"second_entity_filter,omitempty" yaml:"second_entity_filter,omitempty"`
	Action             string             `json:"action,omitempty" yaml:"action,omitempty"`
}

// QuarantineRule for NSR
type QuarantineRule struct {
	TargetGroup       TargetGroup         `json:"target_group,omitempty" yaml:"target_group,omitempty"`
	InboundAllowList  []InboundAllowList  `json:"inbound_allow_list,omitempty" yaml:"inbound_allow_list,omitempty"`
	OutboundAllowList []OutboundAllowList `json:"outbound_allow_list,omitempty" yaml:"outbound_allow_list,omitempty"`
	Action            string              `json:"action,omitempty" yaml:"action,omitempty"`
}

// EntityList used for NSR import dry run request
type EntityList struct {
	EntityType       string `json:"entity_type,omitempty" yaml:"entity_type,omitempty"`
	EntityName       string `json:"entity_name,omitempty" yaml:"entity_name,omitempty"`
	EntityUpdateType string `json:"entity_update_type,omitempty" yaml:"entity_update_type,omitempty"`
}

// Body used for batch submit request
type Body struct {
}

// APIRequestList used for batch submit request
type APIRequestList struct {
	PathAndParams string `json:"path_and_params,omitempty" yaml:"path_and_params,omitempty"`
	Operation     string `json:"operation,omitempty" yaml:"operation,omitempty"`
	Body          Body   `json:"body,omitempty" yaml:"body,omitempty"`
}

// APIResponse used for batch response
type APIResponse struct {
}

// APIResponseList used for batch response
type APIResponseList struct {
	Status        string      `json:"status,omitempty" yaml:"status,omitempty"`
	PathAndParams string      `json:"path_and_params,omitempty" yaml:"path_and_params,omitempty"`
	APIResponse   APIResponse `json:"api_response,omitempty" yaml:"api_response,omitempty"`
}
