package collector

type ContainerMetrics struct {
	Name            string
	Namespace       string
	Pod             string
	CPUUsage        float64
	CPURequest      float64
	MemoryUsageMB   float64
	MemoryRequestMB float64
}