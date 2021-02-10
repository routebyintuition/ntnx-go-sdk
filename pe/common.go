package pe

// Metadata is the metadata in request results
type Metadata struct {
	GrandTotalEntities *int    `json:"grand_total_entities,omitempty"`
	TotalEntities      *int    `json:"total_entities,omitempty"`
	FilterCriteria     *string `json:"filter_criteria,omitempty"`
	SortCriteria       *string `json:"sort_criteria,omitempty"`
	Page               *int    `json:"page,omitempty"`
	Count              *int    `json:"count,omitempty"`
	StartIndex         *int    `json:"start_index,omitempty"`
	EndIndex           *int    `json:"end_index,omitempty"`
}

// DiskHardwareConfig provides HW config for disk-related requests
type DiskHardwareConfig struct {
	SerialNumber           *string      `json:"serial_number,omitempty"`
	DiskID                 *string      `json:"disk_id,omitempty"`
	DiskUUID               *string      `json:"disk_uuid,omitempty"`
	Location               *int         `json:"location,omitempty"`
	Bad                    *bool        `json:"bad,omitempty"`
	Mounted                *bool        `json:"mounted,omitempty"`
	MountPath              *string      `json:"mount_path,omitempty"`
	Model                  *string      `json:"model,omitempty"`
	Vendor                 *string      `json:"vendor,omitempty"`
	BootDisk               *bool        `json:"boot_disk,omitempty"`
	OnlyBootDisk           *bool        `json:"only_boot_disk,omitempty"`
	UnderDiagnosis         *bool        `json:"under_diagnosis,omitempty"`
	BackgroundOperation    *interface{} `json:"background_operation,omitempty"`
	CurrentFirmwareVersion *string      `json:"current_firmware_version,omitempty"`
	TargetFirmwareVersion  *string      `json:"target_firmware_version,omitempty"`
	CanAddAsNewDisk        *bool        `json:"can_add_as_new_disk,omitempty"`
	CanAddAsOldDisk        *bool        `json:"can_add_as_old_disk,omitempty"`
}

