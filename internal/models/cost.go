package models

// CostData rappresenta costo per container
type CostData struct {
	Container   string  `json:"container"`
	HourlyCost  float64 `json:"hourly_cost"`
	DailyCost   float64 `json:"daily_cost"`
	MonthlyCost float64 `json:"monthly_cost"`
}