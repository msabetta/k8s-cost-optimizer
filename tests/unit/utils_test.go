package unit

import (
	"testing"

	"k8s-cost-optimizer/pkg/utils"
)

func TestRoundFloat(t *testing.T) {
	val := 1.23456
	res := utils.RoundFloat(val, 2)
	if res != 1.23 {
		t.Errorf("Expected 1.23, got %f", res)
	}
}

func TestPercentDifference(t *testing.T) {
	diff := utils.PercentDifference(120, 100)
	if diff != 20 {
		t.Errorf("Expected 20, got %f", diff)
	}
}