// Stats are the statistics/metrics for each returned entity
type Stats struct {
	ControllerRandomOpsPpm                          *string `json:"controller.random_ops_ppm,omitempty"`
	ControllerStorageTierSsdUsageBytes              *string `json:"controller.storage_tier.ssd.usage_bytes,omitempty"`
	ReadIoPpm                                       *string `json:"read_io_ppm,omitempty"`
	ControllerFrontendReadLatencyHistogram1000Us    *string `json:"controller.frontend_read_latency_histogram_1000us,omitempty"`
	ControllerNumIops                               *string `json:"controller_num_iops,omitempty"`
	ControllerFrontendWriteOps                      *string `json:"controller.frontend_write_ops,omitempty"`
	ControllerFrontendWriteLatencyHistogram10000Us  *string `json:"controller.frontend_write_latency_histogram_10000us,omitempty"`
	ControllerReadSizeHistogram1024KB               *string `json:"controller.read_size_histogram_1024kB,omitempty"`
	TotalReadIoTimeUsecs                            *string `json:"total_read_io_time_usecs,omitempty"`
	ControllerTotalReadIoTimeUsecs                  *string `json:"controller_total_read_io_time_usecs,omitempty"`
	ControllerWss3600SWriteMB                       *string `json:"controller.wss_3600s_write_MB,omitempty"`
	ControllerFrontendReadLatencyHistogram50000Us   *string `json:"controller.frontend_read_latency_histogram_50000us,omitempty"`
	ControllerFrontendReadLatencyHistogram2000Us    *string `json:"controller.frontend_read_latency_histogram_2000us,omitempty"`
	ControllerNumWriteIo                            *string `json:"controller_num_write_io,omitempty"`
	ControllerReadSourceCacheSsdBytes               *string `json:"controller.read_source_cache_ssd_bytes,omitempty"`
	ControllerReadSourceOplogBytes                  *string `json:"controller.read_source_oplog_bytes,omitempty"`
	ControllerReadSourceCacheDramBytes              *string `json:"controller.read_source_cache_dram_bytes,omitempty"`
	ControllerRandomReadOps                         *string `json:"controller.random_read_ops,omitempty"`
	ControllerTotalIoTimeUsecs                      *string `json:"controller_total_io_time_usecs,omitempty"`
	ControllerNumSeqIo                              *string `json:"controller_num_seq_io,omitempty"`
	ControllerTotalIoSizeKbytes                     *string `json:"controller_total_io_size_kbytes,omitempty"`
	ControllerWss120SWriteMB                        *string `json:"controller.wss_120s_write_MB,omitempty"`
	ControllerReadSourceBlockStoreBytes             *string `json:"controller.read_source_block_store_bytes,omitempty"`
	ControllerNumIo                                 *string `json:"controller_num_io,omitempty"`
	ControllerReadSourceEstoreZeroBytes             *string `json:"controller.read_source_estore_zero_bytes,omitempty"`
	ControllerNumRandomIo                           *string `json:"controller_num_random_io,omitempty"`
	HypervisorNumReadIo                             *string `json:"hypervisor_num_read_io,omitempty"`
	HypervisorTotalReadIoTimeUsecs                  *string `json:"hypervisor_total_read_io_time_usecs,omitempty"`
	NumIo                                           *string `json:"num_io,omitempty"`
	HypervisorNumWriteIo                            *string `json:"hypervisor_num_write_io,omitempty"`
	ControllerWriteSizeHistogram32KB                *string `json:"controller.write_size_histogram_32kB,omitempty"`
	ControllerFrontendReadLatencyHistogram20000Us   *string `json:"controller.frontend_read_latency_histogram_20000us,omitempty"`
	ControllerReadSizeHistogram32KB                 *string `json:"controller.read_size_histogram_32kB,omitempty"`
	HypervisorNumWriteIops                          *string `json:"hypervisor_num_write_iops,omitempty"`
	AvgIoLatencyUsecs                               *string `json:"avg_io_latency_usecs,omitempty"`
	ControllerWriteIoPpm                            *string `json:"controller_write_io_ppm,omitempty"`
	ControllerReadSourceEstoreSsdBytes              *string `json:"controller.read_source_estore_ssd_bytes,omitempty"`
	HypervisorTotalReadIoSizeKbytes                 *string `json:"hypervisor_total_read_io_size_kbytes,omitempty"`
	ControllerNumWriteIops                          *string `json:"controller_num_write_iops,omitempty"`
	TotalIoTimeUsecs                                *string `json:"total_io_time_usecs,omitempty"`
	ControllerWss3600SReadMB                        *string `json:"controller.wss_3600s_read_MB,omitempty"`
	ControllerSummaryReadSourceSsdBytesPerSec       *string `json:"controller.summary_read_source_ssd_bytes_per_sec,omitempty"`
	ControllerWriteSizeHistogram16KB                *string `json:"controller.write_size_histogram_16kB,omitempty"`
	TotalTransformedUsageBytes                      *string `json:"total_transformed_usage_bytes,omitempty"`
	AvgWriteIoLatencyUsecs                          *string `json:"avg_write_io_latency_usecs,omitempty"`
	ControllerCseTarget90PercentWriteMB             *string `json:"controller.cse_target_90_percent_write_MB,omitempty"`
	NumReadIo                                       *string `json:"num_read_io,omitempty"`
	HypervisorReadIoBandwidthKBps                   *string `json:"hypervisor_read_io_bandwidth_kBps,omitempty"`
	HypervisorTotalIoTimeUsecs                      *string `json:"hypervisor_total_io_time_usecs,omitempty"`
	NumRandomIo                                     *string `json:"num_random_io,omitempty"`
	ControllerWriteDestEstoreBytes                  *string `json:"controller.write_dest_estore_bytes,omitempty"`
	ControllerFrontendWriteLatencyHistogram5000Us   *string `json:"controller.frontend_write_latency_histogram_5000us,omitempty"`
	ControllerStorageTierDasSataPinnedUsageBytes    *string `json:"controller.storage_tier.das-sata.pinned_usage_bytes,omitempty"`
	NumWriteIo                                      *string `json:"num_write_io,omitempty"`
	ControllerFrontendWriteLatencyHistogram2000Us   *string `json:"controller.frontend_write_latency_histogram_2000us,omitempty"`
	ControllerRandomWriteOpsPerSec                  *string `json:"controller.random_write_ops_per_sec,omitempty"`
	ControllerFrontendWriteLatencyHistogram20000Us  *string `json:"controller.frontend_write_latency_histogram_20000us,omitempty"`
	IoBandwidthKBps                                 *string `json:"io_bandwidth_kBps,omitempty"`
	ControllerWriteSizeHistogram512KB               *string `json:"controller.write_size_histogram_512kB,omitempty"`
	ControllerReadSizeHistogram16KB                 *string `json:"controller.read_size_histogram_16kB,omitempty"`
	WriteIoPpm                                      *string `json:"write_io_ppm,omitempty"`
	ControllerAvgWriteIoLatencyUsecs                *string `json:"controller_avg_write_io_latency_usecs,omitempty"`
	ControllerFrontendReadLatencyHistogram100000Us  *string `json:"controller.frontend_read_latency_histogram_100000us,omitempty"`
	NumReadIops                                     *string `json:"num_read_iops,omitempty"`
	ControllerSummaryReadSourceHddBytesPerSec       *string `json:"controller.summary_read_source_hdd_bytes_per_sec,omitempty"`
	ControllerReadSourceExtentCacheBytes            *string `json:"controller.read_source_extent_cache_bytes,omitempty"`
	TimespanUsecs                                   *string `json:"timespan_usecs,omitempty"`
	ControllerNumReadIops                           *string `json:"controller_num_read_iops,omitempty"`
	ControllerFrontendReadLatencyHistogram10000Us   *string `json:"controller.frontend_read_latency_histogram_10000us,omitempty"`
	ControllerWriteSizeHistogram64KB                *string `json:"controller.write_size_histogram_64kB,omitempty"`
	ControllerFrontendWriteLatencyHistogram0Us      *string `json:"controller.frontend_write_latency_histogram_0us,omitempty"`
	ControllerFrontendWriteLatencyHistogram100000Us *string `json:"controller.frontend_write_latency_histogram_100000us,omitempty"`
	HypervisorNumIo                                 *string `json:"hypervisor_num_io,omitempty"`
	ControllerTotalTransformedUsageBytes            *string `json:"controller_total_transformed_usage_bytes,omitempty"`
	AvgReadIoLatencyUsecs                           *string `json:"avg_read_io_latency_usecs,omitempty"`
	ControllerTotalReadIoSizeKbytes                 *string `json:"controller_total_read_io_size_kbytes,omitempty"`
	ControllerReadIoPpm                             *string `json:"controller_read_io_ppm,omitempty"`
	ControllerFrontendOps                           *string `json:"controller.frontend_ops,omitempty"`
	ControllerWss120SReadMB                         *string `json:"controller.wss_120s_read_MB,omitempty"`
	ControllerReadSizeHistogram512KB                *string `json:"controller.read_size_histogram_512kB,omitempty"`
	HypervisorAvgReadIoLatencyUsecs                 *string `json:"hypervisor_avg_read_io_latency_usecs,omitempty"`
	ControllerWriteSizeHistogram1024KB              *string `json:"controller.write_size_histogram_1024kB,omitempty"`
	ControllerWriteDestBlockStoreBytes              *string `json:"controller.write_dest_block_store_bytes,omitempty"`
	ControllerReadSizeHistogram4KB                  *string `json:"controller.read_size_histogram_4kB,omitempty"`
	NumWriteIops                                    *string `json:"num_write_iops,omitempty"`
	ControllerRandomOpsPerSec                       *string `json:"controller.random_ops_per_sec,omitempty"`
	NumIops                                         *string `json:"num_iops,omitempty"`
	ControllerStorageTierCloudPinnedUsageBytes      *string `json:"controller.storage_tier.cloud.pinned_usage_bytes,omitempty"`
	ControllerAvgIoLatencyUsecs                     *string `json:"controller_avg_io_latency_usecs,omitempty"`
	ControllerReadSizeHistogram8KB                  *string `json:"controller.read_size_histogram_8kB,omitempty"`
	ControllerNumReadIo                             *string `json:"controller_num_read_io,omitempty"`
	ControllerSeqIoPpm                              *string `json:"controller_seq_io_ppm,omitempty"`
	ControllerReadIoBandwidthKBps                   *string `json:"controller_read_io_bandwidth_kBps,omitempty"`
	ControllerIoBandwidthKBps                       *string `json:"controller_io_bandwidth_kBps,omitempty"`
	ControllerReadSizeHistogram0KB                  *string `json:"controller.read_size_histogram_0kB,omitempty"`
	ControllerRandomOps                             *string `json:"controller.random_ops,omitempty"`
	HypervisorTimespanUsecs                         *string `json:"hypervisor_timespan_usecs,omitempty"`
	TotalReadIoSizeKbytes                           *string `json:"total_read_io_size_kbytes,omitempty"`
	HypervisorTotalIoSizeKbytes                     *string `json:"hypervisor_total_io_size_kbytes,omitempty"`
	ControllerFrontendOpsPerSec                     *string `json:"controller.frontend_ops_per_sec,omitempty"`
	ControllerWriteDestOplogBytes                   *string `json:"controller.write_dest_oplog_bytes,omitempty"`
	ControllerFrontendWriteLatencyHistogram1000Us   *string `json:"controller.frontend_write_latency_histogram_1000us,omitempty"`
	HypervisorNumReadIops                           *string `json:"hypervisor_num_read_iops,omitempty"`
	ControllerSummaryReadSourceCacheBytesPerSec     *string `json:"controller.summary_read_source_cache_bytes_per_sec,omitempty"`
	ControllerWriteIoBandwidthKBps                  *string `json:"controller_write_io_bandwidth_kBps,omitempty"`
	ControllerUserBytes                             *string `json:"controller_user_bytes,omitempty"`
	HypervisorAvgWriteIoLatencyUsecs                *string `json:"hypervisor_avg_write_io_latency_usecs,omitempty"`
	ControllerStorageTierSsdPinnedUsageBytes        *string `json:"controller.storage_tier.ssd.pinned_usage_bytes,omitempty"`
	ReadIoBandwidthKBps                             *string `json:"read_io_bandwidth_kBps,omitempty"`
	ControllerFrontendReadOps                       *string `json:"controller.frontend_read_ops,omitempty"`
	HypervisorNumIops                               *string `json:"hypervisor_num_iops,omitempty"`
	HypervisorIoBandwidthKBps                       *string `json:"hypervisor_io_bandwidth_kBps,omitempty"`
	ControllerWss120SUnionMB                        *string `json:"controller.wss_120s_union_MB,omitempty"`
	ControllerReadSourceEstoreHddBytes              *string `json:"controller.read_source_estore_hdd_bytes,omitempty"`
	ControllerRandomIoPpm                           *string `json:"controller_random_io_ppm,omitempty"`
	ControllerCseTarget90PercentReadMB              *string `json:"controller.cse_target_90_percent_read_MB,omitempty"`
	ControllerStorageTierDasSataUsageBytes          *string `json:"controller.storage_tier.das-sata.usage_bytes,omitempty"`
	ControllerFrontendReadLatencyHistogram5000Us    *string `json:"controller.frontend_read_latency_histogram_5000us,omitempty"`
	ControllerAvgReadIoSizeKbytes                   *string `json:"controller_avg_read_io_size_kbytes,omitempty"`
	WriteIoBandwidthKBps                            *string `json:"write_io_bandwidth_kBps,omitempty"`
	ControllerRandomReadOpsPerSec                   *string `json:"controller.random_read_ops_per_sec,omitempty"`
	ControllerReadSizeHistogram64KB                 *string `json:"controller.read_size_histogram_64kB,omitempty"`
	ControllerWss3600SUnionMB                       *string `json:"controller.wss_3600s_union_MB,omitempty"`
	RandomIoPpm                                     *string `json:"random_io_ppm,omitempty"`
	TotalUntransformedUsageBytes                    *string `json:"total_untransformed_usage_bytes,omitempty"`
	ControllerFrontendReadLatencyHistogram0Us       *string `json:"controller.frontend_read_latency_histogram_0us,omitempty"`
	ControllerRandomWriteOps                        *string `json:"controller.random_write_ops,omitempty"`
	ControllerAvgWriteIoSizeKbytes                  *string `json:"controller_avg_write_io_size_kbytes,omitempty"`
	ControllerAvgReadIoLatencyUsecs                 *string `json:"controller_avg_read_io_latency_usecs,omitempty"`
	TotalIoSizeKbytes                               *string `json:"total_io_size_kbytes,omitempty"`
	ControllerStorageTierCloudUsageBytes            *string `json:"controller.storage_tier.cloud.usage_bytes,omitempty"`
	ControllerFrontendWriteLatencyHistogram50000Us  *string `json:"controller.frontend_write_latency_histogram_50000us,omitempty"`
	ControllerWriteSizeHistogram8KB                 *string `json:"controller.write_size_histogram_8kB,omitempty"`
	ControllerTimespanUsecs                         *string `json:"controller_timespan_usecs,omitempty"`
	NumSeqIo                                        *string `json:"num_seq_io,omitempty"`
	ControllerWriteSizeHistogram4KB                 *string `json:"controller.write_size_histogram_4kB,omitempty"`
	SeqIoPpm                                        *string `json:"seq_io_ppm,omitempty"`
	ControllerWriteSizeHistogram0KB                 *string `json:"controller.write_size_histogram_0kB,omitempty"`
}

