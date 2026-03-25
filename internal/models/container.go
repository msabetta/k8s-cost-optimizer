package models

// ContainerMetrics rappresenta metriche di un container
type ContainerMetrics struct {
	Name            string  `json:"name"`
	Namespace       string  `json:"namespace"`
	Pod             string  `json:"pod"`
	Node            string  `json:"node,omitempty"`
	CPUUsage        float64 `json:"cpu_usage"`
	CPURequest      float64 `json:"cpu_request"`
	MemoryUsageMB   float64 `json:"memory_usage_mb"`
	MemoryRequestMB float64 `json:"memory_request_mb"`
}