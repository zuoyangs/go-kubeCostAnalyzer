package kubernetes

// ResourceInfo 包含了 Kubernetes 工作负载的资源使用信息。
type ResourceInfo struct {
	Name                string  `json:"name"`
	ClusterName         string  `json:"clusterName"`
	Replicas            int32   `json:"replicas"`
	TotalRequestsCpu    float64 `json:"totalRequestsCpu"`
	TotalRequestsMemory float64 `json:"totalRequestsMemory"`
	TotalLimitsCpu      float64 `json:"totalLimitsCpu"`
	TotalLimitsMemory   float64 `json:"totalLimitsMemory"`
}

// ClusterResources 代表跨多个集群的资源使用情况
type ClusterResources map[string]EnvResources

// EnvResources 代表在特定环境（集群）中的资源使用情况
type EnvResources map[string]NamespaceResources

// NamespaceResources 代表在 Kubernetes 命名空间中的资源使用情况
type NamespaceResources map[string]ResourceInfo

// Deployment 代表 Kubernetes Deployment 并实现了 ResourceGetter 接口
type Deployment struct {
	Name       string
	Replicas   int32
	Containers []Container
}

// StatefulSet 代表 Kubernetes StatefulSet 并实现了 ResourceGetter 接口
type StatefulSet struct {
	Name       string
	Replicas   int32
	Containers []Container
}

// Container 代表 Kubernetes 工作负载中的一个容器
type Container struct {
	RequestsCpu    float64
	RequestsMemory float64
	LimitsCpu      float64
	LimitsMemory   float64
}