// UsageStats provides usage stats for the associated response entities
type UsageStats struct {
	StorageLogicalUsageBytes *string `json:"storage.logical_usage_bytes,omitempty"`
	StorageCapacityBytes     *string `json:"storage.capacity_bytes,omitempty"`
	StorageFreeBytes         *string `json:"storage.free_bytes,omitempty"`
	StorageUsageBytes        *string `json:"storage.usage_bytes,omitempty"`
}

// Entities are common types between most response results
type Entities struct {
	ID                      *string             `json:"id,omitempty"`
	DiskUUID                *string             `json:"disk_uuid,omitempty"`
	ClusterUUID             *string             `json:"cluster_uuid,omitempty"`
	StorageTierName         *string             `json:"storage_tier_name,omitempty"`
	ServiceVmid             *string             `json:"service_vmid,omitempty"`
	NodeUUID                *string             `json:"node_uuid,omitempty"`
	LastServiceVmid         *interface{}        `json:"last_service_vmid,omitempty"`
	LastNodeUUID            *interface{}        `json:"last_node_uuid,omitempty"`
	HostName                *string             `json:"host_name,omitempty"`
	CvmIPAddress            *string             `json:"cvm_ip_address,omitempty"`
	NodeName                *string             `json:"node_name,omitempty"`
	MountPath               *string             `json:"mount_path,omitempty"`
	DiskSize                *int64              `json:"disk_size,omitempty"`
	MarkedForRemoval        *bool               `json:"marked_for_removal,omitempty"`
	DataMigrated            *bool               `json:"data_migrated,omitempty"`
	Online                  *bool               `json:"online,omitempty"`
	DiskStatus              *string             `json:"disk_status,omitempty"`
	Location                *int                `json:"location,omitempty"`
	SelfManagedNvme         *bool               `json:"self_managed_nvme,omitempty"`
	SelfEncryptingDrive     *bool               `json:"self_encrypting_drive,omitempty"`
	DiskHardwareConfig      *DiskHardwareConfig `json:"disk_hardware_config,omitempty"`
	DynamicRingChangingNode *interface{}        `json:"dynamic_ring_changing_node,omitempty"`
	Stats                   *Stats              `json:"stats,omitempty"`
	UsageStats              *UsageStats         `json:"usage_stats,omitempty"`
	VirtualDiskID           *string             `json:"virtual_disk_id,omitempty"`
	UUID                    *string             `json:"uuid,omitempty"`
	DeviceUUID              *string             `json:"device_uuid,omitempty"`
	NutanixNfsfilePath      *string             `json:"nutanix_nfsfile_path,omitempty"`
	DiskAddress             *string             `json:"disk_address,omitempty"`
	AttachedVMID            *string             `json:"attached_vm_id,omitempty"`
	AttachedVMUUID          *string             `json:"attached_vm_uuid,omitempty"`
	AttachedVmname          *string             `json:"attached_vmname,omitempty"`
	AttachedVolumeGroupID   *string             `json:"attached_volume_group_id,omitempty"`
	DiskCapacityInBytes     *int64              `json:"disk_capacity_in_bytes,omitempty"`
	StorageContainerID      *string             `json:"storage_container_id,omitempty"`
	StorageContainerUUID    *string             `json:"storage_container_uuid,omitempty"`
	FlashModeEnabled        *interface{}        `json:"flash_mode_enabled,omitempty"`
	DataSourceURL           *interface{}        `json:"data_source_url,omitempty"`
}

