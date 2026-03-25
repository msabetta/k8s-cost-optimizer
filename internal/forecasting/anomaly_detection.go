package forecasting

import (
	"k8s-cost-optimizer/internal/models"
	"math"
)

// Anomaly rappresenta un'anomalia rilevata
type Anomaly struct {
	Container string  `json:"container"`
	CurrentCost float64 `json:"current_cost"`
	ExpectedCost float64 `json:"expected_cost"`
	Deviation   float64 `json:"deviation"`
}

// DetectAnomalies identifica costi anomali
func DetectAnomalies(costs []models.CostData) []Anomaly {

	var anomalies []Anomaly

	// 🔹 calcolo media
	var total float64
	for _, c := range costs {
		total += c.HourlyCost
	}
	avg := total / float64(len(costs))

	// 🔹 deviazione standard semplice
	var variance float64
	for _, c := range costs {
		variance += math.Pow(c.HourlyCost-avg, 2)
	}
	stdDev := math.Sqrt(variance / float64(len(costs)))

	// 🔹 detection (threshold = 2 deviazioni standard)
	for _, c := range costs {
		deviation := math.Abs(c.HourlyCost - avg)

		if deviation > 2*stdDev {
			anomalies = append(anomalies, Anomaly{
				Container:   c.Container,
				CurrentCost: c.HourlyCost,
				ExpectedCost: avg,
				Deviation:   deviation,
			})
		}
	}

	return anomalies
}
