package pe

// Metadata is the metadata in request results
type Metadata struct {
	GrandTotalEntities int    `json:"grand_total_entities"`
	TotalEntities      int    `json:"total_entities"`
	FilterCriteria     string `json:"filter_criteria"`
	SortCriteria       string `json:"sort_criteria"`
	Page               int    `json:"page"`
	Count              int    `json:"count"`
	StartIndex         int    `json:"start_index"`
	EndIndex           int    `json:"end_index"`
}

// DiskHardwareConfig provides HW config for disk-related requests
type DiskHardwareConfig struct {
	SerialNumber           string      `json:"serial_number"`
	DiskID                 string      `json:"disk_id"`
	DiskUUID               string      `json:"disk_uuid"`
	Location               int         `json:"location"`
	Bad                    bool        `json:"bad"`
	Mounted                bool        `json:"mounted"`
	MountPath              string      `json:"mount_path"`
	Model                  string      `json:"model"`
	Vendor                 string      `json:"vendor"`
	BootDisk               bool        `json:"boot_disk"`
	OnlyBootDisk           bool        `json:"only_boot_disk"`
	UnderDiagnosis         bool        `json:"under_diagnosis"`
	BackgroundOperation    interface{} `json:"background_operation"`
	CurrentFirmwareVersion string      `json:"current_firmware_version"`
	TargetFirmwareVersion  string      `json:"target_firmware_version"`
	CanAddAsNewDisk        bool        `json:"can_add_as_new_disk"`
	CanAddAsOldDisk        bool        `json:"can_add_as_old_disk"`
}

