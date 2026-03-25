package services

import (
	"k8s-cost-optimizer/internal/optimizer"
	"sort"
)

type SavingsSummary struct {
	TotalSavings float64                    `json:"total_savings"`
	TopActions   []optimizer.Recommendation `json:"top_actions"`
}

// GetSavingsSummary ritorna summary completo
func GetSavingsSummary(recs []optimizer.Recommendation) SavingsSummary {

	var total float64
	for _, r := range recs {
		total += r.Savings
	}

	// 🔹 ordina per savings (desc)
	sort.Slice(recs, func(i, j int) bool {
		return recs[i].Savings > recs[j].Savings
	})

	// 🔹 prendi top 5
	top := recs
	if len(recs) > 5 {
		top = recs[:5]
	}

	return SavingsSummary{
		TotalSavings: total,
		TopActions:   top,
	}
}