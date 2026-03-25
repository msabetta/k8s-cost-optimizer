package models

// Node rappresenta un nodo Kubernetes
type Node struct {
	Name          string  `json:"name"`
	CPUCapacity   float64 `json:"cpu_capacity"`
	MemoryGB      float64 `json:"memory_gb"`
	CPUAllocated  float64 `json:"cpu_allocated"`
	MemoryUsedGB  float64 `json:"memory_used_gb"`
}