// Stats are the statistics/metrics for each returned entity
type Stats struct {
	ControllerRandomOpsPpm                          string `json:"controller.random_ops_ppm"`
	ControllerStorageTierSsdUsageBytes              string `json:"controller.storage_tier.ssd.usage_bytes"`
	ReadIoPpm                                       string `json:"read_io_ppm"`
	ControllerFrontendReadLatencyHistogram1000Us    string `json:"controller.frontend_read_latency_histogram_1000us"`
	ControllerNumIops                               string `json:"controller_num_iops"`
	ControllerFrontendWriteOps                      string `json:"controller.frontend_write_ops"`
	ControllerFrontendWriteLatencyHistogram10000Us  string `json:"controller.frontend_write_latency_histogram_10000us"`
	ControllerReadSizeHistogram1024KB               string `json:"controller.read_size_histogram_1024kB"`
	TotalReadIoTimeUsecs                            string `json:"total_read_io_time_usecs"`
	ControllerTotalReadIoTimeUsecs                  string `json:"controller_total_read_io_time_usecs"`
	ControllerWss3600SWriteMB                       string `json:"controller.wss_3600s_write_MB"`
	ControllerFrontendReadLatencyHistogram50000Us   string `json:"controller.frontend_read_latency_histogram_50000us"`
	ControllerFrontendReadLatencyHistogram2000Us    string `json:"controller.frontend_read_latency_histogram_2000us"`
	ControllerNumWriteIo                            string `json:"controller_num_write_io"`
	ControllerReadSourceCacheSsdBytes               string `json:"controller.read_source_cache_ssd_bytes"`
	ControllerReadSourceOplogBytes                  string `json:"controller.read_source_oplog_bytes"`
	ControllerReadSourceCacheDramBytes              string `json:"controller.read_source_cache_dram_bytes"`
	ControllerRandomReadOps                         string `json:"controller.random_read_ops"`
	ControllerTotalIoTimeUsecs                      string `json:"controller_total_io_time_usecs"`
	ControllerNumSeqIo                              string `json:"controller_num_seq_io"`
	ControllerTotalIoSizeKbytes                     string `json:"controller_total_io_size_kbytes"`
	ControllerWss120SWriteMB                        string `json:"controller.wss_120s_write_MB"`
	ControllerReadSourceBlockStoreBytes             string `json:"controller.read_source_block_store_bytes"`
	ControllerNumIo                                 string `json:"controller_num_io"`
	ControllerReadSourceEstoreZeroBytes             string `json:"controller.read_source_estore_zero_bytes"`
	ControllerNumRandomIo                           string `json:"controller_num_random_io"`
	HypervisorNumReadIo                             string `json:"hypervisor_num_read_io"`
	HypervisorTotalReadIoTimeUsecs                  string `json:"hypervisor_total_read_io_time_usecs"`
	NumIo                                           string `json:"num_io"`
	HypervisorNumWriteIo                            string `json:"hypervisor_num_write_io"`
	ControllerWriteSizeHistogram32KB                string `json:"controller.write_size_histogram_32kB"`
	ControllerFrontendReadLatencyHistogram20000Us   string `json:"controller.frontend_read_latency_histogram_20000us"`
	ControllerReadSizeHistogram32KB                 string `json:"controller.read_size_histogram_32kB"`
	HypervisorNumWriteIops                          string `json:"hypervisor_num_write_iops"`
	AvgIoLatencyUsecs                               string `json:"avg_io_latency_usecs"`
	ControllerWriteIoPpm                            string `json:"controller_write_io_ppm"`
	ControllerReadSourceEstoreSsdBytes              string `json:"controller.read_source_estore_ssd_bytes"`
	HypervisorTotalReadIoSizeKbytes                 string `json:"hypervisor_total_read_io_size_kbytes"`
	ControllerNumWriteIops                          string `json:"controller_num_write_iops"`
	TotalIoTimeUsecs                                string `json:"total_io_time_usecs"`
	ControllerWss3600SReadMB                        string `json:"controller.wss_3600s_read_MB"`
	ControllerSummaryReadSourceSsdBytesPerSec       string `json:"controller.summary_read_source_ssd_bytes_per_sec"`
	ControllerWriteSizeHistogram16KB                string `json:"controller.write_size_histogram_16kB"`
	TotalTransformedUsageBytes                      string `json:"total_transformed_usage_bytes"`
	AvgWriteIoLatencyUsecs                          string `json:"avg_write_io_latency_usecs"`
	ControllerCseTarget90PercentWriteMB             string `json:"controller.cse_target_90_percent_write_MB"`
	NumReadIo                                       string `json:"num_read_io"`
	HypervisorReadIoBandwidthKBps                   string `json:"hypervisor_read_io_bandwidth_kBps"`
	HypervisorTotalIoTimeUsecs                      string `json:"hypervisor_total_io_time_usecs"`
	NumRandomIo                                     string `json:"num_random_io"`
	ControllerWriteDestEstoreBytes                  string `json:"controller.write_dest_estore_bytes"`
	ControllerFrontendWriteLatencyHistogram5000Us   string `json:"controller.frontend_write_latency_histogram_5000us"`
	ControllerStorageTierDasSataPinnedUsageBytes    string `json:"controller.storage_tier.das-sata.pinned_usage_bytes"`
	NumWriteIo                                      string `json:"num_write_io"`
	ControllerFrontendWriteLatencyHistogram2000Us   string `json:"controller.frontend_write_latency_histogram_2000us"`
	ControllerRandomWriteOpsPerSec                  string `json:"controller.random_write_ops_per_sec"`
	ControllerFrontendWriteLatencyHistogram20000Us  string `json:"controller.frontend_write_latency_histogram_20000us"`
	IoBandwidthKBps                                 string `json:"io_bandwidth_kBps"`
	ControllerWriteSizeHistogram512KB               string `json:"controller.write_size_histogram_512kB"`
	ControllerReadSizeHistogram16KB                 string `json:"controller.read_size_histogram_16kB"`
	WriteIoPpm                                      string `json:"write_io_ppm"`
	ControllerAvgWriteIoLatencyUsecs                string `json:"controller_avg_write_io_latency_usecs"`
	ControllerFrontendReadLatencyHistogram100000Us  string `json:"controller.frontend_read_latency_histogram_100000us"`
	NumReadIops                                     string `json:"num_read_iops"`
	ControllerSummaryReadSourceHddBytesPerSec       string `json:"controller.summary_read_source_hdd_bytes_per_sec"`
	ControllerReadSourceExtentCacheBytes            string `json:"controller.read_source_extent_cache_bytes"`
	TimespanUsecs                                   string `json:"timespan_usecs"`
	ControllerNumReadIops                           string `json:"controller_num_read_iops"`
	ControllerFrontendReadLatencyHistogram10000Us   string `json:"controller.frontend_read_latency_histogram_10000us"`
	ControllerWriteSizeHistogram64KB                string `json:"controller.write_size_histogram_64kB"`
	ControllerFrontendWriteLatencyHistogram0Us      string `json:"controller.frontend_write_latency_histogram_0us"`
	ControllerFrontendWriteLatencyHistogram100000Us string `json:"controller.frontend_write_latency_histogram_100000us"`
	HypervisorNumIo                                 string `json:"hypervisor_num_io"`
	ControllerTotalTransformedUsageBytes            string `json:"controller_total_transformed_usage_bytes"`
	AvgReadIoLatencyUsecs                           string `json:"avg_read_io_latency_usecs"`
	ControllerTotalReadIoSizeKbytes                 string `json:"controller_total_read_io_size_kbytes"`
	ControllerReadIoPpm                             string `json:"controller_read_io_ppm"`
	ControllerFrontendOps                           string `json:"controller.frontend_ops"`
	ControllerWss120SReadMB                         string `json:"controller.wss_120s_read_MB"`
	ControllerReadSizeHistogram512KB                string `json:"controller.read_size_histogram_512kB"`
	HypervisorAvgReadIoLatencyUsecs                 string `json:"hypervisor_avg_read_io_latency_usecs"`
	ControllerWriteSizeHistogram1024KB              string `json:"controller.write_size_histogram_1024kB"`
	ControllerWriteDestBlockStoreBytes              string `json:"controller.write_dest_block_store_bytes"`
	ControllerReadSizeHistogram4KB                  string `json:"controller.read_size_histogram_4kB"`
	NumWriteIops                                    string `json:"num_write_iops"`
	ControllerRandomOpsPerSec                       string `json:"controller.random_ops_per_sec"`
	NumIops                                         string `json:"num_iops"`
	ControllerStorageTierCloudPinnedUsageBytes      string `json:"controller.storage_tier.cloud.pinned_usage_bytes"`
	ControllerAvgIoLatencyUsecs                     string `json:"controller_avg_io_latency_usecs"`
	ControllerReadSizeHistogram8KB                  string `json:"controller.read_size_histogram_8kB"`
	ControllerNumReadIo                             string `json:"controller_num_read_io"`
	ControllerSeqIoPpm                              string `json:"controller_seq_io_ppm"`
	ControllerReadIoBandwidthKBps                   string `json:"controller_read_io_bandwidth_kBps"`
	ControllerIoBandwidthKBps                       string `json:"controller_io_bandwidth_kBps"`
	ControllerReadSizeHistogram0KB                  string `json:"controller.read_size_histogram_0kB"`
	ControllerRandomOps                             string `json:"controller.random_ops"`
	HypervisorTimespanUsecs                         string `json:"hypervisor_timespan_usecs"`
	TotalReadIoSizeKbytes                           string `json:"total_read_io_size_kbytes"`
	HypervisorTotalIoSizeKbytes                     string `json:"hypervisor_total_io_size_kbytes"`
	ControllerFrontendOpsPerSec                     string `json:"controller.frontend_ops_per_sec"`
	ControllerWriteDestOplogBytes                   string `json:"controller.write_dest_oplog_bytes"`
	ControllerFrontendWriteLatencyHistogram1000Us   string `json:"controller.frontend_write_latency_histogram_1000us"`
	HypervisorNumReadIops                           string `json:"hypervisor_num_read_iops"`
	ControllerSummaryReadSourceCacheBytesPerSec     string `json:"controller.summary_read_source_cache_bytes_per_sec"`
	ControllerWriteIoBandwidthKBps                  string `json:"controller_write_io_bandwidth_kBps"`
	ControllerUserBytes                             string `json:"controller_user_bytes"`
	HypervisorAvgWriteIoLatencyUsecs                string `json:"hypervisor_avg_write_io_latency_usecs"`
	ControllerStorageTierSsdPinnedUsageBytes        string `json:"controller.storage_tier.ssd.pinned_usage_bytes"`
	ReadIoBandwidthKBps                             string `json:"read_io_bandwidth_kBps"`
	ControllerFrontendReadOps                       string `json:"controller.frontend_read_ops"`
	HypervisorNumIops                               string `json:"hypervisor_num_iops"`
	HypervisorIoBandwidthKBps                       string `json:"hypervisor_io_bandwidth_kBps"`
	ControllerWss120SUnionMB                        string `json:"controller.wss_120s_union_MB"`
	ControllerReadSourceEstoreHddBytes              string `json:"controller.read_source_estore_hdd_bytes"`
	ControllerRandomIoPpm                           string `json:"controller_random_io_ppm"`
	ControllerCseTarget90PercentReadMB              string `json:"controller.cse_target_90_percent_read_MB"`
	ControllerStorageTierDasSataUsageBytes          string `json:"controller.storage_tier.das-sata.usage_bytes"`
	ControllerFrontendReadLatencyHistogram5000Us    string `json:"controller.frontend_read_latency_histogram_5000us"`
	ControllerAvgReadIoSizeKbytes                   string `json:"controller_avg_read_io_size_kbytes"`
	WriteIoBandwidthKBps                            string `json:"write_io_bandwidth_kBps"`
	ControllerRandomReadOpsPerSec                   string `json:"controller.random_read_ops_per_sec"`
	ControllerReadSizeHistogram64KB                 string `json:"controller.read_size_histogram_64kB"`
	ControllerWss3600SUnionMB                       string `json:"controller.wss_3600s_union_MB"`
	RandomIoPpm                                     string `json:"random_io_ppm"`
	TotalUntransformedUsageBytes                    string `json:"total_untransformed_usage_bytes"`
	ControllerFrontendReadLatencyHistogram0Us       string `json:"controller.frontend_read_latency_histogram_0us"`
	ControllerRandomWriteOps                        string `json:"controller.random_write_ops"`
	ControllerAvgWriteIoSizeKbytes                  string `json:"controller_avg_write_io_size_kbytes"`
	ControllerAvgReadIoLatencyUsecs                 string `json:"controller_avg_read_io_latency_usecs"`
	TotalIoSizeKbytes                               string `json:"total_io_size_kbytes"`
	ControllerStorageTierCloudUsageBytes            string `json:"controller.storage_tier.cloud.usage_bytes"`
	ControllerFrontendWriteLatencyHistogram50000Us  string `json:"controller.frontend_write_latency_histogram_50000us"`
	ControllerWriteSizeHistogram8KB                 string `json:"controller.write_size_histogram_8kB"`
	ControllerTimespanUsecs                         string `json:"controller_timespan_usecs"`
	NumSeqIo                                        string `json:"num_seq_io"`
	ControllerWriteSizeHistogram4KB                 string `json:"controller.write_size_histogram_4kB"`
	SeqIoPpm                                        string `json:"seq_io_ppm"`
	ControllerWriteSizeHistogram0KB                 string `json:"controller.write_size_histogram_0kB"`
}

