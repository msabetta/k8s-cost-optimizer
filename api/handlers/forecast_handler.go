package handlers

import (
	"encoding/json"
	"net/http"

	"k8s-cost-optimizer/internal/collector"
	"k8s-cost-optimizer/internal/cost"
	"k8s-cost-optimizer/internal/forecasting"
	"k8s-cost-optimizer/internal/models"
)

// ForecastResponse struttura risposta API
type ForecastResponse struct {
	HourlyCost   float64                `json:"hourly_cost"`
	DailyCost    float64                `json:"daily_cost"`
	MonthlyCost  float64                `json:"monthly_cost"`
	Anomalies    []forecasting.Anomaly   `json:"anomalies"`
}

// GetForecast restituisce forecast + anomaly detection
func GetForecast(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 🔹 1. fetch metrics
	rawMetrics := collector.FetchMetrics()

	if rawMetrics == nil {
		http.Error(w, "Failed to fetch metrics", http.StatusInternalServerError)
		return
	}

	// 🔹 2. convert → models
	var metrics []models.ContainerMetrics
	for _, m := range rawMetrics {
		metrics = append(metrics, models.ContainerMetrics{
			Name:            m.Name,
			Namespace:       m.Namespace,
			Pod:             m.Pod,
			CPUUsage:        m.CPUUsage,
			CPURequest:      m.CPURequest,
			MemoryUsageMB:   m.MemoryUsageMB,
			MemoryRequestMB: m.MemoryRequestMB,
		})
	}

	// 🔹 3. cost calculation
	costs := cost.CalculateCosts(metrics)

	// 🔹 4. forecast
	fc := forecasting.ForecastCosts(costs)

	// 🔹 5. anomaly detection
	anomalies := forecasting.DetectAnomalies(costs)

	response := ForecastResponse{
		HourlyCost:  fc.CurrentHourly,
		DailyCost:   fc.DailyEstimate,
		MonthlyCost: fc.MonthlyEstimate,
		Anomalies:   anomalies,
	}

	// ✅ errcheck fix
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
