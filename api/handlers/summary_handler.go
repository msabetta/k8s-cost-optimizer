package handlers

import (
	"encoding/json"
	"net/http"

	"k8s-cost-optimizer/internal/services"
)

// SummaryResponse rappresenta overview globale
type SummaryResponse struct {
	TotalSavings   float64 `json:"total_savings"`
	NumActions     int     `json:"num_actions"`
	TopSuggestion  string  `json:"top_suggestion"`
}

// GetSummary restituisce una panoramica completa
func GetSummary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 🔹 1. esegui analisi
	recs := services.RunAnalysis()

	if recs == nil {
		http.Error(w, "Failed to compute recommendations", http.StatusInternalServerError)
		return
	}

	// 🔹 2. calcola summary
	var totalSavings float64
	var topSuggestion string

	for i, r := range recs {
		totalSavings += r.Savings

		if i == 0 {
			topSuggestion = r.Action
		}
	}

	response := SummaryResponse{
		TotalSavings:  totalSavings,
		NumActions:    len(recs),
		TopSuggestion: topSuggestion,
	}

	// ✅ errcheck fix
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
