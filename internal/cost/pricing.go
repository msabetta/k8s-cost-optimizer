package cost

// Pricing rappresenta il costo unitario delle risorse
type Pricing struct {
	CPUPerCoreHour float64
	MemoryPerGBHour float64
}

// DefaultPricing (MVP - simile AWS)
func DefaultPricing() Pricing {
	return Pricing{
		CPUPerCoreHour:  0.031, // €/core/ora (approx AWS)
		MemoryPerGBHour: 0.004, // €/GB/ora
	}
}