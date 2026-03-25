package integration

import (
	"testing"

	"k8s-cost-optimizer/internal/cost"
	"k8s-cost-optimizer/internal/models"
)

func TestCalculateCosts(t *testing.T) {
	metrics := []models.ContainerMetrics{
		{CPUUsage: 1.0, MemoryUsageMB: 512},
	}
	pricing := cost.Pricing{CPUPerCoreHour: 0.05, 	MemoryPerGBHour: 0.01}

	result := cost.CalculateCostsWithPricing(metrics,pricing)
	if len(result) != 1 {
		t.Errorf("Expected 1 cost entry, got %d", len(result))
	}
	if result[0].DailyCost != 0.05 {
		t.Errorf("Expected CPU cost 0.05, got %f", result[0].DailyCost)
	}
}
