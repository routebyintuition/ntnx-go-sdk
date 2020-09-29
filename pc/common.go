package pc

import "time"

// Metadata is the Metadata associated with an API responsee
// this same type can be used across the calls as it includes
// type entries for other API calls
type Metadata struct {
	TotalMatches         int               `json:"total_matches,omitempty"`
	Kind                 string            `json:"kind,omitempty"`
	Length               int               `json:"length,omitempty"`
	Offset               int               `json:"offset,omitempty"`
	LastUpdateTime       time.Time         `json:"last_update_time,omitempty"`
	UUID                 string            `json:"uuid,omitempty"`
	ProjectReference     ProjectReference  `json:"project_reference,omitempty"`
	SpecVersion          int               `json:"spec_version,omitempty"`
	CreationTime         time.Time         `json:"creation_time,omitempty"`
	OwnerReference       OwnerReference    `json:"owner_reference,omitempty"`
	Categories           Categories        `json:"categories,omitempty"`
	SpecHash             string            `json:"spec_hash,omitempty"`
	ShouldForceTranslate bool              `json:"should_force_translate,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Filter               string            `json:"filter,omitempty"`
	SortOrder            string            `json:"sort_order,omitempty"`
	SortAttribute        string            `json:"sort_attribute,omitempty"`
	CategoriesMapping    CategoriesMapping `json:"categories_mapping,omitempty"`
	UseCategoriesMapping bool              `json:"use_categories_mapping,omitempty"`
	EntityVersion        string            `json:"entity_version,omitempty"`
	GroupMemberCount     int               `json:"group_member_count,omitempty"`
	GroupMemberOffset    int               `json:"group_member_offset,omitempty"`
	UsageType            string            `json:"usage_type,omitempty"`
}

// Config is the API response for Config used across all API requests
type Config struct {
	AuthorizedPublicKeyList       []AuthorizedPublicKeyList `json:"authorized_public_key_list,omitempty"`
	Build                         Build                     `json:"build,omitempty"`
	CaCertificateList             []CaCertificateList       `json:"ca_certificate_list,omitempty"`
	CertificationSigningInfo      CertificationSigningInfo  `json:"certification_signing_info,omitempty"`
	ClientAuth                    ClientAuth                `json:"client_auth,omitempty"`
	ClusterArch                   string                    `json:"cluster_arch,omitempty"`
	DomainAwarenessLevel          string                    `json:"domain_awareness_level,omitempty"`
	EnabledFeatureList            []string                  `json:"enabled_feature_list,omitempty"`
	EncryptionStatus              string                    `json:"encryption_status,omitempty"`
	ExternalConfigurations        ExternalConfigurations    `json:"external_configurations,omitempty"`
	GpuDriverVersion              string                    `json:"gpu_driver_version,omitempty"`
	IsAvailable                   bool                      `json:"is_available,omitempty"`
	ManagementServerList          []ManagementServerList    `json:"management_server_list,omitempty"`
	OperationMode                 string                    `json:"operation_mode,omitempty"`
	RedundancyFactor              int                       `json:"redundancy_factor,omitempty"`
	ServiceList                   []string                  `json:"service_list,omitempty"`
	SoftwareMap                   SoftwareMap               `json:"software_map,omitempty"`
	SslKey                        SslKey                    `json:"ssl_key,omitempty"`
	SupportedInformationVerbosity string                    `json:"supported_information_verbosity,omitempty"`
	Timezone                      string                    `json:"timezone,omitempty"`
}

// Resources is the main Resources type across all Response calls
type Resources struct {
	RetrievalURIList                 []string                           `json:"retrieval_uri_list,omitempty"`
	SizeBytes                        int                                `json:"size_bytes,omitempty"`
	CurrentClusterReferenceList      []CurrentClusterReferenceList      `json:"current_cluster_reference_list,omitempty"`
	ImagePlacementPolicyStateList    []ImagePlacementPolicyStateList    `json:"image_placement_policy_state_list,omitempty"`
	ImageType                        string                             `json:"image_type,omitempty"`
	Checksum                         string                             `json:"checksum,omitempty"`
	SourceURI                        string                             `json:"source_uri,omitempty"`
	SourceOptions                    SourceOptions                      `json:"source_options,omitempty"`
	InitialPlacementRefList          []InitialPlacementRefList          `json:"initial_placement_ref_list,omitempty"`
	DataSourceReference              DataSourceReference                `json:"data_source_reference,omitempty"`
	Version                          Version                            `json:"version,omitempty"`
	Architecture                     string                             `json:"architecture,omitempty"`
	BootConfig                       BootConfig                         `json:"boot_config,omitempty"`
	DiskList                         []DiskList                         `json:"disk_list,omitempty"`
	GuestCustomization               GuestCustomization                 `json:"guest_customization,omitempty"`
	HardwareClockTimezone            string                             `json:"hardware_clock_timezone,omitempty"`
	HostReference                    HostReference                      `json:"host_reference,omitempty"`
	HypervisorType                   string                             `json:"hypervisor_type,omitempty"`
	MemorySizeMib                    int                                `json:"memory_size_mib,omitempty"`
	NicList                          []NicList                          `json:"nic_list,omitempty"`
	NumSockets                       int                                `json:"num_sockets,omitempty"`
	NumThreadsPerCore                int                                `json:"num_threads_per_core,omitempty"`
	NumVcpusPerSocket                int                                `json:"num_vcpus_per_socket,omitempty"`
	PowerState                       string                             `json:"power_state,omitempty"`
	PowerStateMechanism              PowerStateMechanism                `json:"power_state_mechanism,omitempty"`
	ProtectionType                   string                             `json:"protection_type,omitempty"`
	SerialPortList                   []interface{}                      `json:"serial_port_list,omitempty"`
	VgaConsoleEnabled                bool                               `json:"vga_console_enabled,omitempty"`
	VnumaConfig                      VnumaConfig                        `json:"vnuma_config,omitempty"`
	Nodes                            Nodes                              `json:"nodes,omitempty"`
	Config                           Config                             `json:"config,omitempty"`
	Network                          Network                            `json:"network,omitempty"`
	Analysis                         Analysis                           `json:"analysis,omitempty"`
	RuntimeStatusList                []string                           `json:"runtime_status_list,omitempty"`
	CPUModel                         string                             `json:"cpu_model,omitempty"`
	CPUCapacityHz                    int                                `json:"cpu_capacity_hz,omitempty"`
	NumCPUCores                      int                                `json:"num_cpu_cores,omitempty"`
	NumCPUSockets                    int                                `json:"num_cpu_sockets,omitempty"`
	MemoryCapacityMib                int                                `json:"memory_capacity_mib,omitempty"`
	SerialNumber                     string                             `json:"serial_number,omitempty"`
	MonitoringState                  string                             `json:"monitoring_state,omitempty"`
	HostNicsIDList                   []string                           `json:"host_nics_id_list,omitempty"`
	WindowsDomain                    WindowsDomain                      `json:"windows_domain,omitempty"`
	FailoverCluster                  string                             `json:"failover_cluster,omitempty"`
	Ipmi                             Ipmi                               `json:"ipmi,omitempty"`
	ControllerVM                     ControllerVM                       `json:"controller_vm,omitempty"`
	Hypervisor                       Hypervisor                         `json:"hypervisor,omitempty"`
	Block                            Block                              `json:"block,omitempty"`
	RackableUnitReference            RackableUnitReference              `json:"rackable_unit_reference,omitempty"`
	HostDisksReferenceList           []HostDisksReferenceList           `json:"host_disks_reference_list,omitempty"`
	GpuList                          []GpuList                          `json:"gpu_list,omitempty"`
	GpuDriverVersion                 string                             `json:"gpu_driver_version,omitempty"`
	HostType                         string                             `json:"host_type,omitempty"`
	GuestOsID                        string                             `json:"guest_os_id"`
	IsVcpuHardPinned                 bool                               `json:"is_vcpu_hard_pinned"`
	IsAgentVM                        bool                               `json:"is_agent_vm"`
	DisableBranding                  bool                               `json:"disable_branding"`
	EnableCPUPassthrough             bool                               `json:"enable_cpu_passthrough"`
	MachineType                      string                             `json:"machine_type"`
	HardwareVirtualizationEnabled    bool                               `json:"hardware_virtualization_enabled"`
	ParentReference                  ParentReference                    `json:"parent_reference"`
	GuestTools                       GuestTools                         `json:"guest_tools"`
	StorageConfig                    StorageConfig                      `json:"storage_config"`
	DirectoryServiceUser             DirectoryServiceUser               `json:"directory_service_user"`
	IdentityProviderUser             IdentityProviderUser               `json:"identity_provider_user"`
	UserType                         string                             `json:"user_type"`
	DisplayName                      string                             `json:"display_name"`
	ProjectsReferenceList            []ProjectsReferenceList            `json:"projects_reference_list"`
	AccessControlPolicyReferenceList []AccessControlPolicyReferenceList `json:"access_control_policy_reference_list"`
	ResourceUsageSummary             ResourceUsageSummary               `json:"resource_usage_summary"`
	SubnetType                       string                             `json:"subnet_type"`
	VlanID                           int                                `json:"vlan_id"`
	IPConfig                         IPConfig                           `json:"ip_config"`
	VswitchName                      string                             `json:"vswitch_name"`
	NetworkFunctionChainReference    NetworkFunctionChainReference      `json:"network_function_chain_reference"`
	VirtualNetworkReference          VirtualNetworkReference            `json:"virtual_network_reference"`
	VpcReference                     VpcReference                       `json:"vpc_reference"`
	AvailabilityZoneReferenceList    []AvailabilityZoneReferenceList    `json:"availability_zone_reference_list"`
	ExternalConnectivityState        string                             `json:"external_connectivity_state"`
}

// Spec is the main Spec type across all Response calls
type Spec struct {
	ClusterReference          ClusterReference          `json:"cluster_reference,omitempty"`
	Description               string                    `json:"description,omitempty"`
	Resources                 Resources                 `json:"resources,omitempty"`
	Name                      string                    `json:"name,omitempty"`
	APIVersion                string                    `json:"api_version,omitempty"`
	AvailabilityZoneReference AvailabilityZoneReference `json:"availability_zone_reference"`
}

// ImageReferenceList is used for image service details
type ImageReferenceList struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type ProjectsReferenceList struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type AccessControlPolicyReferenceList struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

// SourceOptions used for Image Update
type SourceOptions struct {
	AllowInsecureConnection bool `json:"allow_insecure_connection"`
}

type ResourceUsageSummary struct {
	ResourceDomain string `json:"resource_domain"`
}

// InitialPlacementRefList used or image update and placement
type InitialPlacementRefList struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

// DataSourceReference used for image update
type DataSourceReference struct {
	Kind           string `json:"kind"`
	UUID           string `json:"uuid"`
	Name           string `json:"name"`
	URL            string `json:"url"`
	IsDirectAttach bool   `json:"is_direct_attach"`
}

type VolumeGroupReference struct {
	Kind string `json:"kind"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
type StorageContainerReference struct {
	Kind string `json:"kind"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
type StorageConfig struct {
	StorageContainerReference StorageContainerReference `json:"storage_container_reference"`
	FlashMode                 string                    `json:"flash_mode"`
	QosPolicy                 QosPolicy                 `json:"qos_policy"`
}

type DirectoryServiceReference struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
type DirectoryServiceUser struct {
	UserPrincipalName         string                    `json:"user_principal_name"`
	DefaultUserPrincipalName  string                    `json:"default_user_principal_name"`
	DirectoryServiceReference DirectoryServiceReference `json:"directory_service_reference"`
}
type IdentityProviderReference struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
type IdentityProviderUser struct {
	Username                  string                    `json:"username"`
	IdentityProviderReference IdentityProviderReference `json:"identity_provider_reference"`
}

// ProjectReference is the main ProjectReference type across all Response calls
type ProjectReference struct {
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

// ProjectResourceDomainList used for user resource/domain get information
type ProjectResourceDomainList struct {
	ProjectReference ProjectReference `json:"project_reference"`
	ResourceDomain   string           `json:"resource_domain"`
}

// OwnerReference is the main OwnerReference type across all Response calls
type OwnerReference struct {
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

// Categories is the main Categories type across all Response calls
type Categories struct{}

// State is the various states associated with a cluster
type State struct {
	ReadOnly    bool   `json:"readOnly,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

type AvailabilityZoneReference struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type ImagePlacementPolicyStateList struct {
	ComplianceStatus string             `json:"compliance_status"`
	EnforcementMode  string             `json:"enforcement_mode"`
	PolicyReference  PolicyReference    `json:"policy_reference"`
	PolicyInfo       PolicyInfo         `json:"policy_info"`
	ErrorMessageList []ErrorMessageList `json:"error_message_list"`
}

type CurrentClusterReferenceList struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
type PolicyReference struct {
	Kind string `json:"kind"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
type EnforcedPolicyClusterReferenceList struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
type CompletePolicyClusterReferenceList struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
type ConflictingImagePlacementPolicyReferenceList struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
type PolicyInfo struct {
	EnforcedPolicyClusterReferenceList           []EnforcedPolicyClusterReferenceList           `json:"enforced_policy_cluster_reference_list"`
	CompletePolicyClusterReferenceList           []CompletePolicyClusterReferenceList           `json:"complete_policy_cluster_reference_list"`
	ConflictingImagePlacementPolicyReferenceList []ConflictingImagePlacementPolicyReferenceList `json:"conflicting_image_placement_policy_reference_list"`
}

type ErrorMessageList struct {
	Message string  `json:"message"`
	Reason  string  `json:"reason"`
	Details Details `json:"details"`
}

// Status is the main Status type across all Response calls
type Status struct {
	State                     string                    `json:"state,omitempty"`
	ClusterReference          ClusterReference          `json:"cluster_reference,omitempty"`
	ExecutionContext          ExecutionContext          `json:"execution_context,omitempty"`
	Description               string                    `json:"description,omitempty"`
	Resources                 Resources                 `json:"resources,omitempty"`
	Name                      string                    `json:"name,omitempty"`
	MessageList               []MessageList             `json:"message_list,omitempty"`
	AvailabilityZoneReference AvailabilityZoneReference `json:"availability_zone_reference"`
}

// Entities is the main Entities type across all Response calls
type Entities struct {
	Status        Status       `json:"status,omitempty"`
	Spec          Spec         `json:"spec,omitempty"`
	APIVersion    string       `json:"api_version,omitempty"`
	Metadata      Metadata     `json:"metadata,omitempty"`
	Name          string       `json:"name,omitempty"`
	Capabilities  Capabilities `json:"capabilities,omitempty"`
	Description   string       `json:"description,omitempty"`
	SystemDefined bool         `json:"system_defined,omitempty"`
}

// MessageList is the list of messages
type MessageList struct {
	Message string  `json:"message,omitempty"`
	Reason  string  `json:"reason,omitempty"`
	Details Details `json:"details,omitempty"`
}

// Properties are tied to response properties
type Properties struct {
	Kind                 Kind                 `json:"kind,omitempty"`
	ProjectReference     ProjectReference     `json:"project_reference,omitempty"`
	UUID                 UUID                 `json:"uuid,omitempty"`
	Name                 Name                 `json:"name,omitempty"`
	SpecVersion          SpecVersion          `json:"spec_version,omitempty"`
	SpecHash             SpecHash             `json:"spec_hash,omitempty"`
	Categories           Categories           `json:"categories,omitempty"`
	CategoriesMapping    CategoriesMapping    `json:"categories_mapping,omitempty"`
	UseCategoriesMapping UseCategoriesMapping `json:"use_categories_mapping,omitempty"`
	CreationTime         CreationTime         `json:"creation_time,omitempty"`
	LastUpdateTime       LastUpdateTime       `json:"last_update_time,omitempty"`
	OwnerReference       OwnerReference       `json:"owner_reference,omitempty"`
	ShouldForceTranslate ShouldForceTranslate `json:"should_force_translate,omitempty"`
	EntityVersion        EntityVersion        `json:"entity_version,omitempty"`
	Offset               Offset               `json:"offset,omitempty"`
	Length               Length               `json:"length,omitempty"`
	TotalMatches         TotalMatches         `json:"total_matches,omitempty"`
	Filter               Filter               `json:"filter,omitempty"`
	SortOrder            SortOrder            `json:"sort_order,omitempty"`
	SortAttribute        SortAttribute        `json:"sort_attribute,omitempty"`
	MessageList          MessageList          `json:"message_list,omitempty"`
}

type Kind struct {
	Type        string   `json:"type,omitempty"`
	ReadOnly    bool     `json:"readOnly,omitempty"`
	Description string   `json:"description,omitempty"`
	Default     string   `json:"default,omitempty"`
	XNtnxEnum   []string `json:"x-ntnx-enum,omitempty"`
}

type ShouldForceTranslate struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
}

type EntityVersion struct {
	Type        string `json:"type,omitempty"`
	ReadOnly    bool   `json:"readOnly,omitempty"`
	Description string `json:"description,omitempty"`
}

type UUID struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Format      string `json:"format,omitempty"`
	Pattern     string `json:"pattern,omitempty"`
}

type Offset struct {
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}
type Length struct {
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}
type TotalMatches struct {
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}
type Filter struct {
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}
type SortOrder struct {
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Type        string   `json:"type,omitempty"`
	XNtnxEnum   []string `json:"x-ntnx-enum,omitempty"`
}
type SortAttribute struct {
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}
type LastUpdateTime struct {
	Type        string `json:"type,omitempty"`
	ReadOnly    bool   `json:"readOnly,omitempty"`
	Format      string `json:"format,omitempty"`
	Description string `json:"description,omitempty"`
}

type CreationTime struct {
	Type        string `json:"type,omitempty"`
	Format      string `json:"format,omitempty"`
	ReadOnly    bool   `json:"readOnly,omitempty"`
	Description string `json:"description,omitempty"`
}

type Name struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	ReadOnly    bool   `json:"readOnly,omitempty"`
	MaxLength   int    `json:"maxLength,omitempty"`
}

type SpecVersion struct {
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

type SpecHash struct {
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

type UseCategoriesMapping struct {
	Type        string `json:"type,omitempty"`
	Default     bool   `json:"default,omitempty"`
	Description string `json:"description,omitempty"`
}

// Results is used for category query results
type Results struct {
	Kind                string              `json:"kind"`
	TotalEntityCount    int                 `json:"total_entity_count"`
	FilteredEntityCount int                 `json:"filtered_entity_count"`
	KindReferenceList   []KindReferenceList `json:"kind_reference_list"`
}

// KindReferenceList is used for category query results
type KindReferenceList struct {
	Kind       string     `json:"kind"`
	UUID       string     `json:"uuid"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Categories Categories `json:"categories"`
}

type CategoryFilter struct {
	Type     string   `json:"type"`
	KindList []string `json:"kind_list"`
	Params   Params   `json:"params"`
}

type Params struct {
}

// HypervisorServerList is the hyprevisor list of servers
type HypervisorServerList struct {
	IP      string `json:"ip,omitempty"`
	Version string `json:"version,omitempty"`
	Type    string `json:"type,omitempty"`
}

// Nodes are the nodes
type Nodes struct {
	HypervisorServerList []HypervisorServerList `json:"hypervisor_server_list,omitempty"`
}

// ClientAuth is the PKE auth info
type ClientAuth struct {
	Status  string `json:"status,omitempty"`
	CaChain string `json:"ca_chain,omitempty"`
	Name    string `json:"name,omitempty"`
}

// SoftwareMap is the software map
type SoftwareMap struct {
}

// SigningInfo is the SigningInfo
type SigningInfo struct {
	City             string `json:"city,omitempty"`
	State            string `json:"state,omitempty"`
	CountryCode      string `json:"country_code,omitempty"`
	CommonName       string `json:"common_name,omitempty"`
	Organization     string `json:"organization,omitempty"`
	EmailAddress     string `json:"email_address,omitempty"`
	CommonNameSuffix string `json:"common_name_suffix,omitempty"`
}

// Details are the details
type Details struct {
}

// SslKey is the SSL Key
type SslKey struct {
	KeyType        string      `json:"key_type,omitempty"`
	KeyName        string      `json:"key_name,omitempty"`
	SigningInfo    SigningInfo `json:"signing_info,omitempty"`
	ExpireDatetime time.Time   `json:"expire_datetime,omitempty"`
}

// CertificationSigningInfo is the signing info PKE
type CertificationSigningInfo struct {
	City             string `json:"city,omitempty"`
	State            string `json:"state,omitempty"`
	CountryCode      string `json:"country_code,omitempty"`
	CommonName       string `json:"common_name,omitempty"`
	Organization     string `json:"organization,omitempty"`
	EmailAddress     string `json:"email_address,omitempty"`
	CommonNameSuffix string `json:"common_name_suffix,omitempty"`
}

// CitrixVMReferenceList is the CitrixVMReferenceList
type CitrixVMReferenceList struct {
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

// ResourceLocation is the ResourceLocation
type ResourceLocation struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// CitrixConnectorConfig is the CitrixConnectorConfig
type CitrixConnectorConfig struct {
	CitrixVMReferenceList []CitrixVMReferenceList `json:"citrix_vm_reference_list,omitempty"`
	ClientSecret          string                  `json:"client_secret,omitempty"`
	CustomerID            string                  `json:"customer_id,omitempty"`
	ClientID              string                  `json:"client_id,omitempty"`
	ResourceLocation      ResourceLocation        `json:"resource_location,omitempty"`
}

// ExternalConfigurations is the ExternalConfigurations
type ExternalConfigurations struct {
	CitrixConnectorConfig CitrixConnectorConfig `json:"citrix_connector_config,omitempty"`
}

// CaCertificateList is the CaCertificateList
type CaCertificateList struct {
	CaName      string `json:"ca_name,omitempty"`
	Certificate string `json:"certificate,omitempty"`
}

// ManagementServerList is the ManagementServerList
type ManagementServerList struct {
	DrsEnabled bool     `json:"drs_enabled,omitempty"`
	IP         string   `json:"ip,omitempty"`
	Type       string   `json:"type,omitempty"`
	StatusList []string `json:"status_list,omitempty"`
}

// AuthorizedPublicKeyList is the AuthorizedPublicKeyList
type AuthorizedPublicKeyList struct {
	Name string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
}

// Build is the Build
type Build struct {
	CommitID      string    `json:"commit_id,omitempty"`
	Version       string    `json:"version,omitempty"`
	ShortCommitID string    `json:"short_commit_id,omitempty"`
	FullVersion   string    `json:"full_version,omitempty"`
	CommitDate    time.Time `json:"commit_date,omitempty"`
	BuildType     string    `json:"build_type,omitempty"`
}

// HTTPProxyWhitelist is the HTTPProxyWhitelist
type HTTPProxyWhitelist struct {
	Target     string `json:"target,omitempty"`
	TargetType string `json:"target_type,omitempty"`
}

// Credentials are the Credentials
type Credentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type NutanixGuestTools struct {
	NgtState              string      `json:"ngt_state"`
	Credentials           Credentials `json:"credentials"`
	State                 string      `json:"state"`
	IsoMountState         string      `json:"iso_mount_state"`
	EnabledCapabilityList []string    `json:"enabled_capability_list"`
	Version               string      `json:"version"`
}

type GuestTools struct {
	NutanixGuestTools NutanixGuestTools `json:"nutanix_guest_tools"`
}

// Address is the address
type Address struct {
	IP       string `json:"ip,omitempty"`
	Ipv6     string `json:"ipv6,omitempty"`
	IsBackup bool   `json:"is_backup,omitempty"`
	Port     int    `json:"port,omitempty"`
	Fqdn     string `json:"fqdn,omitempty"`
}

// HTTPProxyList is the HTTPProxyList
type HTTPProxyList struct {
	Credentials   Credentials `json:"credentials,omitempty"`
	Address       Address     `json:"address,omitempty"`
	ProxyTypeList []string    `json:"proxy_type_list,omitempty"`
}

// Server is the server
type Server struct {
	Credentials   Credentials `json:"credentials,omitempty"`
	Address       Address     `json:"address,omitempty"`
	ProxyTypeList []string    `json:"proxy_type_list,omitempty"`
}

// SMTPServer is the SMTP server
type SMTPServer struct {
	EmailAddress string `json:"email_address,omitempty"`
	Type         string `json:"type,omitempty"`
	Server       Server `json:"server,omitempty"`
}

// DomainCredentials are the domain credentials
type DomainCredentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// DomainServer is the domain server
type DomainServer struct {
	Nameserver        string            `json:"nameserver,omitempty"`
	Name              string            `json:"name,omitempty"`
	DomainCredentials DomainCredentials `json:"domain_credentials,omitempty"`
}

// Network is the network
type Network struct {
	HTTPProxyWhitelist     []HTTPProxyWhitelist `json:"http_proxy_whitelist,omitempty"`
	MasqueradingPort       int                  `json:"masquerading_port,omitempty"`
	MasqueradingIP         string               `json:"masquerading_ip,omitempty"`
	ExternalIP             string               `json:"external_ip,omitempty"`
	HTTPProxyList          []HTTPProxyList      `json:"http_proxy_list,omitempty"`
	SMTPServer             SMTPServer           `json:"smtp_server,omitempty"`
	NtpServerIPList        []string             `json:"ntp_server_ip_list,omitempty"`
	ExternalSubnet         string               `json:"external_subnet,omitempty"`
	ExternalDataServicesIP string               `json:"external_data_services_ip,omitempty"`
	DomainServer           DomainServer         `json:"domain_server,omitempty"`
	NameServerIPList       []string             `json:"name_server_ip_list,omitempty"`
	NfsSubnetWhitelist     []string             `json:"nfs_subnet_whitelist,omitempty"`
	InternalSubnet         string               `json:"internal_subnet,omitempty"`
}

// VMEfficiencyMap is the VMEfficiencyMap
type VMEfficiencyMap struct {
}

// Analysis is the Analysis
type Analysis struct {
	VMEfficiencyMap VMEfficiencyMap `json:"vm_efficiency_map,omitempty"`
}

// CategoriesMapping is mapping for categories
type CategoriesMapping struct {
}

type ACPMetadata struct {
	Kind                 string            `json:"kind"`
	ProjectReference     ProjectReference  `json:"project_reference"`
	UUID                 string            `json:"uuid"`
	Name                 string            `json:"name"`
	SpecVersion          int               `json:"spec_version"`
	SpecHash             string            `json:"spec_hash"`
	Categories           Categories        `json:"categories"`
	CategoriesMapping    CategoriesMapping `json:"categories_mapping"`
	UseCategoriesMapping bool              `json:"use_categories_mapping"`
	CreationTime         string            `json:"creation_time"`
	LastUpdateTime       string            `json:"last_update_time"`
	OwnerReference       OwnerReference    `json:"owner_reference"`
	ShouldForceTranslate bool              `json:"should_force_translate"`
	EntityVersion        string            `json:"entity_version"`
}
type UserReferenceList struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
type UserGroupReferenceList struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
type RoleReference struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
type RightHandSide struct {
	Collection string     `json:"collection"`
	Categories Categories `json:"categories"`
	UUIDList   []string   `json:"uuid_list"`
}
type ScopeFilterExpressionList struct {
	LeftHandSide  string        `json:"left_hand_side"`
	Operator      string        `json:"operator"`
	RightHandSide RightHandSide `json:"right_hand_side"`
}
type LeftHandSide struct {
	EntityType string `json:"entity_type"`
}
type EntityFilterExpressionList struct {
	LeftHandSide  LeftHandSide  `json:"left_hand_side"`
	Operator      string        `json:"operator"`
	RightHandSide RightHandSide `json:"right_hand_side"`
}
type ContextList struct {
	ScopeFilterExpressionList  []ScopeFilterExpressionList  `json:"scope_filter_expression_list"`
	EntityFilterExpressionList []EntityFilterExpressionList `json:"entity_filter_expression_list"`
}
type FilterList struct {
	ContextList []ContextList `json:"context_list"`
}
type ACPResources struct {
	UserReferenceList      []UserReferenceList      `json:"user_reference_list"`
	UserGroupReferenceList []UserGroupReferenceList `json:"user_group_reference_list"`
	RoleReference          RoleReference            `json:"role_reference"`
	FilterList             FilterList               `json:"filter_list"`
}
type ACPSpec struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Resources   ACPResources `json:"resources"`
}

type AutoGenerated struct {
	APIVersion string     `json:"api_version"`
	Metadata   Metadata   `json:"metadata"`
	Entities   []Entities `json:"entities"`
}

// OplogUsage produces the CVM OpLog utilization with a very precise result
type OplogUsage struct {
	OplogDiskPct  float64 `json:"oplog_disk_pct"`
	OplogDiskSize int64   `json:"oplog_disk_size"`
}

// ControllerVM returns information about CVM and utilization
type ControllerVM struct {
	IP         string     `json:"ip"`
	NatIP      string     `json:"nat_ip"`
	NatPort    int        `json:"nat_port"`
	OplogUsage OplogUsage `json:"oplog_usage"`
}
type DomainCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type WindowsDomain struct {
	Name                 string           `json:"name"`
	DomainName           string           `json:"domain_name"`
	OrganizationUnitPath string           `json:"organization_unit_path"`
	NamePrefix           string           `json:"name_prefix"`
	NameServerIP         string           `json:"name_server_ip"`
	DomainCredential     DomainCredential `json:"domain_credential"`
}

type ClusterReference struct {
	Kind string `json:"kind"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Ipmi struct {
	IP string `json:"ip"`
}
type Hypervisor struct {
	IP                 string `json:"ip"`
	HypervisorFullName string `json:"hypervisor_full_name"`
	NumVms             int    `json:"num_vms"`
}
type Block struct {
	BlockSerialNumber string `json:"block_serial_number"`
	BlockModel        string `json:"block_model"`
}
type RackableUnitReference struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
type HostDisksReferenceList struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
type ConsumerReference struct {
	Kind string `json:"kind"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
type GpuList struct {
	Name                   string            `json:"name"`
	Mode                   string            `json:"mode"`
	DeviceID               int               `json:"device_id"`
	Vendor                 string            `json:"vendor"`
	Status                 string            `json:"status"`
	Index                  int               `json:"index"`
	NumVgpusAllocated      int               `json:"num_vgpus_allocated"`
	Assignable             bool              `json:"assignable"`
	PciAddress             string            `json:"pci_address"`
	NumaNode               int               `json:"numa_node"`
	GuestDriverVersion     string            `json:"guest_driver_version"`
	FrameBufferSizeMib     int               `json:"frame_buffer_size_mib"`
	MaxInstancesPerVM      int               `json:"max_instances_per_vm"`
	NumVirtualDisplayHeads int               `json:"num_virtual_display_heads"`
	MaxResolution          string            `json:"max_resolution"`
	Fraction               int               `json:"fraction"`
	UUID                   string            `json:"uuid"`
	LicenseList            []string          `json:"license_list"`
	ConsumerReference      ConsumerReference `json:"consumer_reference"`
}

type ParentReference struct {
	Kind string `json:"kind"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
	URL  string `json:"url"`
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
	IP                 string   `json:"ip,omitempty"`
	Type               string   `json:"type,omitempty"`
	IPType             string   `json:"ip_type"`
	PrefixLength       int      `json:"prefix_length"`
	GatewayAddressList []string `json:"gateway_address_list"`
}

// SubnetReference is the SubnetReference
type SubnetReference struct {
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

type SerialPortList struct {
	IsConnected bool `json:"is_connected"`
	Index       int  `json:"index"`
}
type QosPolicy struct {
	ThrottledIops int `json:"throttled_iops"`
}

type NetworkFunctionChainReference struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

// NicList is the NicList
type NicList struct {
	NicType                       string                        `json:"nic_type,omitempty"`
	UUID                          string                        `json:"uuid,omitempty"`
	IPEndpointList                []IPEndpointList              `json:"ip_endpoint_list,omitempty"`
	MacAddress                    string                        `json:"mac_address,omitempty"`
	SubnetReference               SubnetReference               `json:"subnet_reference,omitempty"`
	IsConnected                   bool                          `json:"is_connected,omitempty"`
	VlanMode                      string                        `json:"vlan_mode"`
	TrunkedVlanList               []int                         `json:"trunked_vlan_list"`
	NetworkFunctionChainReference NetworkFunctionChainReference `json:"network_function_chain_reference"`
	NetworkFunctionNicType        string                        `json:"network_function_nic_type"`
	Model                         string                        `json:"model"`
}

// HostReference is the HostReference
type HostReference struct {
	Kind string `json:"kind,omitempty"`
	UUID string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
}

// DiskList is the DiskList
type DiskList struct {
	DataSourceReference  DataSourceReference  `json:"data_source_reference,omitempty"`
	DeviceProperties     DeviceProperties     `json:"device_properties,omitempty"`
	UUID                 string               `json:"uuid,omitempty"`
	DiskSizeBytes        int                  `json:"disk_size_bytes,omitempty"`
	DiskSizeMib          int                  `json:"disk_size_mib,omitempty"`
	VolumeGroupReference VolumeGroupReference `json:"volume_group_reference"`
	StorageConfig        StorageConfig        `json:"storage_config"`
}

// ExecutionContext is the ExecutionContext
type ExecutionContext struct {
	TaskUuids []string `json:"task_uuids,omitempty"`
}

// Version is used for image version
type Version struct {
	ProductName    string `json:"product_name"`
	ProductVersion string `json:"product_version"`
}

// BootDevice is the BootDevice
type BootDevice struct {
	DiskAddress DiskAddress `json:"disk_address,omitempty"`
	MacAddress  string      `json:"mac_address"`
}
type CustomKeyValues struct {
}
type Sysprep struct {
	InstallType     string          `json:"install_type"`
	UnattendXML     string          `json:"unattend_xml"`
	CustomKeyValues CustomKeyValues `json:"custom_key_values"`
}
type CloudInit struct {
	MetaData        string          `json:"meta_data"`
	UserData        string          `json:"user_data"`
	CustomKeyValues CustomKeyValues `json:"custom_key_values"`
}
type GuestCustomization struct {
	IsOverridable bool      `json:"is_overridable"`
	Sysprep       Sysprep   `json:"sysprep"`
	CloudInit     CloudInit `json:"cloud_init"`
}

// BootConfig is the BootConfig
type BootConfig struct {
	BootDevice          BootDevice          `json:"boot_device,omitempty"`
	BootDeviceOrderList []string            `json:"boot_device_order_list"`
	BootType            string              `json:"boot_type"`
	DataSourceReference DataSourceReference `json:"data_source_reference"`
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

// Capabilities are used for categories
type Capabilities struct {
	Cardinality int `json:"cardinality,omitempty"`
}

// ExpressionList used for category get value results
type ExpressionList struct {
	PropertyName       string `json:"property_name,omitempty"`
	Operator           string `json:"operator,omitempty"`
	Value              string `json:"value,omitempty"`
	DisplayForOperator string `json:"display_for_operator,omitempty"`
	DisplayForValue    string `json:"display_for_value,omitempty"`
}

// SelectionCriteriaList used for category get value results
type SelectionCriteriaList struct {
	EntityType     string           `json:"entity_type,omitempty"`
	ExpressionList []ExpressionList `json:"expression_list,omitempty"`
}

// ExclusionList used for category get value results
type ExclusionList struct {
	Kind string `json:"kind,omitempty"`
	UUID string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// InclusionList used for category get value results
type InclusionList struct {
	Kind string `json:"kind,omitempty"`
	UUID string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// AssignmentRule used for category get value results
type AssignmentRule struct {
	Name                  string                  `json:"name,omitempty"`
	Description           string                  `json:"description,omitempty"`
	SelectionCriteriaList []SelectionCriteriaList `json:"selection_criteria_list,omitempty"`
	ExclusionList         []ExclusionList         `json:"exclusion_list,omitempty"`
	InclusionList         []InclusionList         `json:"inclusion_list,omitempty"`
}

type DhcpOptions struct {
	BootFileName         string   `json:"boot_file_name"`
	DomainName           string   `json:"domain_name"`
	DomainNameServerList []string `json:"domain_name_server_list"`
	DomainSearchList     []string `json:"domain_search_list"`
	TftpServerName       string   `json:"tftp_server_name"`
}
type DhcpServerAddress struct {
	IP       string `json:"ip"`
	Ipv6     string `json:"ipv6"`
	Fqdn     string `json:"fqdn"`
	Port     int    `json:"port"`
	IsBackup bool   `json:"is_backup"`
}

type PoolList struct {
	Range string `json:"range"`
}
type IPConfig struct {
	SubnetIP          string            `json:"subnet_ip"`
	PrefixLength      int               `json:"prefix_length"`
	DefaultGatewayIP  string            `json:"default_gateway_ip"`
	DhcpOptions       DhcpOptions       `json:"dhcp_options"`
	DhcpServerAddress DhcpServerAddress `json:"dhcp_server_address"`
	PoolList          []PoolList        `json:"pool_list"`
}

type VirtualNetworkReference struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
type VpcReference struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
type AvailabilityZoneReferenceList struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
