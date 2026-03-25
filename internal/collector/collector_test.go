package collector

import "testing"

// func TestFetchMetrics(t *testing.T) {
// 	metrics := FetchMetrics()

// 	if metrics == nil {
// 		t.Errorf("Expected metrics, got nil")
// 	}

// 	if len(metrics) == 0 {
// 		t.Log("Warning: no metrics returned (check Prometheus)")
// 	}
// }

// func TestFetchMetrics(t *testing.T) {
// 	metrics := FetchMetrics()

// 	if metrics == nil {
// 		t.Fatalf("Expected non-nil slice")
// 	}

// 	if len(metrics) == 0 {
// 		t.Log("No metrics returned (Prometheus may be empty)")
// 	}
// }

func TestFetchMetrics(t *testing.T) {
	metrics := FetchMetrics()

	if metrics == nil {
		t.Fatalf("Expected non-nil slice")
	}

	if len(metrics) == 0 {
		t.Log("No metrics returned (Prometheus may be empty)")
	}
}
