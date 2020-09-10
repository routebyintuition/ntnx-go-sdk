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
