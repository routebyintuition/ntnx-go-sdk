package karbon

type EtcdConfig struct {
	NodePools []string `json:"node_pools"`
}
type MasterConfig struct {
	DeploymentType string   `json:"deployment_type"`
	NodePools      []string `json:"node_pools"`
}
type WorkerConfig struct {
	NodePools []string `json:"node_pools"`
}

// DeleteSkipCheck is used within the delete cluster call to pass into query parameters
type DeleteSkipCheck struct {
	SkipCheck bool `url:"skip-prechecks"`
}

type IPPoolConfigs struct {
	Cidr string `json:"cidr"`
}
type CalicoConfig struct {
	IPPoolConfigs []IPPoolConfigs `json:"ip_pool_configs"`
}
type FlannelConfig struct {
}
type CniConfig struct {
	CalicoConfig     CalicoConfig  `json:"calico_config"`
	FlannelConfig    FlannelConfig `json:"flannel_config"`
	NodeCidrMaskSize int           `json:"node_cidr_mask_size"`
	PodIpv4Cidr      string        `json:"pod_ipv4_cidr"`
	ServiceIpv4Cidr  string        `json:"service_ipv4_cidr"`
}
type AhvConfig struct {
	CPU                     int    `json:"cpu"`
	DiskMib                 int    `json:"disk_mib"`
	MemoryMib               int    `json:"memory_mib"`
	NetworkUUID             string `json:"network_uuid"`
	PrismElementClusterUUID string `json:"prism_element_cluster_uuid"`
}
type NodePools struct {
	AhvConfig     AhvConfig `json:"ahv_config"`
	Name          string    `json:"name"`
	NodeOsVersion string    `json:"node_os_version"`
	NumInstances  int       `json:"num_instances"`
}
type CreateEtcdConfig struct {
	NodePools []NodePools `json:"node_pools"`
}
type ActivePassiveConfig struct {
	ExternalIpv4Address string `json:"external_ipv4_address"`
}
type MasterNodesConfig struct {
	Ipv4Address  string `json:"ipv4_address"`
	NodePoolName string `json:"node_pool_name"`
}
type ExternalLbConfig struct {
	ExternalIpv4Address string              `json:"external_ipv4_address"`
	MasterNodesConfig   []MasterNodesConfig `json:"master_nodes_config"`
}
type SingleMasterConfig struct {
}
type MastersConfig struct {
	ActivePassiveConfig ActivePassiveConfig `json:"active_passive_config"`
	ExternalLbConfig    ExternalLbConfig    `json:"external_lb_config"`
	NodePools           []NodePools         `json:"node_pools"`
	SingleMasterConfig  SingleMasterConfig  `json:"single_master_config"`
}
type Metadata struct {
	APIVersion string `json:"api_version"`
}
type VolumesConfig struct {
	FileSystem              string `json:"file_system"`
	FlashMode               bool   `json:"flash_mode"`
	Password                string `json:"password"`
	PrismElementClusterUUID string `json:"prism_element_cluster_uuid"`
	StorageContainer        string `json:"storage_container"`
	Username                string `json:"username"`
}
type StorageClassConfig struct {
	DefaultStorageClass bool          `json:"default_storage_class"`
	Name                string        `json:"name"`
	ReclaimPolicy       string        `json:"reclaim_policy"`
	VolumesConfig       VolumesConfig `json:"volumes_config"`
}
type WorkersConfig struct {
	NodePools []NodePools `json:"node_pools"`
}
