package handlers

import (
	"encoding/json"
	"net/http"

	"k8s-cost-optimizer/internal/collector"
	"k8s-cost-optimizer/internal/cost"
	"k8s-cost-optimizer/internal/models"
)

// CostResponse rappresenta risposta API costi
type CostResponse struct {
	TotalHourly  float64           `json:"total_hourly"`
	TotalDaily   float64           `json:"total_daily"`
	TotalMonthly float64           `json:"total_monthly"`
	Breakdown    []models.CostData `json:"breakdown"`
}

// GetCosts restituisce costi per container + aggregati
func GetCosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 🔹 1. Fetch metrics (collector)
	rawMetrics := collector.FetchMetrics()

	if rawMetrics == nil {
		http.Error(w, "Failed to fetch metrics", http.StatusInternalServerError)
		return
	}

	// 🔹 2. Convert → models
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

	// 🔹 3. Calcolo costi
	costs := cost.CalculateCosts(metrics)

	// 🔹 4. Aggregazione
	var totalHourly float64
	for _, c := range costs {
		totalHourly += c.HourlyCost
	}

	totalDaily := totalHourly * 24
	totalMonthly := totalDaily * 30

	// 🔹 5. Response
	response := CostResponse{
		TotalHourly:  totalHourly,
		TotalDaily:   totalDaily,
		TotalMonthly: totalMonthly,
		Breakdown:    costs,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
