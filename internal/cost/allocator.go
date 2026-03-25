package cost

import "k8s-cost-optimizer/internal/models"

// // CostData rappresenta il costo per container
// type CostData struct {
// 	Container string  `json:"container"`
// 	HourlyCost float64 `json:"hourly_cost"`
// }

// allocateCost calcola costo per singolo container
func allocateCost(m models.ContainerMetrics, pricing Pricing) models.CostData {

	// CPU cost
	cpuCost := m.CPUUsage * pricing.CPUPerCoreHour

	// Memory cost (MB → GB)
	memoryGB := m.MemoryUsageMB / 1024
	memCost := memoryGB * pricing.MemoryPerGBHour

  total := cpuCost + memCost
	daily := total * 24
	monthly := daily * 30

	return models.CostData{
		Container:   m.Namespace + "/" + m.Pod + "/" + m.Name,
		HourlyCost:  total,
		DailyCost:   daily,
		MonthlyCost: monthly,
	}

}
