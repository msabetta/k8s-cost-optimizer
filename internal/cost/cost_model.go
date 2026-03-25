package cost

import "k8s-cost-optimizer/internal/models"

// CalculateCosts calcola costi per tutti i container
func CalculateCosts(metrics []models.ContainerMetrics) []models.CostData {

	pricing := DefaultPricing()

	var costs []models.CostData

	for _, m := range metrics {
		cost := allocateCost(m, pricing)
		costs = append(costs, cost)
	}

	return costs
}


func CalculateCostsWithPricing(metrics []models.ContainerMetrics, pricing Pricing) []models.CostData {

	//pricing := DefaultPricing()

	var costs []models.CostData

	for _, m := range metrics {
		cost := allocateCost(m, pricing)
		costs = append(costs, cost)
	}

	return costs
}

