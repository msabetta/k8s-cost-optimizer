package api

import (
	"k8s-cost-optimizer/api/handlers"
	"net/http"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/recommendations", handlers.GetRecommendations)
	mux.HandleFunc("/summary", http.HandlerFunc(handlers.GetSummary))
	mux.HandleFunc("/forecast", http.HandlerFunc(handlers.GetForecast))
	mux.HandleFunc("/costs", handlers.GetCosts)

	// 🔥 applica middleware globali
	return ChainMiddleware(
		mux,
		RecoveryMiddleware,
		LoggingMiddleware,
		JSONMiddleware,
	)
}
