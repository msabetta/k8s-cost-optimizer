package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"k8s-cost-optimizer/internal/collector"
)

func TestFetchMetricsIntegration(t *testing.T) {
	// server fake Prometheus
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"success","data":{"result":[]}}`))
	}))
	defer ts.Close()

	client := collector.NewPrometheusClient(ts.URL)
	metrics, err := client.FetchMetrics("container_cpu_usage_seconds_total")
	if err != nil {
		t.Fatalf("Failed to fetch metrics: %v", err)
	}

	if metrics == nil {
		t.Errorf("Expected metrics, got nil")
	}
}
