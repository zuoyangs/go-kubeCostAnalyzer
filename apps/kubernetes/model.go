package kubernetes

type ResourceInfo struct {
	Name                string
	ClusterName         string
	Replicas            int32
	TotalRequestsCpu    float64
	TotalRequestsMemory float64
	TotalLimitsCpu      float64
	TotalLimitsMemory   float64
}

type NamespaceResources map[string]ResourceInfo // namespace -> ResourceInfo
type EnvResources map[string]NamespaceResources // env -> NamespaceResources
type ClusterResources map[string]EnvResources   // cluster -> EnvResources