// AlertSummaries used in vdisk get request
type AlertSummaries struct {
}

// AlertSummary used in vdisk get request
type AlertSummary struct {
	AlertSummaries *AlertSummaries `json:"alert_summaries,omitempty"`
	Count          *int            `json:"count,omitempty"`
}

// ChecksInError used in vdisk get request
type ChecksInError struct {
}

// DetailedCheckSummary used in vdisk get request
type DetailedCheckSummary struct {
}

// EntityTypeSummaries used in vdisk get request
type EntityTypeSummaries struct {
	ChecksInError        *ChecksInError        `json:"checks_in_error,omitempty"`
	DetailedCheckSummary *DetailedCheckSummary `json:"detailed_check_summary,omitempty"`
	EntityType           *string               `json:"entity_type,omitempty"`
	FilterCriteria       *string               `json:"filter_criteria,omitempty"`
	HealthSummary        *HealthSummary        `json:"health_summary,omitempty"`
}

// HealthCheckSummaries used in vdisk get request
type HealthCheckSummaries struct {
}

// HealthSummary is used in vdisk get request
type HealthSummary struct {
	EntityTypeSummaries  *[]EntityTypeSummaries `json:"entity_type_summaries,omitempty"`
	HealthCheckSummaries *HealthCheckSummaries  `json:"health_check_summaries,omitempty"`
	HealthStatus         *string                `json:"health_status,omitempty"`
}

