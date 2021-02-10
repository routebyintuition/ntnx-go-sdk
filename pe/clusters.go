package pe

import (
	"fmt"
	"net/http"
)

// ClusterService handles information on the PE cluster
type ClusterService Service

// ClusterGetRequest provides nothing
type ClusterGetRequest struct{}

// ClusterGetResponse provides details on a PE cluster
type ClusterGetResponse struct {
	AlertSummary                          *AlertSummary                        `json:"alert_summary,omitempty"`
	AllHypervNodesInFailoverCluster       *bool                                `json:"all_hyperv_nodes_in_failover_cluster,omitempty"`
	BlockSerials                          *[]string                            `json:"block_serials,omitempty"`
	Cloudcluster                          *bool                                `json:"cloudcluster,omitempty"`
	ClusterArch                           *string                              `json:"cluster_arch,omitempty"`
	ClusterDataState                      *ClusterDataState                    `json:"cluster_data_state,omitempty"`
	ClusterExternalAddress                *[]ClusterExternalAddress            `json:"cluster_external_address,omitempty"`
	ClusterExternalDataServicesAddress    *ClusterExternalDataServicesAddress  `json:"cluster_external_data_services_address,omitempty"`
	ClusterExternalDataServicesIpaddress  *string                              `json:"cluster_external_data_services_ipaddress,omitempty"`
	ClusterExternalIpaddress              *string                              `json:"cluster_external_ipaddress,omitempty"`
	ClusterFullyQualifiedDomainName       *string                              `json:"cluster_fully_qualified_domain_name,omitempty"`
	ClusterFunctions                      *[]string                            `json:"cluster_functions,omitempty"`
	ClusterGpus                           *[]string                            `json:"cluster_gpus,omitempty"`
	ClusterIncarnationID                  *int                                 `json:"cluster_incarnation_id,omitempty"`
	ClusterMasqueradingAddress            *ClusterMasqueradingAddress          `json:"cluster_masquerading_address,omitempty"`
	ClusterMasqueradingIpaddress          *string                              `json:"cluster_masquerading_ipaddress,omitempty"`
	ClusterMasqueradingPort               *int                                 `json:"cluster_masquerading_port,omitempty"`
	ClusterRedundancyState                *ClusterRedundancyState              `json:"cluster_redundancy_state,omitempty"`
	ClusterUsageCriticalAlertThresholdPct *int                                 `json:"cluster_usage_critical_alert_threshold_pct,omitempty"`
	ClusterUsageWarningAlertThresholdPct  *int                                 `json:"cluster_usage_warning_alert_threshold_pct,omitempty"`
	ClusterUUID                           *string                              `json:"cluster_uuid,omitempty"`
	CommonCriteriaMode                    *bool                                `json:"common_criteria_mode,omitempty"`
	Credential                            *Credential                          `json:"credential,omitempty"`
	DisableDegradedNodeMonitoring         *bool                                `json:"disable_degraded_node_monitoring,omitempty"`
	Domain                                *string                              `json:"domain,omitempty"`
	EnableEfficientMetricSync             *bool                                `json:"enable_efficient_metric_sync,omitempty"`
	EnableLockDown                        *bool                                `json:"enable_lock_down,omitempty"`
	EnableOnDiskDedup                     *bool                                `json:"enable_on_disk_dedup,omitempty"`
	EnablePasswordRemoteLoginToCluster    *bool                                `json:"enable_password_remote_login_to_cluster,omitempty"`
	EnableRf1Container                    *bool                                `json:"enable_rf1_container,omitempty"`
	EnableShadowClones                    *bool                                `json:"enable_shadow_clones,omitempty"`
	Encrypted                             *bool                                `json:"encrypted,omitempty"`
	EnforceRackableUnitAwarePlacement     *bool                                `json:"enforce_rackable_unit_aware_placement,omitempty"`
	ExternalAddress                       *[]ExternalAddress                   `json:"external_address,omitempty"`
	ExternalSubnet                        *string                              `json:"external_subnet,omitempty"`
	FaultToleranceDomainType              *string                              `json:"fault_tolerance_domain_type,omitempty"`
	FingerprintContentCachePercentage     *int                                 `json:"fingerprint_content_cache_percentage,omitempty"`
	FullVersion                           *string                              `json:"full_version,omitempty"`
	GlobalNfsWhiteList                    *[]string                            `json:"global_nfs_white_list,omitempty"`
	GlobalNfsWhiteListAddress             *[]GlobalNfsWhiteListAddress         `json:"global_nfs_white_list_address,omitempty"`
	GpuDriverVersion                      *string                              `json:"gpu_driver_version,omitempty"`
	HasSelfEncryptingDrive                *bool                                `json:"has_self_encrypting_drive,omitempty"`
	HealthSummary                         *HealthSummary                       `json:"health_summary,omitempty"`
	HTTPProxies                           *[]HTTPProxies                       `json:"http_proxies,omitempty"`
	HypervisorLldpConfig                  *HypervisorLldpConfig                `json:"hypervisor_lldp_config,omitempty"`
	HypervisorSecurityComplianceConfig    *HypervisorSecurityComplianceConfig  `json:"hypervisor_security_compliance_config,omitempty"`
	HypervisorTypes                       *[]string                            `json:"hypervisor_types,omitempty"`
	ID                                    *string                              `json:"id,omitempty"`
	InternalAddress                       *[]InternalAddress                   `json:"internal_address,omitempty"`
	InternalSubnet                        *string                              `json:"internal_subnet,omitempty"`
	IsLts                                 *bool                                `json:"is_lts,omitempty"`
	IsNsenabled                           *bool                                `json:"is_nsenabled,omitempty"`
	IsRegisteredToPc                      *bool                                `json:"is_registered_to_pc,omitempty"`
	IsSspEnabled                          *bool                                `json:"is_ssp_enabled,omitempty"`
	IsUpgradeInProgress                   *bool                                `json:"is_upgrade_in_progress,omitempty"`
	IscsiConfig                           *IscsiConfig                         `json:"iscsi_config,omitempty"`
	ManagementServers                     *[]ManagementServers                 `json:"management_servers,omitempty"`
	Multicluster                          *bool                                `json:"multicluster,omitempty"`
	Name                                  *string                              `json:"name,omitempty"`
	NameServers                           *[]string                            `json:"name_servers,omitempty"`
	NameServersList                       *[]NameServersList                   `json:"name_servers_list,omitempty"`
	NccVersion                            *string                              `json:"ncc_version,omitempty"`
	NosClusterAndHostsDomainJoined        *bool                                `json:"nos_cluster_and_hosts_domain_joined,omitempty"`
	NtpServers                            *[]string                            `json:"ntp_servers,omitempty"`
	NtpServersList                        *[]NtpServersList                    `json:"ntp_servers_list,omitempty"`
	NumNodes                              *int                                 `json:"num_nodes,omitempty"`
	OperationMode                         *string                              `json:"operation_mode,omitempty"`
	PublicKeys                            *[]PublicKeys                        `json:"public_keys,omitempty"`
	RackableUnits                         *[]RackableUnits                     `json:"rackable_units,omitempty"`
	SecurityComplianceConfig              *SecurityComplianceConfig            `json:"security_compliance_config,omitempty"`
	SegmentedIscsiDataServicesAddress     *[]SegmentedIscsiDataServicesAddress `json:"segmented_iscsi_data_services_address,omitempty"`
	SegmentedIscsiDataServicesIpaddress   *string                              `json:"segmented_iscsi_data_services_ipaddress,omitempty"`
	ServiceCenters                        *[]ServiceCenters                    `json:"service_centers,omitempty"`
	SMTPServer                            *SMTPServer                          `json:"smtp_server,omitempty"`
	SsdPinningPercentageLimit             *int                                 `json:"ssd_pinning_percentage_limit,omitempty"`
	Stats                                 *Stats                               `json:"stats,omitempty"`
	StorageType                           *string                              `json:"storage_type,omitempty"`
	SupportVerbosityType                  *string                              `json:"support_verbosity_type,omitempty"`
	TargetVersion                         *string                              `json:"target_version,omitempty"`
	ThresholdForStorageThinProvision      *float64                             `json:"threshold_for_storage_thin_provision,omitempty"`
	Timezone                              *string                              `json:"timezone,omitempty"`
	UsageStats                            *UsageStats                          `json:"usage_stats,omitempty"`
	UUID                                  *string                              `json:"uuid,omitempty"`
	Version                               *string                              `json:"version,omitempty"`
}

// Get makes the call to get one virtual machine by UUID
func (cls *ClusterService) Get(reqdata *ClusterGetRequest) (*ClusterGetResponse, *http.Response, error) {

	u := fmt.Sprintf("cluster")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cls.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var data *ClusterGetResponse
	resp, err := cls.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}
