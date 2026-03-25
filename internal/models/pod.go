package models

// Pod rappresenta un pod Kubernetes
type Pod struct {
	Name        string             `json:"name"`
	Namespace   string             `json:"namespace"`
	Node        string             `json:"node"`
	Containers  []ContainerMetrics `json:"containers"`
}