// DiskAddress provides response on VM GET
type DiskAddress struct {
	DeviceBus       *string `json:"device_bus,omitempty"`
	DeviceIndex     *int    `json:"device_index,omitempty"`
	DeviceUUID      *string `json:"device_uuid,omitempty"`
	DiskLabel       *string `json:"disk_label,omitempty"`
	IsCdrom         *bool   `json:"is_cdrom,omitempty"`
	NdfsFilepath    *string `json:"ndfs_filepath,omitempty"`
	VmdiskUUID      *string `json:"vmdisk_uuid,omitempty"`
	VolumeGroupUUID *string `json:"volume_group_uuid,omitempty"`
}

// Boot provides response on VM GET
type Boot struct {
	BootDeviceOrder        *[]string    `json:"boot_device_order,omitempty"`
	BootDeviceType         *string      `json:"boot_device_type,omitempty"`
	DiskAddress            *DiskAddress `json:"disk_address,omitempty"`
	HardwareVirtualization *bool        `json:"hardware_virtualization,omitempty"`
	MacAddr                *string      `json:"mac_addr,omitempty"`
	SecureBoot             *bool        `json:"secure_boot,omitempty"`
	UefiBoot               *bool        `json:"uefi_boot,omitempty"`
}

// SerialPorts provides response on VM GET
type SerialPorts struct {
	Index *int    `json:"index,omitempty"`
	Type  *string `json:"type,omitempty"`
}