// UsageStats provides usage stats for the associated response entities
type UsageStats struct {
	StorageLogicalUsageBytes string `json:"storage.logical_usage_bytes"`
	StorageCapacityBytes     string `json:"storage.capacity_bytes"`
	StorageFreeBytes         string `json:"storage.free_bytes"`
	StorageUsageBytes        string `json:"storage.usage_bytes"`
}

// Entities are common types between most response results
type Entities struct {
	ID                      string             `json:"id"`
	DiskUUID                string             `json:"disk_uuid"`
	ClusterUUID             string             `json:"cluster_uuid"`
	StorageTierName         string             `json:"storage_tier_name"`
	ServiceVmid             string             `json:"service_vmid"`
	NodeUUID                string             `json:"node_uuid"`
	LastServiceVmid         interface{}        `json:"last_service_vmid"`
	LastNodeUUID            interface{}        `json:"last_node_uuid"`
	HostName                string             `json:"host_name"`
	CvmIPAddress            string             `json:"cvm_ip_address"`
	NodeName                string             `json:"node_name"`
	MountPath               string             `json:"mount_path"`
	DiskSize                int64              `json:"disk_size"`
	MarkedForRemoval        bool               `json:"marked_for_removal"`
	DataMigrated            bool               `json:"data_migrated"`
	Online                  bool               `json:"online"`
	DiskStatus              string             `json:"disk_status"`
	Location                int                `json:"location"`
	SelfManagedNvme         bool               `json:"self_managed_nvme"`
	SelfEncryptingDrive     bool               `json:"self_encrypting_drive"`
	DiskHardwareConfig      DiskHardwareConfig `json:"disk_hardware_config"`
	DynamicRingChangingNode interface{}        `json:"dynamic_ring_changing_node"`
	Stats                   Stats              `json:"stats"`
	UsageStats              UsageStats         `json:"usage_stats"`
	VirtualDiskID           string             `json:"virtual_disk_id"`
	UUID                    string             `json:"uuid"`
	DeviceUUID              interface{}        `json:"device_uuid"`
	NutanixNfsfilePath      string             `json:"nutanix_nfsfile_path"`
	DiskAddress             string             `json:"disk_address"`
	AttachedVMID            interface{}        `json:"attached_vm_id"`
	AttachedVMUUID          interface{}        `json:"attached_vm_uuid"`
	AttachedVmname          interface{}        `json:"attached_vmname"`
	AttachedVolumeGroupID   string             `json:"attached_volume_group_id"`
	DiskCapacityInBytes     int64              `json:"disk_capacity_in_bytes"`
	StorageContainerID      string             `json:"storage_container_id"`
	StorageContainerUUID    string             `json:"storage_container_uuid"`
	FlashModeEnabled        interface{}        `json:"flash_mode_enabled"`
	DataSourceURL           interface{}        `json:"data_source_url"`
}
