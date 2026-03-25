package services

import (
	"k8s-cost-optimizer/internal/collector"
	"k8s-cost-optimizer/internal/models"
	"k8s-cost-optimizer/internal/cost"
	"k8s-cost-optimizer/internal/optimizer"
	"k8s-cost-optimizer/internal/forecasting"
)

func RunAnalysis() []optimizer.Recommendation {
	// 1. Fetch metrics da Prometheus
	metricsCollector := collector.FetchMetrics()

	// 2. Conversione esplicita in models.ContainerMetrics
	var metricsModels []models.ContainerMetrics
	for _, m := range metricsCollector {
		metricsModels = append(metricsModels, models.ContainerMetrics{
			Name:            m.Name,
			Namespace:       m.Namespace,
			Pod:             m.Pod,
			CPUUsage:        m.CPUUsage,
			CPURequest:      m.CPURequest,
			MemoryUsageMB:   m.MemoryUsageMB,
			MemoryRequestMB: m.MemoryRequestMB,
		})
	}

	// 3. Calcolo dei costi usando models
	costData := cost.CalculateCosts(metricsModels)

	// 4. Genera raccomandazioni
	recommendations := optimizer.GenerateRecommendations(metricsModels, costData)

	return recommendations
}


func RunFullAnalysis() interface{} {

	metrics := collector.FetchMetrics()

	// conversion models...
	var modelsMetrics []models.ContainerMetrics
	for _, m := range metrics {
		modelsMetrics = append(modelsMetrics, models.ContainerMetrics{
			Name:            m.Name,
			Namespace:       m.Namespace,
			Pod:             m.Pod,
			CPUUsage:        m.CPUUsage,
			CPURequest:      m.CPURequest,
			MemoryUsageMB:   m.MemoryUsageMB,
			MemoryRequestMB: m.MemoryRequestMB,
		})
	}

	costs := cost.CalculateCosts(modelsMetrics)
	recs := optimizer.GenerateRecommendations(modelsMetrics, costs)

	forecast := forecasting.ForecastCosts(costs)
	anomalies := forecasting.DetectAnomalies(costs)

	return map[string]interface{}{
		"recommendations": recs,
		"forecast":        forecast,
		"anomalies":       anomalies,
	}
}