// FilesToInjectList provides response on VM GET
type FilesToInjectList struct {
	DestinationPath *string `json:"destination_path,omitempty"`
	SourcePath      *string `json:"source_path,omitempty"`
}

// VMCustomizationConfig provides response on VM GET
type VMCustomizationConfig struct {
	DatasourceType    *string              `json:"datasource_type,omitempty"`
	FilesToInjectList *[]FilesToInjectList `json:"files_to_inject_list,omitempty"`
	FreshInstall      *bool                `json:"fresh_install,omitempty"`
	Userdata          *string              `json:"userdata,omitempty"`
	UserdataPath      *string              `json:"userdata_path,omitempty"`
}

// SourceDiskAddress provides response on VM GET
type SourceDiskAddress struct {
	DeviceBus       *string `json:"device_bus,omitempty"`
	DeviceIndex     *int    `json:"device_index,omitempty"`
	DeviceUUID      *string `json:"device_uuid,omitempty"`
	DiskLabel       *string `json:"disk_label,omitempty"`
	IsCdrom         *bool   `json:"is_cdrom,omitempty"`
	NdfsFilepath    *string `json:"ndfs_filepath,omitempty"`
	VmdiskUUID      *string `json:"vmdisk_uuid,omitempty"`
	VolumeGroupUUID *string `json:"volume_group_uuid,omitempty"`
}

