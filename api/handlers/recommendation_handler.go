package handlers

import (
	"encoding/json"
	"net/http"

	"k8s-cost-optimizer/internal/services"
)

// GetRecommendations restituisce tutte le raccomandazioni
func GetRecommendations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	recs := services.RunAnalysis()

	if recs == nil {
		http.Error(w, "No recommendations available", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(recs); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
