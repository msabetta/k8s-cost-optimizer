package forecasting

import "k8s-cost-optimizer/internal/models"

// ForecastResult rappresenta previsione costo
type ForecastResult struct {
	CurrentHourly float64 `json:"current_hourly"`
	DailyEstimate float64 `json:"daily_estimate"`
	MonthlyEstimate float64 `json:"monthly_estimate"`
}

// ForecastCosts stima costi futuri
func ForecastCosts(costs []models.CostData) ForecastResult {

	var totalHourly float64

	for _, c := range costs {
		totalHourly += c.HourlyCost
	}

	daily := totalHourly * 24
	monthly := daily * 30

	return ForecastResult{
		CurrentHourly:  totalHourly,
		DailyEstimate:  daily,
		MonthlyEstimate: monthly,
	}
}