// VMDiskInfo provides response on VM GET
type VMDiskInfo struct {
	DataSourceURL        *string            `json:"data_source_url,omitempty"`
	DiskAddress          *DiskAddress       `json:"disk_address,omitempty"`
	FlashModeEnabled     *bool              `json:"flash_mode_enabled,omitempty"`
	IsCdrom              *bool              `json:"is_cdrom,omitempty"`
	IsEmpty              *bool              `json:"is_empty,omitempty"`
	IsHotRemoveEnabled   *bool              `json:"is_hot_remove_enabled,omitempty"`
	IsScsiPassthrough    *bool              `json:"is_scsi_passthrough,omitempty"`
	IsThinProvisioned    *bool              `json:"is_thin_provisioned,omitempty"`
	Shared               *bool              `json:"shared,omitempty"`
	Size                 *int               `json:"size,omitempty"`
	SourceDiskAddress    *SourceDiskAddress `json:"source_disk_address,omitempty"`
	StorageContainerUUID *string            `json:"storage_container_uuid,omitempty"`
}

// VMDiskClone provides response on VM GET
type VMDiskClone struct {
	DiskAddress          *DiskAddress `json:"disk_address,omitempty"`
	MinimumSize          *int         `json:"minimum_size,omitempty"`
	SnapshotGroupUUID    *string      `json:"snapshot_group_uuid,omitempty"`
	StorageContainerUUID *string      `json:"storage_container_uuid,omitempty"`
}

// VMDiskCloneExternal provides response from VM GET
type VMDiskCloneExternal struct {
	ExternalDiskURL      *string `json:"external_disk_url,omitempty"`
	Size                 *int    `json:"size,omitempty"`
	StorageContainerUUID *string `json:"storage_container_uuid,omitempty"`
}

