package pc

import "time"

// Metadata is the Metadata associated with an API responsee
// this same type can be used across the calls as it includes
// type entries for other API calls
type Metadata struct {
	TotalMatches         int              `json:"total_matches,omitempty"`
	Kind                 string           `json:"kind,omitempty"`
	Length               int              `json:"length,omitempty"`
	Offset               int              `json:"offset,omitempty"`
	LastUpdateTime       time.Time        `json:"last_update_time,omitempty"`
	UUID                 string           `json:"uuid,omitempty"`
	ProjectReference     ProjectReference `json:"project_reference,omitempty"`
	SpecVersion          int              `json:"spec_version,omitempty"`
	CreationTime         time.Time        `json:"creation_time,omitempty"`
	OwnerReference       OwnerReference   `json:"owner_reference,omitempty"`
	Categories           Categories       `json:"categories,omitempty"`
	SpecHash             string           `json:"spec_hash,omitempty"`
	ShouldForceTranslate bool             `json:"should_force_translate,omitempty"`
	Name                 string           `json:"name,omitempty"`
	Filter               string           `json:"filter,omitempty"`
	SortOrder            string           `json:"sort_order,omitempty"`
	SortAttribute        string           `json:"sort_attribute,omitempty"`
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
	BootConfig            BootConfig          `json:"boot_config,omitempty"`
	DiskList              []DiskList          `json:"disk_list,omitempty"`
	GpuList               []interface{}       `json:"gpu_list,omitempty"`
	GuestCustomization    GuestCustomization  `json:"guest_customization,omitempty"`
	HardwareClockTimezone string              `json:"hardware_clock_timezone,omitempty"`
	HostReference         HostReference       `json:"host_reference,omitempty"`
	HypervisorType        string              `json:"hypervisor_type,omitempty"`
	MemorySizeMib         int                 `json:"memory_size_mib,omitempty"`
	NicList               []NicList           `json:"nic_list,omitempty"`
	NumSockets            int                 `json:"num_sockets,omitempty"`
	NumThreadsPerCore     int                 `json:"num_threads_per_core,omitempty"`
	NumVcpusPerSocket     int                 `json:"num_vcpus_per_socket,omitempty"`
	PowerState            string              `json:"power_state,omitempty"`
	PowerStateMechanism   PowerStateMechanism `json:"power_state_mechanism,omitempty"`
	ProtectionType        string              `json:"protection_type,omitempty"`
	SerialPortList        []interface{}       `json:"serial_port_list,omitempty"`
	VgaConsoleEnabled     bool                `json:"vga_console_enabled,omitempty"`
	VnumaConfig           VnumaConfig         `json:"vnuma_config,omitempty"`
	Nodes                 Nodes               `json:"nodes,omitempty"`
	Config                Config              `json:"config,omitempty"`
	Network               Network             `json:"network,omitempty"`
	Analysis              Analysis            `json:"analysis,omitempty"`
	RuntimeStatusList     []string            `json:"runtime_status_list,omitempty"`
}

// Spec is the main Spec type across all Response calls
type Spec struct {
	ClusterReference ClusterReference `json:"cluster_reference,omitempty"`
	Description      string           `json:"description,omitempty"`
	Resources        Resources        `json:"resources,omitempty"`
	Name             string           `json:"name,omitempty"`
	APIVersion       string           `json:"api_version,omitempty"`
}

// ProjectReference is the main ProjectReference type across all Response calls
type ProjectReference struct {
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
	UUID string `json:"uuid,omitempty"`
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
	ReadOnly    bool   `json:"readOnly"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// Status is the main Status type across all Response calls
type Status struct {
	State            string           `json:"state,omitempty"`
	ClusterReference ClusterReference `json:"cluster_reference,omitempty"`
	ExecutionContext ExecutionContext `json:"execution_context,omitempty"`
	Description      string           `json:"description,omitempty"`
	Resources        Resources        `json:"resources,omitempty"`
	Name             string           `json:"name,omitempty"`
	MessageList      []MessageList    `json:"message_list,omitempty"`
}

// Entities is the main Entities type across all Response calls
type Entities struct {
	Status     Status   `json:"status,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
}

// MessageList is the list of messages
type MessageList struct {
	Message string  `json:"message,omitempty"`
	Reason  string  `json:"reason,omitempty"`
	Details Details `json:"details,omitempty"`
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
