package models

// Recommendation rappresenta suggerimento di ottimizzazione
type Recommendation struct {
	Container string  `json:"container"`
	Action    string  `json:"action"`
	Savings   float64 `json:"savings"`
	Priority  string  `json:"priority,omitempty"` // HIGH, MEDIUM, LOW
	Type      string  `json:"type,omitempty"`     // cpu, memory, idle, overprovision
}