// VMDiskCreate provides response from VM GET
type VMDiskCreate struct {
	Size                 *int    `json:"size,omitempty"`
	StorageContainerUUID *string `json:"storage_container_uuid,omitempty"`
}

// VMDiskPassthruExternal provides response from VM GET
type VMDiskPassthruExternal struct {
	ExternalDiskURL      *string `json:"external_disk_url,omitempty"`
	StorageContainerUUID *string `json:"storage_container_uuid,omitempty"`
}

// VMDisks provides VDISK information during VM GET
type VMDisks struct {
	DiskAddress            *DiskAddress            `json:"disk_address,omitempty"`
	FlashModeEnabled       *bool                   `json:"flash_mode_enabled,omitempty"`
	IsCdrom                *bool                   `json:"is_cdrom,omitempty"`
	IsEmpty                *bool                   `json:"is_empty,omitempty"`
	IsScsiPassThrough      *bool                   `json:"is_scsi_pass_through,omitempty"`
	IsThinProvisioned      *bool                   `json:"is_thin_provisioned,omitempty"`
	VMDiskClone            *VMDiskClone            `json:"vm_disk_clone,omitempty"`
	VMDiskCloneExternal    *VMDiskCloneExternal    `json:"vm_disk_clone_external,omitempty"`
	VMDiskCreate           *VMDiskCreate           `json:"vm_disk_create,omitempty"`
	VMDiskPassthruExternal *VMDiskPassthruExternal `json:"vm_disk_passthru_external,omitempty"`
}

// VMFeatures response from VM GET
type VMFeatures struct {
}

// VMGpus provides VM GPU information during VM GET
type VMGpus struct {
	Assignable             *bool     `json:"assignable,omitempty"`
	DeviceID               *int      `json:"device_id,omitempty"`
	DeviceName             *string   `json:"device_name,omitempty"`
	Fraction               *int      `json:"fraction,omitempty"`
	FrameBufferSizeBytes   *int      `json:"frame_buffer_size_bytes,omitempty"`
	GpuMode                *string   `json:"gpu_mode,omitempty"`
	GpuProfile             *string   `json:"gpu_profile,omitempty"`
	GpuType                *string   `json:"gpu_type,omitempty"`
	GpuVendor              *string   `json:"gpu_vendor,omitempty"`
	GuestDriverVersion     *string   `json:"guest_driver_version,omitempty"`
	InUse                  *bool     `json:"in_use,omitempty"`
	Licenses               *[]string `json:"licenses,omitempty"`
	MaxInstancesPerVM      *int      `json:"max_instances_per_vm,omitempty"`
	MaxResolution          *string   `json:"max_resolution,omitempty"`
	NumVirtualDisplayHeads *int      `json:"num_virtual_display_heads,omitempty"`
	NumaNode               *int      `json:"numa_node,omitempty"`
	Sbdf                   *string   `json:"sbdf,omitempty"`
	VMUuids                *[]string `json:"vm_uuids,omitempty"`
}

// VMNics provides VM NIC information during VM calls
type VMNics struct {
	AdapterType        *string   `json:"adapter_type,omitempty"`
	IPAddress          *string   `json:"ip_address,omitempty"`
	IPAddresses        *[]string `json:"ip_addresses,omitempty"`
	IsConnected        *bool     `json:"is_connected,omitempty"`
	MacAddress         *string   `json:"mac_address,omitempty"`
	Model              *string   `json:"model,omitempty"`
	NetworkUUID        *string   `json:"network_uuid,omitempty"`
	NicUUID            *string   `json:"nic_uuid,omitempty"`
	PortID             *string   `json:"port_id,omitempty"`
	RequestIP          *bool     `json:"request_ip,omitempty"`
	RequestedIPAddress *string   `json:"requested_ip_address,omitempty"`
	VlanMode           *string   `json:"vlan_mode,omitempty"`
}
