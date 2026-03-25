package models

// CostAnomaly rappresenta un'anomalia nei costi
type CostAnomaly struct {
	Container   string  `json:"container"`
	Namespace   string  `json:"namespace"`
	Pod         string  `json:"pod"`
	Value       float64 `json:"value"`
	Threshold   float64 `json:"threshold"`
	Difference  float64 `json:"difference"`
	Percentage  float64 `json:"percentage"`
	Description string  `json:"description"`
}
